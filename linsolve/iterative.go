// Copyright ©2017 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package linsolve

import (
	"errors"

	"gonum.org/v1/gonum/floats"
)

const defaultTolerance = 1e-8

var (
	// ErrIterationLimit is returned when a maximum number of iterations were done
	// without converging to a solution.
	ErrIterationLimit = errors.New("linsolve: iteration limit reached")

	// ErrNoProgress is returned when a Method stagnates, that is, when the
	// difference between two consecutive iterates is too small.
	ErrNoProgress = errors.New("linsolve: no progress")
)

// MulVecToer represents a square matrix A by means of a matrix-vector
// multiplication.
type MulVecToer interface {
	// MulVecTo computes A*x or A^T*x and stores the result into dst.
	MulVecTo(dst []float64, trans bool, x []float64)
}

// Settings holds settings for solving a linear system.
type Settings struct {
	// InitX holds the initial guess. If it is nil, the zero vector will be
	// used, otherwise its length must be equal to the dimension of the
	// system.
	InitX []float64

	// Dst, if not nil, will be used for storing the approximate solution,
	// otherwise a new slice will be allocated. In both cases the slice will
	// also be returned in Result. If Dst is not nil, its length must be
	// equal to the dimension of the system.
	Dst []float64

	// Tolerance specifies error tolerance for the final (approximate)
	// solution produced by the iterative method. If Tolerance is zero, a
	// default value of 1e-8 will be used, otherwise it must be positive and
	// less than 1.
	//
	// If NormA is not zero, the stopping criterion used will be
	//  |r_i| < Tolerance * (|A|*|x_i| + |b|),
	// where r_i is the residual at i-th iteration.
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
	//  M dst = rhs, or M^T dst = rhs,
	// where M is the preconditioning matrix. If PreconSolve is nil, no
	// preconditioning will be used (M is the identity).
	PreconSolve func(dst, rhs []float64, trans bool) error

	// Work context can be provided to reduce memory allocation when solving
	// multiple linear systems.
	Work *Context
}

// DefaultSettings returns default settings for solving a general linear system
// of dimension dim. It will not allocate fields like InitX, Dst, or Work.
func DefaultSettings(dim int) *Settings {
	var s Settings
	defaultSettings(&s, dim)
	return &s
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

func checkSettings(s *Settings, dim int) {
	if s.InitX != nil && len(s.InitX) != dim {
		panic("linsolve: mismatched length of initial guess")
	}
	if s.Dst != nil && len(s.Dst) != dim {
		panic("linsolve: mismatched destination length")
	}
	if s.Tolerance <= 0 || 1 <= s.Tolerance {
		panic("linsolve: invalid tolerance")
	}
	if s.MaxIterations <= 0 {
		panic("linsolve: negative iteration limit")
	}
}

// Result holds the result of an iterative solve.
type Result struct {
	// X is the approximate solution.
	X []float64

	// ResidualNorm is an approximation to the norm of the final residual.
	ResidualNorm float64

	// Stats holds statistics about the iterative solve.
	Stats Stats
}

// Stats holds statistics about an iterative solve.
type Stats struct {
	// Iterations is the number of iterations performed by Method.
	Iterations int

	// MulVec is the number of MulVec operations commanded by Method.
	MulVec int

	// PreconSolve is the number of PreconSolve operations commanded by Method.
	PreconSolve int
}

// Iterative finds an approximate solution of the system of n linear equations
//  A*x = b,
// where A is a nonsingular square matrix of order n and b is the right-hand
// side vector, using an iterative method m. If m is nil, default GMRES will be
// used.
//
// settings provide means for adjusting parameters of the iterative process. The
// nil value of settings is equivalent to using DefaultSettings. See the Settings
// documentation for more information. Iterative will not modify the fields of
// Settings.
//
// Note that the default choices of Method and Settings were chosen to provide
// accuracy and robustness, rather than speed. There are many algorithms for
// iterative linear solutions that have different tradeoffs, and can exploit
// special structure in the A matrix. Similarly, in many cases the number of
// iterations can be significantly reduced by using an appropriate
// preconditioner or increasing the error tolerance. Combined, these choices can
// significantly reduce computation time. Thus, while Iterative has supplied
// defaults, users are strongly encouraged to adjust these defaults for their
// problem.
func Iterative(a MulVecToer, b []float64, m Method, settings *Settings) (*Result, error) {
	n := len(b)
	if n == 0 {
		panic("linsolve: dimension is zero")
	}

	var s Settings
	if settings != nil {
		s = *settings
	}
	defaultSettings(&s, n)
	if s.Dst == nil {
		s.Dst = make([]float64, n)
	}
	checkSettings(&s, n)

	var stats Stats
	ctx := reuseContext(s.Work, n)
	if s.InitX != nil {
		copy(ctx.X, s.InitX)
		computeResidual(ctx.Residual, a, b, ctx.X, &stats)
	} else {
		// Initial x is the zero vector.
		// Residual b-A*x is then equal to b.
		copy(ctx.Residual, b)
	}

	if m == nil {
		m = &GMRES{}
	}

	var err error
	ctx.ResidualNorm = floats.Norm(ctx.Residual, 2)
	if ctx.ResidualNorm >= s.Tolerance {
		err = iterate(a, b, ctx, s, m, &stats)
	}

	return &Result{
		X:            s.Dst,
		ResidualNorm: ctx.ResidualNorm,
		Stats:        stats,
	}, err
}

func iterate(a MulVecToer, b []float64, ctx *Context, settings Settings, method Method, stats *Stats) error {
	bNorm := floats.Norm(b, 2)
	if bNorm == 0 {
		bNorm = 1
	}
	copy(settings.Dst, ctx.X)

	n := len(b)
	method.Init(n)

	for {
		op, err := method.Iterate(ctx)
		if err != nil {
			return err
		}

		switch op {
		case NoOperation:
		case MulVec:
			stats.MulVec++
			a.MulVecTo(ctx.Dst, op&Trans == Trans, ctx.Src)
		case PreconSolve:
			stats.PreconSolve++
			err = settings.PreconSolve(ctx.Dst, ctx.Src, op&Trans == Trans)
			if err != nil {
				return err
			}
		case CheckResidualNorm:
			ctx.Converged = ctx.ResidualNorm < settings.Tolerance*bNorm
		case ComputeResidual:
			computeResidual(ctx.Residual, a, b, ctx.X, stats)
		case MajorIteration:
			stats.Iterations++
			rNorm := floats.Norm(ctx.Residual, 2)
			var converged bool
			if settings.NormA != 0 {
				xNorm := floats.Norm(ctx.X, 2)
				converged = rNorm < settings.Tolerance*(settings.NormA*xNorm+bNorm)
			} else {
				converged = rNorm < settings.Tolerance*bNorm
			}
			if converged {
				copy(settings.Dst, ctx.X)
				ctx.ResidualNorm = rNorm
				return nil
			}
			if floats.Distance(settings.Dst, ctx.X, 2) <= eps*floats.Norm(ctx.X, 2) {
				return ErrNoProgress
			}
			copy(settings.Dst, ctx.X)
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

func computeResidual(dst []float64, a MulVecToer, b, x []float64, stats *Stats) {
	stats.MulVec++
	a.MulVecTo(dst, false, x)
	floats.AddScaledTo(dst, b, -1, dst)
}
