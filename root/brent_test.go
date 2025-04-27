// Copyright ©2025 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package root_test

import (
	"fmt"
	"math"
	"testing"

	"gonum.org/v1/exp/root"
	"gonum.org/v1/gonum/floats/scalar"
)

var eps = math.Nextafter(1, 2) - 1

func TestBrent(t *testing.T) {
	t.Parallel()

	type testcase struct {
		name string
		f    func(float64) float64
		a    float64
		b    float64
		tol  float64
		want float64
	}
	var tests []testcase
	for _, f := range tstutilsFns {
		tests = append(tests, testcase{name: f.name, f: f.f, a: 0.5, b: math.Sqrt(3), tol: 4 * eps, want: 1})
	}
	for _, tc := range aps {
		tests = append(tests, testcase{name: tc.name, f: tc.f, a: tc.a, b: tc.b, tol: 4 * eps, want: tc.root})
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%s", test.name), func(t *testing.T) {
			got, err := root.Brent(test.f, test.a, test.b, test.tol)
			if !scalar.EqualWithinAbsOrRel(got, test.want, 2e-12, test.tol) {
				if err != nil {
					t.Fatalf("%s: %s", test.name, err)
				}
				t.Fatalf("%s: got %f want %f", test.name, got, test.want)
			}
		})
	}
}

func TestScipyIssue5557(t *testing.T) {
	t.Parallel()

	f := func(x float64) float64 {
		if x < 0.5 {
			return -0.1
		}
		return x - 0.6
	}
	var a, b float64 = 0, 1
	tol := 4 * eps
	got, err := root.Brent(f, a, b, tol)
	if err != nil {
		t.Fatalf("Brent(%f, %f, %f): %s", a, b, tol, err)
	}
	const want = 0.6
	if !scalar.EqualWithinRel(got, want, tol) {
		t.Fatalf("Brent(%f, %f, %f): got %f want %f", a, b, tol, got, want)
	}
}

func TestScipyIssue13737(t *testing.T) {
	t.Parallel()

	// Constants based on https://github.com/scipy/scipy/blob/main/scipy/optimize/tests/test_zeros.py
	tests := []struct {
		a    float64
		b    float64
		root float64
	}{
		{a: -450, b: -350, root: -400},
		{a: 350, b: 450, root: 400},
	}
	for i, test := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			c := math.Exp(test.root)
			f := func(x float64) float64 { return math.Exp(x) - c }
			tol := eps
			got, err := root.Brent(f, test.a, test.b, tol)
			if err != nil {
				t.Fatalf("Brent(%f, %f, %f): %s", test.a, test.b, tol, err)
			}
			if !scalar.EqualWithinRel(got, test.root, tol) {
				t.Fatalf("Brent(%f, %f, %f): got %f want %f", test.a, test.b, tol, got, test.root)
			}
		})
	}
}

func TestScipyIssue5584(t *testing.T) {
	t.Parallel()

	f := func(x float64) float64 { return 1e-200 * x }

	// Report failure when signs are the same.
	a, b := -0.5, -0.4
	tol := eps
	x0, err := root.Brent(f, a, b, tol)
	if err == nil {
		t.Fatalf("Brent(%f, %f, %f) should error", a, b, tol)
	}

	// Solve successfully when signs are different.
	a, b = -0.5, 0.4
	x0, err = root.Brent(f, a, b, tol)
	if err != nil {
		t.Fatalf("Brent(%f, %f, %f): %s", a, b, tol, err)
	}
	const want = 0.
	if !scalar.EqualWithinRel(x0, want, tol) {
		t.Fatalf("Brent(%f, %f, %f): got %f want %f", a, b, tol, x0, want)
	}

	// Solve successfully when one side is negative zero.
	a, b = -0.5, math.Copysign(0, -1)
	x0, err = root.Brent(f, a, b, tol)
	if err != nil {
		t.Fatalf("Brent(%f, %f, %f): %s", a, b, tol, err)
	}
	if !scalar.EqualWithinRel(x0, want, tol) {
		t.Fatalf("Brent(%f, %f, %f): got %f want %f", a, b, tol, x0, want)
	}
}

