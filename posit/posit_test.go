// Copyright ©2021 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package posit

import (
	"math"
	"testing"
)

func TestRegime(t *testing.T) {
	for _, test := range []struct {
		p      Posit
		expect int
	}{
		// posit8
		{Posit{0b_100_0000}, 0},
		{Posit{0b_111_0000}, 2},
		{Posit{0b_010_0000}, -1},
		{Posit{0b_111_1000}, 3},
		{Posit{0b_000_0100}, -4},

		// posit16
		{Posit{0, 0b0010_0000}, -9},
		{Posit{0xff, 0b0010_0000}, 6},

		// posit 32
		{Posit{0, 0, 0b1000_0001, 0}, -7 - 8},
		{Posit{0, 0, 0xff, 0}, -7 - 8},
		{Posit{0xff, 0xff, 0xff, 0}, 7 + 8 + 8 - 1},
		{Posit{0xff, 0, 0xff, 0}, 7 - 1},
	} {
		got := test.p.regime()
		if got != test.expect {
			t.Errorf("posit %s expected regime %v, got %v", test.p, test.expect, got)
		}
	}
}

func TestSign(t *testing.T) {
	for _, test := range []struct {
		p      Posit
		expect bool
	}{
		{Posit{0b_1100_0000}, true},
		{Posit{0b_0111_0000}, false},
		{Posit{0b1010_0000}, true},
		{Posit{0b_0111_1000}, false},
		{Posit{0b_1000_0100}, true},
		{Posit{0, 0b0010_0000}, false},
	} {
		got := test.p.sign()
		if got != test.expect {
			t.Errorf("posit %s expected sign %t, got %t", test.p, test.expect, got)
		}
	}
}

func TestUseed(t *testing.T) {
	for _, test := range []struct {
		p Posit
	}{
		{Posit{0b1100_0000}},

		{Posit{0, 0b0010_0000}},
		{Posit{0, 0b0010_0000, 0}},
	} {
		es := test.p.es()
		expect := int(math.RoundToEven(math.Pow(2, math.Pow(2, float64(es)))))
		got := test.p.useed()
		if expect != got {
			t.Errorf("posit %s expected useed %v, got %v", test.p, expect, got)
		}
	}
}

func TestExp(t *testing.T) {
	for _, test := range []struct {
		p      Posit
		expect int
	}{
		// posit16 has es == 1
		{Posit{0b_101_0000, 0}, 1},
		{Posit{0b_100_0000, 0}, 0},
		{Posit{0b_111_1001, 0}, 0},
		{Posit{0b_111_1011, 0}, 1},
		// posit 32 has es == 3
		{Posit{0, 0, 0b10111, 0}, 3},
		{Posit{0, 0, 0b11111, 0}, 7},
	} {
		got := test.p.exp()
		expect := test.expect

		if expect != got {
			t.Errorf("posit %s expected exp %v, got %v", test.p, expect, got)
		}
	}
}
func TestFraction(t *testing.T) {
	for _, test := range []struct {
		p      Posit
		expect int
	}{
		// posit16 has es == 1
		{Posit{0b_101_0000, 0}, 1},
		{Posit{0b_100_0000, 0}, 0},
		{Posit{0b_111_1001, 0}, 0},
		{Posit{0b_111_1011, 0}, 1},
		// posit 32 has es == 3
		{Posit{0, 0, 0b10111, 0}, 3},
		{Posit{0, 0, 0b11111, 0}, 7},
	} {
		got := test.p.fraction()
		expect := test.expect

		if expect != got {
			t.Errorf("posit %s expected exp %v, got %v", test.p, expect, got)
		}
	}
}
