// Copyright ©2025 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package root

import "math"

// FindBracketMono finds a bracket interval [a, b] where f(a)f(b) < 0.
// f must be a monotonically increasing function.
func FindBracketMono(f func(float64) float64, guess float64) (a, b float64) {
	// Make sure initial guess has the same sign as the root.
	f0 := f(0)
	if (guess < 0 && f0 < 0) || (guess > 0 && f0 > 0) {
		guess *= -1
	}

	// r is the rate in which we adjust the interval.
	var r float64
	a, fa := guess, f(guess)
	if (a > 0) == (fa < 0) {
		r = 2
	} else {
		r = 0.5
	}

	// Expand bracket until x-axis is crossed.
	// maxiter value is based on https://github.com/boostorg/math/blob/boost-1.88.0/include/boost/math/policies/policy.hpp#L130
	const maxiter = 200
	crossed := false
	b = a * r
	fb := f(b)
	for range maxiter {
		if math.Signbit(fa) != math.Signbit(fb) || fa == 0 || fb == 0 {
			crossed = true
			break
		}
		a, fa = b, fb
		b *= r
		fb = f(b)
	}
	// If unable to cross x-axis, return the largest possible bracket.
	if !crossed {
		if r > 1 {
			return a, math.Inf(int(math.Copysign(1, b)))
		}
		return a, 0
	}
	return a, b
}