type function struct {
	name string
	f    func(float64) float64
}

// Based on https://github.com/scipy/scipy/blob/main/scipy/optimize/_tstutils.py
var tstutilsFns = []function{
	// f2 is a symmetric parabola.
	{name: "f2", f: func(x float64) float64 { return math.Pow(x, 2) - 1 }},
	// f3 is a quartic with roots at 0, 1, 2, and 3.
	{name: "f3", f: func(x float64) float64 { return x * (x - 1) * (x - 2) * (x - 3) }},
	// f4 is piecewise linear, left- and right- discontinuous at x=1, the root.
	{name: "f4", f: func(x float64) float64 {
		switch {
		case x > 1:
			return 1 + 0.1*x
		case x < 1:
			return -1 + 0.1*x
		default:
			return 0
		}
	}},
	// f5 is a hyperbola with a pole at x=1, but pole replaced with 0. Not continuous at root.
	{name: "f5", f: func(x float64) float64 {
		if x != 1 {
			return 1 / (1 - x)
		}
		return 0
	}},
}

type rootTestcase struct {
	name string
	f    func(float64) float64
	a    float64
	b    float64
	root float64
}

// aps02 has poles at x=n**2, 1st and 2nd derivatives at root are also close to 0.
func aps02(x float64) float64 {
	var sum float64
	for i := 1; i < 21; i++ {
		ii := float64(i)
		sum += math.Pow(2*ii-5, 2) / math.Pow(x-ii*ii, 3)
	}
	return -2 * sum

}

// aps03 is rapidly changing at the root.
func aps03(x, a, b float64) float64 {
	return a * x * math.Exp(b*x)
}

// aps04 is a medium-degree polynomial.
func aps04(x, n, a float64) float64 {
	return math.Pow(x, n) - a
}

// aps06's exponential is rapidly changing from -1 to 1 at x=0.
func aps06(x, n float64) float64 {
	return 2*x*math.Exp(-n) - 2*math.Exp(-n*x) + 1
}

// aps07 is upside down parabola with parametrizable height.
func aps07(x, n float64) float64 {
	return (1+math.Pow(1-n, 2))*x - math.Pow(1-n*x, 2)
}

// aps08 is a degree n polynomial.
func aps08(x, n float64) float64 {
	return x*x - math.Pow(1-x, n)
}

// aps09 is an upside down quartic with parametrizable height.
func aps09(x, n float64) float64 {
	return (1+math.Pow(1-n, 4))*x - math.Pow(1-n*x, 4)
}

// aps10 is an exponential plus a polynomial.
func aps10(x, n float64) float64 {
	return math.Exp(-n*x)*(x-1) + math.Pow(x, n)
}

// aps11 is a rational function with a zero at x=1/n and a pole at x=0.
func aps11(x, n float64) float64 {
	return (n*x - 1) / ((n - 1) * x)
}

// aps12 is nth root of x, with a zero at x=n.
func aps12(x, n float64) float64 {
	return math.Pow(x, 1/n) - math.Pow(n, 1/n)
}

// aps14 returns 0 for negative x-values, trigonometric+linear for x positive.
func aps14(x, n float64) float64 {
	if x <= 0 {
		return -n / 20
	}
	return n / 20 * (x/1.5 + math.Sin(x) - 1)
}

// aps15 is piecewise linear, constant outside of [0, 0.002/(1+n)].
func aps15(x, n float64) float64 {
	if x < 0 {
		return -0.859
	}
	if x > 2e-3/(1+n) {
		return math.E - 1.859
	}
	return math.Exp((n+1)*x/2*1000) - 1.859
}

