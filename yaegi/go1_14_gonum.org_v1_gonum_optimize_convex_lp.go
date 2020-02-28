// Code generated by 'goexports gonum.org/v1/gonum/optimize/convex/lp'. DO NOT EDIT.

// Copyright ©2019 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build go1.14,!go1.15

package yaegi

import (
	"reflect"

	"gonum.org/v1/gonum/optimize/convex/lp"
)

func init() {
	Symbols["gonum.org/v1/gonum/optimize/convex/lp"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"Convert":       reflect.ValueOf(lp.Convert),
		"ErrBland":      reflect.ValueOf(&lp.ErrBland).Elem(),
		"ErrInfeasible": reflect.ValueOf(&lp.ErrInfeasible).Elem(),
		"ErrLinSolve":   reflect.ValueOf(&lp.ErrLinSolve).Elem(),
		"ErrSingular":   reflect.ValueOf(&lp.ErrSingular).Elem(),
		"ErrUnbounded":  reflect.ValueOf(&lp.ErrUnbounded).Elem(),
		"ErrZeroColumn": reflect.ValueOf(&lp.ErrZeroColumn).Elem(),
		"ErrZeroRow":    reflect.ValueOf(&lp.ErrZeroRow).Elem(),
		"Simplex":       reflect.ValueOf(lp.Simplex),
	}
}
