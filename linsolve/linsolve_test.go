// Copyright ©2017 The Gonum authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package linsolve

import (
	"compress/gzip"
	"math/rand"
	"os"

	"gonum.org/v1/exp/linsolve/internal/mmarket"
	"gonum.org/v1/gonum/blas"
	"gonum.org/v1/gonum/blas/blas64"
)

type testCase struct {
	name   string
	n      int
	iters  int
	tol    float64
	mulvec func(dst, x []float64, trans bool)
}

// randomSPD returns a random symmetric positive-definite matrix of order n.
func randomSPD(n int, rnd *rand.Rand) testCase {
	a := make([]float64, n*n)
	lda := n
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			a[i*lda+j] = rnd.Float64()
		}
	}
	for i := 0; i < n; i++ {
		a[i*lda+i] += float64(n)
	}
	return testCase{
		name:  "randomSPD",
		n:     n,
		iters: 2 * n,
		mulvec: func(dst, x []float64, _ bool) {
			bi := blas64.Implementation()
			bi.Dsymv(blas.Upper, n, 1, a, lda, x, 1, 0, dst, 1)
		},
	}
}

// market returns a test matrix from the Matrix Market.
func market(name string, tol float64) testCase {
	f, err := os.Open("testdata/" + name + ".mtx.gz")
	if err != nil {
		panic(err)
	}
	gz, err := gzip.NewReader(f)
	if err != nil {
		panic(err)
	}
	m, err := mmarket.NewReader(gz).Read()
	if err != nil {
		panic(err)
	}
	n, _ := m.Dims()
	return testCase{
		name:   name,
		n:      n,
		iters:  10 * n,
		tol:    tol,
		mulvec: m.MulVec,
	}
}
