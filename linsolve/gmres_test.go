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
	testCases := spdTestCases(rnd)
	testCases = append(testCases, unsymTestCases()...)

testLoop:
	for _, tc := range testCases {
		n := tc.n
		// Compute the right-hand side b so that a predetermined vector
		// is the solution.
		want := make([]float64, n)
		for i := range want {
			want[i] = 2 + 0.1*float64(i%10)
		}
		b := make([]float64, n)
		tc.MulVecTo(b, false, want)
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
		tc.MulVecTo(ctx.Residual, false, ctx.X)
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
				tc.MulVecTo(ctx.Dst, false, ctx.Src)
			case PreconSolve:
				copy(ctx.Dst, ctx.Src)
			case ComputeResidual:
				tc.MulVecTo(ctx.Residual, false, ctx.X)
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

		wantTol := tc.tol
		if wantTol == 0 {
			wantTol = defaultWantTol
		}
		dist := floats.Distance(ctx.X, want, math.Inf(1))
		if dist > wantTol {
			t.Errorf("Case %v (n=%v): unexpected solution, |want-got|=%v", tc.name, n, dist)
		}
	}
}
