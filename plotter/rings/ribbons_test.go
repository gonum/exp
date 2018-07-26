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
)

func TestRibbons(t *testing.T) {
	rand.Seed(1)
	b, err := NewGappedBlocks(randomFeatures(3, 100000, 1000000, false, plotter.DefaultLineStyle),
		Arc{0, Complete * Clockwise},
		80, 100, 0.01,
	)
	if err != nil {
		t.Fatalf("unexpected error for NewGappedBlocks: %v", err)
	}

	redSty := plotter.DefaultLineStyle
	redSty.Width *= 2
	redSty.Color = color.RGBA{R: 255, G: 0, B: 0, A: 255}
	blueSty := plotter.DefaultLineStyle
	blueSty.Width *= 2
	blueSty.Color = color.RGBA{R: 0, G: 0, B: 255, A: 255}

	for i, test := range []struct {
		orient   []Orientation
		pairs    []Pair
		segments int
		twist    Twist
	}{
		{
			orient: []Orientation{NotOriented, NotOriented, NotOriented},
			pairs: []Pair{
				fp{
					feats: [2]*fs{
						{
							start:    b.Set[0].Start(),
							end:      b.Set[0].Start() + lengthOf(b.Set[0])/5,
							orient:   Backward,
							location: b.Set[0],
							style:    redSty,
						},
						{
							start:    b.Set[1].Start(),
							end:      b.Set[1].End() - lengthOf(b.Set[1])/5,
							orient:   Backward,
							location: b.Set[1],
							style:    blueSty,
						},
					},
					sty: plotter.DefaultLineStyle,
				},
				fp{
					feats: [2]*fs{
						{
							start:    b.Set[2].Start(),
							end:      b.Set[2].Start() + lengthOf(b.Set[2])/3,
							orient:   Forward,
							location: b.Set[2],
							style:    redSty,
						},
						{
							start:    b.Set[2].End() - lengthOf(b.Set[2])/8,
							end:      b.Set[2].End(),
							orient:   Backward,
							location: b.Set[2],
							style:    blueSty,
						},
					},
					sty: plotter.DefaultLineStyle,
				},
			},
			segments: 5,
			twist:    Individual | Flat,
		},
		{
			orient: []Orientation{Forward, Backward, NotOriented},
			pairs: []Pair{
				fp{
					feats: [2]*fs{
						{
							start:    b.Set[0].Start(),
							end:      b.Set[0].Start() + lengthOf(b.Set[0])/5,
							orient:   Backward,
							location: b.Set[0],
							style:    redSty,
						},
						{
							start:    b.Set[1].Start(),
							end:      b.Set[1].End() - lengthOf(b.Set[1])/5,
							orient:   Backward,
							location: b.Set[1],
							style:    blueSty,
						},
					},
					sty: plotter.DefaultLineStyle,
				},
				fp{
					feats: [2]*fs{
						{
							start:    b.Set[2].Start(),
							end:      b.Set[2].Start() + lengthOf(b.Set[2])/3,
							orient:   Forward,
							location: b.Set[2],
							style:    redSty,
						},
						{
							start:    b.Set[2].End() - lengthOf(b.Set[2])/8,
							end:      b.Set[2].End(),
							orient:   Backward,
							location: b.Set[2],
							style:    blueSty,
						},
					},
					sty: plotter.DefaultLineStyle,
				},
			},
			segments: 5,
			twist:    Individual | Flat,
		},
		{
			orient: []Orientation{NotOriented, NotOriented, Backward},
			pairs: []Pair{
				fp{
					feats: [2]*fs{
						{
							start:    b.Set[0].Start(),
							end:      b.Set[0].Start() + lengthOf(b.Set[0])/5,
							orient:   Backward,
							location: b.Set[0],
							style:    redSty,
						},
						{
							start:    b.Set[1].Start(),
							end:      b.Set[1].End() - lengthOf(b.Set[1])/5,
							orient:   Backward,
							location: b.Set[1],
							style:    blueSty,
						},
					},
					sty: plotter.DefaultLineStyle,
				},
				fp{
					feats: [2]*fs{
						{
							start:    b.Set[2].Start(),
							end:      b.Set[2].Start() + lengthOf(b.Set[2])/3,
							orient:   Forward,
							location: b.Set[2],
							style:    redSty,
						},
						{
							start:    b.Set[2].End() - lengthOf(b.Set[2])/8,
							end:      b.Set[2].End(),
							orient:   Backward,
							location: b.Set[2],
							style:    blueSty,
						},
					},
					sty: plotter.DefaultLineStyle,
				},
			},
			segments: 5,
			twist:    Individual | Flat,
		},
		{
			orient: []Orientation{NotOriented, NotOriented, Forward},
			pairs: []Pair{
				fp{
					feats: [2]*fs{
						{
							start:    b.Set[0].Start(),
							end:      b.Set[0].Start() + lengthOf(b.Set[0])/5,
							orient:   Backward,
							location: b.Set[0],
							style:    redSty,
						},
						{
							start:    b.Set[1].Start(),
							end:      b.Set[1].End() - lengthOf(b.Set[1])/5,
							orient:   Backward,
							location: b.Set[1],
							style:    blueSty,
						},
					},
					sty: plotter.DefaultLineStyle,
				},
				fp{
					feats: [2]*fs{
						{
							start:    b.Set[2].Start(),
							end:      b.Set[2].Start() + lengthOf(b.Set[2])/3,
							orient:   Forward,
							location: b.Set[2],
							style:    redSty,
						},
						{
							start:    b.Set[2].End() - lengthOf(b.Set[2])/8,
							end:      b.Set[2].End(),
							orient:   Forward,
							location: b.Set[2],
							style:    blueSty,
						},
					},
					sty: plotter.DefaultLineStyle,
				},
			},
			segments: 5,
			twist:    Individual | Flat,
		},
	} {
		p, err := plot.New()
		if err != nil {
			t.Fatalf("unexpected error for plot.New: %v", err)
		}

		for j, o := range test.orient {
			b.Set[j].(*fs).orient = o
		}
		b.Base = NewGappedArcs(b.Base, b.Set, 0.01)

		r, err := NewRibbons(test.pairs, [2]ArcOfer{b, b}, [2]vg.Length{70, 70})
		if err != nil {
			t.Fatalf("unexpected error for NewRibbons: %v", err)
		}
		r.Bezier = &Bezier{Segments: test.segments}
		r.Twist = test.twist
		r.LineStyle = plotter.DefaultLineStyle
		r.Color = color.RGBA{R: 0xc4, G: 0x18, B: 0x80, A: 0xff}

		p.Add(r)
		p.HideAxes()
		p.Add(b)

		checkImage(t, fmt.Sprintf("ribbons-%d", i), p, *allPics)
	}
}
