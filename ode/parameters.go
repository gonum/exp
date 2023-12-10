// Copyright ©2022 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ode

import "math"

// Parameters is a generic configuration object for adaptive methods.
type Parameters struct {
	// Permissible error tolerance given an adaptive method.
	AbsTolerance float64
	// Minimum/Maximum step allowed for a single iteration.
	MinStep, MaxStep float64
}

func (p Parameters) mustValidate() {
	if math.IsInf(p.MaxStep, 0) || math.IsNaN(p.MaxStep) {
		panic("ode: max step must be finite and not NaN")
	}
	if math.IsInf(p.MinStep, 0) || math.IsNaN(p.MinStep) {
		panic("ode: min step must be finite and not NaN")
	}
	if math.IsInf(p.AbsTolerance, 0) || math.IsNaN(p.AbsTolerance) {
		panic("ode: tolerance must be finite and not NaN")
	}
	if p.AbsTolerance < 0 || p.MinStep < 0 || p.MaxStep < 0 {
		panic("ode: negative step parameters or tolerance")
	}

	isadaptive := p.AbsTolerance > 0
	if isadaptive && (p.MaxStep == 0 || p.MinStep == 0) {
		panic("ode: max and min step must be set when using adaptive methods")
	}
	if isadaptive && p.MinStep >= p.MaxStep {
		panic("ode: min step must be greater than max step")
	}
}
