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

func TestFindBracketMonoSpecialCases(t *testing.T) {
	// Bracket to positive infinity.
	f := func(x float64) float64 { return math.Atan(x) - 2 }
	var guess float64 = 3
	a, b := root.FindBracketMono(f, guess)
	if !(a > 0 && math.IsInf(b, 1)) {
		t.Errorf("FindBracketMono(Atan-2, %f): got (%f, %f) want (+a, +Inf)", guess, a, b)
	}

	// Bracket to negative infinity.
	f = func(x float64) float64 { return math.Atan(x) + 2 }
	guess = 3
	a, b = root.FindBracketMono(f, guess)
	if !(a < 0 && math.IsInf(b, -1)) {
		t.Errorf("FindBracketMono(Atan+2, %f): got (%f, %f) want (-a, -Inf)", guess, a, b)
	}

	// Root is a tiny positive value.
	f = func(x float64) float64 {
		rt := math.SmallestNonzeroFloat64
		switch {
		case x > rt:
			return 1
		case x < rt:
			return -1
		default:
			return 0
		}
	}
	guess = 3
	a, b = root.FindBracketMono(f, guess)
	if !(a > 0 && b == 0) {
		t.Errorf("FindBracketMono(f, %f): got (%g, %g) want (+a, 0)", guess, a, b)
	}

	// Root is a tiny negative value.
	f = func(x float64) float64 {
		rt := -math.SmallestNonzeroFloat64
		switch {
		case x > rt:
			return 1
		case x < rt:
			return -1
		default:
			return 0
		}
	}
	guess = -3
	a, b = root.FindBracketMono(f, guess)
	if !(a < 0 && b == 0) {
		t.Errorf("FindBracketMono(f, %f): got (%g, %g) want (-a, 0)", guess, a, b)
	}
}
