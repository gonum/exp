// Code generated by 'goexports gonum.org/v1/gonum/dsp/fourier'. DO NOT EDIT.

// Copyright ©2019 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build go1.14 && !go1.15
// +build go1.14,!go1.15

package yaegi

import (
	"reflect"

	"gonum.org/v1/gonum/dsp/fourier"
)

func init() {
	Symbols["gonum.org/v1/gonum/dsp/fourier"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"NewCmplxFFT":       reflect.ValueOf(fourier.NewCmplxFFT),
		"NewDCT":            reflect.ValueOf(fourier.NewDCT),
		"NewDST":            reflect.ValueOf(fourier.NewDST),
		"NewFFT":            reflect.ValueOf(fourier.NewFFT),
		"NewQuarterWaveFFT": reflect.ValueOf(fourier.NewQuarterWaveFFT),

		// type definitions
		"CmplxFFT":       reflect.ValueOf((*fourier.CmplxFFT)(nil)),
		"DCT":            reflect.ValueOf((*fourier.DCT)(nil)),
		"DST":            reflect.ValueOf((*fourier.DST)(nil)),
		"FFT":            reflect.ValueOf((*fourier.FFT)(nil)),
		"QuarterWaveFFT": reflect.ValueOf((*fourier.QuarterWaveFFT)(nil)),
	}
}
