// Copyright ©2018 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package dframe provides a set of tools to work with data-frames.
//
// This is still a WIP package, building on the experience from:
// - https://github.com/kniren/gota
// - https://github.com/tobgu/qframe
// Ultimately, dframe should also allow for a good inter-operability with
// Apache Arrow:
// - https://godoc.org/github.com/apache/arrow/go/arrow
package dframe

import (
	"fmt"
	"sync"
	"sync/atomic"

	"github.com/apache/arrow/go/arrow"
	"github.com/apache/arrow/go/arrow/array"
	"github.com/apache/arrow/go/arrow/memory"
	"github.com/pkg/errors"
)

// Frame is a Go-based data frame built on top of Apache Arrow.
type Frame struct {
	mu sync.RWMutex // serialize creation of transactions

	refs   int64 // reference count
	err    error // first error encountered
	mem    memory.Allocator
	schema *arrow.Schema

	cols []array.Column
	rows int64
}

// FromCols creates a new data frame from the provided columns.
func FromCols(cols []array.Column, opts ...Option) (*Frame, error) {
	df := &Frame{
		refs: 1,
		mem:  memory.NewGoAllocator(),
		cols: cols,
		rows: -1,
	}

	for _, opt := range opts {
		err := opt(df)
		if err != nil {
			return nil, err
		}
	}

	if df.rows < 0 {
		switch len(df.cols) {
		case 0:
			df.rows = 0
		default:
			df.rows = int64(df.cols[0].Len())
		}
	}

	if df.schema == nil {
		fields := make([]arrow.Field, len(cols))
		for i, col := range cols {
			fields[i].Name = col.Name()
			fields[i].Type = col.DataType()
		}
		df.schema = arrow.NewSchema(fields, nil)
	}

	// validate the data frame and its constituents.
	// note we retain the columns after having validated the data frame
	// in case the validation fails and panics (and would otherwise leak
	// a ref-count on the columns.)
	df.validate()

	for i := range df.cols {
		df.cols[i].Retain()
	}

	return df, nil
}

// FromTable creates a new data frame from the provided arrow table.
func FromTable(tbl array.Table, opts ...Option) (*Frame, error) {
	df := &Frame{
		refs:   1,
		mem:    memory.NewGoAllocator(),
		schema: tbl.Schema(),
		rows:   tbl.NumRows(),
	}

	for _, opt := range opts {
		err := opt(df)
		if err != nil {
			return nil, err
		}
	}

	df.cols = make([]array.Column, tbl.NumCols())
	for i := range df.cols {
		df.cols[i] = *tbl.Column(i).NewSlice(0, -1)
	}

	return df, nil
}

func (df *Frame) validate() {
	if len(df.cols) != len(df.schema.Fields()) {
		panic(errors.New("dframe: table schema mismatch"))
	}
	for i, col := range df.cols {
		if !col.Field().Equal(df.schema.Field(i)) {
			panic(fmt.Errorf("dframe: column field %q is inconsistent with schema", col.Name()))
		}

		if int64(col.Len()) < df.rows {
			panic(fmt.Errorf("dframe: column %q expected length >= %d but got length %d", col.Name(), df.rows, col.Len()))
		}
	}
}

// Option configures an aspect of a data frame.
type Option func(*Frame) error

// WithMemAllocator configures a data frame to use the provided memory allocator.
func WithMemAllocator(mem memory.Allocator) Option {
	return func(df *Frame) error {
		df.mem = mem
		return nil
	}
}

// Err returns the first error encountered during operations on a Frame.
func (df *Frame) Err() error {
	return df.err
}

// Retain increases the reference count by 1.
// Retain may be called simultaneously from multiple goroutines.
func (df *Frame) Retain() {
	atomic.AddInt64(&df.refs, 1)
}

// Release decreases the reference count by 1.
// When the reference count goes to zero, the memory is freed.
// Release may be called simultaneously from multiple goroutines.
func (df *Frame) Release() {
	if atomic.LoadInt64(&df.refs) <= 0 {
		panic("dframe: too many releases")
	}

	if atomic.AddInt64(&df.refs, -1) == 0 {
		for i := range df.cols {
			df.cols[i].Release()
		}
		df.cols = nil
	}
}

// Schema returns the schema of this Frame.
func (df *Frame) Schema() *arrow.Schema {
	return df.schema
}

// NumRows returns the number of rows of this Frame.
func (df *Frame) NumRows() int64 {
	return df.rows
}

