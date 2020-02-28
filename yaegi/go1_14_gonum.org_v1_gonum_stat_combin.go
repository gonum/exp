// Code generated by 'goexports gonum.org/v1/gonum/stat/combin'. DO NOT EDIT.

// Copyright ©2019 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build go1.14,!go1.15

package yaegi

import (
	"reflect"

	"gonum.org/v1/gonum/stat/combin"
)

func init() {
	Symbols["gonum.org/v1/gonum/stat/combin"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"Binomial":                reflect.ValueOf(combin.Binomial),
		"Card":                    reflect.ValueOf(combin.Card),
		"Cartesian":               reflect.ValueOf(combin.Cartesian),
		"CombinationIndex":        reflect.ValueOf(combin.CombinationIndex),
		"Combinations":            reflect.ValueOf(combin.Combinations),
		"GeneralizedBinomial":     reflect.ValueOf(combin.GeneralizedBinomial),
		"IdxFor":                  reflect.ValueOf(combin.IdxFor),
		"IndexToCombination":      reflect.ValueOf(combin.IndexToCombination),
		"IndexToPermutation":      reflect.ValueOf(combin.IndexToPermutation),
		"LogGeneralizedBinomial":  reflect.ValueOf(combin.LogGeneralizedBinomial),
		"NewCartesianGenerator":   reflect.ValueOf(combin.NewCartesianGenerator),
		"NewCombinationGenerator": reflect.ValueOf(combin.NewCombinationGenerator),
		"NewPermutationGenerator": reflect.ValueOf(combin.NewPermutationGenerator),
		"NumPermutations":         reflect.ValueOf(combin.NumPermutations),
		"PermutationIndex":        reflect.ValueOf(combin.PermutationIndex),
		"Permutations":            reflect.ValueOf(combin.Permutations),
		"SubFor":                  reflect.ValueOf(combin.SubFor),

		// type definitions
		"CartesianGenerator":   reflect.ValueOf((*combin.CartesianGenerator)(nil)),
		"CombinationGenerator": reflect.ValueOf((*combin.CombinationGenerator)(nil)),
		"PermutationGenerator": reflect.ValueOf((*combin.PermutationGenerator)(nil)),
	}
}
