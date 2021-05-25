package ode

import (
	"math"

	"gonum.org/v1/gonum/mat"
)

type DoPri5 struct {
	param    parameters
	adaptive bool
	x, aux   *mat.VecDense
	// solutions
	y5, err45 *mat.VecDense
	// integration coefficients (for Butcher Tableau)
	k1, k2, k3, k4, k5, k6, k7 *mat.VecDense
	dom                        float64
	fx                         func(y *mat.VecDense, t float64, x mat.Vector)
}

// Set implements the Integrator interface. Initializes a Dormand Prince method
// for use as a solver
func (dp *DoPri5) Set(ivp IVP) {

	dp.fx = ivp.Func

	t0, x0 := ivp.IV()
	dp.dom = t0
	nx := x0.Len()
	dp.k1, dp.k2, dp.k3 = mat.NewVecDense(nx, nil), mat.NewVecDense(nx, nil), mat.NewVecDense(nx, nil)
	dp.k4, dp.k5, dp.k6 = mat.NewVecDense(nx, nil), mat.NewVecDense(nx, nil), mat.NewVecDense(nx, nil)
	dp.k7 = mat.NewVecDense(nx, nil)
	dp.aux = mat.NewVecDense(nx, nil)
	dp.x = mat.NewVecDense(nx, nil)
	dp.x.CloneFromVec(x0)

	dp.y5 = mat.NewVecDense(nx, nil)
	dp.err45 = mat.NewVecDense(nx, nil)

}

// Step implements Integrator interface. Advances solution by step h. If algorithm
// is set to adaptive then h is just a suggestion
func (dp *DoPri5) Step(y4 *mat.VecDense, h float64) (float64, error) {
	const c20, c21 = 1. / 5., 1. / 5.
	const c30, c31, c32 = 3. / 10., 3. / 40., 9. / 40.
	const c40, c41, c42, c43 = 4. / 5., 44. / 45., -56. / 15., 32. / 9.
	const c50, c51, c52, c53, c54 = 8. / 9., 19372. / 6561., -25360. / 2187., 64448. / 6561., -212. / 729.
	const c60, c61, c62, c63, c64, c65 = 1., 9017. / 3168., -355. / 33., 46732. / 5247., 49. / 176., -5103. / 18656.
	const c70, c71, c72, c73, c74, c75, c76 = 1., 35. / 384., 0., 500. / 1113., 125. / 192., -2187. / 6784., 11. / 84.
	// Alternate solution for error calculation
	const a1, a3, a4, a5, a6, a7 = 5179. / 57600., 7571. / 16695., 393. / 640., -92097. / 339200., 187. / 2100., 1. / 40.
	// Fourth order solution is used to advance
	const b1, b3, b4, b5, b6 = 35. / 384., 500. / 1113., 125. / 192., -2187. / 6784., 11. / 84.

	// prettier variable names, also someone once said pointers are equivalent to JMP
	x := dp.x
	// auxiliary vector for storing results of calling ivp.Func
	aux := dp.aux
	F := dp.fx
	t := dp.dom
	// if adaptive then this tag will be used until h==DoPri5.MinStep
SOLVE:
	F(aux, t, x)
	dp.k1.ScaleVec(h, aux)

	dp.k2.AddScaledVec(x, c21, dp.k1)
	F(aux, t+c20*h, dp.k2)
	dp.k2.ScaleVec(h, aux)

	dp.k3.AddScaledVec(x, c31, dp.k1)
	dp.k3.AddScaledVec(dp.k3, c32, dp.k2)
	F(aux, t+h*c30, dp.k3)
	dp.k3.ScaleVec(h, aux)

	dp.k4.AddScaledVec(x, c41, dp.k1)
	dp.k4.AddScaledVec(dp.k4, c42, dp.k2)
	dp.k4.AddScaledVec(dp.k4, c43, dp.k3)
	F(aux, t+h*c40, dp.k4)
	dp.k4.ScaleVec(h, aux)

	dp.k5.AddScaledVec(x, c51, dp.k1)
	dp.k5.AddScaledVec(dp.k5, c52, dp.k2)
	dp.k5.AddScaledVec(dp.k5, c53, dp.k3)
	dp.k5.AddScaledVec(dp.k5, c54, dp.k4)
	F(aux, t+h*c50, dp.k5)
	dp.k5.ScaleVec(h, aux)

	dp.k6.AddScaledVec(x, c61, dp.k1)
	dp.k6.AddScaledVec(dp.k6, c62, dp.k2)
	dp.k6.AddScaledVec(dp.k6, c63, dp.k3)
	dp.k6.AddScaledVec(dp.k6, c64, dp.k4)
	dp.k6.AddScaledVec(dp.k6, c65, dp.k5)
	F(aux, t+h*c60, dp.k6)
	dp.k6.ScaleVec(h, aux)

	y4.AddScaledVec(x, b1, dp.k1)
	y4.AddScaledVec(y4, b3, dp.k3)
	y4.AddScaledVec(y4, b4, dp.k4)
	y4.AddScaledVec(y4, b5, dp.k5)
	y4.AddScaledVec(y4, b6, dp.k6)
	// fourth order approximation used to advance solution

	if dp.adaptive {
		y5 := dp.y5
		dp.k7.AddScaledVec(x, c71, dp.k1)
		dp.k7.AddScaledVec(dp.k7, c72, dp.k2)
		dp.k7.AddScaledVec(dp.k7, c73, dp.k3)
		dp.k7.AddScaledVec(dp.k7, c74, dp.k4)
		dp.k7.AddScaledVec(dp.k7, c75, dp.k5)
		dp.k7.AddScaledVec(dp.k7, c76, dp.k6)
		// calculation of order 5 coefficients

		F(aux, t+h*c70, dp.k7)
		dp.k7.ScaleVec(h, aux)

		y5.AddScaledVec(x, a1, dp.k1)
		y5.AddScaledVec(y5, a3, dp.k3)
		y5.AddScaledVec(y5, a4, dp.k4)
		y5.AddScaledVec(y5, a5, dp.k5)
		y5.AddScaledVec(y5, a6, dp.k6)
		y5.AddScaledVec(y5, a7, dp.k7)
		dp.err45.SubVec(y4, y5)

		// compute error between fifth order solution and fourth order solution
		errRatio := dp.param.atol / mat.Norm(dp.err45, math.Inf(1))
		hnew := math.Min(math.Max(0.9*h*math.Pow(errRatio, .2), dp.param.minStep), dp.param.maxStep)
		// given "bad" error ratio, algorithm will recompute with a smaller step, given it is not the minimum permissible step or less
		if errRatio < 1 && h > dp.param.minStep {
			h = hnew
			goto SOLVE
		}
	}
	// advance solution with fourth order solution
	dp.x.CloneFromVec(y4)
	dp.dom += h
	return h, nil
}

// NewDormandPrince5 returns a adaptive-ready Dormand Prince solver of order 5
//
// To enable step size adaptivity minimum step size must be set and
// absolute tolerance must be set. i.e:
//
//  NewDormandPrince5(ConfigScalarTolerance(0, 0.1), ConfigStepLimits(1, 1e-3))
//
// If a invalid configuration is passed the function panics.
func NewDormandPrince5(configs ...Configuration) *DoPri5 {
	dp := new(DoPri5)
	for i := range configs {
		err := configs[i](&dp.param)
		if err != nil {
			panic(err)
		}
	}

	if dp.param.atol > 0 && dp.param.minStep > 0 {
		dp.adaptive = true
	}
	return dp
}
