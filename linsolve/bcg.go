// Copyright ©2017 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package linsolve

import (
	"errors"
	"math"

	"github.com/gonum/floats"
)

// BCG implements the biconjugate gradient iterative method with
// preconditioning for solving the system of linear equations
//  A*x = b,
// where A is a non-symmetric matrix. For symmetric positive definite
// systems use CG.
//
// BCG needs MatVec, MatTransVec, PSolve, and PSolveTrans matrix
// operations.
type BCG struct {
	first  bool
	resume int

	rho, rhoPrev float64
	alpha        float64

	rt    []float64
	z, zt []float64
	p, pt []float64
}

// Init implements the Method interface.
func (b *BCG) Init(dim int) {
	if dim <= 0 {
		panic("bcg: dimension not positive")
	}

	b.rt = reuse(b.rt, dim)
	b.z = reuse(b.z, dim)
	b.zt = reuse(b.zt, dim)
	b.p = reuse(b.p, dim)
	b.pt = reuse(b.pt, dim)

	b.first = true
	b.resume = 1
}

// Iterate implements the Method interface.
func (b *BCG) Iterate(ctx *Context) (Operation, error) {
	switch b.resume {
	case 1:
		if b.first {
			copy(b.rt, ctx.Residual)
		}
		ctx.Src = ctx.Residual
		ctx.Dst = b.z
		b.resume = 2
		return PSolve, nil
		// Solve M*z = r_{i-1}
	case 2:
		ctx.Src = b.rt
		ctx.Dst = b.zt
		b.resume = 3
		return PSolveTrans, nil
		// Solve M^T*zt = rt_{i-1}
	case 3:
		b.rho = floats.Dot(b.z, b.rt)
		if math.Abs(b.rho) < rhoBreakdownTol {
			// Make sure that calling Iterate again without Init will panic.
			b.resume = 0
			return NoOperation, errors.New("bcg: rho breakdown")
		}
		if b.first {
			copy(b.p, b.z)
			copy(b.pt, b.zt)
		} else {
			beta := b.rho / b.rhoPrev
			floats.AddScaledTo(b.p, b.z, beta, b.p)
			floats.AddScaledTo(b.pt, b.zt, beta, b.pt)
		}
		// z and zt will now be overwritten and reused.
		ctx.Src = b.p
		ctx.Dst = b.z
		b.resume = 4
		return MatVec, nil
		// z <- A*p
	case 4:
		ctx.Src = b.pt
		ctx.Dst = b.zt
		b.resume = 5
		return MatTransVec, nil
		// zt <- A^T*pt
	case 5:
		b.alpha = b.rho / floats.Dot(b.pt, b.z)
		floats.AddScaled(ctx.X, b.alpha, b.p)
		floats.AddScaled(ctx.Residual, -b.alpha, b.z)
		ctx.Src = nil
		ctx.Dst = nil
		ctx.ResidualNorm = floats.Norm(ctx.Residual, 2)
		ctx.Converged = false
		b.resume = 6
		return CheckConvergence, nil
	case 6:
		if ctx.Converged {
			// Make sure calling Iterate again without Init will panic.
			b.resume = 0
			return EndIteration, nil
		}
		// Prepare for the next iteration.
		floats.AddScaled(b.rt, -b.alpha, b.zt)
		b.rhoPrev = b.rho
		b.first = false
		b.resume = 1
		return EndIteration, nil

	default:
		panic("bcg: Init not called")
	}
}
