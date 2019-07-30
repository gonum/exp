// Copyright ©2019 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package yaegi provides symbol lookups for Gonum packages for use
// with the yaegi interpreter.
package yaegi

//go:generate ./make_symbols.bash

import "reflect"

// Symbols variable stores the map of all gonum.org/v1/gonum symbols
// keyed by package path and symbol name.
var Symbols = map[string]map[string]reflect.Value{}

func init() {
	Symbols["gonum.org/v1/exp/yaegi"] = map[string]reflect.Value{
		"Symbols": reflect.ValueOf(Symbols),
	}
}
