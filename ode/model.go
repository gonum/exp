// Copyright ©2021 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ode

import (
	"errors"
	"math"

	"gonum.org/v1/gonum/mat"
)

// Model implements IVP interface. X vector and equations can not be nil or zero length.
type Model struct {
	x0  mat.Vector
	t0  float64
	xeq func(dst *mat.VecDense, dom float64, x mat.Vector)
}

// IV returns initial values of the IVP. First returned parameter is the
// starting x vector and second parameter are inputs when solving non-autonomous
// ODEs.
func (m *Model) IV() (t0 float64, x0 mat.Vector) { return m.t0, m.x0 }

// Equations returns differential equations relating to state vector x and input functions
// for non-autonomous ODEs.
//
// Input functions may be nil (ueq).
func (m *Model) Func(y *mat.VecDense, dom float64, x mat.Vector) {
	m.xeq(y, dom, x)
}

// NewModel returns a IVP given initial conditions (x0,u0), differential equations (xeq) and
// input functions for non-autonomous ODEs (ueq).
func NewModel(t0 float64, x0 mat.Vector, xeq func(y *mat.VecDense, dom float64, x mat.Vector)) (*Model, error) {
	if x0 == nil || math.IsNaN(t0) || xeq == nil {
		return nil, errors.New("bad model value")
	}
	return &Model{xeq: xeq, x0: x0, t0: t0}, nil
}
