// Copyright ©2016 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package linsolve_test

import (
	"fmt"
	"math"

	"gonum.org/v1/exp/linsolve"
)

// L2Projection returns a linear system whose solution is the L2 projection of u
// into the space of piecewise linear functions defined on a uniform grid on the
// interval (x0,x1). The number of cells in the grid is determined by n, and the
// dimension of the problem will be n+1.
//
// x1 must be greater than x0, and n must be a positive integer.
func L2Projection(x0, x1 float64, n int, u func(float64) float64) linsolve.System {
	h := (x1 - x0) / float64(n)

	matvec := func(dst, x []float64, trans bool) {
		h := h
		dst[0] = h / 3 * (x[0] + x[1]/2)
		for i := 1; i < n; i++ {
			dst[i] = h / 3 * (x[i-1]/2 + 2*x[i] + x[i+1]/2)
		}
		dst[n] = h / 3 * (x[n-1]/2 + x[n])
	}

	b := make([]float64, n+1)
	b[0] = u(x0) * h / 2
	for i := 1; i < n; i++ {
		b[i] = u(x0+float64(i)*h) * h
	}
	b[n] = u(x1) * h / 2

	return linsolve.System{
		MatVec: matvec,
		B:      b,
	}
}

func ExampleCG() {
	sys := L2Projection(0, 1, 10, func(x float64) float64 {
		return x * math.Sin(x)
	})
	res, err := linsolve.Iterative(sys, nil, &linsolve.CG{}, linsolve.Settings{})
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("# iterations: %v\n", res.Stats.Iterations)
		fmt.Printf("Final residual: %.6e\n", res.ResidualNorm)
		fmt.Printf("Solution: %.6f\n", res.X)
	}

	// Output:
	// # iterations: 10
	// Final residual: 6.495861e-08
	// Solution: [-0.003341 0.006678 0.036530 0.085606 0.152981 0.237072 0.337006 0.447616 0.578244 0.682719 0.920847]
}
