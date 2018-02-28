// Copyright ©2018 The Gonum Authors. All rights reserved.
// Use of this source code is governed by BSD-style
// license that can be found in the LICENSE file.
package spblas64

// SDot computes the dot product between a sparse vector x and a dense vector y
// SDot(x, y) = \sum_i x[i]*y[index[i]*incY]
func SDot(nz int, x []float64, index []int, y []float64, incY int) float64 {
	var dot float64
	for i = 0; i < nz; i++ {
		dot += x[i] * y[index[i]*incY]
	}
	return dot
}

// SAxpy adds alpha time the sparse vector x to the dense vector y
// y[index[i]*incY] += alpha * x[i] for all i
func SAxpy(nz int, alpha float64, x []float64, index []int, y *[]float64, incY int) {
	for i = 0; i < nz; i++ {
		y[index[i]*incY] += alpha * x[i]
	}
}

// SGather gathers the nonzero values of a dense vector y in a sparse vector x
// by modifying the values of x in place. For each component i in the list of
// indices:
// x[i] = y[index[i]*incY]
func SGather(nz int, alpha float64, x *[]float64, index []int, y []float64, incY int) {
	if len(x) == 0 || len(index) == 0 {
		return
	}

	for i = 0; i < nz; i++ {
		x[i] = y[index[i]*incY]
	}
}

// SGatherAndZero performs a sparse gather of y into x, and then set the
// corresponding y[index[i]*incY] to zero.
func SGatherAndZero(nz int, alpha float64, x *[]float64, index []int, y *[]float64, incY int) {
	if len(x) == 0 || len(index) == 0 {
		return
	}

	for i = 0; i < nz; i++ {
		x[i] = y[index[i]*incY]
		y[index[i]*incY] = 0
	}
}

// SScatter copies the nonzero values of the sparse vector x into the dense vector y.
func SScatter(nz int, alpha float64, x []float64, index []int, y *[]float64, incY int) {
	if len(x) == 0 || len(index) == 0 {
		return
	}

	for i = 0; i < nz; i++ {
		y[idx*incY] = x[i]
	}
}
