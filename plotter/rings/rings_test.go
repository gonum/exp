// Copyright ©2018 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rings

import (
	"fmt"
	"image/color"
	"math/rand/v2"
	"path/filepath"
	"testing"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/cmpimg"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
)

// fs is a Feature implementation for testing.
type fs struct {
	start, end float64
	name       string
	parent     Feature
	orient     Orientation
	style      draw.LineStyle
	scores     []float64
}

func (f *fs) Start() float64            { return f.start }
func (f *fs) End() float64              { return f.end }
func (f *fs) Name() string              { return f.name }
func (f *fs) Parent() Feature           { return f.parent }
func (f *fs) Orientation() Orientation  { return f.orient }
func (f *fs) LineStyle() draw.LineStyle { return f.style }
func (f *fs) Scores() []float64         { return f.scores }

type fp struct {
	feats [2]*fs
	sty   draw.LineStyle
}

func (p fp) Features() [2]Feature { return [2]Feature{p.feats[0], p.feats[1]} }
func (p fp) LineStyle() draw.LineStyle {
	var col color.RGBA
	for _, f := range p.feats {
		r, g, b, a := f.style.Color.RGBA()
		col.R += byte(r / 2)
		col.G += byte(g / 2)
		col.B += byte(b / 2)
		col.A += byte(a / 2)
	}
	p.sty.Color = col
	return p.sty
}

func randomFeatures(rnd *rand.Rand, n int, min, max float64, single bool, sty draw.LineStyle) []Feature {
	data := make([]Feature, n)
	for i := range data {
		// IntN is used here to avoid drastic random
		// sequence changes at this stage.
		start := float64(rnd.IntN(int(max-min))) + min
		var end float64
		if !single {
			end = float64(rnd.IntN(int(max - start)))
		}
		data[i] = &fs{
			start: start,
			end:   start + end,
			name:  fmt.Sprintf("feature%v", i),
			style: sty,
		}
	}
	return data
}

func floatPtr(f float64) *float64 { return &f }

// makeScorers returns n Scorers each with m scores.
func makeScorers(f *fs, n, m int, fn func(i, j int) float64) []Scorer {
	s := make([]Scorer, n)
	for i := 0; i < n; i++ {
		cs := &fs{
			start:  f.Start() + float64(i)*(lengthOf(f)/float64(n)),
			end:    f.Start() + float64(i+1)*(lengthOf(f)/float64(n)),
			name:   fmt.Sprintf("%s#%d", f.Name(), i),
			parent: f,
			scores: make([]float64, m),
		}
		for j := range cs.scores {
			cs.scores[j] = fn(i, j)
		}
		s[i] = cs
	}
	return s
}

// checkImage compares the plot in p to the image in testdata/name_golden.png.
func checkImage(t *testing.T, p *plot.Plot) {
	name := filepath.FromSlash(t.Name())
	fct := func() {
		err := p.Save(vg.Length(300), vg.Length(300), filepath.Join("testdata", name+".png"))
		if err != nil {
			t.Fatalf("could not generate plot for %q: %v", name, err)
		}
	}

	cmpimg.CheckPlotApprox(fct, t, 0.01, name+".png")
}
