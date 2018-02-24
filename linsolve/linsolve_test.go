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

func (tc testCase) Order() int {
	return tc.n
}

func (tc testCase) MulVec(dst, x []float64, trans bool) {
	tc.mulvec(dst, x, trans)
}

func spdTestCases(rnd *rand.Rand) []testCase {
	return []testCase{
		randomSPD(1, rnd),
		randomSPD(2, rnd),
		randomSPD(3, rnd),
		randomSPD(4, rnd),
		randomSPD(5, rnd),
		randomSPD(10, rnd),
		randomSPD(20, rnd),
		randomSPD(50, rnd),
		randomSPD(100, rnd),
		randomSPD(200, rnd),
		randomSPD(500, rnd),
		market("spd_100_nos4", 1e-11),
		market("spd_138_bcsstm22", 1e-11),
		market("spd_237_nos1", 1e-9),
		market("spd_468_nos5", 1e-9),
		market("spd_485_bcsstm20", 1e-8),
		market("spd_900_gr_30_30", 1e-11),
	}
}

func unsymTestCases() []testCase {
	return []testCase{
		market("gen_236_e05r0100", 1e-9),
		market("gen_236_e05r0500", 1e-11),
		market("gen_434_hor__131", 1e-8),
		market("gen_886_orsirr_2", 1e-9),
	}
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
		iters: 40 * n,
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
		iters:  40 * n,
		tol:    tol,
		mulvec: m.MulVec,
	}
}
