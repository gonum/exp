// Copyright ©2017 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package linsolve

import (
	"math"

	"gonum.org/v1/gonum/blas"
	"gonum.org/v1/gonum/blas/blas64"
	"gonum.org/v1/gonum/floats"
)

// GMRES implements the Generalized Minimum Residual method with the modified
// Gram-Schmidt orthogonalization for solving systems of linear equations
//  A*x = b,
// where A is a square matrix (not necessarily symmetric). It uses restarts to control
// memory requirements.
//
// References:
//  - Barrett, R. et al. (1994). Section 2.3.4 Generalized Minimal Residual
//    (GMRES). In Templates for the Solution of Linear Systems: Building Blocks
//    for Iterative Methods (2nd ed.) (pp. 17-19). Philadelphia, PA: SIAM.
//    Retrieved from http://www.netlib.org/templates/templates.pdf
//  - Saad, Y., and Schultz, M. (1986). GMRES: A generalized minimal residual
//    algorithm for solving nonsymmetric linear systems. SIAM J. Sci. Stat.
//    Comput., 7(3), 856. doi:10.6028/jres.049.044
//    Retrieved from https://web.stanford.edu/class/cme324/saad-schultz.pdf
type GMRES struct {
	// Restart is the restart parameter which limits the computation and
	// storage costs. It must hold that
	//  1 <= Restart <= n
	// where n is the dimension of the problem. If Restart is 0, n will be
	// used instead. This guarantess convergence of GMRES and increases
	// robustness. Many specific problems however, particularly for large
	// n, will benefit in efficiency by setting Restart to
	// a problem-dependent value less than n.
	Restart int

	// m is the used value of Restart.
	m int
	// vt is an (m+1)×n matrix. It corresponds to V^T in
	// standard descriptions of GMRES. Its rows form an orthonormal basis of the
	// Krylov subspace.
	vt []float64
	// ht is an m×(m+1) lower Hessenberg matrix. It corresponds to H^T in
	// standard descriptions of GMRES.
	ht []float64
	// givs holds Givens rotations that are used to reduce H to
	// upper-triangular form.
	givs []givens

	x  []float64
	av []float64
	s  []float64
	y  []float64

	k      int // Loop variable for inner iterations.
	resume int
}

// givens is a Givens rotation.
type givens struct {
	c, s float64
}

// Init initializes the data for a linear solve. See the Method interface for more details.
func (g *GMRES) Init(x, residual []float64) {
	dim := len(x)
	if dim == 0 {
		panic("gmres: dimension not positive")
	}
	if len(residual) != dim {
		panic("gmres: slice length mismatch")
	}

	g.m = g.Restart
	if g.m == 0 {
		g.m = dim
	}
	if g.m <= 0 || dim < g.m {
		panic("gmres: invalid value of Restart")
	}

	g.x = reuse(g.x, dim)
	copy(g.x, x)

	ldv := dim
	g.vt = reuse(g.vt, (g.m+1)*ldv)
	ldh := g.m + 1
	g.ht = reuse(g.ht, g.m*ldh)

	if cap(g.givs) < g.m {
		g.givs = make([]givens, g.m)
	} else {
		g.givs = g.givs[:g.m]
		for i := range g.givs {
			g.givs[i].c = 0
			g.givs[i].s = 0
		}
	}

	g.s = reuse(g.s, g.m+1)
	g.y = reuse(g.y, dim)
	// Use g.y for storing the initial residual to avoid having and
	// allocating an extra slice.
	copy(g.y, residual)
	g.av = reuse(g.av, dim)

	g.resume = 1
}

