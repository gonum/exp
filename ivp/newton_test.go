package ivp_test

import (
	"math"
	"testing"

	"gonum.org/v1/exp/ivp"
	"gonum.org/v1/gonum/mat"
)

// This is a stiff equation https://en.wikipedia.org/wiki/Stiff_equation
//  y'(t) = -15*y(t)
//  solution: y(t) = exp(-15*t)
func exponential1DTestModel(t *testing.T) *TestModel {
	tau := -15.
	Stiff := new(TestModel)
	stiff, err := ivp.NewModel(mat.NewVecDense(1, []float64{1.}),
		nil, func(y []float64, dom float64, x []float64) {
			y[0] = tau * x[0]
		}, nil)
	if err != nil {
		t.Fatal(err)
	}
	Stiff.Model = stiff
	stiffsol, err := ivp.NewModel(mat.NewVecDense(1, []float64{0}),
		nil, func(y []float64, dom float64, x []float64) {
			y[0] = math.Exp(tau * dom)
		}, nil)
	if err != nil {
		t.Fatal(err)
	}
	Stiff.solution = stiffsol
	Stiff.err = func(h, i float64) float64 { return 2 * h * i }
	return Stiff
}

func TestNewton1DStiff(t *testing.T) {
	stiff1D := exponential1DTestModel(t)
	solver := ivp.NewNewtonRaphson(
		ivp.ConfigScalarTolerance(0, 1e-5),
	)
	solver.Set(0.0, stiff1D)
	stepsize := 1. / 30.
	results, err := ivp.Solve(solver, stepsize, 0.5)
	if err != nil {
		t.Fatal(err)
	}

	quadsol := stiff1D.solution
	sol := quadsol.Equations()
	nxs, _ := quadsol.Dims()
	solresults := make([]float64, nxs)
	for i := range results {
		sol(solresults, results[i].DomainStartOffset, results[i].X)
		for j := range solresults {
			got := math.Abs(solresults[j] - results[i].X[j])
			expect := stiff1D.err(stepsize, float64(i))
			if got > expect {
				t.Errorf("error %g is greater than permitted tolerance %g", got, expect)
			}
		}
	}
}
