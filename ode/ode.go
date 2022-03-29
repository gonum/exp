// Copyright ©2021 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ode

import (
	"errors"
	"math"

	"gonum.org/v1/gonum/mat"
)

// Integrator can integrate an initial-value problem (IVP) for a first-order
// system of ordinary differential equations (ODEs).
type Integrator interface {
	// Init initializes the integrator and sets the initial condition.
	Init(IVP)

	// Step advances the current state by taking at most the given step. It returns a proposed step size
	// for the next step and an error indicating whether the step was successful.
	Step(step float64) (stepNext float64, err error)

	// State stores the current state of the integrator in-place in dst.
	State(dst *State)
}

// State represents the state of a system Y at a domain point T.
type State struct {
	T float64
	Y *mat.VecDense
}

// State represents the state of a system Y
// and the derivative of the state DY at a domain point T.
type State2 struct {
	T  float64
	Y  *mat.VecDense
	DY *mat.VecDense
}

// SolveIVP solves an already initialized Integrator returning state vector results.
func SolveIVP(p IVP, solver Integrator, stepsize, tend float64) (results []State, err error) {
	t0, x0 := p.T0, mat.VecDenseCopyOf(p.Y0)
	nx := x0.Len()
	if nx == 0 {
		return nil, errors.New("state vector length can not be equal to zero. Has ivp been set?")
	}
	// calculate expected size of results, these may differ
	domainLength := tend - t0
	expectedLength := int(domainLength/stepsize) + 1
	results = make([]State, 0, expectedLength)

	// t stores current domain
	t := t0
	for t < tend {
		res := State{Y: mat.NewVecDense(nx, nil)}

		if t-tend > 1e-10 {
			stepsize = math.Min(stepsize, (t-tend)*(1+1e-3))
		}
		stepsize, err = solver.Step(stepsize)
		if err != nil {
			return results, err
		}
		if stepsize <= 0 {
			return results, errors.New("got zero or negative step size from Integrator")
		}
		solver.State(&res)
		t += stepsize
		results = append(results, res)
	}
	return results, nil
}
