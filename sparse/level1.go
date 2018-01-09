// Copyright 2018 The Gonum Authors. All rights reserved.
// Use of this source code is governed by BSD-style
// license that can be found in the LICENSE file.
//
// Functions in this file implement the Level 1 BLAS specifications
// for sparse vectors
//
// Dot: dot product between sparse vectors
// Nrm2: norm 2 of the vector
// Asum: compute the sum of amplitudes of the vector' elements
// Idmax: returns the index of the element with maximum absolute value
// Swap: swap two vectors x <-> y
// Scal: product of a vector by a scalar a*x
// Axpy: compute z = a*x + y
// Copy: copy an array y <- x
// Scatter: y|x <- x
// Gather: x <- y|x and y|x = 0

package blas64

import (
	"math"
)

// SparseVector represents a sparse vector.
type SparseVector struct {
	Size int // Number of elements in the (non sparse) vector.
	Ind  []int
	Data []float64
}

type DenseVector struct {
	Data []float64
}

// Level 1 --------------------------------------------------------------------

//  Dot computes the dot product between two sparse vectors
//  Dot(x, y) = \sum_i x[i]*y[i]
//  Will panic if vectors are not the same size
func Dot(x, y SparseVector) float64 {
	if x.Size != y.Size {
		panic("X and Y must have the same length.")
	}
	if x.Size == 0 {
		return 0
	}

	var dot float64
	ity := 0
	j := y.Ind[ity]
	for itx, i := range x.Ind {
		for j < i {
			ity += 1
			j = y.Ind[ity]
		}
		if i == j {
			dot += y.Data[itx] * y.Data[ity]
		}
	}
	return dot
}

// Nrm2 computes the L2 norm of the vector x:
// Nrm2(x) = \sum_i x[i]*x[i]
func Nrm2(x SparseVector) float64 {
	var nrm2 float64
	for _, value := range x.Data {
		nrm2 += value * value
	}
	return math.Sqrt(nrm2)
}

// Asum computes the sum of the absolute values of the elemnts of x:
// Asum(x) = \sum_i |x[i]|
func Asum(x SparseVector) float64 {
	var asum float64
	for _, value := range x.Data {
		asum += math.Abs(value)
	}
	return asum
}

// Idmax returns the index of an element of x with the largest aboslute value
// If there are multiple such indices the maximum is returned.
// Idmax returns -1 if n == 0
func Idmax(x SparseVector) int {
	if x.Size == 0 {
		return 0
	}
	imax := x.Ind[0]
	max := x.Data[0]
	for i, ind := range x.Ind {
		if x.Data[i] > max {
			imax = ind
			max = x.Data[i]
		}
	}
	return imax
}

// Swap exchanges the elements of the two vectors x and y
// x[i], y[i] = y[i], x[i] for all i
// Will panic if the vectors are not of the same length
func Swap(x, y *SparseVector) {
	if x.Size != y.Size {
		panic("input vectors are not of the same size")
	}
	x, y = y, x
}

// Scal scales the vector x by \alpha:
// x[i] *= \alpha for all i
func Scal(alpha float64, x SparseVector) SparseVector {
	if alpha == 0 {
		return SparseVector{x.Size, []int{}, []float64{}}
	}

	for i, _ := range x.Data {
		x.Data[i] *= alpha
	}
	return x
}

// Copy copies the elements of x into the elements of y
// y[i] = x[i] for all i
// Will panic if x and y are not of same length
func Copy(x, y *SparseVector) {
	if x.Size != y.Size {
		panic("X and Y must have the same length.")
	}
	y.Ind = y.Ind[:0]
	y.Data = y.Data[:0]
	for i, ind := range x.Ind {
		y.Ind = append(y.Ind, ind)
		y.Data = append(y.Data, x.Data[i])
	}
}

// Axpy add alpha time x to y
// y[i] += alpha * x[i] for all i
// Current version does not keep an ordered vec
func Axpy(x, y *SparseVector, alpha float64) {
	ity := 0
	for itx, i := range x.Ind {
		j := y.Ind[ity]
		for j < i {
			y.Ind = append(y.Ind, i)
			y.Data = append(y.Data, alpha*x.Data[itx])
			ity += 1
			j = y.Ind[ity]
		}
		if i == j {
			y.Data[ity] += alpha * x.Data[itx]
		}
		ity += 1
	}
}

// Gather takes a DenseVector as an input and
// transforms it in a sparse vector
func Gather(x DenseVector) SparseVector {
	var y SparseVector
	for i, datum := range x.Data {
		if datum != 0 {
			y.Ind = append(y.Ind, i)
			y.Data = append(x.Data, datum)
		}
	}
	return y
}

// Scatter builds a DenseVector from a SparseVector
func Scatter(y SparseVector) DenseVector {
	x := DenseVector{Data: []float64{}}
	ity := 0
	j := y.Ind[ity]
	for i := 0; i < y.Size; i++ {
		if i == j {
			x.Data = append(x.Data, y.Data[j])
			ity += 1
			j = y.Ind[ity]
		} else {
			x.Data = append(x.Data, 0)
			ity += 1
			j = y.Ind[ity]
		}
	}
	return x
}
