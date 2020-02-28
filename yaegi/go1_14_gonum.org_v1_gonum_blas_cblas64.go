// Code generated by 'goexports gonum.org/v1/gonum/blas/cblas64'. DO NOT EDIT.

// Copyright ©2019 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build go1.14,!go1.15

package yaegi

import (
	"reflect"

	"gonum.org/v1/gonum/blas/cblas64"
)

func init() {
	Symbols["gonum.org/v1/gonum/blas/cblas64"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"Asum":           reflect.ValueOf(cblas64.Asum),
		"Axpy":           reflect.ValueOf(cblas64.Axpy),
		"Copy":           reflect.ValueOf(cblas64.Copy),
		"Dotc":           reflect.ValueOf(cblas64.Dotc),
		"Dotu":           reflect.ValueOf(cblas64.Dotu),
		"Dscal":          reflect.ValueOf(cblas64.Dscal),
		"Gbmv":           reflect.ValueOf(cblas64.Gbmv),
		"Gemm":           reflect.ValueOf(cblas64.Gemm),
		"Gemv":           reflect.ValueOf(cblas64.Gemv),
		"Gerc":           reflect.ValueOf(cblas64.Gerc),
		"Geru":           reflect.ValueOf(cblas64.Geru),
		"Hbmv":           reflect.ValueOf(cblas64.Hbmv),
		"Hemm":           reflect.ValueOf(cblas64.Hemm),
		"Hemv":           reflect.ValueOf(cblas64.Hemv),
		"Her":            reflect.ValueOf(cblas64.Her),
		"Her2":           reflect.ValueOf(cblas64.Her2),
		"Her2k":          reflect.ValueOf(cblas64.Her2k),
		"Herk":           reflect.ValueOf(cblas64.Herk),
		"Hpmv":           reflect.ValueOf(cblas64.Hpmv),
		"Hpr":            reflect.ValueOf(cblas64.Hpr),
		"Hpr2":           reflect.ValueOf(cblas64.Hpr2),
		"Iamax":          reflect.ValueOf(cblas64.Iamax),
		"Implementation": reflect.ValueOf(cblas64.Implementation),
		"Nrm2":           reflect.ValueOf(cblas64.Nrm2),
		"Scal":           reflect.ValueOf(cblas64.Scal),
		"Swap":           reflect.ValueOf(cblas64.Swap),
		"Symm":           reflect.ValueOf(cblas64.Symm),
		"Syr2k":          reflect.ValueOf(cblas64.Syr2k),
		"Syrk":           reflect.ValueOf(cblas64.Syrk),
		"Tbmv":           reflect.ValueOf(cblas64.Tbmv),
		"Tbsv":           reflect.ValueOf(cblas64.Tbsv),
		"Tpmv":           reflect.ValueOf(cblas64.Tpmv),
		"Tpsv":           reflect.ValueOf(cblas64.Tpsv),
		"Trmm":           reflect.ValueOf(cblas64.Trmm),
		"Trmv":           reflect.ValueOf(cblas64.Trmv),
		"Trsm":           reflect.ValueOf(cblas64.Trsm),
		"Trsv":           reflect.ValueOf(cblas64.Trsv),
		"Use":            reflect.ValueOf(cblas64.Use),

		// type definitions
		"Band":               reflect.ValueOf((*cblas64.Band)(nil)),
		"BandCols":           reflect.ValueOf((*cblas64.BandCols)(nil)),
		"General":            reflect.ValueOf((*cblas64.General)(nil)),
		"GeneralCols":        reflect.ValueOf((*cblas64.GeneralCols)(nil)),
		"Hermitian":          reflect.ValueOf((*cblas64.Hermitian)(nil)),
		"HermitianBand":      reflect.ValueOf((*cblas64.HermitianBand)(nil)),
		"HermitianBandCols":  reflect.ValueOf((*cblas64.HermitianBandCols)(nil)),
		"HermitianCols":      reflect.ValueOf((*cblas64.HermitianCols)(nil)),
		"HermitianPacked":    reflect.ValueOf((*cblas64.HermitianPacked)(nil)),
		"Symmetric":          reflect.ValueOf((*cblas64.Symmetric)(nil)),
		"SymmetricBand":      reflect.ValueOf((*cblas64.SymmetricBand)(nil)),
		"SymmetricPacked":    reflect.ValueOf((*cblas64.SymmetricPacked)(nil)),
		"Triangular":         reflect.ValueOf((*cblas64.Triangular)(nil)),
		"TriangularBand":     reflect.ValueOf((*cblas64.TriangularBand)(nil)),
		"TriangularBandCols": reflect.ValueOf((*cblas64.TriangularBandCols)(nil)),
		"TriangularCols":     reflect.ValueOf((*cblas64.TriangularCols)(nil)),
		"TriangularPacked":   reflect.ValueOf((*cblas64.TriangularPacked)(nil)),
		"Vector":             reflect.ValueOf((*cblas64.Vector)(nil)),
	}
}
