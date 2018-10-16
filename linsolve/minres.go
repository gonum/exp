// Copyright ©2018 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package linsolve

import (
	"errors"
	"math"

	"gonum.org/v1/gonum/blas/blas64"
	"gonum.org/v1/gonum/floats"
)

const indefTol = 1e-50

// MINRES implements the Minimum Residual method for solving linear systems with
// a symmetric (generally indefinite) matrix. The preconditioner must be
// symmetric positive definite.
//
// Reference:
//  - C. C. Paige and M. A. Saunders (1975). Solution of sparse indefinite
//    systems of linear equations, SIAM J. Numerical Analysis 12, 617-629.
//  - A. Greenbaum (1997). Section 2.5 in Iterative Methods for Solving Linear
//    Systems. SIAM.
type MINRES struct {
	x         []float64
	r         []float64
	z         []float64
	u, u1     []float64
	v, v1     []float64
	w, w1, w2 []float64

	// Recurrence coefficients from the Lanczos algorithm
	alpha       float64
	beta, beta1 float64
	// Current element eta[k] of the vector eta
	eta float64
	// Residual norm
	rnorm float64
	// Orthogonal matrices Q_k, Q_{k-1}, Q_{k-2} represented by Givens
	// rotations
	q, q1, q2 givens
	// Indicates that the preconditioner might be indefinite.
	maybeIndef bool

	resume int
}

// Init initializes the data for a linear solve. See the Method interface for more details.
func (m *MINRES) Init(x, residual []float64) {
	dim := len(x)
	if dim == 0 {
		panic("minres: dimension not positive")
	}
	if len(residual) != dim {
		panic("minres: slice length mismatch")
	}

	// Reuse and zero out work vectors.
	m.x = reuse(m.x, dim)
	copy(m.x, x)
	m.r = reuse(m.r, dim)
	copy(m.r, residual)
	m.z = reuse(m.z, dim)
	m.u = reuse(m.u, dim)
	m.u1 = reuse(m.u1, dim)
	m.v = reuse(m.v, dim)
	m.v1 = reuse(m.v1, dim)
	m.w = reuse(m.w, dim)
	m.w1 = reuse(m.w1, dim)
	m.w2 = reuse(m.w2, dim)

	// Initialize Q_k and Q_{k-1} to the identity matrix. Q_{k-2} will be
	// overwritten in the first iteration.
	m.q.c = 1
	m.q.s = 0
	m.q1.c = 1
	m.q1.s = 0

	m.resume = 1
}

// Iterate performs an iteration of the linear solve. See the Method interface
// for more details.
//
// MINRES will command the following operations:
//  MulVec
//  PreconSolve
//  CheckResidualNorm
//  ComputeResidual
//  MajorIteration
//  NoOperation
func (m *MINRES) Iterate(ctx *Context) (Operation, error) {
	switch m.resume {
	case 1:
		// Solve M^{-1} * r.
		copy(ctx.Src, m.r)
		m.resume = 2
		return PreconSolve, nil
	case 2:
		copy(m.z, ctx.Dst)
		m.rnorm = floats.Norm(m.z, 2)
		rz := floats.Dot(m.r, m.z)
		if rz < indefTol && m.rnorm > indefTol {
			m.resume = 0
			return NoOperation, errors.New("minres: indefinite preconditioner")
		}
		m.beta = math.Sqrt(math.Abs(rz))
		m.eta = m.beta
		ctx.ResidualNorm = m.rnorm
		m.resume = 3
		return CheckResidualNorm, nil
	case 3:
		if ctx.Converged {
			m.resume = 0
			return MajorIteration, nil
		}
		copy(m.v, m.r)
		floats.Scale(1/m.beta, m.v)
		copy(m.u, m.z)
		floats.Scale(1/m.beta, m.u)
		fallthrough
	case 4:
		// Compute A * u.
		copy(ctx.Src, m.u)
		m.resume = 5
		return MulVec, nil
	case 5:
		copy(m.r, ctx.Dst)
		// Compute M^{-1} * r, that is, M^{-1} * A * u.
		copy(ctx.Src, m.r)
		m.resume = 6
		return PreconSolve, nil
	case 6:
		copy(m.z, ctx.Dst)

		// Lanczos algorithm
		m.alpha = floats.Dot(m.u, m.r)
		floats.AddScaled(m.r, -m.alpha, m.v)
		floats.AddScaled(m.r, -m.beta, m.v1)
		floats.AddScaled(m.z, -m.alpha, m.u)
		floats.AddScaled(m.z, -m.beta, m.u1)
		m.beta1 = m.beta
		rz := math.Abs(floats.Dot(m.r, m.z))
		m.beta = math.Sqrt(rz)

		// QR factorization of T. Apply Q_{k-1}*Q_{k-2} to the affected
		// elements T[k-2,k]=0, T[k-1,k]=beta_{k-1}, T[k,k]=alpha_k of
		// the k-th column of the symmetric tridiagonal matrix T.
		m.q2, m.q1 = m.q1, m.q
		tk2 := m.q2.s * m.beta1
		tk1p := m.q2.c * m.beta1
		tk1 := m.q1.c*tk1p + m.q1.s*m.alpha
		tk0 := -m.q1.s*tk1p + m.q1.c*m.alpha

		// Generate a Givens rotation Q_k defined to zero out T[k+1,k]=beta_k.
		m.q.c, m.q.s, _, _ = blas64.Implementation().Drotg(tk0, m.beta)
		// Apply Q_k to update T[k,k].
		tk0 = m.q.c*tk0 + m.q.s*m.beta

		// Update the solution vector.
		copy(m.w2, m.w1)
		copy(m.w1, m.w)
		floats.AddScaledTo(m.w, m.u, -tk1, m.w1)
		floats.AddScaled(m.w, -tk2, m.w2)
		floats.Scale(1/tk0, m.w)
		floats.AddScaled(ctx.X, m.q.c*m.eta, m.w)

		m.maybeIndef = false
		if rz < indefTol {
			m.maybeIndef = true
			// We have either convergence or an indefinite
			// preconditioner. Compute true residual norm to check
			// for convergence and if convergence is not detected,
			// return an error below.
			m.resume = 7
			return ComputeResidual, nil
		}
		fallthrough
	case 7:
		if m.maybeIndef {
			m.rnorm = floats.Norm(ctx.Dst, 2)
		} else {
			m.rnorm *= math.Abs(m.q.s)
		}
		ctx.ResidualNorm = m.rnorm
		m.resume = 8
		return CheckResidualNorm, nil
	case 8:
		if !ctx.Converged && m.maybeIndef {
			m.resume = 0
			return NoOperation, errors.New("minres: indefinite preconditioner")
		}
		if ctx.Converged {
			m.resume = 0
			return MajorIteration, nil
		}
		// Apply Q_k to the vector eta to update its current element.
		m.eta *= -m.q.s
		copy(m.v1, m.v)
		copy(m.u1, m.u)
		copy(m.v, m.r)
		copy(m.u, m.z)
		floats.Scale(1/m.beta, m.v)
		floats.Scale(1/m.beta, m.u)
		m.resume = 4
		return MajorIteration, nil

	default:
		panic("linsolve: MINRES.Init not called")
	}
}
