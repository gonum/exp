// Copyright ©2017 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package linsolve

import (
	"errors"
	"time"

	"gonum.org/v1/gonum/floats"
)

const defaultTolerance = 1e-8

// ErrIterationLimit is returned when a maximum number of iterations were done
// without converging to a solution.
var ErrIterationLimit = errors.New("linsolve: iteration limit reached")

// System describes a linear system in terms of A*x and A^T*x operations, and
// the right-hand side.
type System struct {
	// MulVec computes A*x or A^T*x and stores the result into dst.
	// MulVec must not be nil.
	MulVec func(dst, x []float64, trans bool)

	// B is the right-hand side vector. Its length determines the dimension
	// of the system.
	B []float64
}

// Settings holds various settings for solving a linear system.
type Settings struct {
	// InitX holds the initial guess. If it is nil, the zero vector will be
	// used, otherwise its length must be equal to the dimension of the
	// system.
	InitX []float64

	// Tolerance specifies error tolerance for the final approximate
	// solution produced by the iterative method. Tolerance must be positive
	// and smaller than 1. If it is zero, a default value of 1e-8 will be
	// used.
	//
	// If NormA is not zero, the stopping criterion used will be
	//  |r_i| < Tolerance * (|A|*|x_i| + |b|),
	// If NormA is zero (not available), the stopping criterion will be
	//  |r_i| < Tolerance * |b|.
	Tolerance float64

	// NormA is an estimate of a norm of the matrix A. For example,
	// an approximation of the absolute value of the largest entry of A can be
	// used. Zero value means that the norm is not known.
	NormA float64

	// MaxIterations is the limit on the number of iterations. If it is
	// zero, a default value of twice the dimension of the system will be
	// used.
	MaxIterations int

	// PreconSolve describes a preconditioner solve that stores into dst the
	// solution of the system
	//  M dst = rhs, or M^T dst = rhs.
	// If it is nil, no preconditioning will be used (M is the identity).
	PreconSolve func(dst, rhs []float64, trans bool) error
}

// defaultSettings fills zero fields of s with default values.
func defaultSettings(s *Settings, dim int) {
	if s.Tolerance == 0 {
		s.Tolerance = defaultTolerance
	}
	if s.MaxIterations == 0 {
		s.MaxIterations = 2 * dim
	}
	if s.PreconSolve == nil {
		s.PreconSolve = NoPreconditioner
	}
}

// Result holds the result of an iterative solve.
type Result struct {
	// X is the approximate solution.
	X []float64

	// ResidualNorm is the norm of the final residual. It may not be equal
	// to the actual residual b-A*x.
	ResidualNorm float64

	// Stats holds statistics about the iterative solve.
	Stats Stats
}

// Stats holds statistics about an iterative solve.
type Stats struct {
	// Iterations is the number of iterations done by Method.
	Iterations int

	// MulVec is the number of MulVec operations commanded by Method.
	MulVec int

	// PreconSolve is the number of PreconSolve operations commanded by Method.
	PreconSolve int

	// StartTime is an approximate time when the solve was started.
	StartTime time.Time

	// Runtime is an approximate duration of the solve.
	Runtime time.Duration
}

// Iterative solves the system of n linear equations
//  A*x = b,
// where the n×n matrix A and the right-hand side b are represented by sys. The
// dimension n of the system is determined by the length of sys.B.
//
// If x is not nil, its length must be equal to n, and it will be used as an
// initial guess. On return it will contain the approximate solution. If it is
// nil, the zero vector will be used as an initial guess, and the solution will
// be returned in Result.X.
//
// method is an iterative method used for finding an approximate solution of the
// linear system. It must not be nil.
//
// settings provide means for adjusting parameters of the iterative process.
// Zero values of the fields mean default values.
func Iterative(dst []float64, sys System, method Method, settings Settings) (*Result, error) {
	stats := Stats{StartTime: time.Now()}

	if sys.MulVec == nil {
		panic("linsolve: nil matrix-vector multiplication")
	}
	dim := len(sys.B)
	if dim == 0 {
		return nil, errors.New("linsolve: dimension not positive")
	}

	if dst == nil {
		dst = make([]float64, dim)
	}
	if len(dst) != dim {
		panic("linsolve: mismatched length of dst")
	}

	ctx := &Context{
		X:        dst,
		Residual: make([]float64, dim),
		Src:      make([]float64, dim),
		Dst:      make([]float64, dim),
	}
	if settings.InitX != nil {
		if len(settings.InitX) != dim {
			panic("linsolve: mismatched length of initial guess")
		}
		copy(ctx.X, settings.InitX)
		computeResidual(ctx.Residual, sys, ctx.X, &stats)
	} else {
		// Initial x is the zero vector.
		for i := range ctx.X {
			ctx.X[i] = 0
		}
		// Residual b-A*x is then equal to b.
		copy(ctx.Residual, sys.B)
	}

	defaultSettings(&settings, dim)
	if settings.Tolerance <= 0 || 1 <= settings.Tolerance {
		panic("linsolve: invalid tolerance")
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

		case MulVec:
			stats.MulVec++
			sys.MulVec(ctx.Dst, ctx.Src, op&Trans == Trans)

		case PreconSolve:
			stats.PreconSolve++
			err = settings.PreconSolve(ctx.Dst, ctx.Src, op&Trans == Trans)
			if err != nil {
				return err
			}
		case CheckResidualNorm:
			ctx.Converged = ctx.ResidualNorm < settings.Tolerance*bnorm

		case ComputeResidual:
			computeResidual(ctx.Residual, sys, ctx.X, stats)

		case MajorIteration:
			stats.Iterations++
			rnorm := floats.Norm(ctx.Residual, 2)
			var converged bool
			if settings.NormA != 0 {
				xnorm := floats.Norm(ctx.X, 2)
				converged = rnorm < settings.Tolerance*(settings.NormA*xnorm+bnorm)
			} else {
				converged = rnorm < settings.Tolerance*bnorm
			}
			if converged {
				ctx.ResidualNorm = rnorm
				return nil
			}
			if stats.Iterations == settings.MaxIterations {
				return ErrIterationLimit
			}

		default:
			panic("linsolve: invalid operation")
		}
	}
}

// NoPreconditioner implements the identity preconditioner.
func NoPreconditioner(dst, rhs []float64, trans bool) error {
	if len(dst) != len(rhs) {
		panic("linsolve: mismatched slice length")
	}
	copy(dst, rhs)
	return nil
}

func computeResidual(dst []float64, sys System, x []float64, stats *Stats) {
	stats.MulVec++
	sys.MulVec(dst, x, false)
	floats.AddScaledTo(dst, sys.B, -1, dst)
}
