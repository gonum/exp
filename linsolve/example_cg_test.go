// Copyright ©2016 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package linsolve_test

import (
	"fmt"
	"math"

	"gonum.org/v1/exp/linsolve"
	"gonum.org/v1/gonum/mat"
)

type system struct {
	a *mat.SymBandDense
	b []float64
}

func (sys system) MulVecTo(dst []float64, trans bool, x []float64) {
	n := len(x)
	ax := mat.NewVecDense(n, dst)
	ax.MulVec(sys.a, mat.NewVecDense(n, x))
}

// L2Projection returns a linear system whose solution is the L2 projection of f
// into the space of piecewise linear functions defined on the given grid.
//
// References:
//  - M. Larson, F. Bengzon, The Finite Element Method: Theory,
//    Implementations, and Applications. Springer (2013), Section 1.3, also
//    available at:
//    http://www.springer.com/cda/content/document/cda_downloaddocument/9783642332869-c1.pdf
func L2Projection(grid []float64, f func(float64) float64) system {
	n := len(grid)

	// Assemble the mass matrix by iterating over all elements.
	lda := 2
	a := make([]float64, n*lda)
	for i := 0; i < n-1; i++ {
		// h is the length of the i-th element.
		h := grid[i+1] - grid[i]
		// Add contribution from the i-th element, first the two diagonal
		// elements, then the one off-diagonal element.
		a[i*lda] += h / 3
		a[(i+1)*lda] += h / 3
		a[i*lda+1] += h / 6
	}

	// Allocate the mass matrix.
	A := mat.NewSymBandDense(n, 1, a)

	// Assemble the load vector by iterating over all elements.
	b := make([]float64, n)
	for i := 0; i < n-1; i++ {
		h := grid[i+1] - grid[i]
		b[i] += f(grid[i]) * h / 2
		b[i+1] += f(grid[i+1]) * h / 2
	}

	return system{A, b}
}

// UniformGrid returns a slice of n+1 evenly spaced values between
// x0 and x1, inclusive.
func UniformGrid(x0, x1 float64, n int) []float64 {
	h := (x1 - x0) / float64(n)
	grid := make([]float64, n+1)
	for i := range grid {
		grid[i] = x0 + float64(i)*h
	}
	grid[n] = x1
	return grid
}

func ExampleIterative() {
	grid := UniformGrid(0, 1, 10)
	sys := L2Projection(grid, func(x float64) float64 {
		return x * math.Sin(x)
	})

	result, err := linsolve.Iterative(sys, sys.b, &linsolve.CG{}, nil)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("# iterations: %v\n", result.Stats.Iterations)
	fmt.Printf("Final solution: %.6f\n", result.X)

	// Output:
	// # iterations: 11
	// Final solution: [-0.003339 0.006677 0.036530 0.085606 0.152981 0.237072 0.337006 0.447616 0.578244 0.682719 0.920847]
}
