// Copyright ©2016 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package linsolve_test

import (
	"fmt"
	"math"

	"gonum.org/v1/exp/linsolve"
	"gonum.org/v1/gonum/floats"
	"gonum.org/v1/gonum/mat"
)

type system struct {
	A *mat.SymBandDense
	b []float64
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

	// Allocate the mass matrix.
	lda := 2
	a := make([]float64, n*lda)
	// Assemble the mass matrix by iterating over all elements.
	for i := 0; i < n-1; i++ {
		// h is the length of the i-th element.
		h := grid[i+1] - grid[i]
		// Add contribution from the i-th element, first the two diagonal
		// elements, then the one off-diagonal element.
		a[i*lda] += h / 3
		a[(i+1)*lda] += h / 3
		a[i*lda+1] += h / 6
	}
	A := mat.NewSymBandDense(n, 1, a)

	// Allocate the load vector.
	b := make([]float64, n)
	// Assemble the load vector by iterating over all elements.
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

func ExampleCG() {
	const tol = 1e-6

	grid := UniformGrid(0, 1, 10)
	sys := L2Projection(grid, func(x float64) float64 {
		return x * math.Sin(x)
	})
	n := sys.A.Symmetric()

	ctx := linsolve.Context{
		X:        make([]float64, n),
		Residual: make([]float64, n),
		Src:      make([]float64, n),
		Dst:      make([]float64, n),
	}
	copy(ctx.Residual, sys.b)

	if floats.Norm(ctx.Residual, 2) < tol {
		fmt.Println("Initial estimate is sufficiently accurate")
		return
	}

	bnorm := floats.Norm(sys.b, 2)
	var (
		numiter int
		rnorms  []float64
		cg      linsolve.CG
	)
	cg.Init(n)
MainLoop:
	for {
		op, err := cg.Iterate(&ctx)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		switch op {
		case linsolve.MulVec:
			dst := mat.NewVecDense(n, ctx.Dst)
			dst.MulVec(sys.A, mat.NewVecDense(n, ctx.Src))
		case linsolve.PreconSolve:
			copy(ctx.Dst, ctx.Src)
		case linsolve.CheckResidual:
			rnorm := floats.Norm(ctx.Residual, 2) / bnorm
			rnorms = append(rnorms, rnorm)
			if rnorm < tol {
				ctx.Converged = true
			}
		case linsolve.EndIteration:
			numiter++
			if ctx.Converged {
				break MainLoop
			}
		}
	}

	fmt.Printf("# iterations: %v\n", numiter)
	fmt.Printf("Residual history: %.6g\n", rnorms)
	fmt.Printf("Final solution: %.6f\n", ctx.X)

	// Output:
	// # iterations: 10
	// Residual history: [0.136572 0.0674303 0.0249177 0.00703348 0.00186369 0.000477831 0.000122098 2.98577e-05 6.44331e-06 5.46512e-07]
	// Final solution: [-0.003341 0.006678 0.036530 0.085606 0.152981 0.237072 0.337006 0.447616 0.578244 0.682719 0.920847]
}
