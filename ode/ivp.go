// Copyright ©2021 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ode

import (
	"errors"
	"math"

	"gonum.org/v1/gonum/mat"
)

// IVP defines a multivariable, initial value problem represented by a system of ordinary differential equations.
//
// These problems have the form
//  y'(t) = f(t, y(t))
//  y(0) = y_0
//
// Where:
// t is a scalar representing the integration domain, which is time for most physical problems.
// y is the state vector.
// y' is the vector of first derivatives of the state vector y.
// f are the differential equations represented by Func.
// An initial value problem is characterized by the initial conditions imposed
// on the state vector y at the beginning of the integration domain. These
// initial conditions are returned by the IV() method for the state vector
// as y0.
//
// The term "state vector" and "state variables" are used interchangeably
// throughout the code and refer to y vector of independent variables.
type IVP struct {
	// Initial values for the state vector
	Y0 mat.Vector
	// Independent variable point at which Y0 is evaluated
	T0 float64
	// Func are the differential equations f(t,y(t)).
	// The result is y'(t) and is stored in dst.
	Func func(dst *mat.VecDense, t float64, y mat.Vector)
}

// NewModel returns a IVP given initial conditions (x0,u0), differential equations (xeq) and
// input functions for non-autonomous ODEs (ueq).
func NewIVP(t0 float64, y0 mat.Vector, f func(y *mat.VecDense, dom float64, x mat.Vector)) (IVP, error) {
	if y0 == nil || math.IsNaN(t0) || f == nil {
		return IVP{}, errors.New("bad model value")
	}
	return IVP{Func: f, Y0: y0, T0: t0}, nil
}
