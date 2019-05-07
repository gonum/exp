// Copyright ©2017 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package linsolve

import (
	"math"

	"gonum.org/v1/gonum/floats"
)

// BiCG implements the Bi-Conjugate Gradient method with preconditioning for
// solving systems of linear equations
//  A * x = b,
// where A is a nonsymmetric, nonsingular matrix. It uses limited memory storage
// but the convergence may be irregular and the method may break down. BiCG
// requires a multiplication with A and A^T at each iteration.
//
// References:
//  - Barrett, R. et al. (1994). Section 2.3.5 BiConjugate Gradient (BiCG).
//    In Templates for the Solution of Linear Systems: Building Blocks
//    for Iterative Methods (2nd ed.) (pp. 19-20). Philadelphia, PA: SIAM.
//    Retrieved from http://www.netlib.org/templates/templates.pdf
type BiCG struct {
	x     []float64
	r, rt []float64
	p, pt []float64
	z, zt []float64

	rho, rhoPrev float64

	resume int
}

// Init initializes the data for a linear solve. See the Method interface for more details.
func (b *BiCG) Init(x, residual []float64) {
	dim := len(x)
	if dim == 0 {
		panic("bicg: dimension not positive")
	}
	if len(residual) != dim {
		panic("bicg: slice length mismatch")
	}

	b.x = reuse(b.x, dim)
	copy(b.x, x)
	b.r = reuse(b.r, dim)
	copy(b.r, residual)
	b.rt = reuse(b.rt, dim)
	copy(b.rt, b.r)
	b.p = reuse(b.p, dim)
	b.pt = reuse(b.pt, dim)
	b.z = reuse(b.z, dim)
	b.zt = reuse(b.zt, dim)

	b.rhoPrev = 1

	b.resume = 1
}

// Iterate performs an iteration of the linear solve. See the Method interface for more details.
//
// BiCG will command the following operations:
//  MulVec
//  MulVec|Trans
//  PreconSolve
//  PreconSolve|Trans
//  CheckResidualNorm
//  MajorIteration
//  NoOperation
func (b *BiCG) Iterate(ctx *Context) (Operation, error) {
	switch b.resume {
	case 1:
		// Solve M^{-1} * r_{i-1}.
		copy(ctx.Src, b.r)
		b.resume = 2
		return PreconSolve, nil
	case 2:
		copy(b.z, ctx.Dst)
		// Solve M^{-T} * rt_{i-1}.
		copy(ctx.Src, b.rt)
		b.resume = 3
		return PreconSolve | Trans, nil
	case 3:
		copy(b.zt, ctx.Dst)
		b.rho = floats.Dot(b.z, b.rt)
		if math.Abs(b.rho) < breakdownTol {
			b.resume = 0
			return NoOperation, &BreakdownError{math.Abs(b.rho), breakdownTol}
		}
		beta := b.rho / b.rhoPrev
		floats.AddScaledTo(b.p, b.z, beta, b.p)
		floats.AddScaledTo(b.pt, b.zt, beta, b.pt)
		// Compute A * p.
		copy(ctx.Src, b.p)
		b.resume = 4
		return MulVec, nil
	case 4:
		// z is overwritten and reused.
		copy(b.z, ctx.Dst)
		// Compute A^T * pt.
		copy(ctx.Src, b.pt)
		b.resume = 5
		return MulVec | Trans, nil
	case 5:
		// zt is overwritten and reused.
		copy(b.zt, ctx.Dst)
		alpha := b.rho / floats.Dot(b.pt, b.z)
		floats.AddScaled(ctx.X, alpha, b.p)
		floats.AddScaled(b.rt, -alpha, b.zt)
		floats.AddScaled(b.r, -alpha, b.z)
		ctx.ResidualNorm = floats.Norm(b.r, 2)
		b.resume = 6
		return CheckResidualNorm, nil
	case 6:
		if ctx.Converged {
			b.resume = 0
			return MajorIteration, nil
		}
		b.rhoPrev = b.rho
		b.resume = 1
		return MajorIteration, nil

	default:
		panic("bicg: Init not called")
	}
}
