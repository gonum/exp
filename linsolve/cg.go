// Copyright ©2016 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package linsolve

import (
	"gonum.org/v1/gonum/floats"
)

// CG implements the Conjugate Gradient iterative method with
// preconditioning for solving systems of linear equations
//  A*x = b,
// where A is a symmetric positive definite matrix.
//
// References:
//  - Barrett, Richard et al. (1994). Section 2.3.1 Conjugate Gradient Method (CG).
//    In Templates for the Solution of Linear Systems: Building Blocks for
//    Iterative Methods (2nd ed.) (pp. 12-15). Philadelphia, PA: SIAM.
//    Retrieved from http://www.netlib.org/templates/templates.pdf
//  - Hestenes, M., and Stiefel, E. (1952). Methods of conjugate gradients for
//    solving linear systems. Journal of Research of the National Bureau of
//    Standards, 49(6), 409. doi:10.6028/jres.049.044
//  - Málek, J. and Strakoš, Z. (2015). Preconditioning and the Conjugate Gradient
//    Method in the Context of Solving PDEs. Philadelphia, PA: SIAM.
type CG struct {
	rho, rhoPrev float64

	p []float64

	first  bool
	resume int
}

// Init initializes the data for a linear solve. See the Method interface for more details.
func (cg *CG) Init(dim int) {
	if dim <= 0 {
		panic("cg: dimension not positive")
	}

	cg.p = reuse(cg.p, dim)
	cg.first = true
	cg.resume = 1
}

// Iterate performs an iteration of the linear solve. See the Method interface for more details.
//
// CG will command the following operations:
//  MulVec
//  PreconSolve
//  CheckResidual
//  MajorIteration
func (cg *CG) Iterate(ctx *Context) (Operation, error) {
	switch cg.resume {
	case 1:
		copy(ctx.Src, ctx.Residual)
		cg.resume = 2
		// Compute z_{i-1} = M^{-1} * r_{i-1}.
		return PreconSolve, nil
	case 2:
		z := ctx.Dst
		cg.rho = floats.Dot(ctx.Residual, z) // ρ_{i-1} = r_{i-1} · z_{i-1}
		if cg.first {
			copy(cg.p, z) // p_1 = z_0
		} else {
			beta := cg.rho / cg.rhoPrev             // β_{i-1} = ρ_{i-1} / ρ_{i-2}
			floats.AddScaledTo(cg.p, z, beta, cg.p) // p_i = z_{i-1} + β p_{i-1}
		}
		copy(ctx.Src, cg.p)
		cg.resume = 3
		// Compute A * p_i.
		return MulVec, nil
	case 3:
		ap := ctx.Dst
		alpha := cg.rho / floats.Dot(cg.p, ap)     // α_i = ρ_{i-1} / (p_i · A p_i)
		floats.AddScaled(ctx.Residual, -alpha, ap) // r_i = r_{i-1} - α A p_i
		floats.AddScaled(ctx.X, alpha, cg.p)       // x_i = x_{i-1} + α p_i
		cg.resume = 4
		return CheckResidual, nil
	case 4:
		cg.rhoPrev = cg.rho
		cg.first = false
		cg.resume = 1
		return MajorIteration, nil

	default:
		panic("cg: Init not called")
	}
}
