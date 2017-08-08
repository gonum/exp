// Copyright ©2017 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package linsolve

import (
	"errors"
	"math"

	"github.com/gonum/floats"
)

// BCGSTAB implements the Biconjugate Gradient Stabilized iterative method with
// preconditioning for solving systems of linear equations
//  Ax = b,
// where A is a non-symmetric matrix. For symmetric positive definite systems
// use CG.
//
// BCGSTAB needs MatVec and PSolve matrix operations.
type BCGSTAB struct {
	first  bool
	resume int

	rho, rhoPrev float64
	alpha        float64
	omega        float64

	rt   []float64
	p    []float64
	v    []float64
	t    []float64
	phat []float64
	s    []float64
	shat []float64
}

// Init implements the Method interface.
func (b *BCGSTAB) Init(dim int) {
	if dim <= 0 {
		panic("bcgstab: dimension not positive")
	}

	b.rt = reuse(b.rt, dim)
	b.p = reuse(b.p, dim)
	b.v = reuse(b.v, dim)
	b.t = reuse(b.t, dim)
	b.phat = reuse(b.phat, dim)
	b.s = reuse(b.s, dim)
	b.shat = reuse(b.shat, dim)
	b.first = true
	b.resume = 1
}

// Iterate implements the Method interface.
func (b *BCGSTAB) Iterate(ctx *Context) (Operation, error) {
	switch b.resume {
	case 1:
		if b.first {
			copy(b.rt, ctx.Residual)
		}
		b.rho = floats.Dot(b.rt, ctx.Residual)
		if math.Abs(b.rho) < rhoBreakdownTol {
			// Make sure that calling Iterate again without Init will panic.
			b.resume = 0
			return NoOperation, errors.New("bcgstab: rho breakdown")
		}
		if b.first {
			copy(b.p, ctx.Residual)
		} else {
			beta := (b.rho / b.rhoPrev) * (b.alpha / b.omega)
			floats.AddScaled(b.p, -b.omega, b.v)
			floats.Scale(beta, b.p)
			floats.Add(b.p, ctx.Residual)
		}
		ctx.Src = b.p
		ctx.Dst = b.phat
		b.resume = 2
		return PSolve, nil
		// Solve M*p^_i = p_i.
	case 2:
		ctx.Src = b.phat
		ctx.Dst = b.v
		b.resume = 3
		return MatVec, nil
		// Compute v_i <- A*p^_i.
	case 3:
		b.alpha = b.rho / floats.Dot(b.rt, b.v)
		// Early check for tolerance.
		floats.AddScaled(ctx.Residual, -b.alpha, b.v)
		copy(b.s, ctx.Residual)
		ctx.Src = nil
		ctx.Dst = nil
		ctx.ResidualNorm = floats.Norm(ctx.Residual, 2)
		ctx.Converged = false
		b.resume = 4
		return CheckConvergence, nil
	case 4:
		if ctx.Converged {
			floats.AddScaled(ctx.X, b.alpha, b.phat)
			// Make sure that calling Iterate again without Init will panic.
			b.resume = 0
			return EndIteration, nil
		}
		ctx.Src = b.s
		ctx.Dst = b.shat
		b.resume = 5
		return PSolve, nil
		// Solve M*s^_i = r_i.
	case 5:
		ctx.Src = b.shat
		ctx.Dst = b.t
		b.resume = 6
		return MatVec, nil
		// Compute t_i <- A*s^_i.
	case 6:
		b.omega = floats.Dot(b.t, b.s) / floats.Dot(b.t, b.t)
		floats.AddScaled(ctx.X, b.alpha, b.phat)
		floats.AddScaled(ctx.X, b.omega, b.shat)
		floats.AddScaled(ctx.Residual, -b.omega, b.t)
		ctx.Src = nil
		ctx.Dst = nil
		ctx.ResidualNorm = floats.Norm(ctx.Residual, 2)
		ctx.Converged = false
		b.resume = 7
		return CheckConvergence, nil
	case 7:
		if ctx.Converged {
			// Make sure that calling Iterate again without Init will panic.
			b.resume = 0
			return EndIteration, nil
		}
		if math.Abs(b.omega) < omegaBreakdownTol {
			return NoOperation, errors.New("bcgstab: omega breakdown")
		}
		b.rhoPrev = b.rho
		b.first = false
		b.resume = 1
		return EndIteration, nil

	default:
		panic("bcgstab: Init not called")
	}
}
