// Code generated by 'github.com/containous/yaegi/extract gonum.org/v1/gonum/integrate/testquad'. DO NOT EDIT.

// Copyright ©2019 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build go1.15 && !go1.16
// +build go1.15,!go1.16

package yaegi

import (
	"reflect"

	"gonum.org/v1/gonum/integrate/testquad"
)

func init() {
	Symbols["gonum.org/v1/gonum/integrate/testquad"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"Constant":       reflect.ValueOf(testquad.Constant),
		"ExpOverX2Plus1": reflect.ValueOf(testquad.ExpOverX2Plus1),
		"Poly":           reflect.ValueOf(testquad.Poly),
		"Sin":            reflect.ValueOf(testquad.Sin),
		"Sqrt":           reflect.ValueOf(testquad.Sqrt),
		"XExpMinusX":     reflect.ValueOf(testquad.XExpMinusX),

		// type definitions
		"Integral": reflect.ValueOf((*testquad.Integral)(nil)),
	}
}