// Based on Alefeld, G. E. and Potra, F. A. and Shi, Yixun,
// Algorithm 748: Enclosing Zeros of Continuous Functions",
// ACM Trans. Math. Softw. Volume 221(1995)
// doi = {10.1145/210089.210111},
var aps = []rootTestcase{
	{name: "aps.01.00", f: func(x float64) float64 { return math.Sin(x) - x/2 }, a: math.Pi / 2, b: math.Pi, root: 1.89549426703398094},
	{name: "aps.02.00", f: aps02, a: 1 + 1e-9, b: 4 - 1e-9, root: 3.02291534727305677},
	{name: "aps.02.01", f: aps02, a: 4 + 1e-9, b: 9 - 1e-9, root: 6.68375356080807848},
	{name: "aps.02.02", f: aps02, a: 9 + 1e-9, b: 16 - 1e-9, root: 1.12387016550022114e1},
	{name: "aps.02.03", f: aps02, a: 16 + 1e-9, b: 25 - 1e-9, root: 1.96760000806234103e1},
	{name: "aps.02.04", f: aps02, a: 25 + 1e-9, b: 36 - 1e-9, root: 2.98282273265047557e1},
	{name: "aps.02.05", f: aps02, a: 36 + 1e-9, b: 49 - 1e-9, root: 4.19061161952894139e1},
	{name: "aps.02.06", f: aps02, a: 49 + 1e-9, b: 64 - 1e-9, root: 5.59535958001430913e1},
	{name: "aps.02.07", f: aps02, a: 64 + 1e-9, b: 81 - 1e-9, root: 7.19856655865877997e1},
	{name: "aps.02.08", f: aps02, a: 81 + 1e-9, b: 100 - 1e-9, root: 9.00088685391666701e1},
	{name: "aps.02.09", f: aps02, a: 100 + 1e-9, b: 121 - 1e-9, root: 1.10026532748330197e2},
	{name: "aps.03.00", f: func(x float64) float64 { return aps03(x, -40, -1) }, a: -9, b: 31, root: 0},
	{name: "aps.03.01", f: func(x float64) float64 { return aps03(x, -100, -2) }, a: -9, b: 31, root: 0},
	{name: "aps.03.02", f: func(x float64) float64 { return aps03(x, -200, -3) }, a: -9, b: 31, root: 0},
	{name: "aps.04.00", f: func(x float64) float64 { return aps04(x, 4, 0.2) }, a: 0, b: 5, root: 6.68740304976422006e-1},
	{name: "aps.04.01", f: func(x float64) float64 { return aps04(x, 6, 0.2) }, a: 0, b: 5, root: 7.64724491331730039e-01},
	{name: "aps.04.02", f: func(x float64) float64 { return aps04(x, 8, 0.2) }, a: 0, b: 5, root: 8.17765433957942545e-1},
	{name: "aps.04.03", f: func(x float64) float64 { return aps04(x, 10, 0.2) }, a: 0, b: 5, root: 8.51339922520784609e-1},
	{name: "aps.04.04", f: func(x float64) float64 { return aps04(x, 12, 0.2) }, a: 0, b: 5, root: 8.74485272221167897e-1},
	{name: "aps.04.05", f: func(x float64) float64 { return aps04(x, 4, 1) }, a: 0, b: 5, root: 1},
	{name: "aps.04.06", f: func(x float64) float64 { return aps04(x, 6, 1) }, a: 0, b: 5, root: 1},
	{name: "aps.04.07", f: func(x float64) float64 { return aps04(x, 8, 1) }, a: 0, b: 5, root: 1},
	{name: "aps.04.08", f: func(x float64) float64 { return aps04(x, 10, 1) }, a: 0, b: 5, root: 1},
	{name: "aps.04.09", f: func(x float64) float64 { return aps04(x, 12, 1) }, a: 0, b: 5, root: 1},
	{name: "aps.04.10", f: func(x float64) float64 { return aps04(x, 8, 1) }, a: -0.95, b: 4.05, root: 1},
	{name: "aps.04.11", f: func(x float64) float64 { return aps04(x, 10, 1) }, a: -0.95, b: 4.05, root: 1},
	{name: "aps.04.12", f: func(x float64) float64 { return aps04(x, 12, 1) }, a: -0.95, b: 4.05, root: 1},
	{name: "aps.04.13", f: func(x float64) float64 { return aps04(x, 14, 1) }, a: -0.95, b: 4.05, root: 1},
	{name: "aps.05.00", f: func(x float64) float64 { return math.Sin(x) - 1./2 }, a: 0, b: 1.5, root: math.Pi / 6},
	{name: "aps.06.00", f: func(x float64) float64 { return aps06(x, 1) }, a: 0, b: 1, root: 4.22477709641236709e-1},
	{name: "aps.06.01", f: func(x float64) float64 { return aps06(x, 2) }, a: 0, b: 1, root: 3.06699410483203705e-1},
	{name: "aps.06.02", f: func(x float64) float64 { return aps06(x, 3) }, a: 0, b: 1, root: 2.23705457654662959e-1},
	{name: "aps.06.03", f: func(x float64) float64 { return aps06(x, 4) }, a: 0, b: 1, root: 1.71719147519508369e-1},
	{name: "aps.06.04", f: func(x float64) float64 { return aps06(x, 5) }, a: 0, b: 1, root: 1.38257155056824066e-1},
	{name: "aps.06.05", f: func(x float64) float64 { return aps06(x, 20) }, a: 0, b: 1, root: 3.46573590208538521e-2},
	{name: "aps.06.06", f: func(x float64) float64 { return aps06(x, 40) }, a: 0, b: 1, root: 1.73286795139986315e-2},
	{name: "aps.06.07", f: func(x float64) float64 { return aps06(x, 60) }, a: 0, b: 1, root: 1.15524530093324210e-2},
	{name: "aps.06.08", f: func(x float64) float64 { return aps06(x, 80) }, a: 0, b: 1, root: 8.66433975699931573e-3},
	{name: "aps.06.09", f: func(x float64) float64 { return aps06(x, 100) }, a: 0, b: 1, root: 6.93147180559945415e-3},
	{name: "aps.07.00", f: func(x float64) float64 { return aps07(x, 5) }, a: 0, b: 1, root: 3.84025518406218985e-2},
	{name: "aps.07.01", f: func(x float64) float64 { return aps07(x, 10) }, a: 0, b: 1, root: 9.90000999800049949e-3},
	{name: "aps.07.02", f: func(x float64) float64 { return aps07(x, 20) }, a: 0, b: 1, root: 2.49375003906201174e-3},
	{name: "aps.08.00", f: func(x float64) float64 { return aps08(x, 2) }, a: 0, b: 1, root: 0.5},
	{name: "aps.08.01", f: func(x float64) float64 { return aps08(x, 5) }, a: 0, b: 1, root: 3.45954815848242059e-1},
	{name: "aps.08.02", f: func(x float64) float64 { return aps08(x, 10) }, a: 0, b: 1, root: 2.45122333753307220e-1},
	{name: "aps.08.03", f: func(x float64) float64 { return aps08(x, 15) }, a: 0, b: 1, root: 1.95547623536565629e-1},
	{name: "aps.08.04", f: func(x float64) float64 { return aps08(x, 20) }, a: 0, b: 1, root: 1.64920957276440960e-1},
	{name: "aps.09.00", f: func(x float64) float64 { return aps09(x, 1) }, a: 0, b: 1, root: 2.75508040999484394e-1},
	{name: "aps.09.01", f: func(x float64) float64 { return aps09(x, 2) }, a: 0, b: 1, root: 1.37754020499742197e-1},
	{name: "aps.09.02", f: func(x float64) float64 { return aps09(x, 4) }, a: 0, b: 1, root: 1.03052837781564422e-2},
	{name: "aps.09.03", f: func(x float64) float64 { return aps09(x, 5) }, a: 0, b: 1, root: 3.61710817890406339e-3},
	{name: "aps.09.04", f: func(x float64) float64 { return aps09(x, 8) }, a: 0, b: 1, root: 4.10872918496395375e-4},
	{name: "aps.09.05", f: func(x float64) float64 { return aps09(x, 15) }, a: 0, b: 1, root: 2.59895758929076292e-5},
	{name: "aps.09.06", f: func(x float64) float64 { return aps09(x, 20) }, a: 0, b: 1, root: 7.66859512218533719e-6},
	{name: "aps.10.00", f: func(x float64) float64 { return aps10(x, 1) }, a: 0, b: 1, root: 4.01058137541547011e-1},
	{name: "aps.10.01", f: func(x float64) float64 { return aps10(x, 5) }, a: 0, b: 1, root: 5.16153518757933583e-1},
	{name: "aps.10.02", f: func(x float64) float64 { return aps10(x, 10) }, a: 0, b: 1, root: 5.39522226908415781e-1},
	{name: "aps.10.03", f: func(x float64) float64 { return aps10(x, 15) }, a: 0, b: 1, root: 5.48182294340655241e-1},
	{name: "aps.10.04", f: func(x float64) float64 { return aps10(x, 20) }, a: 0, b: 1, root: 5.52704666678487833e-1},
	{name: "aps.11.00", f: func(x float64) float64 { return aps11(x, 2) }, a: 0.01, b: 1, root: 1. / 2},
	{name: "aps.11.01", f: func(x float64) float64 { return aps11(x, 5) }, a: 0.01, b: 1, root: 1. / 5},
	{name: "aps.11.02", f: func(x float64) float64 { return aps11(x, 15) }, a: 0.01, b: 1, root: 1. / 15},
	{name: "aps.11.03", f: func(x float64) float64 { return aps11(x, 20) }, a: 0.01, b: 1, root: 1. / 20},
	{name: "aps.12.00", f: func(x float64) float64 { return aps12(x, 2) }, a: 1, b: 100, root: 2},
	{name: "aps.12.01", f: func(x float64) float64 { return aps12(x, 3) }, a: 1, b: 100, root: 3},
	{name: "aps.12.02", f: func(x float64) float64 { return aps12(x, 4) }, a: 1, b: 100, root: 4},
	{name: "aps.12.03", f: func(x float64) float64 { return aps12(x, 5) }, a: 1, b: 100, root: 5},
	{name: "aps.12.04", f: func(x float64) float64 { return aps12(x, 6) }, a: 1, b: 100, root: 6},
	{name: "aps.12.05", f: func(x float64) float64 { return aps12(x, 7) }, a: 1, b: 100, root: 7},
	{name: "aps.12.06", f: func(x float64) float64 { return aps12(x, 9) }, a: 1, b: 100, root: 9},
	{name: "aps.12.07", f: func(x float64) float64 { return aps12(x, 11) }, a: 1, b: 100, root: 11},
	{name: "aps.12.08", f: func(x float64) float64 { return aps12(x, 13) }, a: 1, b: 100, root: 13},
	{name: "aps.12.09", f: func(x float64) float64 { return aps12(x, 15) }, a: 1, b: 100, root: 15},
	{name: "aps.12.10", f: func(x float64) float64 { return aps12(x, 17) }, a: 1, b: 100, root: 17},
	{name: "aps.12.11", f: func(x float64) float64 { return aps12(x, 19) }, a: 1, b: 100, root: 19},
	{name: "aps.12.12", f: func(x float64) float64 { return aps12(x, 21) }, a: 1, b: 100, root: 21},
	{name: "aps.12.13", f: func(x float64) float64 { return aps12(x, 23) }, a: 1, b: 100, root: 23},
	{name: "aps.12.14", f: func(x float64) float64 { return aps12(x, 25) }, a: 1, b: 100, root: 25},
	{name: "aps.12.15", f: func(x float64) float64 { return aps12(x, 27) }, a: 1, b: 100, root: 27},
	{name: "aps.12.16", f: func(x float64) float64 { return aps12(x, 29) }, a: 1, b: 100, root: 29},
	{name: "aps.12.17", f: func(x float64) float64 { return aps12(x, 31) }, a: 1, b: 100, root: 31},
	{name: "aps.12.18", f: func(x float64) float64 { return aps12(x, 33) }, a: 1, b: 100, root: 33},
	{name: "aps.14.00", f: func(x float64) float64 { return aps14(x, 1) }, a: -1000, b: math.Pi / 2, root: 6.23806518961612433e-1},
	{name: "aps.14.01", f: func(x float64) float64 { return aps14(x, 2) }, a: -1000, b: math.Pi / 2, root: 6.23806518961612433e-1},
	{name: "aps.14.02", f: func(x float64) float64 { return aps14(x, 3) }, a: -1000, b: math.Pi / 2, root: 6.23806518961612433e-1},
	{name: "aps.14.03", f: func(x float64) float64 { return aps14(x, 4) }, a: -1000, b: math.Pi / 2, root: 6.23806518961612433e-1},
	{name: "aps.14.04", f: func(x float64) float64 { return aps14(x, 5) }, a: -1000, b: math.Pi / 2, root: 6.23806518961612433e-1},
	{name: "aps.14.05", f: func(x float64) float64 { return aps14(x, 6) }, a: -1000, b: math.Pi / 2, root: 6.23806518961612433e-1},
	{name: "aps.14.06", f: func(x float64) float64 { return aps14(x, 7) }, a: -1000, b: math.Pi, root: 6.23806518961612433e-1},
	{name: "aps.14.07", f: func(x float64) float64 { return aps14(x, 8) }, a: -1000, b: math.Pi, root: 6.23806518961612433e-1},
	{name: "aps.14.08", f: func(x float64) float64 { return aps14(x, 9) }, a: -1000, b: math.Pi, root: 6.23806518961612433e-1},
	{name: "aps.14.09", f: func(x float64) float64 { return aps14(x, 10) }, a: -1000, b: math.Pi, root: 6.23806518961612433e-1},
	{name: "aps.14.10", f: func(x float64) float64 { return aps14(x, 11) }, a: -1000, b: math.Pi, root: 6.23806518961612433e-1},
	{name: "aps.14.11", f: func(x float64) float64 { return aps14(x, 12) }, a: -1000, b: math.Pi, root: 6.23806518961612433e-1},
	{name: "aps.14.12", f: func(x float64) float64 { return aps14(x, 13) }, a: -1000, b: math.Pi, root: 6.23806518961612433e-1},
	{name: "aps.14.13", f: func(x float64) float64 { return aps14(x, 14) }, a: -1000, b: math.Pi, root: 6.23806518961612433e-1},
	{name: "aps.14.14", f: func(x float64) float64 { return aps14(x, 15) }, a: -1000, b: math.Pi, root: 6.23806518961612433e-1},
	{name: "aps.14.15", f: func(x float64) float64 { return aps14(x, 16) }, a: -1000, b: math.Pi, root: 6.23806518961612433e-1},
	{name: "aps.14.16", f: func(x float64) float64 { return aps14(x, 17) }, a: -1000, b: math.Pi, root: 6.23806518961612433e-1},
	{name: "aps.14.17", f: func(x float64) float64 { return aps14(x, 18) }, a: -1000, b: math.Pi, root: 6.23806518961612433e-1},
	{name: "aps.14.18", f: func(x float64) float64 { return aps14(x, 19) }, a: -1000, b: math.Pi, root: 6.23806518961612433e-1},
	{name: "aps.14.19", f: func(x float64) float64 { return aps14(x, 20) }, a: -1000, b: math.Pi, root: 6.23806518961612433e-1},
	{name: "aps.14.20", f: func(x float64) float64 { return aps14(x, 21) }, a: -1000, b: math.Pi, root: 6.23806518961612433e-1},
	{name: "aps.14.21", f: func(x float64) float64 { return aps14(x, 22) }, a: -1000, b: math.Pi, root: 6.23806518961612433e-1},
	{name: "aps.14.22", f: func(x float64) float64 { return aps14(x, 23) }, a: -1000, b: math.Pi, root: 6.23806518961612433e-1},
	{name: "aps.14.23", f: func(x float64) float64 { return aps14(x, 24) }, a: -1000, b: math.Pi, root: 6.23806518961612433e-1},
	{name: "aps.14.24", f: func(x float64) float64 { return aps14(x, 25) }, a: -1000, b: math.Pi, root: 6.23806518961612433e-1},
	{name: "aps.14.25", f: func(x float64) float64 { return aps14(x, 26) }, a: -1000, b: math.Pi, root: 6.23806518961612433e-1},
	{name: "aps.14.26", f: func(x float64) float64 { return aps14(x, 27) }, a: -1000, b: math.Pi, root: 6.23806518961612433e-1},
	{name: "aps.14.27", f: func(x float64) float64 { return aps14(x, 28) }, a: -1000, b: math.Pi, root: 6.23806518961612433e-1},
	{name: "aps.14.28", f: func(x float64) float64 { return aps14(x, 29) }, a: -1000, b: math.Pi, root: 6.23806518961612433e-1},
	{name: "aps.14.29", f: func(x float64) float64 { return aps14(x, 30) }, a: -1000, b: math.Pi, root: 6.23806518961612433e-1},
	{name: "aps.14.30", f: func(x float64) float64 { return aps14(x, 31) }, a: -1000, b: math.Pi, root: 6.23806518961612433e-1},
	{name: "aps.14.31", f: func(x float64) float64 { return aps14(x, 32) }, a: -1000, b: math.Pi, root: 6.23806518961612433e-1},
	{name: "aps.14.32", f: func(x float64) float64 { return aps14(x, 33) }, a: -1000, b: math.Pi, root: 6.23806518961612433e-1},
	{name: "aps.14.33", f: func(x float64) float64 { return aps14(x, 34) }, a: -1000, b: math.Pi, root: 6.23806518961612433e-1},
	{name: "aps.14.34", f: func(x float64) float64 { return aps14(x, 35) }, a: -1000, b: math.Pi, root: 6.23806518961612433e-1},
	{name: "aps.14.35", f: func(x float64) float64 { return aps14(x, 36) }, a: -1000, b: math.Pi, root: 6.23806518961612433e-1},
	{name: "aps.14.36", f: func(x float64) float64 { return aps14(x, 37) }, a: -1000, b: math.Pi, root: 6.23806518961612433e-1},
	{name: "aps.14.37", f: func(x float64) float64 { return aps14(x, 38) }, a: -1000, b: math.Pi, root: 6.23806518961612433e-1},
	{name: "aps.14.38", f: func(x float64) float64 { return aps14(x, 39) }, a: -1000, b: math.Pi, root: 6.23806518961612433e-1},
	{name: "aps.14.39", f: func(x float64) float64 { return aps14(x, 40) }, a: -1000, b: math.Pi, root: 6.23806518961612433e-1},
	{name: "aps.15.00", f: func(x float64) float64 { return aps15(x, 20) }, a: -1000, b: 1e-4, root: 5.90513055942197166e-5},
	{name: "aps.15.01", f: func(x float64) float64 { return aps15(x, 21) }, a: -1000, b: 1e-4, root: 5.63671553399369967e-5},
	{name: "aps.15.02", f: func(x float64) float64 { return aps15(x, 22) }, a: -1000, b: 1e-4, root: 5.39164094555919196e-5},
	{name: "aps.15.03", f: func(x float64) float64 { return aps15(x, 23) }, a: -1000, b: 1e-4, root: 5.16698923949422470e-5},
	{name: "aps.15.04", f: func(x float64) float64 { return aps15(x, 24) }, a: -1000, b: 1e-4, root: 4.96030966991445609e-5},
	{name: "aps.15.05", f: func(x float64) float64 { return aps15(x, 25) }, a: -1000, b: 1e-4, root: 4.76952852876389951e-5},
	{name: "aps.15.06", f: func(x float64) float64 { return aps15(x, 26) }, a: -1000, b: 1e-4, root: 4.59287932399486662e-5},
	{name: "aps.15.07", f: func(x float64) float64 { return aps15(x, 27) }, a: -1000, b: 1e-4, root: 4.42884791956647841e-5},
	{name: "aps.15.08", f: func(x float64) float64 { return aps15(x, 28) }, a: -1000, b: 1e-4, root: 4.27612902578832391e-5},
	{name: "aps.15.09", f: func(x float64) float64 { return aps15(x, 29) }, a: -1000, b: 1e-4, root: 4.13359139159538030e-5},
	{name: "aps.15.10", f: func(x float64) float64 { return aps15(x, 30) }, a: -1000, b: 1e-4, root: 4.00024973380198076e-5},
	{name: "aps.15.11", f: func(x float64) float64 { return aps15(x, 31) }, a: -1000, b: 1e-4, root: 3.87524192962066869e-5},
	{name: "aps.15.12", f: func(x float64) float64 { return aps15(x, 32) }, a: -1000, b: 1e-4, root: 3.75781035599579910e-5},
	{name: "aps.15.13", f: func(x float64) float64 { return aps15(x, 33) }, a: -1000, b: 1e-4, root: 3.64728652199592355e-5},
	{name: "aps.15.14", f: func(x float64) float64 { return aps15(x, 34) }, a: -1000, b: 1e-4, root: 3.54307833565318273e-5},
	{name: "aps.15.15", f: func(x float64) float64 { return aps15(x, 35) }, a: -1000, b: 1e-4, root: 3.44465949299614980e-5},
	{name: "aps.15.16", f: func(x float64) float64 { return aps15(x, 36) }, a: -1000, b: 1e-4, root: 3.35156058778003705e-5},
	{name: "aps.15.17", f: func(x float64) float64 { return aps15(x, 37) }, a: -1000, b: 1e-4, root: 3.26336162494372125e-5},
	{name: "aps.15.18", f: func(x float64) float64 { return aps15(x, 38) }, a: -1000, b: 1e-4, root: 3.17968568584260013e-5},
	{name: "aps.15.19", f: func(x float64) float64 { return aps15(x, 39) }, a: -1000, b: 1e-4, root: 3.10019354369653455e-5},
	{name: "aps.15.20", f: func(x float64) float64 { return aps15(x, 40) }, a: -1000, b: 1e-4, root: 3.02457906702100968e-5},
	{name: "aps.15.21", f: func(x float64) float64 { return aps15(x, 100) }, a: -1000, b: 1e-4, root: 1.22779942324615231e-5},
	{name: "aps.15.22", f: func(x float64) float64 { return aps15(x, 200) }, a: -1000, b: 1e-4, root: 6.16953939044086617e-6},
	{name: "aps.15.23", f: func(x float64) float64 { return aps15(x, 300) }, a: -1000, b: 1e-4, root: 4.11985852982928163e-6},
	{name: "aps.15.24", f: func(x float64) float64 { return aps15(x, 400) }, a: -1000, b: 1e-4, root: 3.09246238772721682e-6},
	{name: "aps.15.25", f: func(x float64) float64 { return aps15(x, 500) }, a: -1000, b: 1e-4, root: 2.47520442610501789e-6},
	{name: "aps.15.26", f: func(x float64) float64 { return aps15(x, 600) }, a: -1000, b: 1e-4, root: 2.06335676785127107e-6},
	{name: "aps.15.27", f: func(x float64) float64 { return aps15(x, 700) }, a: -1000, b: 1e-4, root: 1.76901200781542651e-6},
	{name: "aps.15.28", f: func(x float64) float64 { return aps15(x, 800) }, a: -1000, b: 1e-4, root: 1.54816156988591016e-6},
	{name: "aps.15.29", f: func(x float64) float64 { return aps15(x, 900) }, a: -1000, b: 1e-4, root: 1.37633453660223511e-6},
	{name: "aps.15.30", f: func(x float64) float64 { return aps15(x, 1000) }, a: -1000, b: 1e-4, root: 1.23883857889971403e-6},
}
