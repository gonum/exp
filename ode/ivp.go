// Copyright ©2021 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ode

import (
	"errors"
	"math"

	"gonum.org/v1/gonum/mat"
)

// IVP defines a multivariable, non-autonomous initial value problem.
// It is worth mentioning the system need not be non-autonomous. https://en.wikipedia.org/wiki/Autonomous_system_(mathematics)
//
// These problems have the form (capital letters are vectors in this example)
//  X'(t) = F(t, X)
//  X(0) = X_0
//
// Where:
// t is a scalar representing the integration domain, which is time for most physical problems.
// X' is the vector of first derivatives of the state vector X.
// F would be xequations as returned by Equations().
// An initial value problem is characterized by the initial conditions imposed
// on the state vector X at the beginning of the integration domain. These
// initial conditions are returned by the IV() method for the state vector
// as x0 and for the input vector as u0.
//
// The term "state vector" and "state variables" are used interchangeably
// throughout the code and refer to X vector of independent variables.
type IVP interface {
	// Initial values vector for state vector x. x0 defines
	// the first values the state vector takes when integration begins.
	IV() (t0 float64, x0 mat.Vector)
	// Func is the coupled, non-linear algebraic differential
	// equations for the state vector (x)
	// Results are stored in dst which are the length of x.
	// The scalar (float64) argument is the domain over which the
	// problem is integrated, which is usually time for most physical problems.
	Func(dst *mat.VecDense, t float64, x mat.Vector)
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
	Set(IVP)
	// Step integrates IVP and stores result in dst. step is a suggested step
	// for the algorithm to take. The algorithm may decide that it is not sufficiently
	// small or big enough (these are adaptive algorithms) and take a different step.
	// The resulting step is returned as the first parameter
	Step(dst *mat.VecDense, step float64) (float64, error)
}

type result = struct {
	T float64
	X *mat.VecDense
}

// Solve solves an already initialized Integrator returning state vector results.
func Solve(p IVP, solver Integrator, stepsize, tend float64) (results []result, err error) {
	t0, x0 := p.IV()
	nx := x0.Len()
	if nx == 0 {
		return nil, errors.New("state vector length can not be equal to zero. Has ivp been set?")
	}
	// calculate expected size of results, these may differ
	domainLength := tend - t0
	expectedLength := int(domainLength/stepsize) + 1
	results = make([]result, 0, expectedLength)

	// t stores current domain
	t := t0
	for t < tend {
		res := mat.NewVecDense(nx, nil)

		if t-tend > 1e-10 {
			stepsize = math.Min(stepsize, (t-tend)*(1+1e-3))
		}
		stepsize, err = solver.Step(res, stepsize)
		if err != nil {
			return results, err
		}
		if stepsize <= 0 {
			return results, errors.New("got zero or negative step size from Integrator")
		}
		t += stepsize
		results = append(results, result{T: t, X: res})
	}
	return results, nil
}
