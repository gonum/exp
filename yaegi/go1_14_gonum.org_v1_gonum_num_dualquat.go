// Code generated by 'goexports gonum.org/v1/gonum/num/dualquat'. DO NOT EDIT.

// Copyright ©2019 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build go1.14 && !go1.15
// +build go1.14,!go1.15

package yaegi

import (
	"reflect"

	"gonum.org/v1/gonum/num/dualquat"
)

func init() {
	Symbols["gonum.org/v1/gonum/num/dualquat"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"Abs":      reflect.ValueOf(dualquat.Abs),
		"Add":      reflect.ValueOf(dualquat.Add),
		"Conj":     reflect.ValueOf(dualquat.Conj),
		"ConjDual": reflect.ValueOf(dualquat.ConjDual),
		"ConjQuat": reflect.ValueOf(dualquat.ConjQuat),
		"Exp":      reflect.ValueOf(dualquat.Exp),
		"Inv":      reflect.ValueOf(dualquat.Inv),
		"Log":      reflect.ValueOf(dualquat.Log),
		"Mul":      reflect.ValueOf(dualquat.Mul),
		"Pow":      reflect.ValueOf(dualquat.Pow),
		"PowReal":  reflect.ValueOf(dualquat.PowReal),
		"Scale":    reflect.ValueOf(dualquat.Scale),
		"Sqrt":     reflect.ValueOf(dualquat.Sqrt),
		"Sub":      reflect.ValueOf(dualquat.Sub),

		// type definitions
		"Number": reflect.ValueOf((*dualquat.Number)(nil)),
	}
}
