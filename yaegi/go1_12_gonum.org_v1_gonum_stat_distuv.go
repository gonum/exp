// Code generated by 'goexports gonum.org/v1/gonum/stat/distuv'. DO NOT EDIT.

// Copyright ©2019 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build go1.12,!go1.13

package yaegi

import (
	"reflect"

	"gonum.org/v1/gonum/stat/distuv"
)

func init() {
	Symbols["gonum.org/v1/gonum/stat/distuv"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"NewCategorical": reflect.ValueOf(distuv.NewCategorical),
		"NewTriangle":    reflect.ValueOf(distuv.NewTriangle),
		"UnitNormal":     reflect.ValueOf(&distuv.UnitNormal).Elem(),
		"UnitUniform":    reflect.ValueOf(&distuv.UnitUniform).Elem(),

		// type definitions
		"Bernoulli":       reflect.ValueOf((*distuv.Bernoulli)(nil)),
		"Beta":            reflect.ValueOf((*distuv.Beta)(nil)),
		"Bhattacharyya":   reflect.ValueOf((*distuv.Bhattacharyya)(nil)),
		"Binomial":        reflect.ValueOf((*distuv.Binomial)(nil)),
		"Categorical":     reflect.ValueOf((*distuv.Categorical)(nil)),
		"ChiSquared":      reflect.ValueOf((*distuv.ChiSquared)(nil)),
		"Exponential":     reflect.ValueOf((*distuv.Exponential)(nil)),
		"F":               reflect.ValueOf((*distuv.F)(nil)),
		"Gamma":           reflect.ValueOf((*distuv.Gamma)(nil)),
		"GumbelRight":     reflect.ValueOf((*distuv.GumbelRight)(nil)),
		"Hellinger":       reflect.ValueOf((*distuv.Hellinger)(nil)),
		"InverseGamma":    reflect.ValueOf((*distuv.InverseGamma)(nil)),
		"KullbackLeibler": reflect.ValueOf((*distuv.KullbackLeibler)(nil)),
		"Laplace":         reflect.ValueOf((*distuv.Laplace)(nil)),
		"LogNormal":       reflect.ValueOf((*distuv.LogNormal)(nil)),
		"LogProber":       reflect.ValueOf((*distuv.LogProber)(nil)),
		"Normal":          reflect.ValueOf((*distuv.Normal)(nil)),
		"Parameter":       reflect.ValueOf((*distuv.Parameter)(nil)),
		"Pareto":          reflect.ValueOf((*distuv.Pareto)(nil)),
		"Poisson":         reflect.ValueOf((*distuv.Poisson)(nil)),
		"Quantiler":       reflect.ValueOf((*distuv.Quantiler)(nil)),
		"RandLogProber":   reflect.ValueOf((*distuv.RandLogProber)(nil)),
		"Rander":          reflect.ValueOf((*distuv.Rander)(nil)),
		"StudentsT":       reflect.ValueOf((*distuv.StudentsT)(nil)),
		"Triangle":        reflect.ValueOf((*distuv.Triangle)(nil)),
		"Uniform":         reflect.ValueOf((*distuv.Uniform)(nil)),
		"Weibull":         reflect.ValueOf((*distuv.Weibull)(nil)),

		// interface wrapper definitions
		"_LogProber":     reflect.ValueOf((*_gonum_org_v1_gonum_stat_distuv_LogProber)(nil)),
		"_Quantiler":     reflect.ValueOf((*_gonum_org_v1_gonum_stat_distuv_Quantiler)(nil)),
		"_RandLogProber": reflect.ValueOf((*_gonum_org_v1_gonum_stat_distuv_RandLogProber)(nil)),
		"_Rander":        reflect.ValueOf((*_gonum_org_v1_gonum_stat_distuv_Rander)(nil)),
	}
}

// _gonum_org_v1_gonum_stat_distuv_LogProber is an interface wrapper for LogProber type
type _gonum_org_v1_gonum_stat_distuv_LogProber struct {
	WLogProb func(a0 float64) float64
}

func (W _gonum_org_v1_gonum_stat_distuv_LogProber) LogProb(a0 float64) float64 {
	return W.WLogProb(a0)
}

// _gonum_org_v1_gonum_stat_distuv_Quantiler is an interface wrapper for Quantiler type
type _gonum_org_v1_gonum_stat_distuv_Quantiler struct {
	WQuantile func(p float64) float64
}

func (W _gonum_org_v1_gonum_stat_distuv_Quantiler) Quantile(p float64) float64 {
	return W.WQuantile(p)
}

// _gonum_org_v1_gonum_stat_distuv_RandLogProber is an interface wrapper for RandLogProber type
type _gonum_org_v1_gonum_stat_distuv_RandLogProber struct {
	WLogProb func(a0 float64) float64
	WRand    func() float64
}

func (W _gonum_org_v1_gonum_stat_distuv_RandLogProber) LogProb(a0 float64) float64 {
	return W.WLogProb(a0)
}
func (W _gonum_org_v1_gonum_stat_distuv_RandLogProber) Rand() float64 { return W.WRand() }

// _gonum_org_v1_gonum_stat_distuv_Rander is an interface wrapper for Rander type
type _gonum_org_v1_gonum_stat_distuv_Rander struct {
	WRand func() float64
}

func (W _gonum_org_v1_gonum_stat_distuv_Rander) Rand() float64 { return W.WRand() }
