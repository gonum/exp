// Copyright ©2017 The gonum Authors. All rights reserved.
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
// Gram-Schmidt orthogonalization. It uses restarts to control storage
// requirements.
type GMRES struct {
	// Restart is the restart parameter.
	// It must be 0 <= Restart <= dim.
	// If it is 0, it will be set to dim.
	Restart int

	resume int

	s  []float64
	y  []float64
	av []float64

	j    int       // Counter for inner iterations.
	v    []float64 // dim×(Restart+1) matrix V.
	ldv  int
	h    []float64 // (Restart+1)×Restart matrix H.
	ldh  int
	givs []givens // Givens rotations.
}

type givens struct {
	c, s float64
}

// Init implements the Method interface.
func (g *GMRES) Init(dim int) {
	if dim <= 0 {
		panic("gmres: dimension not positive")
	}

	if g.Restart == 0 {
		g.Restart = dim
	}
	if g.Restart <= 0 || dim < g.Restart {
		panic("gmres: invalid value of Restart")
	}
	k := g.Restart

	g.s = reuse(g.s, k+1)
	g.y = reuse(g.y, dim)
	g.av = reuse(g.av, dim)

	g.ldv = dim
	g.v = reuse(g.v, g.ldv*(k+1))
	g.ldh = k + 1
	g.h = reuse(g.h, g.ldh*k)

	if cap(g.givs) < k {
		g.givs = make([]givens, k)
	} else {
		g.givs = g.givs[:k]
	}

	g.resume = 1
}

// Iterate implements the Method interface.
func (g *GMRES) Iterate(ctx *Context) (Operation, error) {
	n := len(ctx.X)

	switch g.resume {
	case 1:
		// Construct the first column of V.
		ctx.Src = ctx.Residual
		ctx.Dst = g.v[:n]
		g.resume = 2
		return PSolve, nil
		// Solve M*V[:,0] = r.
	case 2:
		// Normalize V[:,0].
		v0 := g.v[:n]
		norm := floats.Norm(v0, 2)
		floats.Scale(1/norm, v0)
		// Initialize s to the elementary vector e_1 scaled by norm.
		for i := range g.s {
			g.s[i] = 0
		}
		g.s[0] = norm

		// for j := 0; j < Restart; j++ {
		g.j = 0
		fallthrough
	case 3:
		ctx.Src = g.v[g.j*g.ldv : g.j*g.ldv+n] // j-th column of V
		ctx.Dst = g.av
		g.resume = 4
		return MatVec, nil
		// Compute A*V[:,j].
	case 4:
		ctx.Src = g.av
		ctx.Dst = g.v[(g.j+1)*g.ldv : (g.j+1)*g.ldv+n] // (j+1)-th column f V
		g.resume = 5
		return PSolve, nil
		// Solve M*w = A*V[:,j].
	case 5:
		j := g.j
		ldv := g.ldv
		w := g.v[(j+1)*ldv : (j+1)*ldv+n]
		H := g.h
		ldh := g.ldh
		Hj := H[j*ldh : j*ldh+g.Restart+1] // j-th column of H.

		// Construct j-th column of the upper Hessenberg matrix using
		// the Gram-Schmidt process on V and w so that it is orthonormal
		// to the previous j-1 columns.
		for k := 0; k <= j; k++ {
			vk := g.v[k*ldv : k*ldv+n] // k-th column pf V.
			hkj := floats.Dot(vk, w)
			Hj[k] = hkj                   // H[k,j] = V[:,k]^T * V[:,j+1]
			floats.AddScaled(w, -hkj, vk) // w -= H[k,j] * V[:,k]
		}
		wnorm := floats.Norm(w, 2)
		Hj[j+1] = wnorm          // H[j+1,j] = |w|
		floats.Scale(1/wnorm, w) // Normalize V[:,j+1].

		// Apply j Givens rotation matrices to the j-th
		// column of H.
		for i := 0; i < j; i++ {
			Hj[i], Hj[i+1] = rotvec(g.givs[i], Hj[i], Hj[i+1])
		}
		// Compute the (j+1)st Givens rotation that zeroes H[j+1,j].
		g.givs[j] = drotg(Hj[j], Hj[j+1])
		// Apply the (j+1)st Givens rotation.
		Hj[j], Hj[j+1] = rotvec(g.givs[j], Hj[j], Hj[j+1])

		// Apply the (j+1)st Givens rotation to (s[j], s[j+1]).
		s := g.s
		s[j], s[j+1] = rotvec(g.givs[j], s[j], s[j+1])
		// Approximate the residual norm and check for convergence.
		ctx.ResidualNorm = math.Abs(s[j+1])
		ctx.Src = nil
		ctx.Dst = nil
		ctx.Converged = false
		g.resume = 6
		return CheckConvergence, nil
	case 6:
		if ctx.Converged {
			// Compute final approximate solution x and finish.
			g.update(ctx.X)
			// TODO: Should we also call ComputeResidual? It depends
			// on how we specify the reverse-communication protocol.
			// If initially Context.Residual must be valid, then it
			// might make sense to finish the iterations again with
			// an up-to-date Residual.
			g.resume = 0 // Calling Iterate again without Init will panic.
			return EndIteration, nil
		}
		g.j++
		if g.j < g.Restart {
			// Continue the inner for loop.
			g.resume = 3
			return EndIteration, nil
		}
		// End the inner for loop.
		fallthrough
	case 7:
		// Adjust j to point to last valid column of V.
		g.j--
		// We are going to restart, so we need to update the approximate
		// solution vector x and the residual.
		g.update(ctx.X)
		g.resume = 8
		return ComputeResidual, nil
	case 8:
		ctx.Converged = false
		ctx.ResidualNorm = floats.Norm(ctx.Residual, 2)
		g.resume = 9
		return CheckConvergence, nil
	case 9:
		if ctx.Converged {
			g.resume = 0 // Calling Iterate again without Init will panic.
		} else {
			g.resume = 1 // Restart (continue the outer for loop).
		}
		return EndIteration, nil

	default:
		panic("gmres: Init not called")
	}
}

// update computes the current solution vector and stores it in x.
func (g *GMRES) update(x []float64) {
	k := g.j + 1 // Number of valid columns of V.
	y := g.y[:k]
	copy(y, g.s[:k])
	bi := blas64.Implementation()
	// Solve H*y = s for upper triangular H.
	// H is upper triangular but stored in column-major order. Dtrsv
	// expects row-major so adjust the arguments accordingly.
	bi.Dtrsv(blas.Lower, blas.Trans, blas.NonUnit, k, g.h, g.ldh, y, 1)
	// Compute current solution vector x.
	n := len(x)
	for j, yj := range y {
		vj := g.v[j*g.ldv : j*g.ldv+n] // j-th column of V
		floats.AddScaled(x, yj, vj)    // x += y_j * V_j
	}
}

// drotg returns Givens plane rotation.
func drotg(a, b float64) givens {
	if b == 0 {
		return givens{c: 1, s: 0}
	}
	if math.Abs(b) > math.Abs(a) {
		tmp := -a / b
		s := 1 / math.Sqrt(1+tmp*tmp)
		return givens{c: tmp * s, s: s}
	}
	tmp := -b / a
	c := 1 / math.Sqrt(1+tmp*tmp)
	return givens{c: c, s: tmp * c}
}

// rotvec applies Givens rotation g to the vector [x,y] and returns the result.
func rotvec(g givens, x, y float64) (rx, ry float64) {
	rx = g.c*x - g.s*y
	ry = g.s*x + g.c*y
	return
}
