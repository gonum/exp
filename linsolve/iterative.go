// Copyright ©2017 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package linsolve

import (
	"errors"
	"time"

	"github.com/gonum/floats"
)

// System describes a linear system in terms of A*x
// and A^T*x operations, and the right-hand side.
type System struct {
	// MatVec computes A*x or A^T*x
	// and stores the result into
	// dst. It must not be nil.
	MatVec func(dst, x []float64, trans bool)

	// B is the right-hand side
	// vector. Its length determines
	// the dimension of the system.
	B []float64
}

// Settings holds various settings for solving a linear system.
type Settings struct {
	// Tolerance specifies error tolerance for
	// the final approximate solution produced
	// by the iterative method. Tolerance must
	// be smaller than one and greater than
	// the machine epsilon. If it is zero, a
	// default value will be used.
	//
	// If NormA is not zero, the stopping
	// criterion used will be
	//  |r_i| < Tolerance * (|A|*|x_i| + |b|),
	// If NormA is zero (not available), the
	// stopping criterion will be
	//  |r_i| < Tolerance * |b|.
	Tolerance float64

	// NormA is an estimate of a norm |A| of
	// A, for example, an approximation of the
	// largest entry. Zero value means that
	// the norm is not known.
	NormA float64

	// MaxIterations is the limit on the
	// number of iterations. If it is zero, a
	// default value of twice the dimension of
	// the system will be used.
	MaxIterations int

	// PSolve describes the preconditioner
	// solve that stores into dst the solution
	// of the system
	//  M z = rhs, or M^T z = rhs.
	// If it is nil, no preconditioning will
	// be used (M is the identitify).
	PSolve func(dst, rhs []float64, trans bool) error
}

func defaultSettings(s *Settings, dim int) {
	if s.Tolerance == 0 {
		s.Tolerance = 1e-6
	}
	if s.MaxIterations == 0 {
		s.MaxIterations = 2 * dim
	}
}

// Result holds the result of an iterative solve.
type Result struct {
	// X is the approximate solution.
	X []float64
	// ResidualNorm is the norm of
	// the final residual. It may not be equal
	// to the actual residual b-A*X.
	ResidualNorm float64
	// Stats holds statistics about the
	// iterative solve.
	Stats Stats
}

// Stats holds statistics about an iterative solve.
type Stats struct {
	// Iterations is the number of iteration
	// done by Method.
	Iterations int
	// MatVec is the number of MatVec and
	// MatTransVec operations commanded by
	// Method.
	MatVec int
	// PSolve is the number of PSolve and
	// PSolveTrans operations commanded by
	// Method.
	PSolve int
	// StartTime is an approximate time when
	// the solve was started.
	StartTime time.Time
	// Runtime is an approximate duration of
	// the solve.
	Runtime time.Duration
}

// Iterative solves the system of n linear equations
//  A*x = b,
// where the n×n matrix A and the right-hand side be are represented
// by sys. The dimension n of the system is determined by the length
// of sys.B.
//
// If x is not nil, its length must be equal to n, and it will be used
// as an initial guess. On return it will contain the approximate
// solution. If it is nil, the zero vector will be used as an initial
// guess, and the solution will be returned in Result.X.
//
// method is an iterative method used for finding an approximate
// solution of the linear system. It must not be nil.
//
// settings provide means for adjusting parameters of the iterative
// process. Zero values of the fields mean default values.
func Iterative(sys System, x []float64, method Method, settings Settings) (*Result, error) {
	stats := Stats{StartTime: time.Now()}

	b := sys.B
	dim := len(b)
	if dim == 0 {
		return nil, errors.New("linsolve: dimension must be positive")
	}

	if sys.MatVec == nil {
		panic("linsolve: nil matrix-vector multiplication")
	}
	if x != nil && len(x) != dim {
		panic("linsolve: mismatched length of initial guess")
	}

	defaultSettings(&settings, dim)
	if settings.Tolerance < eps || 1 <= settings.Tolerance {
		panic("linsolve: invalid tolerance")
	}

	ctx := &Context{
		Residual: make([]float64, dim),
	}
	if x != nil {
		ctx.X = x
		sys.MatVec(ctx.Residual, ctx.X, false)
		stats.MatVec++
		floats.AddScaledTo(ctx.Residual, b, -1, ctx.Residual) // r = b - Ax
	} else {
		ctx.X = make([]float64, dim)
		copy(ctx.Residual, b) // r = b
	}

	var err error
	ctx.ResidualNorm = floats.Norm(ctx.Residual, 2)
	if ctx.ResidualNorm >= settings.Tolerance {
		err = iterate(sys, ctx, settings, method, &stats)
	}

	stats.Runtime = time.Since(stats.StartTime)
	return &Result{
		X:            ctx.X,
		ResidualNorm: ctx.ResidualNorm,
		Stats:        stats,
	}, err
}

func iterate(sys System, ctx *Context, settings Settings, method Method, stats *Stats) error {
	bnorm := floats.Norm(sys.B, 2)
	if bnorm == 0 {
		bnorm = 1
	}

	dim := len(ctx.X)
	method.Init(dim)

	for {
		op, err := method.Iterate(ctx)
		if err != nil {
			return err
		}

		switch op {
		case NoOperation:

		case MatVec, MatTransVec:
			if op == MatVec {
				sys.MatVec(ctx.Dst, ctx.Src, false)
			} else {
				sys.MatVec(ctx.Dst, ctx.Src, true)
			}
			stats.MatVec++

		case PSolve, PSolveTrans:
			if settings.PSolve == nil {
				copy(ctx.Dst, ctx.Src)
				continue
			}
			if op == PSolve {
				err = settings.PSolve(ctx.Dst, ctx.Src, false)
			} else {
				err = settings.PSolve(ctx.Dst, ctx.Src, true)
			}
			if err != nil {
				return err
			}
			stats.PSolve++

		case CheckConvergence:
			// TODO(vladimir-ch): This is currently not
			// used because ctx.X is not guaranteed to be
			// valid when this operation is requested.
			// There is also the question in which norm
			// x should be measured (and similarly for b).
			//
			// if settings.NormA != 0 {
			// 	xnorm := floats.Norm(ctx.X, 2)
			// 	ctx.Converged = ctx.ResidualNorm/(settings.NormA*xnorm+bnorm) < settings.Tolerance
			// } else {
			// 	ctx.Converged = ctx.ResidualNorm/bnorm < settings.Tolerance
			// }
			ctx.Converged = ctx.ResidualNorm/bnorm < settings.Tolerance

		case ComputeResidual:
			sys.MatVec(ctx.Residual, ctx.X, false)
			stats.MatVec++
			floats.AddScaledTo(ctx.Residual, sys.B, -1, ctx.Residual)

		case EndIteration:
			stats.Iterations++
			if ctx.Converged {
				return nil
			}
			if stats.Iterations == settings.MaxIterations {
				return errors.New("linsolve: iteration limit reached")
			}

		default:
			panic("linsolve: invalid operation")
		}
	}
}
