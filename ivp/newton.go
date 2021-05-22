package ivp

import (
	"math"

	"gonum.org/v1/exp/linsolve"
	"gonum.org/v1/gonum/diff/fd"
	"gonum.org/v1/gonum/floats"
	"gonum.org/v1/gonum/mat"
)

type NewtonRaphson struct {
	param parameters
	x, b  []float64
	//guess contains best guess for next X vector which zeros the residual function
	guess, err []float64

	dom         float64
	fx          func(y []float64, t float64, x []float64)
	residualers func(h, t float64, xnow []float64) func(y, xnext []float64)
	// Residual functions jacobian
	J      *mat.Dense
	Jband  *mat.BandDense
	result *linsolve.Result
}

func (nr *NewtonRaphson) Step(y []float64, h float64) (float64, error) {
	var err error
	if nr.param.atol <= 0 {
		panic("Newton method requires error")
	}
	guess := nr.guess

	nx := len(nr.x)
	copy(guess, nr.x)
	t := nr.dom + h
	iter := 0
	// while step length does not change in iter loop, this may remain outside
	F := nr.residualers(h, t, nr.x)
	aerr := math.Inf(1)
	for iter < 10 && aerr > nr.param.atol {
		// set b vector with "old" x vector
		F(nr.b, guess)
		b := mat.NewVecDense(nx, nr.b)

		fd.Jacobian(nr.J, F, guess, nil)
		for i := 0; i < nx; i++ {
			for j := 0; j < nx; j++ {
				nr.Jband.SetBand(i, j, nr.J.At(i, j)) // we clone Jacobian Dense to our Jacobian BandDense for MulVecToer
			}
		}
		nr.result, err = linsolve.Iterative(nr.Jband, b, &linsolve.GMRES{}, &linsolve.Settings{MaxIterations: 2})
		if err != nil {
			return 0, err
		}
		resx := nr.result.X.RawVector().Data
		// X_(i+1) = X_(i) - alpha * F(X_(g)) / J(X_(g)) where g are guesses, and alpha is the relaxation factor
		// resx now contains the next guess. we proceed to calculate error.
		floats.AddScaledTo(resx, guess, nr.param.relax-1., resx)
		// error calculation if enabled
		if nr.param.atol > 0 {
			floats.SubTo(nr.err, guess, resx)
			aerr = floats.Norm(nr.err, math.Inf(1))
		}
		// advance solution
		copy(guess, resx)
		iter++
	}
	copy(nr.x, guess)
	copy(y, guess)
	return h, nil
}

func (nr *NewtonRaphson) Set(t0 float64, ivp IVP) error {
	nr.dom = t0
	xequations := ivp.Equations()
	nr.fx = xequations
	// nr.fu = ufunc
	nx := ivp.IV().Len()

	nr.x = make([]float64, nx)
	for i := range nr.x {
		nr.x[i] = ivp.IV().AtVec(i)
	}
	nr.err = make([]float64, nx)
	nr.guess = make([]float64, nx)

	nr.b = make([]float64, nx)
	nr.J = mat.NewDense(nx, nx, nil)
	nr.Jband = mat.NewBandDense(nx, nx, nx-1, nx-1, nil)

	// First propose residual functions such that
	// F(X_(i+1)) = 0 = X_(i+1) - X_(i) - step * f(X_(i+1))
	// where f is the vector of differential equations. We
	// seek the zero of the function F evaluated at xnext a.k.a X_(i+1)
	nr.residualers = func(h, t float64, xnow []float64) func(r, xnext []float64) {
		return func(r, xnext []float64) {
			// store results of function in r
			nr.fx(r, t, xnext)
			floats.SubTo(r, xnext, floats.AddTo(r, xnow, floats.ScaleTo(r, h, r)))
		}
	}
	return nil
}
func (nr *NewtonRaphson) XLen() int { return len(nr.x) }
func NewNewtonRaphson(configs ...Configuration) *NewtonRaphson {
	nr := new(NewtonRaphson)
	for i := range configs {
		err := configs[i](&nr.param)
		if err != nil {
			panic(err)
		}
	}
	return nr
}
