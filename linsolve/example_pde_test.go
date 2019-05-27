// Copyright ©2019 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package linsolve_test

import (
	"fmt"
	"log"

	"golang.org/x/exp/rand"

	"gonum.org/v1/exp/linsolve"
	"gonum.org/v1/gonum/mat"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

// AllenCahnFD implements a semi-implicit finite difference scheme for the
// solution of the one-dimensional Allen-Cahn equation
//  ∂_t u = Δu - 1/ξ²·f'(u)  in (0,L)×(0,T)
//   u(0) = u0              on (0,L)
// where f is a double-well potential function
//  f(s) = 1/4·(s²-1)²
//
// The equation arises in materials science in the description of phase
// transitions, e.g. solidification in crystal growth, but also in other areas
// like image processing due to its connection to mean-curvature motion.
// Starting the evolution from an initial distribution u0, the solution u
// develops a thin steep layer, an interface between regions of the domain
// (0,L) where u is constant and close to one of the minima of f.
//
// AllenCahnFD approximates derivatives by finite differences and the solution
// is advanced in time by a semi-implicit Euler scheme where the nonlinear term
// is taken from the previous time step. Therefore, at each time step a linear
// system must be solved.
type AllenCahnFD struct {
	// Xi is the ξ parameter that determines the interface width.
	Xi float64

	// InitCond is the initial condition u0.
	InitCond func(x float64) float64

	h   float64 // Spatial step size
	tau float64 // Time step size

	a *mat.SymBandDense
	b []float64
	u []float64

	ls         linsolve.Method
	lssettings linsolve.Settings
}

// FPrime returns the value of the derivative of the double-well potential f at s.
//  f'(s) = s·(s²-1)
func FPrime(s float64) float64 {
	return s * (s*s - 1)
}

// Setup initializes the receiver for solving the Allen-Cahn equation on a
// uniform grid with n nodes on the spatial interval (0,L) and with the time
// step size tau.
func (ac *AllenCahnFD) Setup(n int, L float64, tau float64) {
	ac.h = L / float64(n-1)
	ac.tau = tau

	// The finite difference scheme can be derived by replacing the derivatives
	// by finite differences:
	//  (u^k_i - u^{k-1}_i) / tau = (u^k_{i-1} - 2*u^k_i + u^k_{i+1}) / h² - 1/ξ² * f'(u^{k-1}_i)
	// Collecting the terms from the same time level on each side of the
	// equation gives:
	//  u^k_i - tau / h² * (u^k_{i-1} - 2*u^k_i + u^k_{i+1}) = u^{k-1}_i - tau/ξ² * f'(u^{k-1}_i)
	// The zero Neumann boundary condition is imposed by reflecting the solution
	// across the boundary:
	//  u_{-1} = u_1,   u_{n+1} = u_{n-1}
	// giving for example at i=0:
	//  u^k_0 / 2 - tau / h² * (u^k_1 - u^k_0) = (u^{k-1}_0 - tau/ξ² * f'(u^{k-1}_0)) / 2

	// Assemble the symmetric tridiagonal system matrix A based on the above
	// discretization scheme.
	const lda = 2
	a := make([]float64, n*lda)
	coef := tau / ac.h / ac.h
	// Boundary condition at the left node
	a[0] = 0.5 + coef
	a[1] = -coef
	// Interior nodes
	for i := 1; i < n-1; i++ {
		a[i*lda] = 1 + 2*coef
		a[i*lda+1] = -coef
	}
	// Boundary condition at the right node
	a[(n-1)*lda] = 0.5 + coef

	// Allocate the matrix A and the right-hand side b.
	ac.a = mat.NewSymBandDense(n, 1, a)
	ac.b = make([]float64, n)

	// Allocate and set up the initial condition.
	ac.u = make([]float64, n)
	for i := range ac.u {
		ac.u[i] = ac.InitCond(float64(i) * ac.h)
	}

	// Allocate the linear solver and the settings.
	ac.ls = &linsolve.CG{}
	ac.lssettings = linsolve.Settings{
		// Solution from the previous time step will be a good initial estimate.
		InitX: ac.u,
		// Store the solution into the existing slice.
		Dst: ac.u,
		// Provide context to reduce memory garbage.
		Work: linsolve.NewContext(n),
	}
}

// Step advances the solution one step in time.
func (ac *AllenCahnFD) Step() error {
	// Assemble the right-hand side vector b.
	coef := ac.tau / ac.Xi / ac.Xi
	n := len(ac.u)
	for i, ui := range ac.u {
		ac.b[i] = ui - coef*FPrime(ui)
		if i == 0 || i == n-1 {
			ac.b[i] *= 0.5
		}
	}
	_, err := linsolve.Iterative(ac, ac.b, ac.ls, &ac.lssettings)
	return err
}

// MulVecTo implements the MulVecToer interface.
func (ac *AllenCahnFD) MulVecTo(dst []float64, _ bool, x []float64) {
	n := len(x)
	ax := mat.NewVecDense(n, dst)
	ax.MulVec(ac.a, mat.NewVecDense(n, x))
}

func (ac *AllenCahnFD) Solution() []float64 {
	return ac.u
}

func output(u []float64, L float64, step int) {
	p, err := plot.New()
	if err != nil {
		log.Fatal(err)
	}

	p.Title.Text = fmt.Sprintf("Step %d", step)
	p.X.Label.Text = "x"
	p.X.Min = 0
	p.X.Max = L
	p.Y.Min = -1.1
	p.Y.Max = 1.1

	n := len(u)
	h := L / float64(n-1)
	pts := make(plotter.XYs, n)
	for i, ui := range u {
		pts[i].X = float64(i) * h
		pts[i].Y = ui
	}
	err = plotutil.AddLinePoints(p, "u", pts)
	if err != nil {
		log.Fatal(err)
	}
	err = p.Save(20*vg.Centimeter, 10*vg.Centimeter, fmt.Sprintf("u%04d.png", step))
	if err != nil {
		log.Fatal(err)
	}
}

func Example_evolutionPDE() {
	const (
		L        = 10.0
		n        = 1000
		tau      = 0.1 * L / n
		xi       = 6 * L / n
		numSteps = 200
	)
	rnd := rand.New(rand.NewSource(1))
	ac := AllenCahnFD{
		Xi: xi,
		InitCond: func(x float64) float64 {
			return 0.01 * rnd.NormFloat64()
		},
	}
	ac.Setup(n, L, tau)
	for i := 1; i <= numSteps; i++ {
		err := ac.Step()
		if err != nil {
			log.Fatal(err)
		}
		// output(ac.Solution(), L, i)
	}

	// Output:
}
