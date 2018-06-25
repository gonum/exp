// Copyright ©2017 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package linsolve

// Method is an iterative method that produces a sequence of vectors that
// converge to the solution of the system of linear equations
//  A x = b,
// where A is non-singular dim×dim matrix, and x and b are vectors of dimension
// dim.
//
// Method uses a reverse-communication interface between the iterative algorithm
// and the caller. Method acts as a client that commands the caller to perform
// needed operations via an Operation returned from the Iterate method. This
// provides independence of Method on representation of the matrix A, and
// enables automation of common operations like checking for convergence and
// maintaining statistics.
type Method interface {
	// Init initializes the method for solving a dim×dim
	// linear system.
	Init(dim int)

	// Iterate performs a step in converging to the
	// solution of a linear system.
	//
	// Iterate retrieves data from Context, updates it,
	// and returns the next operation. The caller must
	// perform the Operation using data in Context, and
	// depending on the state call Iterate again.
	Iterate(*Context) (Operation, error)
}

// Context mediates the communication between the Method and
// the caller. The caller must not modify Context apart from
// the commanded Operations.
type Context struct {
	// X is the current approximate solution. On the
	// first call to Method.Iterate, X must contain the
	// initial estimate. Method will update X with the
	// current estimate when it commands ComputeResidual
	// and EndIteration.
	X []float64

	// Residual is the current residual b-A*x. On the
	// first call to Method.Iterate Residual must
	// contain the initial residual. Method will update
	// it to the current value when it commands
	// MajorIteration.
	Residual []float64

	// ResidualNorm is (an estimate of) a norm of
	// the residual. Method will set it to the current
	// value when it commands CheckResidualNorm.
	ResidualNorm float64

	// Converged indicates to Method whether ResidualNorm
	// satisfies a stopping criterion as a result of
	// CheckResidualNorm operation.
	Converged bool

	// Src and Dst are the source and destination
	// vectors for various Operations. See the Operation
	// documentation for more information.
	Src, Dst []float64
}

// NewContext allocates and returns a Context for solving n-dimensional problems.
func NewContext(n int) *Context {
	if n <= 0 {
		panic("linsolve: context size is not positive")
	}
	return &Context{
		X:        make([]float64, n),
		Residual: make([]float64, n),
		Src:      make([]float64, n),
		Dst:      make([]float64, n),
	}
}

func (ctx *Context) Reset(n int) {
	ctx.X = reuse(ctx.X, n)
	ctx.Residual = reuse(ctx.Residual, n)
	ctx.Src = reuse(ctx.Src, n)
	ctx.Dst = reuse(ctx.Dst, n)
}

// Operation specifies the type of operation.
type Operation uint

// Operations commanded by Method.Iterate.
const (
	NoOperation Operation = 0

	// Compute A*x where x is stored in Context.Src. The
	// result must be placed in Context.Dst.
	MulVec Operation = 1 << (iota - 1)

	// Perform a preconditioner solve
	//  M z = r
	// where r is stored in Context.Src. The solution z
	// must be placed in Context.Dst.
	PreconSolve

	// Trans indicates that MulVec or PrecondSolve
	// operation must be performed wih the transpose,
	// that is, compute A^T*x or solve M^T z = r. Method
	// will command Trans only in bitwise OR combination
	// with MulVec and PreconSolve.
	Trans

	// Compute b-A*x where x is stored in Context.X. The
	// result must be placed into Context.Residual.
	ComputeResidual

	// Check convergence using (an estimate of) a
	// residual norm in Context.ResidualNorm. Context.X
	// does not need to be valid. The caller must set
	// Context.Converged to indicate whether convergence
	// has been determined, and then it must call
	// Method.Iterate again.
	CheckResidualNorm

	// Check convergence using the residual in Context.Residual.
	// Context.X must be valid. The caller must set
	// Context.Converged to indicate whether convergence
	// has been determined, and then it must call
	// Method.Iterate again.
	CheckResidual

	// MajorIteration indicates that Method has finished
	// what it considers to be one iteration. Method
	// will make sure that Context.X and
	// Context.Residual are updated. If Context.Converged is true,
	// the caller should terminate the iterative process,
	// otherwise it should call Method.Iterate again.
	MajorIteration
)

func reuse(v []float64, n int) []float64 {
	if cap(v) < n {
		return make([]float64, n)
	}
	v = v[:n]
	for i := range v {
		v[i] = 0
	}
	return v
}

const (
	// Machine epsilon.
	eps = 1.0 / (1 << 53)

	// Tolerances for BCG and BCGSTAB methods.
	rhoBreakdownTol   = eps * eps
	omegaBreakdownTol = eps * eps
)
