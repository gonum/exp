// Copyright ©2017 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package linsolve provides iterative algorithms for solving linear systems.
package linsolve

// Method is an iterative method that produces a sequence of vectors converging
// to the vector x satisfying a system of linear equations
//  A x = b,
// where A is non-singular dim×dim matrix, and x and b are vectors of dimension
// dim.
//
// Method uses a reverse-communication interface between the iterative algorithm
// and the caller. Method acts as a client that commands the caller to perform
// needed operations via Operation returned from the Iterate method. This
// provides independence of Method on representation of the matrix A, and
// enables automation of common operations like checking for convergence and
// maintaining statistics.
type Method interface {
	// Init initializes the method for solving
	// a dim×dim linear system.
	Init(dim int)

	// Iterate retrieves data from Context,
	// updates it, and returns the next
	// operation. The caller must perform the
	// Operation using data in Context, and
	// depending on the state call Iterate
	// again.
	Iterate(*Context) (Operation, error)
}

// Context mediates the communication between the Method and the caller. It must
// not be modified or accessed apart from the commanded Operations.
type Context struct {
	// X is the current approximate solution.
	// On the first call to Method.Iterate, X
	// must contain the initial estimate.
	// Method must update X with the current
	// estimate when it commands
	// ComputeResidual and EndIteration.
	X []float64
	// Residual is the current residual b-A*x.
	// On the first call to Method.Iterate,
	// Residual must contain the initial
	// residual.
	// TODO(vladimir-ch): Consider whether the
	// behavior should also include: Method
	// must update Residual with the current
	// value of b-A*x when it commands
	// EndIteration.
	Residual []float64
	// ResidualNorm is (an estimate of) the
	// norm of the current residual. Method
	// must update it when it commands
	// CheckConvergence. It is not necessarily
	// equal to the norm of Residual, some
	// methods (e.g., GMRES) can estimate the
	// residual norm without forming the
	// residual itself.
	// TODO(vladimir-ch): Actually the exact
	// behavior here  is something that should
	// be discussed.
	ResidualNorm float64
	// Converged indicates to Method that the
	// ResidualNorm satisfies the stopping
	// criterion as a result of
	// CheckConvergence operation. If a Method
	// commands EndIteration with Converged
	// true, the caller must not call
	// Method.Iterate again without calling
	// Method.Init first.
	Converged bool

	// Src and Dst are the source and
	// destination vectors for various
	// Operations.
	Src, Dst []float64
}

// Operation specifies the type of operation.
type Operation uint64

// Operations commanded by Method.Iterate.
const (
	NoOperation Operation = 0

	// Multiply A*x where x is stored in
	// Context.Src. The result will be
	// stored in Context.Dst.
	MatVec Operation = 1 << (iota - 1)

	// Multiply A^T*x where x is stored in
	// Context.Src. The result will be
	// stored in Context.Dst.
	MatTransVec

	// Do the preconditioner solve
	//  M z = r,
	// where r is stored in Context.Src, and
	// store the solution z in Context.Dst.
	PSolve

	// Do the preconditioner solve
	//  M^T z = r,
	// where r is stored in Context.Src, and
	// store the solution z in Context.Dst.
	PSolveTrans

	// Compute b - A*x where x is stored in
	// Context.X and store the result into
	// Context.Residual.
	ComputeResidual

	// Check convergence using the current
	// approximation in Context.X and the
	// residual norm in Context.ResidualNorm.
	// Context.Converged must be set to
	// indicate whether convergence has been
	// determined, and then Method.Iterate
	// must be called again.
	CheckConvergence

	// EndIteration indicates that Method has
	// finished what it considers to be one
	// iteration. It can be used to update an
	// iteration counter. If Context.Converged
	// is true, Context.X contains the
	// approximate solution and the iterative
	// process must be terminated. Method.Init
	// must be called before calling
	// Method.Iterate again.
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
