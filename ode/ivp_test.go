// Copyright ©2021 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ode_test

import (
	"fmt"
	"log"
	"math"
	"testing"

	"gonum.org/v1/exp/ode"
	"gonum.org/v1/gonum/mat"
)

type TestModel struct {
	*ode.Model
	solution *ode.Model
	err      func(h, i float64) float64
}

func TestSolve(t *testing.T) {
	// domain start
	quad := quadTestModel(t)
	solver := ode.NewDormandPrince5()
	solver.Set(quad)
	stepsize := 0.5 / 15.
	end := 0.5
	results, err := ode.Solve(quad, solver, stepsize, end)
	if err != nil {
		t.Fatal(err)
	}
	// run the solver and corsscheck values
	quadsol := quad.solution
	_, x0 := quadsol.IV()

	sol := quadsol.Func

	solresults := mat.NewVecDense(x0.Len(), nil)
	for i := range results {
		sol(solresults, results[i].T, results[i].X)
		for j := range solresults.RawVector().Data {
			sol := solresults.AtVec(j)
			res := results[i].X.AtVec(j)
			got := math.Abs(solresults.AtVec(j) - results[i].X.AtVec(j))
			expect := quad.err(stepsize, float64(i))
			if got > expect {
				t.Errorf("error %g is greater than permitted tolerance %g. solution:[%0.3g],  result:[%0.3g]", got, expect, sol, res)
			}
		}
	}

}

func Example_solve() {
	const (
		g = -10. // gravity field [m.s^-2]
	)
	// we declare our physical model. First argument is initial time, which is 0 seconds.
	// Next is the initial state vector, which corresponds to 100 meters above the ground
	// with 0 m/s velocity.
	ballModel, err := ode.NewModel(0, mat.NewVecDense(2, []float64{100., 0.}),
		func(yvec *mat.VecDense, _ float64, xvec mat.Vector) {
			// this anonymous function defines the physics.
			// The first variable xvec[0] corresponds to position
			// second variable xvec[1] is velocity.
			Dx := xvec.AtVec(1)
			// yvec represents change in xvec, or derivative with respect to domain
			// Change in position will be equal to velocity, which is the second variable:
			// thus yvec[0] = xvec[1], which is the same as saying "change in xvec[0]" is equal to xvec[1]
			yvec.SetVec(0, Dx)
			// change in velocity is acceleration. We suppose our ball is on earth accelerating at `g`
			yvec.SetVec(1, g)
		})
	if err != nil {
		log.Fatal(err)
	}
	// Here we choose our algorithm. Runge-Kutta 4th order is used
	var solver ode.Integrator = ode.NewDormandPrince5()
	// Solve function makes it easy to integrate a problem without having
	// to implement the `for` loop. This example integrates the IVP with a step size
	// of 0.1 over a domain of 10. arbitrary units, in this case, 10 seconds.
	results, err := ode.Solve(ballModel, solver, 0.1, 10.)
	fmt.Println(results)
}

// Quadratic model may be used for future algorithms
func quadTestModel(t *testing.T) *TestModel {
	t0 := 0.0
	Quadratic := new(TestModel)
	quad, err := ode.NewModel(t0, mat.NewVecDense(2, []float64{0, 0}),
		func(dst *mat.VecDense, t float64, x mat.Vector) {
			dst.SetVec(0, x.AtVec(1))
			dst.SetVec(1, 1.)
		})
	if err != nil {
		t.Fatal(err)
	}
	Quadratic.Model = quad
	quadsol, err := ode.NewModel(t0, mat.NewVecDense(2, []float64{0, 0}),
		func(dst *mat.VecDense, t float64, x mat.Vector) {
			dst.SetVec(0, t*t/2.)
			dst.SetVec(1, t)
		})
	if err != nil {
		t.Fatal(err)
	}
	Quadratic.solution = quadsol
	Quadratic.err = func(h, i float64) float64 { return math.Pow(h*i, 4) + 1e-10 }
	return Quadratic
}

// exponential unidimensional model may be used for future algorithms
//  y'(t) = -15*y(t)
//  y(t=0) = 1
//  solution: y(t) = exp(-15*t)
func exp1DTestModel(t *testing.T) *TestModel {
	tau := -2.
	t0 := 0.0
	Quadratic := new(TestModel)
	quad, err := ode.NewModel(t0, mat.NewVecDense(1, []float64{1.}),
		func(dst *mat.VecDense, t float64, x mat.Vector) {
			dst.SetVec(0, tau*x.AtVec(0))
		})
	if err != nil {
		t.Fatal(err)
	}
	Quadratic.Model = quad
	quadsol, err := ode.NewModel(t0, mat.NewVecDense(1, []float64{0}),
		func(dst *mat.VecDense, t float64, x mat.Vector) {
			dst.SetVec(0, math.Exp(tau*t))

		})
	if err != nil {
		t.Fatal(err)
	}
	Quadratic.solution = quadsol
	Quadratic.err = func(h, i float64) float64 { return math.Pow(h*i, 4) + 1e-10 }
	return Quadratic
}

// func TestQuadratic(t *testing.T) {
// 	Quadratic := quadTestModel(t)
// 	solver := ivp.NewDormandPrince5()

// 	solver.Set(Quadratic.Model)

// 	_, x0 := Quadratic.Model.IV()
// 	steps := 10
// 	dt := 0.1

// 	results := make([]float64, nx)

// 	solmodel := Quadratic.solution
// 	soleq := solmodel.Equations()
// 	solDims, _ := solmodel.Dims()
// 	solution := make([]float64, solDims)
// 	for i := 1.; i < float64(steps+1); i++ {
// 		dom := i * dt
// 		solver.Step(results, dt)
// 		soleq(solution, dom, results)
// 		for j := range results {
// 			got := math.Abs(solution[j] - results[j])
// 			expected := Quadratic.err(dt, i)
// 			if got > expected {
// 				t.Errorf("error %e greater than threshold %e", got, expected)
// 			}

// 		}
// 	}
// }
