// Copyright ©2017 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package linsolve

import (
	"math"
	"testing"

	"golang.org/x/exp/rand"

	"gonum.org/v1/gonum/floats"
)

func TestDefaultMethodDefaultSettings(t *testing.T) {
	rnd := rand.New(rand.NewSource(1))

	testCases := spdTestCases(rnd)
	testCases = append(testCases,
		nonsym3x3(),
		nonsymTridiag(100),
		newGreenbaum54(1, 1, rnd),
		newGreenbaum54(1, 2, rnd),
		newGreenbaum54(2, 4, rnd),
		newGreenbaum54(10, 0, rnd),
		newGreenbaum54(10, 20, rnd),
		newGreenbaum54(50, 3, rnd),
		newGreenbaum73(16, 16, rnd),
		newPDENonsymmetric(16, 16, rnd),
		newPDEYang47(16, 16, rnd),
		newPDEYang48(16, 16, rnd),
		newPDEYang49(16, 16, rnd),
		newPDEYang410(16, 16, rnd),
		newPDEYang412(16, 16, rnd),
		newPDEYang413(16, 16, rnd),
		newPDEYang414(16, 16, rnd),
		newPDEYang415(16, 16, rnd),
	)
	for _, tc := range testCases {
		testMethodWithSettings(t, nil, nil, tc)
	}
}

func TestCG(t *testing.T) {
	rnd := rand.New(rand.NewSource(1))

	testCases := spdTestCases(rnd)
	for _, tc := range testCases {
		s := newTestSettings(rnd, tc)
		testMethodWithSettings(t, &CG{}, s, tc)
	}
}

func TestCGDefaultSettings(t *testing.T) {
	rnd := rand.New(rand.NewSource(1))

	testCases := spdTestCases(rnd)
	for _, tc := range testCases {
		testMethodWithSettings(t, &CG{}, nil, tc)
	}
}

func TestBiCG(t *testing.T) {
	rnd := rand.New(rand.NewSource(1))

	testCases := spdTestCases(rnd)
	testCases = append(testCases,
		nonsym3x3(),
		nonsymTridiag(100),
		newGreenbaum54(1, 1, rnd),
		newGreenbaum54(1, 2, rnd),
		newGreenbaum54(2, 4, rnd),
		newGreenbaum54(10, 0, rnd),
		newGreenbaum54(10, 20, rnd),
		newGreenbaum54(50, 3, rnd),
		newGreenbaum73(16, 16, rnd),
		newPDEYang47(16, 16, rnd),
		newPDEYang48(16, 16, rnd),
		newPDEYang410(16, 16, rnd),
		newPDEYang413(16, 16, rnd),
		newPDEYang414(16, 16, rnd),
		newPDEYang415(16, 16, rnd),
	)
	for _, tc := range testCases {
		s := newTestSettings(rnd, tc)
		s.Tolerance = 1e-10
		testMethodWithSettings(t, &BiCG{}, s, tc)
	}
}

func TestBiCGDefaultSettings(t *testing.T) {
	rnd := rand.New(rand.NewSource(1))

	testCases := spdTestCases(rnd)
	testCases = append(testCases,
		nonsym3x3(),
		nonsymTridiag(100),
		newGreenbaum54(1, 1, rnd),
		newGreenbaum54(1, 2, rnd),
		newGreenbaum54(2, 4, rnd),
		newGreenbaum54(10, 0, rnd),
		newGreenbaum54(10, 20, rnd),
		newGreenbaum54(50, 3, rnd),
		newGreenbaum73(16, 16, rnd),
		newPDENonsymmetric(16, 16, rnd),
		newPDEYang47(16, 16, rnd),
		newPDEYang48(16, 16, rnd),
		newPDEYang410(16, 16, rnd),
		newPDEYang413(16, 16, rnd),
		newPDEYang414(16, 16, rnd),
		newPDEYang415(16, 16, rnd),
	)
	for _, tc := range testCases {
		testMethodWithSettings(t, &BiCG{}, nil, tc)
	}
}

func TestBiCGStab(t *testing.T) {
	rnd := rand.New(rand.NewSource(1))

	testCases := spdTestCases(rnd)
	testCases = append(testCases,
		nonsym3x3(),
		nonsymTridiag(100),
		newGreenbaum54(1, 1, rnd),
		newGreenbaum54(1, 2, rnd),
		newGreenbaum54(2, 4, rnd),
		newGreenbaum54(10, 0, rnd),
		newGreenbaum54(10, 20, rnd),
		newGreenbaum54(50, 3, rnd),
		newGreenbaum73(16, 16, rnd),
	)
	for _, tc := range testCases {
		s := newTestSettings(rnd, tc)
		testMethodWithSettings(t, &BiCGStab{}, s, tc)
	}
}

func TestBiCGStabDefaultSettings(t *testing.T) {
	rnd := rand.New(rand.NewSource(1))

	testCases := spdTestCases(rnd)
	testCases = append(testCases,
		nonsym3x3(),
		nonsymTridiag(100),
		newGreenbaum54(1, 1, rnd),
		newGreenbaum54(1, 2, rnd),
		newGreenbaum54(2, 4, rnd),
		newGreenbaum54(10, 0, rnd),
		newGreenbaum54(10, 20, rnd),
		newGreenbaum54(50, 3, rnd),
		newGreenbaum73(16, 16, rnd),
	)
	for _, tc := range testCases {
		testMethodWithSettings(t, &BiCGStab{}, nil, tc)
	}
}

