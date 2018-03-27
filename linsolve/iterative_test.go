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

func TestIterativeWithDefault(t *testing.T) {
	rnd := rand.New(rand.NewSource(1))
	testCases := spdTestCases(rnd)
	testCases = append(testCases, unsymTestCases()...)
	testIterative(t, nil, testCases)
}

func TestIterativeWithCG(t *testing.T) {
	rnd := rand.New(rand.NewSource(1))
	testCases := spdTestCases(rnd)
	testIterative(t, &CG{}, testCases)
}

func TestIterativeWithGMRES(t *testing.T) {
	rnd := rand.New(rand.NewSource(1))
	testCases := spdTestCases(rnd)
	testCases = append(testCases, unsymTestCases()...)
	testIterative(t, &GMRES{}, testCases)
}

func testIterative(t *testing.T, m Method, testCases []testCase) {
	const convTol = 1e-14
	const defaultWantTol = 1e-10
	rnd := rand.New(rand.NewSource(1))
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

		// Initial guess is a random vector.
		x := make([]float64, n)
		for i := range x {
			x[i] = rnd.NormFloat64()
		}

		_, err := Iterative(x, tc, b, m, Settings{
			InitX:         x,
			Tolerance:     convTol,
			MaxIterations: 40 * n,
		})
		if err != nil {
			t.Errorf("Case %v: unexpected error %v", tc.name, err)
			continue
		}

		wantTol := tc.tol
		if wantTol == 0 {
			wantTol = defaultWantTol
		}
		dist := floats.Distance(x, want, math.Inf(1))
		if dist > wantTol {
			t.Errorf("Case %v: unexpected solution, |want-got|=%v", tc.name, dist)
		}
	}
}
