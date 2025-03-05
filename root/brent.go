package root

import (
	"errors"
	"math"
)

// Brent finds the root of a function using Brent's method.
// The root to be found should lie between [a, b], and will be refined until its accuracy is tol.
// This implementation is based on page 352, Section 9.3, Vol 1: Numerical Recipes in Fortran 77 2nd Ed., William H. Press, Saul A. Teukolsky, William T. Vetterling, Brian P. Flannery.
// https://s3.amazonaws.com/nrbook.com/book_F210.html
//
// See https://en.wikipedia.org/wiki/Brent%27s_method for more details.
func Brent(f func(float64) float64, a, b, tol float64) (float64, error) {
	const itmax = 100
	var eps = math.Nextafter(1, 2) - 1

	var d, e, xm float64

	fa, fb := f(a), f(b)
	if (fa > 0 && fb > 0) || (fa < 0 && fb < 0) {
		return math.NaN(), errors.New("root must be bracketed")
	}

	c, fc := b, fb
	for iter := 0; iter < itmax; iter++ {
		if (fb > 0 && fc > 0) || (fb < 0 && fc < 0) {
			// Rename a, b, c and adjust bounding interval d.
			c, fc = a, fa
			d = b - a
			e = d
		}
		if math.Abs(fc) < math.Abs(fb) {
			a, fa = b, fb
			b, fb = c, fc
			c, fc = a, fa
		}

		// Convergence check.
		tol1 := 2*eps*math.Abs(b) + 0.5*tol
		xm = 0.5 * (c - b)
		if math.Abs(xm) < tol1 || fb == 0 {
			return b, nil
		}

		if math.Abs(e) >= tol1 && math.Abs(fa) > math.Abs(fb) {
			// Attempt inverse quadratic interpolation.
			var p, q float64
			s := fb / fa
			if a == c {
				p, q = 2*xm*s, 1-s
			} else {
				var r float64
				q, r = fa/fc, fb/fc
				p = s * (2*xm*q*(q-r) - (b-a)*(r-1))
				q = (q - 1) * (r - 1) * (s - 1)
			}

			// Check whether in bounds.
			if p > 0 {
				q = -q
			}

			p = math.Abs(p)
			if 2*p < min(3*xm*q-math.Abs(tol1*q), math.Abs(e*q)) {
				// Accept interpolation.
				e = d
				d = p / q
			} else {
				// Interpolation failed, use bisection.
				d = xm
				e = d
			}
		} else {
			// Bounds decreasing too slowly, use bisection.
			d = xm
			e = d
		}

		// Move last best guess to a.
		a, fa = b, fb
		// Evaluate new trial root.
		if math.Abs(d) > tol1 {
			b += d
		} else {
			b += sign(tol1, xm)
		}
		fb = f(b)
	}
	return b, errors.New("maximum iterations exceeded")
}

func sign(x, s float64) float64 {
	if s > 0 {
		return x
	}
	return -x
}

func min(x, y float64) float64 {
	if x < y {
		return x
	}
	return y
}