func TestGMRES(t *testing.T) {
	rnd := rand.New(rand.NewSource(1))

	testCases := spdTestCases(rnd)
	testCases = append(testCases,
		nonsym3x3(),
		nonsymTridiag(100),
		newGreenbaum54(1, 1, rnd),
		newGreenbaum54(1, 2, rnd),
		newGreenbaum54(2, 4, rnd),
		newGreenbaum54(10, 0, rnd),
		newGreenbaum54(10, 20, rnd),
		newGreenbaum54(50, 3, rnd),
		newGreenbaum73(16, 16, rnd),
		newPDENonsymmetric(16, 16, rnd),
		newPDEYang47(16, 16, rnd),
		newPDEYang48(16, 16, rnd),
		newPDEYang49(16, 16, rnd),
		newPDEYang410(16, 16, rnd),
		newPDEYang412(16, 16, rnd),
		newPDEYang413(16, 16, rnd),
		newPDEYang414(16, 16, rnd),
		newPDEYang415(16, 16, rnd),
	)
	for _, tc := range testCases {
		s := newTestSettings(rnd, tc)
		testMethodWithSettings(t, &GMRES{}, s, tc)
	}
}

func TestGMRESDefaultSettings(t *testing.T) {
	rnd := rand.New(rand.NewSource(1))

	testCases := spdTestCases(rnd)
	testCases = append(testCases,
		nonsym3x3(),
		nonsymTridiag(100),
		newGreenbaum54(1, 1, rnd),
		newGreenbaum54(1, 2, rnd),
		newGreenbaum54(2, 4, rnd),
		newGreenbaum54(10, 0, rnd),
		newGreenbaum54(10, 20, rnd),
		newGreenbaum54(50, 3, rnd),
		newGreenbaum73(16, 16, rnd),
		newPDENonsymmetric(16, 16, rnd),
		newPDEYang47(16, 16, rnd),
		newPDEYang48(16, 16, rnd),
		newPDEYang49(16, 16, rnd),
		newPDEYang410(16, 16, rnd),
		newPDEYang412(16, 16, rnd),
		newPDEYang413(16, 16, rnd),
		newPDEYang414(16, 16, rnd),
		newPDEYang415(16, 16, rnd),
	)
	for _, tc := range testCases {
		testMethodWithSettings(t, &GMRES{}, nil, tc)
	}
}

func newTestSettings(rnd *rand.Rand, tc testCase) *Settings {
	n := len(tc.b)

	// Initial guess is a random vector.
	initX := make([]float64, n)
	for i := range initX {
		initX[i] = rnd.NormFloat64()
	}

	// Preallocate a destination slice and fill it with NaN.
	dst := make([]float64, n)
	for i := range dst {
		dst[i] = math.NaN()
	}

	// Preallocate a work context and fill it with NaN.
	work := NewContext(n)
	for i := range work.X {
		work.X[i] = math.NaN()
		work.Src[i] = math.NaN()
		work.Dst[i] = math.NaN()
	}
	work.ResidualNorm = math.NaN()

	return &Settings{
		InitX:         initX,
		Dst:           dst,
		Tolerance:     tc.tol,
		MaxIterations: 5 * n,
		PreconSolve:   tc.PreconSolve,
		Work:          work,
	}
}

func testMethodWithSettings(t *testing.T, m Method, s *Settings, tc testCase) {
	wantTol := 1e-9
	if s == nil {
		// The default value of Settings.Tolerance is not as low as the tolerance in
		// individual test cases, therefore we must use a higher tolerance for
		// the expected accuracy of the computed solution.
		wantTol = 1e-7
	}

	bCopy := make([]float64, len(tc.b))
	copy(bCopy, tc.b)

	result, err := Iterative(&tc, bCopy, m, s)
	if err != nil {
		t.Errorf("%v: unexpected error %v", tc.name, err)
		return
	}

	if !floats.Equal(tc.b, bCopy) {
		t.Errorf("%v: unexpected modification of b", tc.name)
	}

	dist := floats.Distance(result.X, tc.want, 2) / floats.Norm(result.X, 2)
	if dist > wantTol {
		t.Errorf("%v: unexpected solution, |want-got|/|want|=%v", tc.name, dist)
	}

	if s == nil {
		return
	}

	if s.MaxIterations > 0 && result.Stats.Iterations > s.MaxIterations {
		t.Errorf("%v: Result.Stats.Iterations greater than Settings.MaxIterations", tc.name)
	}

	if s.Dst != nil {
		if floats.HasNaN(s.Dst) {
			t.Errorf("%v: Settings.Dst was not used", tc.name)
		}
		if !floats.Equal(s.Dst, result.X) {
			t.Errorf("%v: Settings.Dst and Result.X not equal", tc.name)
		}
		result.X[0] = 123456.7
		if s.Dst[0] != result.X[0] {
			t.Errorf("%v: Settings.Dst and Result.X are not the same slice", tc.name)
		}
	}

	if s.Work != nil {
		if floats.HasNaN(s.Work.X) || floats.HasNaN(s.Work.Src) || floats.HasNaN(s.Work.Dst) {
			t.Errorf("%v: Settings.Work was not used", tc.name)
		}
	}
}
