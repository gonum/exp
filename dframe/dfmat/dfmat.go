// Copyright ©2019 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package dfmat provides a set of tools to easily leverage gonum/mat types
// from exp/dframe (and vice versa.)
//
// This is still a WIP package, building on the experience from:
// - https://github.com/kniren/gota
// - https://github.com/tobgu/qframe
// Ultimately, dframe should also allow for a good inter-operability with
// Apache Arrow:
// - https://godoc.org/github.com/apache/arrow/go/arrow
package dfmat // import "gonum.org/v1/exp/dframe/dfmat"

import (
	"fmt"

	"github.com/apache/arrow/go/arrow"
	"github.com/apache/arrow/go/arrow/array"
	"github.com/apache/arrow/go/arrow/memory"
	"gonum.org/v1/exp/dframe"
	"gonum.org/v1/gonum/mat"
)

type Option func(c *config)

// WithNames configures a dframe.Frame with the provided set of column names.
func WithNames(names ...string) Option {
	return func(c *config) {
		c.names = make([]string, len(names))
		copy(c.names, names)
	}
}

type config struct {
	names []string
}

func newConfig(n int) *config {
	cfg := config{names: make([]string, n)}
	for i := range cfg.names {
		cfg.names[i] = fmt.Sprintf("col-%03d", i+1)
	}
	return &cfg
}

// FromMatrix creates a new dframe.Frame from a gonum/mat.Matrix.
func FromMatrix(m mat.Matrix, opts ...Option) *dframe.Frame {
	var (
		mem = memory.NewGoAllocator()

		r, c   = m.Dims()
		arrs   = make([]array.Interface, c)
		fields = make([]arrow.Field, c)
	)

	cfg := newConfig(c)
	for _, opt := range opts {
		opt(cfg)
	}

	bld := array.NewFloat64Builder(mem)
	defer bld.Release()

	switch m := m.(type) {
	case mat.RawColViewer:
		for i := 0; i < c; i++ {
			col := m.RawColView(i)
			bld.AppendValues(col, nil)
			arrs[i] = bld.NewArray()
			fields[i] = arrow.Field{
				Name: cfg.names[i],
				Type: arrs[i].DataType(),
			}
		}
	default:
		for i := 0; i < c; i++ {
			for j := 0; j < r; j++ {
				bld.Append(m.At(j, i))
			}
			arrs[i] = bld.NewArray()
			fields[i] = arrow.Field{
				Name: cfg.names[i],
				Type: arrs[i].DataType(),
			}
		}
	}

	schema := arrow.NewSchema(fields, nil)
	df, err := dframe.FromArrays(schema, arrs)
	if err != nil {
		panic(err)
	}

	return df
}

// FromVector creates a new dframe.Frame from a gonum/mat.Vector.
func FromVector(vec mat.Vector, opts ...Option) *dframe.Frame {
	var (
		mem = memory.NewGoAllocator()

		rows   = vec.Len()
		arrs   = make([]array.Interface, 1)
		fields = make([]arrow.Field, 1)
	)

	cfg := newConfig(1)
	for _, opt := range opts {
		opt(cfg)
	}

	bld := array.NewFloat64Builder(mem)
	defer bld.Release()

	switch vec := vec.(type) {
	case mat.RawColViewer:
		col := vec.RawColView(0)
		bld.AppendValues(col, nil)
		arrs[0] = bld.NewArray()
		fields[0] = arrow.Field{
			Name: cfg.names[0],
			Type: arrs[0].DataType(),
		}
	default:
		for i := 0; i < rows; i++ {
			bld.Append(vec.AtVec(i))
		}
		arrs[0] = bld.NewArray()
		fields[0] = arrow.Field{
			Name: cfg.names[0],
			Type: arrs[0].DataType(),
		}
	}

	schema := arrow.NewSchema(fields, nil)
	df, err := dframe.FromArrays(schema, arrs)
	if err != nil {
		panic(err)
	}

	return df
}
