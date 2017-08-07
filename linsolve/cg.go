// Copyright ©2016 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package linsolve

import "github.com/gonum/floats"

// CG implements the Conjugate Gradient iterative method with
// preconditioning for solving systems of linear equations
//  A*x = b,
// where A is a symmetric positive definite matrix.
//
// CG needs MatVec and PSolve matrix operations.
type CG struct {
	first  bool
	resume int

	rho, rhoPrev float64

	z  []float64
	p  []float64
	ap []float64
}

// Init implements the Method interface.
func (cg *CG) Init(dim int) {
	if dim <= 0 {
		panic("cg: dimension not positive")
	}

	cg.z = reuse(cg.z, dim)
	cg.p = reuse(cg.p, dim)
	cg.ap = reuse(cg.ap, dim)
	cg.first = true
	cg.resume = 1
}

// Iterate implements the Method interface.
func (cg *CG) Iterate(ctx *Context) (Operation, error) {
	switch cg.resume {
	case 1:
		ctx.Src = ctx.Residual
		ctx.Dst = cg.z
		cg.resume = 2
		return PSolve, nil
		// Solve M*z = r_{i-1}
	case 2:
		cg.rho = floats.Dot(ctx.Residual, cg.z) // ρ_i = r_{i-1} · z
		if !cg.first {
			beta := cg.rho / cg.rhoPrev        // β = ρ_i / ρ_{i-1}
			floats.AddScaled(cg.z, beta, cg.p) // z = z + β p_{i-1}
		}
		copy(cg.p, cg.z) // p_i = z

		ctx.Src = cg.p
		ctx.Dst = cg.ap
		cg.resume = 3
		return MatVec, nil
		// Compute A*p_i
	case 3:
		alpha := cg.rho / floats.Dot(cg.p, cg.ap)     // α = ρ_i / (p_i · Ap_i)
		floats.AddScaled(ctx.Residual, -alpha, cg.ap) // r_i = r_{i-1} - α A p_i
		floats.AddScaled(ctx.X, alpha, cg.p)          // x_i = x_{i-1} + α p_i

		ctx.Src = nil
		ctx.Dst = nil
		ctx.ResidualNorm = floats.Norm(ctx.Residual, 2)
		ctx.Converged = false
		cg.resume = 4
		return CheckConvergence, nil
	case 4:
		if ctx.Converged {
			cg.resume = 0 // Calling Iterate again without Init will panic.
			return EndIteration, nil
		}
		cg.rhoPrev = cg.rho
		cg.first = false
		cg.resume = 1
		return EndIteration, nil

	default:
		panic("cg: Init not called")
	}
}
