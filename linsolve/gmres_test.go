// Copyright ©2017 The Gonum authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package linsolve

import (
	"math"
	"math/rand"
	"testing"

	"gonum.org/v1/gonum/floats"
)

func TestGMRES(t *testing.T) {
	const convTol = 1e-14
	const defaultWantTol = 1e-10

	rnd := rand.New(rand.NewSource(1))
testLoop:
	for _, tc := range []testCase{
		randomSPD(1, rnd),
		randomSPD(2, rnd),
		randomSPD(3, rnd),
		randomSPD(4, rnd),
		randomSPD(5, rnd),
		randomSPD(10, rnd),
		randomSPD(20, rnd),
		randomSPD(50, rnd),
		randomSPD(100, rnd),
		randomSPD(200, rnd),
		randomSPD(500, rnd),
		market("spd_100_nos4", 1e-11),
		market("spd_138_bcsstm22", 1e-11),
		market("spd_237_nos1", 1e-9),
		market("spd_468_nos5", 1e-11),
		market("spd_485_bcsstm20", 1e-8),
		market("spd_900_gr_30_30", 1e-11),
		market("gen_236_e05r0100", 1e-10),
		market("gen_236_e05r0500", 1e-11),
		market("gen_434_hor__131", 1e-11),
		market("gen_886_orsirr_2", 1e-11),
	} {
		n := tc.n
		// Compute the right-hand side b so that a predetermined vector
		// is the solution.
		want := make([]float64, n)
		for i := range want {
			want[i] = 2 + 0.1*float64(i%10)
		}
		b := make([]float64, n)
		tc.mulvec(b, want, false)
		bnorm := floats.Norm(b, 2)

		ctx := Context{
			X:        make([]float64, n),
			Residual: make([]float64, n),
			Src:      make([]float64, n),
			Dst:      make([]float64, n),
		}
		// Initial guess is a random vector.
		for i := range ctx.X {
			ctx.X[i] = rnd.NormFloat64()
		}
		// Compute the initial residual.
		tc.mulvec(ctx.Residual, ctx.X, false)
		floats.AddScaledTo(ctx.Residual, b, -1, ctx.Residual)

		var itercount int
		var g GMRES
		g.Init(n)
	solverLoop:
		for {
			op, err := g.Iterate(&ctx)
			if err != nil {
				t.Errorf("Case %v (n=%v): unexpected error %v", tc.name, n, err)
				continue testLoop
			}
			switch op {
			case MulVec:
				tc.mulvec(ctx.Dst, ctx.Src, false)
			case PreconSolve:
				copy(ctx.Dst, ctx.Src)
			case ComputeResidual:
				tc.mulvec(ctx.Residual, ctx.X, false)
				floats.AddScaledTo(ctx.Residual, b, -1, ctx.Residual)
			case CheckResidualNorm:
				ctx.Converged = ctx.ResidualNorm < convTol*bnorm
			case MajorIteration:
				itercount++
				if floats.Norm(ctx.Residual, 2) < convTol*bnorm {
					break solverLoop
				}
				if itercount == tc.iters {
					t.Logf("Case %v (n=%v): %v exceeded iteration limit (rnorm=%v)", tc.name, n, itercount, floats.Norm(ctx.Residual, 2)/bnorm)
					break solverLoop
				}
			}
		}

		tol := tc.tol
		if tol == 0 {
			tol = defaultWantTol
		}
		dist := floats.Distance(ctx.X, want, math.Inf(1))
		if dist > tol {
			t.Errorf("Case %v (n=%v): unexpected solution, |want-got|=%v", tc.name, n, dist)
		}
	}
}
