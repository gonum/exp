// Code generated by 'github.com/containous/yaegi/extract gonum.org/v1/gonum/num/quat'. DO NOT EDIT.

// Copyright ©2019 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build go1.15 && !go1.16
// +build go1.15,!go1.16

package yaegi

import (
	"reflect"

	"gonum.org/v1/gonum/num/quat"
)

func init() {
	Symbols["gonum.org/v1/gonum/num/quat"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"Abs":   reflect.ValueOf(quat.Abs),
		"Acos":  reflect.ValueOf(quat.Acos),
		"Acosh": reflect.ValueOf(quat.Acosh),
		"Add":   reflect.ValueOf(quat.Add),
		"Asin":  reflect.ValueOf(quat.Asin),
		"Asinh": reflect.ValueOf(quat.Asinh),
		"Atan":  reflect.ValueOf(quat.Atan),
		"Atanh": reflect.ValueOf(quat.Atanh),
		"Conj":  reflect.ValueOf(quat.Conj),
		"Cos":   reflect.ValueOf(quat.Cos),
		"Cosh":  reflect.ValueOf(quat.Cosh),
		"Exp":   reflect.ValueOf(quat.Exp),
		"Inf":   reflect.ValueOf(quat.Inf),
		"Inv":   reflect.ValueOf(quat.Inv),
		"IsInf": reflect.ValueOf(quat.IsInf),
		"IsNaN": reflect.ValueOf(quat.IsNaN),
		"Log":   reflect.ValueOf(quat.Log),
		"Mul":   reflect.ValueOf(quat.Mul),
		"NaN":   reflect.ValueOf(quat.NaN),
		"Parse": reflect.ValueOf(quat.Parse),
		"Pow":   reflect.ValueOf(quat.Pow),
		"Scale": reflect.ValueOf(quat.Scale),
		"Sin":   reflect.ValueOf(quat.Sin),
		"Sinh":  reflect.ValueOf(quat.Sinh),
		"Sqrt":  reflect.ValueOf(quat.Sqrt),
		"Sub":   reflect.ValueOf(quat.Sub),
		"Tan":   reflect.ValueOf(quat.Tan),
		"Tanh":  reflect.ValueOf(quat.Tanh),

		// type definitions
		"Number": reflect.ValueOf((*quat.Number)(nil)),
	}
}
