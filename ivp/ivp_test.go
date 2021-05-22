// Copyright ©2021 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ivp_test

import (
	"fmt"
	"log"
	"math"
	"testing"

	"gonum.org/v1/exp/ivp"

	"gonum.org/v1/gonum/mat"
)

func TestSolve(t *testing.T) {
	quad := quadTestModel(t)
	solver := new(ivp.RK4)
	solver.Set(0.0, quad)
	stepsize := 0.1
	results, err := ivp.Solve(solver, stepsize, 10)
	quadsol := quad.solution
	sol := quadsol.Equations()
	if err != nil {
		t.Fatal(err)
	}
	nxs, _ := quadsol.Dims()
	solresults := make([]float64, nxs)
	for i := range results {
		sol(solresults, results[i].DomainStartOffset, results[i].X)
		for j := range solresults {
			got := math.Abs(solresults[j] - results[i].X[j])
			expect := quad.err(stepsize, float64(i))
			if got > expect {
				t.Errorf("error %g is greater than permitted tolerance %g", got, expect)
			}
		}
	}
}

func Example_solve() {
	const (
		g = -10. // gravity field [m.s^-2]
	)
	// we declare our physical model in the following function
	ballModel, err := ivp.NewModel(mat.NewVecDense(2, []float64{100., 0.}),
		nil, func(yvec []float64, _ float64, xvec []float64) {
			// this anonymous function defines the physics.
			// The first variable xvec[0] corresponds to position
			// second variable xvec[1] is velocity.
			Dx := xvec[1]
			// yvec represents change in xvec, or derivative with respect to domain
			// Change in position will be equal to velocity, which is the second variable:
			// thus yvec[0] = xvec[1], which is the same as saying "change in xvec[0]" is equal to xvec[1]
			yvec[0] = Dx
			// change in velocity is acceleration. We suppose our ball is on earth accelerating at `g`
			yvec[1] = g
		}, nil)
	if err != nil {
		log.Fatal(err)
	}
	// Here we choose our algorithm. Runge-Kutta 4th order is used
	var solver ivp.Integrator = new(ivp.RK4)
	// Before integrating the IVP is passed to the integrator (a.k.a solver). Domain (time) starts at 0
	dom := 0.
	err = solver.Set(dom, ballModel)
	if err != nil {
		log.Fatal(err)
	}
	// Solve function makes it easy to integrate a problem without having
	// to implement the `for` loop. This example integrates the IVP with a step size
	// of 0.1 over a domain of 10. arbitrary units, in this case, 10 seconds.
	results, err := ivp.Solve(solver, 0.1, 10.)
	fmt.Println(results)
}
