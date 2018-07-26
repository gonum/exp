// Copyright ©2013 The bíogo Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rings

import (
	"fmt"
	"image/color"
	"math/rand"
	"testing"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
)

func TestBlocks(t *testing.T) {
	p, err := plot.New()
	if err != nil {
		t.Fatalf("unexpected error for plot.New: %v", err)
	}

	rand.Seed(1)
	b, err := NewGappedBlocks(randomFeatures(3, 100000, 1000000, false, plotter.DefaultLineStyle),
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

	checkImage(t, "blocks", p, *allPics)
}

func TestBlocksScale(t *testing.T) {
	rand.Seed(1)
	b, err := NewGappedBlocks(randomFeatures(3, 100000, 1000000, false, plotter.DefaultLineStyle),
		Arc{0, Complete * Clockwise},
		80, 100, 0.01,
	)
	if err != nil {
		t.Fatalf("unexpected error for plot.New: %v", err)
	}
	font, err := vg.MakeFont("Helvetica", 5)
	if err != nil {
		t.Fatalf("unexpected error for vg.MakeFont: %v", err)
	}

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
		p, err := plot.New()
		if err != nil {
			t.Fatalf("unexpected error for plot.New: %v", err)
		}

		s, err := NewScale(test.feats, b, 110)
		if err != nil {
			t.Fatalf("unexpected error for NewScale: %v", err)
		}
		s.LineStyle = plotter.DefaultLineStyle
		s.Tick.Length = 3
		s.Tick.LineStyle = plotter.DefaultLineStyle
		s.Tick.Label = draw.TextStyle{Color: color.Gray16{0}, Font: font}
		s.Grid.LineStyle = test.grid
		s.Grid.Inner = test.inner
		s.Grid.Outer = test.outer

		p.Add(s)
		p.HideAxes()
		p.Add(b)

		checkImage(t, fmt.Sprintf("scale-%d", i), p, *allPics)
	}
}
