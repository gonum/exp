// Copyright ©2021 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ivp

import (
	"errors"
	"fmt"

	"gonum.org/v1/gonum/mat"
)

// IVP defines a multivariable, non-autonomous initial value problem.
// It is worth mentioning system does not necessarily need
// to be non-autonomous. https://en.wikipedia.org/wiki/Autonomous_system_(mathematics)
//
//
// These problems have the form (capital letters are vectors in this example)
//  X'(t) = F(t, X, U)
//  X(0) = X_0
//
// Where X' is the vector of first derivatives of the state vector X.
// F would be xequations as returned by Equations(). t is
// a scalar representing the integrations domain, which is usually time
// for most physical problems. U is a vector which is a function of the current
// state. Put simply, the next input is a function of all current state variables
// and, possibly, current input as well.
//  U_next = F_u(t, X, U)
//
// Where F_u is ufunc as returned by Equations()
// An initial value problem is characterized by boundary conditions imposed
// on the state vector X at the beginning of the integration domain. These
// boundary conditions are returned by the IV() method for the state vector
// as x0 and for the input vector as u0.
//
// The term "state vector" and "state variables" are used interchangeably
// throughout the code and refer to X vector of independent variables.
type IVP interface {
	// Initial values vector for state variables x and inputs u. x0 defines
	// the first values the state vector takes when integration begins.
	IV() (x0, u0 mat.Vector)
	// Equations returns the coupled, non-linear algebraic differential
	// equations (xequations) for the state variables (x) and the functions for inputs (u).
	// The input functions (ufunc) are not differential equations but rather
	// calculated directly from a given x vector and current input vector.
	// Results are stored in y which are the length of x and u, respectively.
	// The scalar (float64) argument is the domain over which the
	// problem is integrated, which is usually time for most physical problems.
	//
	// If problem has no input functions then u supplied and ufunc returned
	// may be nil. x equations my not be nil.
	Equations() (xequations func(y []float64, t float64, x, u []float64), ufunc func(u_next []float64, t float64, x, u []float64))
	// Dimensions of x state vector and u inputs.
	Dims() (nx, nu int)
}

// Integrator abstracts algorithm specifics. For anyone looking to
// implement it, Set(ivp) should be called first to initialize the IVP with
// initial values. Step will calculate the next x values and store them in y
// u values should not be stored as they can easily be obtained if one has
// x values. Integrator should store 1 or more (depending on algorithm used)
// of previously calculated x values to be able to integrate.
type Integrator interface {
	// Set initializes an initial value problem. First argument
	// is the initial domain integration point, is usually zero.
	Set(float64, IVP) error
	// Dimensions of x state vector and u inputs. Set must be called first.
	Dims() (nx, nu int)
	// Step integrates IVP and stores result in y. step is a suggested step
	// for the algorithm to take. The algorithm may decide that it is not sufficiently
	// small or big enough (these are adaptive algorithms) and take a different step.
	// The resulting step is returned as the first parameter
	Step(y []float64, step float64) (float64, error)
}

type result = struct {
	DomainStartOffset float64
	X                 []float64
}

// Solve solves an already initialized Integrator returning state vector results.
// Returns error upon needing to allocate over 2GB of memory
func Solve(solver Integrator, stepsize, domainLength float64) (results []result, err error) {
	const maxAllocGB = 2
	integrated := 0.
	expectedLength := int(domainLength/stepsize) + 1
	nx, _ := solver.Dims()
	if nx == 0 {
		return nil, errors.New("state vector length can not be equal to zero. Has ivp been set?")
	}
	size := 8 * (nx + 1) * expectedLength
	if size > maxAllocGB*1e9 {
		return nil, fmt.Errorf("solution exceeds %dGB or not initialized (size is %dMB)", maxAllocGB, size/1e6)
	}
	results = make([]result, 0, expectedLength)

	for integrated < domainLength {
		res := make([]float64, nx)
		stepsize, err = solver.Step(res, stepsize)
		if err != nil {
			return results, err
		}
		if stepsize <= 0 {
			return results, errors.New("got zero or negative step size from Integrator")
		}
		integrated += stepsize
		results = append(results, result{DomainStartOffset: integrated, X: res})
	}
	return results, nil
}
