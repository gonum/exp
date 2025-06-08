package root

import "math"

// FindBracketMono finds a bracket interval [a, b] where f(a)f(b) < 0.
// f must be a monotonically increasing function.
func FindBracketMono(f func(float64) float64, guess float64) (float64, float64) {
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
		r = 1. / 2
	}

	// Expand bracket until x-axis is crossed.
	crossed := false
	b := a * r
	fb := f(b)
	for range 200 {
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
		} else {
			return a, 0
		}
	}

	return a, b
}
