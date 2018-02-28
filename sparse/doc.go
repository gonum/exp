// Copyright ©2018 The Gonum Authors. All rights reserved.
// Use of this source code is governed by BSD-style
// license that can be found in the LICENSE file.

/*
	Package sblas64 implements the v0.8 Sparse BLAS specifications for float64 as detailed
	in http://www.netlib.org/blas/blast-forum/chapter3.pdf,

	Level 1 specifications describe the sparse Vector-Vector operations. Sparse vectors
	are represented by two conventional vectors, one holding the nonzero values and the other
	one their indices. Given a dense vector X, its sparse representation is defined by:

	- nz (int) the number of non-zero entries of X;
	- x ([]float64) a slice that contains the non-zero entries of X;
	- index ([]int) as slice that contains he indices of the non-zero entries of X, of X.
*/
package sblas64
