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
	"gonum.org/v1/plot/vg/draw"
)

func TestLabelsBlocks(t *testing.T) {
	rnd := rand.New(rand.NewPCG(1, 1))
	b, err := NewGappedBlocks(randomFeatures(rnd, 3, 100000, 1000000, false, plotter.DefaultLineStyle),
		Arc{0, Complete * Clockwise},
		80, 100, 0.01,
	)
	if err != nil {
		t.Fatalf("unexpected error for NewGappedBlocks: %v", err)
	}
	cache := font.NewCache(liberation.Collection())
	fnt := cache.Lookup(font.Font{Typeface: "Liberation", Variant: "Sans"}, 10)

	for i, test := range []struct {
		feats     []Feature
		placement TextPlacement
	}{
		{
			feats: b.Set,
		},
		{
			feats: b.Set[1:],
		},
		{
			feats:     b.Set,
			placement: Radial,
		},
		{
			feats: b.Set,
			placement: func(a Angle) (rot Angle, xalign, yalign float64) {
				return a, 0, -0.5
			},
		},
		{
			feats:     b.Set,
			placement: Horizontal,
		},
	} {
		t.Run(fmt.Sprintf("labels-%d", i), func(t *testing.T) {
			p := plot.New()
			l, err := NewLabels(b, 110, NameLabels(test.feats)...)
			if err != nil {
				t.Fatalf("unexpected error for NewLabels: %v", err)
			}
			l.TextStyle = draw.TextStyle{Color: color.Gray16{0}, Font: fnt.Font}
			l.Placement = test.placement

			p.Add(l)
			p.HideAxes()
			p.Add(b)

			checkImage(t, p)
		})
	}
}

func TestLabelsArcs(t *testing.T) {
	a := Arc{Theta: -0.031415926535897934, Phi: -1.7009436868899361} // This is feature0 from the blocks test.
	h := NewHighlight(
		color.NRGBA{R: 243, G: 243, B: 21, A: 128},
		a,
		80, 100,
	)
	h.LineStyle = plotter.DefaultLineStyle

	cache := font.NewCache(liberation.Collection())
	fnt := cache.Lookup(font.Font{Typeface: "Liberation", Variant: "Sans"}, 10)

	for i, test := range []struct {
		arc       Arcer
		label     Label
		placement TextPlacement
	}{
		{
			arc:   a,
			label: "Label",
		},
		{
			arc:   h,
			label: "Label",
		},
		{
			arc:       h,
			label:     "Label",
			placement: Radial,
		},
		{
			arc:   h,
			label: "Label",
			placement: func(a Angle) (rot Angle, xalign, yalign float64) {
				return a, 0, -0.5
			},
		},
		{
			arc:       h,
			label:     "Label",
			placement: Horizontal,
		},
	} {
		t.Run(fmt.Sprintf("labelsarcs-%d", i), func(t *testing.T) {
			p := plot.New()
			l, err := NewLabels(test.arc, 110, test.label)
			if err != nil {
				t.Fatalf("unexpected error for NewLabels: %v", err)
			}
			l.TextStyle = draw.TextStyle{Color: color.Gray16{0}, Font: fnt.Font}
			l.Placement = test.placement

			p.Add(l)
			p.HideAxes()
			p.Add(h)

			checkImage(t, p)
		})
	}
}

func TestLabelSpokes(t *testing.T) {
	p := plot.New()
	rnd := rand.New(rand.NewPCG(1, 1))
	b, err := NewGappedBlocks(randomFeatures(rnd, 3, 100000, 1000000, false, plotter.DefaultLineStyle),
		Arc{0, Complete * Clockwise},
		80, 100, 0.01,
	)
	if err != nil {
		t.Fatalf("unexpected error for NewGappedBlocks: %v", err)
	}

	m := randomFeatures(rnd, 10, b.Set[1].Start(), b.Set[1].End(), true, plotter.DefaultLineStyle)
	for _, mf := range m {
		mf.(*fs).parent = b.Set[1]
	}
	ms, err := NewSpokes(m, b, 73, 78)
	if err != nil {
		t.Fatalf("unexpected error for NewSpokes: %v", err)
	}
	ms.LineStyle = plotter.DefaultLineStyle

	cache := font.NewCache(liberation.Collection())
	fnt := cache.Lookup(font.Font{Typeface: "Liberation", Variant: "Sans"}, 10)

	l, err := NewLabels(ms, 125, NameLabels([]Feature{m[1], m[5], m[9]})...)
	if err != nil {
		t.Fatalf("unexpected error for NewLabels: %v", err)
	}
	l.TextStyle = draw.TextStyle{Color: color.Gray16{0}, Font: fnt.Font}
	l.Placement = Radial

	p.Add(l)
	p.HideAxes()

	checkImage(t, p)
}
