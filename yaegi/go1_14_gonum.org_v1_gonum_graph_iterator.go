// Code generated by 'goexports gonum.org/v1/gonum/graph/iterator'. DO NOT EDIT.

// Copyright ©2019 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build go1.14,!go1.15

package yaegi

import (
	"reflect"

	"gonum.org/v1/gonum/graph/iterator"
)

func init() {
	Symbols["gonum.org/v1/gonum/graph/iterator"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"NewImplicitNodes":        reflect.ValueOf(iterator.NewImplicitNodes),
		"NewNodes":                reflect.ValueOf(iterator.NewNodes),
		"NewOrderedEdges":         reflect.ValueOf(iterator.NewOrderedEdges),
		"NewOrderedLines":         reflect.ValueOf(iterator.NewOrderedLines),
		"NewOrderedNodes":         reflect.ValueOf(iterator.NewOrderedNodes),
		"NewOrderedWeightedEdges": reflect.ValueOf(iterator.NewOrderedWeightedEdges),
		"NewOrderedWeightedLines": reflect.ValueOf(iterator.NewOrderedWeightedLines),

		// type definitions
		"ImplicitNodes":        reflect.ValueOf((*iterator.ImplicitNodes)(nil)),
		"Nodes":                reflect.ValueOf((*iterator.Nodes)(nil)),
		"OrderedEdges":         reflect.ValueOf((*iterator.OrderedEdges)(nil)),
		"OrderedLines":         reflect.ValueOf((*iterator.OrderedLines)(nil)),
		"OrderedNodes":         reflect.ValueOf((*iterator.OrderedNodes)(nil)),
		"OrderedWeightedEdges": reflect.ValueOf((*iterator.OrderedWeightedEdges)(nil)),
		"OrderedWeightedLines": reflect.ValueOf((*iterator.OrderedWeightedLines)(nil)),
	}
}
