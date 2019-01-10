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
// - https://github.com/apache/arrow
package dframe

import (
	"sync"

	"github.com/apache/arrow/go/arrow/array"
)

type Frame struct {
	mu  sync.RWMutex
	tbl array.Table
	err error
}

// Err returns the first error encountered during operations on a Frame.
func (df *Frame) Err() error {
	return df.err
}

// NumRows returns the number of rows of this Frame.
func (df *Frame) NumRows() int64 {
	return df.tbl.NumRows()
}

// NumCols returns the number of columns of this Frame.
func (df *Frame) NumCols() int64 {
	return df.tbl.NumCols()
}

// Column returns the i-th column of this Frame.
func (df *Frame) Column(i int) *array.Column {
	return df.tbl.Column(i)
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

	tx := &Tx{df: df}
	return f(tx)
}

// Tx represents a read-only or read/write transaction on a data frame.
type Tx struct {
	df  *Frame
	err error
}

func (tx *Tx) Err() error {
	return tx.err
}
