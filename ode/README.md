# Gonum ivp [![GoDoc](https://godoc.org/gonum.org/v1/gonum/ivp?status.svg)](https://godoc.org/gonum.org/v1/gonum/ivp)

package ode provides numerical methods solving multivariable ODE initial value problems for the Go programming language.

Below is an example of usage
```go
	const (
		g = -10. // gravity field [m.s^-2]
	)
	// we declare our physical model. First argument is initial time, which is 0 seconds.
	// Next is the initial state vector, which corresponds to 100 meters above the ground
	// with 0 m/s velocity.
	ballModel, err := ode.NewIVP(0, mat.NewVecDense(2, []float64{100., 0.}),
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
	var solver ode.Integrator = ode.NewDormandPrince5(ode.DefaultParam)
	// Solve function makes it easy to integrate a problem without having
	// to implement the `for` loop. This example integrates the IVP with a step size
	// of 0.1 over a domain of 10. arbitrary units, in this case, 10 seconds.
	results, err := ode.SolveIVP(ballModel, solver, 0.1, 10.)
	fmt.Println(results)
```