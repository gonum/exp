// Copyright ©2017 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package linsolve provides iterative algorithms for solving linear systems.
package linsolve

// TODO(vladimir-ch): Improve documentation. Write an introduction about
// iterative methods and that they can be more efficient than direct methods
// when we are solving large (sparse) systems, when the solution does not have
// to be known to machine precision. Write that the matrix is accessed only via
// matrix-vector products. Write that the documentation is written from the
// perspective of users who want to call Iterative and designers who want to
// implement Method (or direct users of Method?).

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
	// first call to Method.Iterate, Residual must
	// contain the initial residual. Method will set it
	// to a valid value when it commands CheckResidual.
	// TODO(vladimir-ch): Consider whether the behavior
	// should also include: Method will update Residual
	// with the current value of b-A*x when it commands
	// EndIteration. Probably not because of GMRES.
	Residual []float64

	// ResidualNorm is (an estimate of) the norm of the
	// current residual. Method will set it to a valid
	// value when it commands CheckResidualNorm.
	ResidualNorm float64

	// Converged indicates to Method whether the
	// Residual or the ResidualNorm satisfies the
	// stopping criterion as a result of CheckResidual
	// or CheckResidualNorm operation. If Converged is
	// set to true, Method will then command
	// EndIteration with Converged set to true. After
	// that the caller must not call Method.Iterate
	// again without calling Method.Init first.
	Converged bool

	// Src and Dst are the source and destination
	// vectors for various Operations. See the Operation
	// documentation for more information.
	Src, Dst []float64

	// Trans indicates whether MulVec and PreconSolve
	// operations must be performed with a matrix
	// transpose.
	Trans bool
}

// Operation specifies the type of operation.
type Operation uint

// Operations commanded by Method.Iterate.
const (
	NoOperation Operation = 0

	// Compute
	//  A*x    if Context.Trans==false, or
	//  A^T*x  if Context.Trans==true,
	// where x is stored in Context.Src.
	// The result must be placed in Context.Dst.
	MulVec Operation = 1 << (iota - 1)

	// Perform a preconditioner solve
	//  M z = r    if Context.Trans==false, or
	//  M^T z = r, if Context.Trans==true,
	// where r is stored in Context.Src. The solution z
	// must be placed in Context.Dst.
	PreconSolve

	// Compute b-A*x where x is stored in Context.X. The
	// result must be placed into Context.Residual.
	ComputeResidual

	// Check convergence using the current approximation
	// in Context.X and the current residual
	// Context.Residual. The caller must set
	// Context.Converged to indicate whether convergence
	// has been determined, and then it must call
	// Method.Iterate again.
	CheckResidual

	// Check convergence using the current approximation
	// in Context.X and the residual norm in
	// Context.ResidualNorm. The caller must set
	// Context.Converged to indicate whether convergence
	// has been determined, and then it must call
	// Method.Iterate again.
	CheckResidualNorm

	// EndIteration indicates that Method has finished
	// what it considers to be one iteration. If
	// Context.Converged is true, Context.X contains the
	// approximate solution and the caller must
	// terminate the iterative process. If the caller
	// performs a new iterative run, it must call
	// Method.Init before calling Method.Iterate.
	EndIteration
)

func reuse(v []float64, n int) []float64 {
	if cap(v) < n {
		return make([]float64, n)
	}
	return v[:n]
}

const (
	// Machine epsilon.
	eps = 1.0 / (1 << 53)

	// Tolerances for BCG and BCGSTAB methods.
	rhoBreakdownTol   = eps * eps
	omegaBreakdownTol = eps * eps
)
