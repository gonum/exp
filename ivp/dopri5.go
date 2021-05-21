package ivp

import (
	"math"

	"gonum.org/v1/gonum/floats"
)

type DoPri5 struct {
	maxError, minStep, maxStep float64
	adaptive                   bool
	x, u                       []float64
	// solutions
	y5, err45 []float64
	// integration coefficients (for Butcher Tableau)
	k1, k2, k3, k4, k5, k6, k7 []float64
	dom                        float64
	fx, fu                     func(y []float64, t float64, x, u []float64)
}

// Set implements the Integrator interface. Initializes a Dormand Prince method
// for use as a solver
func (dp *DoPri5) Set(t0 float64, ivp IVP) error {
	dp.dom = t0
	xequations, ufunc := ivp.Equations()
	dp.fx = xequations
	dp.fu = ufunc
	nx, nu := ivp.Dims()
	dp.k1, dp.k2, dp.k3 = make([]float64, nx), make([]float64, nx), make([]float64, nx)
	dp.k4, dp.k5, dp.k6 = make([]float64, nx), make([]float64, nx), make([]float64, nx)
	dp.k7 = make([]float64, nx)
	dp.x = make([]float64, nx)
	dp.y5 = make([]float64, nx)
	dp.err45 = make([]float64, nx)
	dp.u = make([]float64, nu)
	return nil
}

// Step implements Integrator interface. Advances solution by step h. If algorithm
// is set to adaptive then h is just a suggestion
func (dp *DoPri5) Step(y4 []float64, h float64) (float64, error) {
	const c20, c21 = 1. / 5., 1. / 5.
	const c30, c31, c32 = 3. / 10., 3. / 40., 9. / 40.
	const c40, c41, c42, c43 = 4. / 5., 44. / 45., -56. / 15., 32. / 9
	const c50, c51, c52, c53, c54 = 8. / 9., 19372. / 6561., -25360. / 2187., 64448. / 6561., -212. / 729.
	const c60, c61, c62, c63, c64, c65 = 1., 9017. / 3168., -355. / 33., 46732. / 5247., 49. / 176., -5103. / 18656.
	const c70, c71, c72, c73, c74, c75, c76 = 1., 35. / 384., 0., 500. / 1113., 125. / 192., -2187. / 6784., 11. / 84.
	// Alternate solution for error calculation
	const a1, a3, a4, a5, a6, a7 = 5179. / 57600., 7571. / 16695., 393. / 640., -92097. / 339200., 187. / 2100., 1. / 40.
	// Fourth order solution is used to advance
	const b1, b3, b4, b5, b6 = 35. / 384., 500. / 1113., 125. / 192., -2187. / 6784., 11. / 84.

	// prettier variable names, also someone once said pointers are equivalent to JMP
	x, u := dp.x, dp.u
	F := dp.fx
	t := dp.dom
	// if adaptive then this tag will be used until h==DoPri5.MinStep
SOLVE:
	F(dp.k1, t, x, u)
	floats.Scale(h, dp.k1)

	floats.AddScaledTo(dp.k2, x, c21, dp.k1)
	F(dp.k2, t+c20*h, x, u)
	floats.Scale(h, dp.k2)

	floats.AddScaledTo(dp.k3, x, c31, dp.k1)
	floats.AddScaled(dp.k3, c32, dp.k2)
	F(dp.k3, t+c30*h, x, u)
	floats.Scale(h, dp.k3)

	floats.AddScaledTo(dp.k4, x, c41, dp.k1)
	floats.AddScaled(dp.k4, c42, dp.k2)
	floats.AddScaled(dp.k4, c43, dp.k3)
	F(dp.k4, t+c40*h, x, u)
	floats.Scale(h, dp.k4)

	floats.AddScaledTo(dp.k5, x, c51, dp.k1)
	floats.AddScaled(dp.k5, c52, dp.k2)
	floats.AddScaled(dp.k5, c53, dp.k3)
	floats.AddScaled(dp.k5, c54, dp.k4)
	F(dp.k5, t+c50*h, x, u)
	floats.Scale(h, dp.k5)

	floats.AddScaledTo(dp.k6, x, c61, dp.k1)
	floats.AddScaled(dp.k6, c62, dp.k2)
	floats.AddScaled(dp.k6, c63, dp.k3)
	floats.AddScaled(dp.k6, c64, dp.k4)
	floats.AddScaled(dp.k6, c65, dp.k5)
	F(dp.k6, t+c60*h, x, u)
	floats.Scale(h, dp.k6)

	// fourth order approximation used to advance solution
	floats.AddScaledTo(y4, x, b1, dp.k1)
	floats.AddScaled(y4, b3, dp.k3)
	floats.AddScaled(y4, b4, dp.k4)
	floats.AddScaled(y4, b5, dp.k5)
	floats.AddScaled(y4, b6, dp.k6)
	if dp.adaptive {
		// calculation of order 5 coefficients
		floats.AddScaledTo(dp.k7, x, c71, dp.k1)
		floats.AddScaled(dp.k7, c72, dp.k2)
		floats.AddScaled(dp.k7, c73, dp.k3)
		floats.AddScaled(dp.k7, c74, dp.k4)
		floats.AddScaled(dp.k7, c75, dp.k5)
		floats.AddScaled(dp.k7, c76, dp.k6)
		F(dp.k7, t+h*c70, x, u)
		floats.Scale(h, dp.k7)

		floats.AddScaledTo(dp.y5, x, a1, dp.k1)
		floats.AddScaled(dp.y5, a3, dp.k3)
		floats.AddScaled(dp.y5, a4, dp.k4)
		floats.AddScaled(dp.y5, a5, dp.k5)
		floats.AddScaled(dp.y5, a6, dp.k6)
		floats.AddScaled(dp.y5, a7, dp.k7)

		// compute error between fifth order solution and fourth order solution
		errRatio := dp.maxError / floats.Norm(floats.SubTo(dp.err45, y4, dp.y5), math.Inf(1))
		hnew := math.Min(math.Max(0.9*h*math.Pow(errRatio, .2), dp.minStep), dp.maxStep)
		// given "bad" error ratio, algorithm will recompute with a smaller step, given it is not the minimum permissible step or less
		if errRatio < 1 && h > dp.minStep {
			h = hnew
			goto SOLVE
		}
	}
	// advance solution with fourth order solution
	copy(dp.x, y4)
	if dp.fu != nil {
		dp.fu(u, t+h, dp.x, u)
	}
	return h, nil
}

// NewDormandPrince5 returns a adaptive-ready Dormand Prince solver of order 5
func NewDormandPrince5(maxError, maxStep, minStep float64) *DoPri5 {
	dp := new(DoPri5)
	if maxError > 0 && maxStep > minStep && minStep > 0 {
		dp.adaptive = true
		dp.maxError = maxError
		dp.minStep, dp.maxStep = minStep, maxStep
	} else {
		panic("NewDormandPrince5 generator requires proper arguments: maxError > 0 && maxStep > minStep && minStep > 0")
	}
	return dp
}