// Iterate performs an iteration of the linear solve. See the Method interface for more details.
//
// GMRES will command the following operations:
//  MulVec
//  PreconSolve
//  CheckResidualNorm
//  MajorIteration
//  NoOperation
func (g *GMRES) Iterate(ctx *Context) (Operation, error) {
	switch g.resume {
	case 1:
		// g.y contains the initial residual.
		copy(ctx.Src, g.y)
		g.resume = 2
		// Solve M^{-1} * r_0.
		return PreconSolve, nil
	case 2:
		v0 := g.vcol(0)
		// v_0 = M^{-1} * r_0
		copy(v0, ctx.Dst)
		// Normalize v_0.
		norm := floats.Norm(v0, 2)
		floats.Scale(1/norm, v0)
		// Initialize s to the elementary vector e_1 scaled by norm.
		for i := range g.s {
			g.s[i] = 0
		}
		g.s[0] = norm

		// for k := 0; k < m; k++ {
		g.k = 0
		fallthrough
	case 3:
		copy(ctx.Src, g.vcol(g.k))
		g.resume = 4
		// Compute A * v_k.
		return MulVec, nil
	case 4:
		copy(ctx.Src, ctx.Dst)
		g.resume = 5
		// Solve M^{-1} * (A * v_k).
		return PreconSolve, nil
	case 5:
		// v_{k+1} = M^{-1} * (A * v_k)
		copy(g.vcol(g.k+1), ctx.Dst)
		// Construct the k-th column of the upper Hessenberg matrix H
		// using the modified Gram-Schmidt process to make v_{k+1}
		// orthonormal to the first k+1 columns of V.
		g.modifiedGS(g.k+1, g.vt, g.vcol(g.k+1), g.hcol(g.k))
		// Reduce H back to upper triangular form and update
		// the vector s.
		g.qr(g.k, g.givs, g.hcol(g.k), g.s)
		// Check the approximate residual norm.
		ctx.ResidualNorm = math.Abs(g.s[g.k+1])
		g.resume = 6
		return CheckResidualNorm, nil
	case 6:
		g.k++
		if g.k < g.m && !ctx.Converged {
			// Continue the inner for loop.
			g.resume = 3
			return NoOperation, nil
		}
		// Either restart or convergence, we have to update the solution.
		g.updateSolution(g.k, g.x)
		copy(ctx.X, g.x)
		if ctx.Converged {
			g.resume = 0
			return MajorIteration, nil
		}
		// We are restarting, so we have to also compute the residual.
		g.resume = 7
		return ComputeResidual, nil
	case 7:
		copy(g.y, ctx.Dst)
		g.resume = 1
		return MajorIteration, nil

	default:
		panic("gmres: Init not called")
	}
}

// modifiedGS orthonormalizes the vector w with respect to the rows of the k×n
// matrix vt using the modified Gram-Schmidt algorithm, and stores the
// coefficients and scales in the vector hk.
func (g *GMRES) modifiedGS(k int, vt []float64, w, hk []float64) {
	n := len(w)
	for j := 0; j < k; j++ {
		vj := vt[j*n : (j+1)*n]
		hjk := floats.Dot(vj, w)
		hk[j] = hjk                   // H[j,k] = v_j · v_{k+1}
		floats.AddScaled(w, -hjk, vj) // v_{k+1} -= H[j,k] * v_j
	}
	wnorm := floats.Norm(w, 2)
	hk[k] = wnorm            // H[k+1,k] = |v_{k+1}|
	floats.Scale(1/wnorm, w) // Normalize v_{k+1}.
}

// qr applies and computes Givens rotations to zero out (k+1)-th elements of the
// vector hk.
func (g *GMRES) qr(k int, givs []givens, hk, s []float64) {
	bi := blas64.Implementation()
	// Apply previous Givens rotations to the k-th row of H.
	for i := 0; i < k; i++ {
		bi.Drot(1, hk[i:], 1, hk[i+1:], 1, givs[i].c, givs[i].s)
	}
	// Compute the k-th Givens rotation that zeroes H[k+1,k].
	givs[k].c, givs[k].s, _, _ = bi.Drotg(hk[k], hk[k+1])
	// Apply the k-th Givens rotation to (hk[k], hk[k+1]).
	bi.Drot(1, hk[k:], 1, hk[k+1:], 1, givs[k].c, givs[k].s)
	// Apply the k-th Givens rotation to (s[k], s[k+1]).
	bi.Drot(1, s[k:], 1, s[k+1:], 1, givs[k].c, givs[k].s)
}

// vcol returns a view of the j-th column of V.
func (g *GMRES) vcol(j int) []float64 {
	ldv := len(g.av)
	return g.vt[j*ldv : (j+1)*ldv]
}

// hcol returns a view of the j-th column of H.
func (g *GMRES) hcol(j int) []float64 {
	ldh := g.m + 1
	return g.ht[j*ldh : (j+1)*ldh]
}

// updateSolution updates the solution vector x.
func (g *GMRES) updateSolution(k int, x []float64) {
	y := g.y[:k]
	copy(y, g.s)

	// Solve H*y = s for upper-triangular H.
	// Note that we are actually storing H^T which is lower-triangular so we
	// need to adjust the arguments accordingly.
	bi := blas64.Implementation()
	ldh := g.m + 1
	bi.Dtrsv(blas.Lower, blas.Trans, blas.NonUnit, k, g.ht, ldh, y, 1)

	// Update the current solution vector x.
	for j, yj := range y {
		vj := g.vcol(j)
		floats.AddScaled(x, yj, vj) // x += y_j * v_j
	}
}
