// Copyright ©2017 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package linsolve

import (
	"math"

	"gonum.org/v1/gonum/floats"
)

// BiCGStab implements the BiConjugate Gradient Stabilized method with
// preconditioning for solving systems of linear equations
//  A*x = b,
// where A is a square, possibly nonsymmetric matrix.
//
// References:
//  - Barrett, R. et al. (1994). Section 2.3.8 BiConjugate Gradient Stabilized (Bi-CGSTAB).
//    In Templates for the Solution of Linear Systems: Building Blocks
//    for Iterative Methods (2nd ed.) (pp. 24-25). Philadelphia, PA: SIAM.
//    Retrieved from http://www.netlib.org/templates/templates.pdf
type BiCGStab struct {
	p    []float64
	phat []float64
	rt   []float64
	shat []float64
	t    []float64
	v    []float64

	rho, rhoPrev float64
	alpha        float64
	omega        float64

	first  bool
	resume int
}

// Init initializes the data for a linear solve. See the Method interface for more details.
func (b *BiCGStab) Init(dim int) {
	if dim <= 0 {
		panic("bicgstab: dimension not positive")
	}

	b.p = reuse(b.p, dim)
	b.phat = reuse(b.phat, dim)
	b.rt = reuse(b.rt, dim)
	b.shat = reuse(b.shat, dim)
	b.t = reuse(b.t, dim)
	b.v = reuse(b.v, dim)

	b.first = true
	b.resume = 1
}

// Iterate performs an iteration of the linear solve. See the Method interface for more details.
//
// BiCGStab will command the following operations:
//  MulVec
//  PreconSolve
//  CheckResidual
//  MajorIteration
//  NoOperation
func (b *BiCGStab) Iterate(ctx *Context) (Operation, error) {
	switch b.resume {
	case 1:
		if b.first {
			copy(b.rt, ctx.Residual)
		}
		b.rho = floats.Dot(b.rt, ctx.Residual)
		if math.Abs(b.rho) < breakdownTol {
			b.resume = 0
			return NoOperation, &BreakdownError{math.Abs(b.rho), breakdownTol}
		}
		if b.first {
			b.first = false
			copy(b.p, ctx.Residual)
		} else {
			beta := (b.rho / b.rhoPrev) * (b.alpha / b.omega)
			floats.AddScaled(b.p, -b.omega, b.v)
			floats.Scale(beta, b.p)
			floats.Add(b.p, ctx.Residual)
		}
		// Solve M^{-1} * p_i.
		copy(ctx.Src, b.p)
		b.resume = 2
		return PreconSolve, nil
	case 2:
		copy(b.phat, ctx.Dst)
		// Compute A * \hat{p}_i.
		copy(ctx.Src, b.phat)
		b.resume = 3
		return MulVec, nil
	case 3:
		copy(b.v, ctx.Dst)
		rtv := floats.Dot(b.rt, b.v)
		if rtv == 0 {
			b.resume = 0
			return NoOperation, &BreakdownError{}
		}
		b.alpha = b.rho / rtv
		// Form the residual and X so that we can check for tolerance early.
		floats.AddScaled(ctx.X, b.alpha, b.phat)
		floats.AddScaled(ctx.Residual, -b.alpha, b.v)
		b.resume = 4
		return CheckResidual, nil
	case 4:
		if ctx.Converged {
			b.resume = 0
			return MajorIteration, nil
		}
		// Solve M^{-1} * r_i.
		copy(ctx.Src, ctx.Residual)
		b.resume = 5
		return PreconSolve, nil
	case 5:
		copy(b.shat, ctx.Dst)
		// Compute A * \hat{s}_i.
		copy(ctx.Src, b.shat)
		b.resume = 6
		return MulVec, nil
	case 6:
		copy(b.t, ctx.Dst)
		b.omega = floats.Dot(b.t, ctx.Residual) / floats.Dot(b.t, b.t)
		floats.AddScaled(ctx.X, b.omega, b.shat)
		floats.AddScaled(ctx.Residual, -b.omega, b.t)
		b.resume = 7
		return CheckResidual, nil
	case 7:
		if !ctx.Converged && math.Abs(b.omega) < breakdownTol {
			b.resume = 0
			return NoOperation, &BreakdownError{math.Abs(b.omega), breakdownTol}
		}
		b.rhoPrev = b.rho
		b.resume = 1
		return MajorIteration, nil

	default:
		panic("bicgstab: Init not called")
	}
}
