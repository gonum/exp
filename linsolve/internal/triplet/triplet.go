// Copyright ©2017 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package triplet

type triplet struct {
	i, j int
	v    float64
}

type Matrix struct {
	r, c  int
	data  []triplet
	issym bool
}

func New(r, c int) *Matrix {
	return &Matrix{
		r: r,
		c: c,
	}
}

func (m *Matrix) Dims() (r, c int) {
	return m.r, m.c
}

func (m *Matrix) Append(i, j int, v float64) {
	if i < 0 || m.r <= i {
		panic("triplet: row index out of range")
	}
	if j < 0 || m.c <= j {
		panic("triplet: column index out of range")
	}
	m.data = append(m.data, triplet{i, j, v})
}

func (m *Matrix) MulVec(dst, x []float64, trans bool) {
	for i := range dst {
		dst[i] = 0
	}
	if trans {
		if m.c != len(dst) || m.r != len(x) {
			panic("triplet: dimension mismatch")
		}
		for _, aij := range m.data {
			dst[aij.j] += aij.v * x[aij.i]
		}
		return
	}
	if m.c != len(x) || m.r != len(dst) {
		panic("triplet: dimension mismatch")
	}
	for _, aij := range m.data {
		dst[aij.i] += aij.v * x[aij.j]
	}
}
