// Copyright ©2019 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package dframe_test

import (
	"fmt"
	"log"

	"github.com/apache/arrow/go/arrow"
	"github.com/apache/arrow/go/arrow/array"
	"github.com/apache/arrow/go/arrow/memory"
	"gonum.org/v1/exp/dframe"
)

func ExampleFrame_fromTable() {
	pool := memory.NewGoAllocator()

	schema := arrow.NewSchema(
		[]arrow.Field{
			{Name: "f1-i32", Type: arrow.PrimitiveTypes.Int32},
			{Name: "f2-f64", Type: arrow.PrimitiveTypes.Float64},
		},
		nil,
	)

	b := array.NewRecordBuilder(pool, schema)
	defer b.Release()

	b.Field(0).(*array.Int32Builder).AppendValues([]int32{1, 2, 3, 4, 5, 6}, nil)
	b.Field(0).(*array.Int32Builder).AppendValues([]int32{7, 8, 9, 10}, []bool{true, true, false, true})
	b.Field(1).(*array.Float64Builder).AppendValues([]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, nil)

	rec1 := b.NewRecord()
	defer rec1.Release()

	b.Field(0).(*array.Int32Builder).AppendValues([]int32{11, 12, 13, 14, 15, 16, 17, 18, 19, 20}, nil)
	b.Field(1).(*array.Float64Builder).AppendValues([]float64{11, 12, 13, 14, 15, 16, 17, 18, 19, 20}, nil)

	rec2 := b.NewRecord()
	defer rec2.Release()

	tbl := array.NewTableFromRecords(schema, []array.Record{rec1, rec2})
	defer tbl.Release()

	df, err := dframe.FromTable(tbl)
	if err != nil {
		log.Fatal(err)
	}
	defer df.Release()

	fmt.Printf("cols: %v\n", df.ColumnNames())

	err = df.Exec(func(tx *dframe.Tx) error {
		tx.Drop("f1-i32")
		tx.Copy("fx-f64", "f2-f64")
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("cols: %v\n", df.ColumnNames())

	tr := array.NewTableReader(df, 5)
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
	// cols: [f1-i32 f2-f64]
	// cols: [f2-f64 fx-f64]
	// rec[0]["f2-f64"]: [1 2 3 4 5]
	// rec[0]["fx-f64"]: [1 2 3 4 5]
	// rec[1]["f2-f64"]: [6 7 8 9 10]
	// rec[1]["fx-f64"]: [6 7 8 9 10]
	// rec[2]["f2-f64"]: [11 12 13 14 15]
	// rec[2]["fx-f64"]: [11 12 13 14 15]
	// rec[3]["f2-f64"]: [16 17 18 19 20]
	// rec[3]["fx-f64"]: [16 17 18 19 20]
}

func ExampleFrame_fromArrays() {
	pool := memory.NewGoAllocator()

	schema := arrow.NewSchema(
		[]arrow.Field{
			{Name: "f1-i32", Type: arrow.PrimitiveTypes.Int32},
			{Name: "f2-f64", Type: arrow.PrimitiveTypes.Float64},
		},
		nil,
	)

	b := array.NewRecordBuilder(pool, schema)
	defer b.Release()

	b.Field(0).(*array.Int32Builder).AppendValues([]int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, nil)
	b.Field(1).(*array.Float64Builder).AppendValues([]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, nil)

	rec := b.NewRecord()
	defer rec.Release()

	df, err := dframe.FromArrays(schema, rec.Columns())
	if err != nil {
		log.Fatal(err)
	}
	defer df.Release()

	fmt.Printf("cols: %v\n", df.ColumnNames())

	err = df.Exec(func(tx *dframe.Tx) error {
		tx.Drop("f1-i32")
		tx.Copy("fx-f64", "f2-f64")
		tx.Slice(3, 8)
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
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
	// cols: [f1-i32 f2-f64]
	// cols: [f2-f64 fx-f64]
	// rec[0]["f2-f64"]: [4 5 6 7 8]
	// rec[0]["fx-f64"]: [4 5 6 7 8]
}

func ExampleFrame_fromCols() {
	pool := memory.NewGoAllocator()

	schema := arrow.NewSchema(
		[]arrow.Field{
			{Name: "f1-i32", Type: arrow.PrimitiveTypes.Int32},
			{Name: "f2-f64", Type: arrow.PrimitiveTypes.Float64},
		},
		nil,
	)

	b := array.NewRecordBuilder(pool, schema)
	defer b.Release()

	b.Field(0).(*array.Int32Builder).AppendValues([]int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, nil)
	b.Field(1).(*array.Float64Builder).AppendValues([]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, nil)

	rec := b.NewRecord()
	defer rec.Release()

	cols := func() []array.Column {
		var cols []array.Column
		for i, field := range schema.Fields() {
			chunk := array.NewChunked(field.Type, []array.Interface{rec.Column(i)})
			defer chunk.Release()
			col := array.NewColumn(field, chunk)
			cols = append(cols, *col)
		}
		return cols
	}()

	df, err := dframe.FromCols(cols)
	if err != nil {
		log.Fatal(err)
	}
	defer df.Release()

	fmt.Printf("cols: %v\n", df.ColumnNames())

	err = df.Exec(func(tx *dframe.Tx) error {
		tx.Drop("f1-i32")
		tx.Copy("fx-f64", "f2-f64")
		tx.Slice(3, 8)
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("cols: %v\n", df.ColumnNames())

	tr := array.NewTableReader(df, 5)
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
	// cols: [f1-i32 f2-f64]
	// cols: [f2-f64 fx-f64]
	// rec[0]["f2-f64"]: [4 5 6 7 8]
	// rec[0]["fx-f64"]: [4 5 6 7 8]
}

func ExampleFrame_fromFrame() {
	pool := memory.NewGoAllocator()

	schema := arrow.NewSchema(
		[]arrow.Field{
			{Name: "f1-i32", Type: arrow.PrimitiveTypes.Int32},
			{Name: "f2-f64", Type: arrow.PrimitiveTypes.Float64},
		},
		nil,
	)

	b := array.NewRecordBuilder(pool, schema)
	defer b.Release()

	b.Field(0).(*array.Int32Builder).AppendValues([]int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, nil)
	b.Field(1).(*array.Float64Builder).AppendValues([]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, nil)

	rec := b.NewRecord()
	defer rec.Release()

	df, err := dframe.FromArrays(schema, rec.Columns())
	if err != nil {
		log.Fatal(err)
	}
	defer df.Release()

	fmt.Printf("cols: %v\n", df.ColumnNames())

	sub, err := dframe.FromFrame(df, func(tx *dframe.Tx) error {
		tx.Drop("f1-i32")
		tx.Copy("fx-f64", "f2-f64")
		tx.Slice(3, 8)
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
	defer sub.Release()

	fmt.Printf("sub:  %v\n", sub.ColumnNames())
	fmt.Printf("cols: %v\n", df.ColumnNames())

	for i, df := range []*dframe.Frame{df, sub} {
		fmt.Printf("--- frame %d ---\n", i)
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
	// cols: [f1-i32 f2-f64]
	// sub:  [f2-f64 fx-f64]
	// cols: [f1-i32 f2-f64]
	// --- frame 0 ---
	// rec[0]["f1-i32"]: [1 2 3 4 5 6 7 8 9 10]
	// rec[0]["f2-f64"]: [1 2 3 4 5 6 7 8 9 10]
	// --- frame 1 ---
	// rec[0]["f2-f64"]: [4 5 6 7 8]
	// rec[0]["fx-f64"]: [4 5 6 7 8]
}

func ExampleFrame_fromMem() {
	df, err := dframe.FromMem(dframe.Dict{
		"f1-i32": []int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		"f2-f64": []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
	})
	if err != nil {
		log.Fatal(err)
	}
	defer df.Release()

	fmt.Printf("cols: %v\n", df.ColumnNames())

	sub, err := dframe.FromFrame(df, func(tx *dframe.Tx) error {
		tx.Drop("f1-i32")
		tx.Copy("fx-f64", "f2-f64")
		tx.Slice(3, 8)
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
	defer sub.Release()

	fmt.Printf("sub:  %v\n", sub.ColumnNames())
	fmt.Printf("cols: %v\n", df.ColumnNames())

	for i, df := range []*dframe.Frame{df, sub} {
		fmt.Printf("--- frame %d ---\n", i)
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
	// cols: [f1-i32 f2-f64]
	// sub:  [f2-f64 fx-f64]
	// cols: [f1-i32 f2-f64]
	// --- frame 0 ---
	// rec[0]["f1-i32"]: [1 2 3 4 5 6 7 8 9 10]
	// rec[0]["f2-f64"]: [1 2 3 4 5 6 7 8 9 10]
	// --- frame 1 ---
	// rec[0]["f2-f64"]: [4 5 6 7 8]
	// rec[0]["fx-f64"]: [4 5 6 7 8]
}
