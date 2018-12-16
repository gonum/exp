// Copyright ©2017 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package linsolve

import (
	"errors"
	"math"

	"gonum.org/v1/gonum/floats"
)

// BiCG implements the Bi-Conjugate Gradient method with
// preconditioning for solving systems of linear equations
//  A*x = b,
// where A is a square, possibly nonsymmetric matrix.
//
// References:
//  - Barrett, R. et al. (1994). Section 2.3.5 BiConjugate Gradient (BiCG).
//    In Templates for the Solution of Linear Systems: Building Blocks
//    for Iterative Methods (2nd ed.) (pp. 19-20). Philadelphia, PA: SIAM.
//    Retrieved from http://www.netlib.org/templates/templates.pdf
type BiCG struct {
	first  bool
	resume int

	rho, rhoPrev float64

	rt    []float64
	z, zt []float64
	p, pt []float64
}

// Init initializes the data for a linear solve. See the Method interface for more details.
func (b *BiCG) Init(dim int) {
	if dim <= 0 {
		panic("bicg: dimension not positive")
	}

	b.rt = reuse(b.rt, dim)
	b.z = reuse(b.z, dim)
	b.zt = reuse(b.zt, dim)
	b.p = reuse(b.p, dim)
	b.pt = reuse(b.pt, dim)

	b.first = true
	b.resume = 1
}

// Iterate performs an iteration of the linear solve. See the Method interface for more details.
//
// BiCG will command the following operations:
//  MulVec
//  MulVec|Trans
//  PreconSolve
//  PreconSolve|Trans
//  CheckResidual
//  MajorIteration
//  NoOperation
func (b *BiCG) Iterate(ctx *Context) (Operation, error) {
	switch b.resume {
	case 1:
		if b.first {
			copy(b.rt, ctx.Residual)
		}
		// Solve M^{-1} * r_{i-1}.
		copy(ctx.Src, ctx.Residual)
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
		if math.Abs(b.rho) < rhoBreakdownTol {
			b.resume = 0
			return NoOperation, errors.New("bicg: rho breakdown")
		}
		if b.first {
			copy(b.p, b.z)
			copy(b.pt, b.zt)
		} else {
			beta := b.rho / b.rhoPrev
			floats.AddScaledTo(b.p, b.z, beta, b.p)
			floats.AddScaledTo(b.pt, b.zt, beta, b.pt)
		}
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
		floats.AddScaled(b.rt, -alpha, b.zt)
		floats.AddScaled(ctx.X, alpha, b.p)
		floats.AddScaled(ctx.Residual, -alpha, b.z)
		b.resume = 6
		return CheckResidual, nil
	case 6:
		b.rhoPrev = b.rho
		b.first = false
		b.resume = 1
		return MajorIteration, nil

	default:
		panic("bicg: Init not called")
	}
}
