// Code generated by 'github.com/containous/yaegi/extract gonum.org/v1/gonum/graph/multi'. DO NOT EDIT.

// Copyright ©2019 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build go1.15 && !go1.16
// +build go1.15,!go1.16

package yaegi

import (
	"reflect"

	"gonum.org/v1/gonum/graph/multi"
)

func init() {
	Symbols["gonum.org/v1/gonum/graph/multi"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"NewDirectedGraph":           reflect.ValueOf(multi.NewDirectedGraph),
		"NewUndirectedGraph":         reflect.ValueOf(multi.NewUndirectedGraph),
		"NewWeightedDirectedGraph":   reflect.ValueOf(multi.NewWeightedDirectedGraph),
		"NewWeightedUndirectedGraph": reflect.ValueOf(multi.NewWeightedUndirectedGraph),

		// type definitions
		"DirectedGraph":           reflect.ValueOf((*multi.DirectedGraph)(nil)),
		"Edge":                    reflect.ValueOf((*multi.Edge)(nil)),
		"Line":                    reflect.ValueOf((*multi.Line)(nil)),
		"Node":                    reflect.ValueOf((*multi.Node)(nil)),
		"UndirectedGraph":         reflect.ValueOf((*multi.UndirectedGraph)(nil)),
		"WeightedDirectedGraph":   reflect.ValueOf((*multi.WeightedDirectedGraph)(nil)),
		"WeightedEdge":            reflect.ValueOf((*multi.WeightedEdge)(nil)),
		"WeightedLine":            reflect.ValueOf((*multi.WeightedLine)(nil)),
		"WeightedUndirectedGraph": reflect.ValueOf((*multi.WeightedUndirectedGraph)(nil)),
	}
}