// NumCols returns the number of columns of this Frame.
func (df *Frame) NumCols() int64 {
	return int64(len(df.cols))
}

// Column returns the i-th column of this Frame.
func (df *Frame) Column(i int) *array.Column {
	return &df.cols[i]
}

// ColumnNames returns the list of column names of this Frame.
func (df *Frame) ColumnNames() []string {
	names := make([]string, df.NumCols())
	for i := range names {
		names[i] = df.Column(i).Name()
	}
	return names
}

func (df *Frame) Exec(f func(tx *Tx) error) error {
	df.mu.Lock()
	defer df.mu.Unlock()

	if df.err != nil {
		return df.err
	}

	tx := newTx(df)
	defer tx.Close()

	err := f(tx)
	if err != nil {
		return err
	}
	if tx.Err() != nil {
		return tx.Err()
	}

	df.swap(tx.df)
	return nil
}

func (lhs *Frame) swap(rhs *Frame) {
	rhs.refs = atomic.SwapInt64(&lhs.refs, rhs.refs)
	lhs.mem, rhs.mem = rhs.mem, lhs.mem
	lhs.schema, rhs.schema = rhs.schema, lhs.schema
	lhs.rows, rhs.rows = rhs.rows, lhs.rows
	lhs.cols, rhs.cols = rhs.cols, lhs.cols
}

func (df *Frame) clone() *Frame {
	o := &Frame{
		refs:   1,
		mem:    df.mem,
		schema: df.schema,
		cols:   make([]array.Column, len(df.cols)),
		rows:   df.rows,
	}
	copy(o.cols, df.cols)
	for i := range o.cols {
		o.cols[i].Retain()
	}
	return o
}

// Tx represents a read-only or read/write transaction on a data frame.
type Tx struct {
	df  *Frame
	err error
}

func newTx(df *Frame) *Tx {
	tx := &Tx{df: df.clone()}
	return tx
}

func (tx *Tx) Close() error {
	if tx.err != nil {
		return tx.err
	}

	tx.df.Release()
	return nil
}

func (tx *Tx) Err() error {
	return tx.err
}

// Copy copies the content of the column named src to the column named dst.
//
// Copy fails if src does not exist.
// Copy fails if dst already exist.
func (tx *Tx) Copy(dst, src string) *Tx {
	if tx.err != nil {
		return tx
	}

	if tx.df.Schema().HasField(dst) {
		tx.err = errors.Errorf("dframe: column %q already exists", dst)
		return tx
	}
	if !tx.df.Schema().HasField(src) {
		tx.err = errors.Errorf("dframe: no column named %q", src)
		return tx
	}

	isrc := tx.df.Schema().FieldIndex(src)
	idst := len(tx.df.Schema().Fields())

	fields := make([]arrow.Field, len(tx.df.Schema().Fields())+1)
	copy(fields, tx.df.Schema().Fields())

	fields[idst] = fields[isrc]
	fields[idst].Name = dst

	md := tx.df.Schema().Metadata()
	tx.df.schema = arrow.NewSchema(fields, &md)

	col := array.NewColumn(fields[idst], tx.df.cols[isrc].Data())
	tx.df.cols = append(tx.df.cols, *col)
	return tx
}

func (tx *Tx) Slice(beg, end int) *Tx {
	if tx.err != nil {
		return tx
	}

	cols := make([]array.Column, tx.df.NumCols())
	for i := range cols {
		cols[i] = *tx.df.Column(i).NewSlice(int64(beg), int64(end))
	}

	for _, col := range tx.df.cols {
		col.Release()
	}

	tx.df.cols = cols
	return tx
}

func (tx *Tx) Drop(cols ...string) *Tx {
	if tx.err != nil || len(cols) == 0 {
		return tx
	}

	set := make(map[string]struct{}, len(cols))
	for _, col := range cols {
		set[col] = struct{}{}
	}

	cs := make([]array.Column, 0, len(tx.df.cols)-len(cols))
	fs := make([]arrow.Field, 0, len(tx.df.Schema().Fields())-len(cols))

	for i := range tx.df.cols {
		col := &tx.df.cols[i]
		if _, ok := set[col.Name()]; ok {
			col.Release()
			continue
		}
		cs = append(cs, *col)
		fs = append(fs, tx.df.Schema().Field(i))
	}

	md := tx.df.Schema().Metadata() // FIXME(sbinet): also remove metadata of removed cols.
	sc := arrow.NewSchema(fs, &md)

	tx.df.cols = cs
	tx.df.schema = sc
	return tx
}
