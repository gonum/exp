// Copyright ©2017 The Gonum authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package linsolve

import (
	"compress/gzip"
	"math/rand"
	"os"

	"gonum.org/v1/exp/linsolve/internal/mmarket"
	"gonum.org/v1/gonum/mat"
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
	data := make([]float64, n*n)
	for i := range data {
		data[i] = rnd.NormFloat64()
	}
	m := mat.NewDense(n, n, data)
	var a mat.SymDense
	a.SymOuterK(1, m)
	for i := 0; i < n; i++ {
		a.SetSym(i, i, a.At(i, i)+float64(n))
	}
	return testCase{
		name:  "randomSPD",
		n:     n,
		iters: 10 * n,
		mulvec: func(dst, x []float64, _ bool) {
			d := mat.NewVecDense(n, dst)
			d.MulVec(&a, mat.NewVecDense(n, x))
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
