// Copyright ©2025 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package root_test

import (
	"fmt"
	"math"
	"testing"

	"gonum.org/v1/exp/root"
)

var findBracketMonoTests = []struct {
	name  string
	f     func(float64) float64
	guess float64
}{
	// Based on https://github.com/boostorg/math/blob/boost-1.88.0/test/test_toms748_solve.cpp
	{name: "f4.4", f: func(x float64) float64 { return math.Pow(x, 4) - 0.2 }, guess: 2},
	{name: "f4.6", f: func(x float64) float64 { return math.Pow(x, 6) - 0.2 }, guess: 2},
	{name: "f4.8", f: func(x float64) float64 { return math.Pow(x, 8) - 0.2 }, guess: 2},
	{name: "f4.10", f: func(x float64) float64 { return math.Pow(x, 10) - 0.2 }, guess: 2},
	{name: "f4.12", f: func(x float64) float64 { return math.Pow(x, 12) - 0.2 }, guess: 2},

	// Based on https://github.com/boostorg/math/blob/boost-1.88.0/test/test_root_finding_concepts.cpp
	{name: "f1", f: func(x float64) float64 { return x*x*x - 27 }, guess: 27},
}

func TestFindBracketMono(t *testing.T) {
	t.Parallel()

	for _, test := range findBracketMonoTests {
		t.Run(fmt.Sprintf("%s", test.name), func(t *testing.T) {
			a, b := root.FindBracketMono(test.f, test.guess)
			fa, fb := test.f(a), test.f(b)
			if fa*fb > 0 {
				t.Fatalf("%s: invalid bracket (%f, %f)", test.name, fa, fb)
			}
		})
	}
}

var findBracketMonoSpecials = []struct {
	name     string
	f        func(float64) float64
	guess    float64
	criteria func(a, b float64) bool
}{
	{name: "+Inf", f: func(x float64) float64 { return math.Atan(x) - 2 }, guess: 3, criteria: func(a, b float64) bool { return a > 0 && math.IsInf(b, 1) }},
	{name: "-Inf", f: func(x float64) float64 { return math.Atan(x) + 2 }, guess: 3, criteria: func(a, b float64) bool { return a < 0 && math.IsInf(b, -1) }},
	{name: "tiny positive", f: func(x float64) float64 {
		rt := math.SmallestNonzeroFloat64
		switch {
		case x > rt:
			return 1
		case x < rt:
			return -1
		default:
			return 0
		}
	}, guess: 3, criteria: func(a, b float64) bool { return a > 0 && b == 0 }},
	{name: "tiny negative", f: func(x float64) float64 {
		rt := -math.SmallestNonzeroFloat64
		switch {
		case x > rt:
			return 1
		case x < rt:
			return -1
		default:
			return 0
		}
	}, guess: -3, criteria: func(a, b float64) bool { return a < 0 && b == 0 }},
}

func TestFindBracketMonoSpecialCases(t *testing.T) {
	for _, tc := range findBracketMonoSpecials {
		t.Run(tc.name, func(t *testing.T) {
			a, b := root.FindBracketMono(tc.f, tc.guess)
			if !tc.criteria(a, b) {
				t.Fatalf("%s: wrong bracket (%g, %g)", tc.name, a, b)
			}
		})
	}
}
