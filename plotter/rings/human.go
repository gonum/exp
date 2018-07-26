// Copyright ©2013 The bíogo Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ignore

// human.go renders a rings plot of the human hg19 karyotype. It depends on
// packages from biogo.
package main

import (
	"fmt"
	"image/color"
	"math"

	"github.com/biogo/biogo/feat"
	"github.com/biogo/biogo/feat/genome"

	// hg19 provides the karyotype band locations for the plot.
	human "github.com/biogo/biogo/feat/genome/human/hg19"

	"gonum.org/v1/exp/plotter/rings"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
)

func main() {
	p, err := plot.New()
	if err != nil {
		panic(err)
	}

	sty := plotter.DefaultLineStyle
	sty.Width /= 2

	chr := make([]rings.Feature, len(human.Chromosomes))
	for i, c := range human.Chromosomes {
		chr[i] = chromosome{c}
	}
	hs, err := rings.NewGappedBlocks(chr, rings.Arc{rings.Complete / 4 * rings.CounterClockwise, rings.Complete * rings.Clockwise}, 100, 110, 0.005)
	if err != nil {
		panic(err)
	}
	hs.LineStyle = sty
	p.Add(hs)

	bands := make([]rings.Feature, len(human.Bands))
	cens := make([]rings.Feature, 0, len(human.Chromosomes))
	for i, b := range human.Bands {
		bands[i] = colorBand{b}
	}
	b, err := rings.NewBlocks(bands, hs, 100, 110)
	if err != nil {
		panic(err)
	}
	p.Add(b)
	c, err := rings.NewBlocks(cens, hs, 100, 110)
	if err != nil {
		panic(err)
	}
	p.Add(c)

	font, err := vg.MakeFont("Helvetica", 5)
	if err != nil {
		panic(err)
	}
	lb, err := rings.NewLabels(hs, 117, rings.NameLabels(hs.Set)...)
	if err != nil {
		panic(err)
	}
	lb.TextStyle = draw.TextStyle{Color: color.Gray16{0}, Font: font}
	p.Add(lb)

	bfont, err := vg.MakeFont("Helvetica", 0.5)
	if err != nil {
		panic(err)
	}
	blb, err := rings.NewLabels(b, 111, rings.NameLabels(bands)...)
	if err != nil {
		panic(err)
	}
	blb.Placement = func(a rings.Angle) (rot rings.Angle, xalign, yalign float64) {
		return a, 0, 0.75
	}
	blb.TextStyle = draw.TextStyle{Color: color.Gray16{0}, Font: bfont}
	p.Add(blb)

	p.HideAxes()

	if err := p.Save(300, 300, "human.svg"); err != nil {
		panic(err)
	}
}

type chromosome struct {
	feat.Feature
}

func (c chromosome) Start() float64        { return float64(c.Feature.Start()) }
func (c chromosome) End() float64          { return float64(c.Feature.End()) }
func (c chromosome) Parent() rings.Feature { return nil }

type colorBand struct {
	*genome.Band
}

func (b colorBand) Start() float64        { return float64(b.Band.Start()) }
func (b colorBand) End() float64          { return float64(b.Band.End()) }
func (b colorBand) Parent() rings.Feature { return chromosome{b.Band.Location()} }

func (b colorBand) FillColor() color.Color {
	switch b.Giemsa {
	case "acen":
		return color.RGBA{R: 0xff, A: 0xff}
	case "gvar":
		return color.RGBA{R: 0xbc, G: 0xbd, B: 0xdc, A: 0xff}
	case "stalk":
		return color.Gray{0x0}
	case "gneg":
		return color.Gray{0xff}
	case "gpos25":
		return color.Gray{3 * math.MaxUint8 / 4}
	case "gpos33":
		return color.Gray{2 * math.MaxUint8 / 3}
	case "gpos50":
		return color.Gray{math.MaxUint8 / 2}
	case "gpos66":
		return color.Gray{math.MaxUint8 / 3}
	case "gpos75":
		return color.Gray{math.MaxUint8 / 4}
	case "gpos100":
		return color.Gray{0x0}
	default:
		panic(fmt.Sprintf("unexpected giemsa value: %q", b.Giemsa))
	}
}

func (b colorBand) LineStyle() draw.LineStyle {
	switch b.Giemsa {
	case "acen":
		return draw.LineStyle{Color: color.RGBA{R: 0xff, A: 0xff}}
	case "stalk":
		return draw.LineStyle{Color: color.RGBA{R: 0xff, G: 0xff, B: 0xff, A: 0xff}, Width: 0.6}
	case "gneg", "gvar", "gpos25", "gpos33", "gpos50", "gpos66", "gpos75", "gpos100":
		return draw.LineStyle{}
	default:
		panic(fmt.Sprintf("unexpected giemsa value: %q", b.Giemsa))
	}
}
