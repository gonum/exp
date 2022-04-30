// Copyright ©2022 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ode

// Parameters represents the most common configuration parameters
type Parameters struct {
	// Permissible tolerance given an adaptive method
	AbsTolerance float64
	// Minimum/Maximum step allowed for a single iteration
	MinStep, MaxStep float64
}

var DefaultParam = Parameters{}
