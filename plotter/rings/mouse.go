// Copyright ©2018 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ignore

// mouse.go renders a rings plot of the mouse mm10 karyotype. It depends on
// packages from biogo.
package main

import (
	"image/color"
	"math"

	"github.com/biogo/biogo/feat"
	"github.com/biogo/biogo/feat/genome"

	// mm10 provides the karyotype band locations for the plot.
	mouse "github.com/biogo/biogo/feat/genome/mouse/mm10"

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

	chr := make([]rings.Feature, len(mouse.Chromosomes))
	for i, c := range mouse.Chromosomes {
		chr[i] = chromosome{c}
	}
	mm, err := rings.NewGappedBlocks(chr, rings.Arc{rings.Complete / 4 * rings.CounterClockwise, rings.Complete * rings.Clockwise}, 100, 110, 0.005)
	if err != nil {
		panic(err)
	}
	mm.LineStyle = sty
	p.Add(mm)

	bands := make([]rings.Feature, len(mouse.Bands))
	cens := make([]rings.Feature, 0, len(mouse.Chromosomes))
	for i, b := range mouse.Bands {
		bands[i] = colorBand{b}
		s := b.Start()
		// This condition depends on p -> q sort order in the $karyotype.Bands variable. All standard genome package follow this.
		if b.Band[0] == 'q' && (s == 0 || mouse.Bands[i-1].Band[0] == 'p') {
			cens = append(cens, colorBand{&genome.Band{Band: "cen", Desc: "Band", StartPos: s, EndPos: s, Giemsa: "acen", Chr: b.Location()}})
		}
	}
	b, err := rings.NewBlocks(bands, mm, 100, 110)
	if err != nil {
		panic(err)
	}
	p.Add(b)
	c, err := rings.NewBlocks(cens, mm, 100, 110)
	if err != nil {
		panic(err)
	}
	p.Add(c)

	font, err := vg.MakeFont("Helvetica", 7)
	if err != nil {
		panic(err)
	}
	lb, err := rings.NewLabels(mm, 117, rings.NameLabels(mm.Set)...)
	if err != nil {
		panic(err)
	}
	lb.TextStyle = draw.TextStyle{Color: color.Gray16{0}, Font: font}
	p.Add(lb)

	p.HideAxes()

	if err := p.Save(300, 300, "mouse.svg"); err != nil {
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
		panic("unexpected giemsa value")
	}
}

func (b colorBand) LineStyle() draw.LineStyle {
	switch b.Giemsa {
	case "acen":
		return draw.LineStyle{Color: color.RGBA{R: 0xff, A: 0xff}, Width: 1}
	case "gneg", "gpos25", "gpos33", "gpos50", "gpos66", "gpos75", "gpos100":
		return draw.LineStyle{}
	default:
		panic("unexpected giemsa value")
	}
}
