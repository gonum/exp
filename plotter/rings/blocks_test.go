// Copyright ©2018 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rings

import (
	"fmt"
	"image/color"
	"math/rand/v2"
	"testing"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/font"
	"gonum.org/v1/plot/font/liberation"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
)

func TestBlocks(t *testing.T) {
	p := plot.New()
	rnd := rand.New(rand.NewPCG(1, 1))
	b, err := NewGappedBlocks(randomFeatures(rnd, 3, 100000, 1000000, false, plotter.DefaultLineStyle),
		Arc{0, Complete * Clockwise},
		80, 100, 0.01,
	)
	if err != nil {
		t.Fatalf("unexpected error for NewGappedBlocks: %v", err)
	}
	b.LineStyle = plotter.DefaultLineStyle
	b.Color = color.RGBA{R: 0xc4, G: 0x18, B: 0x80, A: 0xff}

	p.Add(b)
	p.HideAxes()

	checkImage(t, p)
}

func TestBlocksScale(t *testing.T) {
	rnd := rand.New(rand.NewPCG(1, 1))
	b, err := NewGappedBlocks(randomFeatures(rnd, 3, 100000, 1000000, false, plotter.DefaultLineStyle),
		Arc{0, Complete * Clockwise},
		80, 100, 0.01,
	)
	if err != nil {
		t.Fatalf("unexpected error for plot.New: %v", err)
	}
	cache := font.NewCache(liberation.Collection())
	fnt := cache.Lookup(font.Font{Typeface: "Liberation", Variant: "Sans"}, 5)

	for i, test := range []struct {
		feats []Feature
		grid  draw.LineStyle
		inner vg.Length
		outer vg.Length
	}{
		{
			feats: b.Set,
		},
		{
			feats: b.Set,
			grid:  plotter.DefaultGridLineStyle,
			inner: b.Inner - 5,
			outer: b.Outer + 5,
		},
	} {
		t.Run(fmt.Sprintf("scale-%d", i), func(t *testing.T) {
			p := plot.New()

			s, err := NewScale(test.feats, b, 110)
			if err != nil {
				t.Fatalf("unexpected error for NewScale: %v", err)
			}
			s.LineStyle = plotter.DefaultLineStyle
			s.Tick.Length = 3
			s.Tick.LineStyle = plotter.DefaultLineStyle
			s.Tick.Label = draw.TextStyle{Color: color.Gray16{0}, Font: fnt.Font}
			s.Grid.LineStyle = test.grid
			s.Grid.Inner = test.inner
			s.Grid.Outer = test.outer

			p.Add(s)
			p.HideAxes()
			p.Add(b)

			checkImage(t, p)
		})
	}
}
