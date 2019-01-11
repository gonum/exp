// Copyright ©2019 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package dfmat_test

import (
	"fmt"

	"github.com/apache/arrow/go/arrow/array"

	"gonum.org/v1/exp/dframe/dfmat"
	"gonum.org/v1/gonum/mat"
)

func Example_fromMatrix() {
	m := mat.NewDense(3, 2, []float64{
		1, 2,
		3, 4,
		5, 6,
	})

	{
		df := dfmat.FromMatrix(m, dfmat.WithNames("x", "y"))
		defer df.Release()

		fmt.Printf("cols: %v\n", df.ColumnNames())

		tr := array.NewTableReader(df, -1)
		defer tr.Release()

		n := 0
		for tr.Next() {
			rec := tr.Record()
			for i, col := range rec.Columns() {
				fmt.Printf("rec[%d][%q]: %v\n", n, rec.ColumnName(i), col)
			}
			n++
		}
	}

	{
		df := dfmat.FromMatrix(m.T(), dfmat.WithNames("x", "y", "z"))
		defer df.Release()

		fmt.Printf("cols: %v\n", df.ColumnNames())

		tr := array.NewTableReader(df, -1)
		defer tr.Release()

		n := 0
		for tr.Next() {
			rec := tr.Record()
			for i, col := range rec.Columns() {
				fmt.Printf("rec[%d][%q]: %v\n", n, rec.ColumnName(i), col)
			}
			n++
		}
	}

	// Output:
	// cols: [x y]
	// rec[0]["x"]: [1 3 5]
	// rec[0]["y"]: [2 4 6]
	// cols: [x y z]
	// rec[0]["x"]: [1 2]
	// rec[0]["y"]: [3 4]
	// rec[0]["z"]: [5 6]
}

func Example_fromVector() {
	m := mat.NewVecDense(6, []float64{
		1, 2,
		3, 4,
		5, 6,
	})

	df := dfmat.FromVector(m, dfmat.WithNames("x"))
	defer df.Release()

	fmt.Printf("cols: %v\n", df.ColumnNames())

	tr := array.NewTableReader(df, -1)
	defer tr.Release()

	n := 0
	for tr.Next() {
		rec := tr.Record()
		for i, col := range rec.Columns() {
			fmt.Printf("rec[%d][%q]: %v\n", n, rec.ColumnName(i), col)
		}
		n++
	}

	// Output:
	// cols: [x]
	// rec[0]["x"]: [1 2 3 4 5 6]
}
