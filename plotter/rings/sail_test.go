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
)

func TestSail(t *testing.T) {
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
		ends     []Feature
		segments int
		twist    Twist
	}{
		{
			orient: []Orientation{NotOriented, NotOriented, NotOriented},
			ends: []Feature{
				&fs{
					start:    b.Set[0].Start(),
					end:      b.Set[0].Start() + lengthOf(b.Set[0])/8,
					orient:   NotOriented,
					location: b.Set[0],
					style:    redSty,
				},
				&fs{
					start:    b.Set[1].Start() + lengthOf(b.Set[0])/8,
					end:      b.Set[1].Start() + lengthOf(b.Set[0])/4,
					orient:   NotOriented,
					location: b.Set[1],
					style:    redSty,
				},
				&fs{
					start:    b.Set[2].Start() + 2*lengthOf(b.Set[2])/5,
					end:      b.Set[2].End() - 2*lengthOf(b.Set[2])/5,
					orient:   Backward,
					location: b.Set[2],
					style:    blueSty,
				},
			},
			segments: 5,
			twist:    Individual | Flat,
		},
		{
			orient: []Orientation{Backward, NotOriented, NotOriented},
			ends: []Feature{
				&fs{
					start:    b.Set[0].Start(),
					end:      b.Set[0].Start() + lengthOf(b.Set[0])/8,
					orient:   NotOriented,
					location: b.Set[0],
					style:    redSty,
				},
				&fs{
					start:    b.Set[1].Start() + lengthOf(b.Set[0])/8,
					end:      b.Set[1].Start() + lengthOf(b.Set[0])/4,
					orient:   NotOriented,
					location: b.Set[1],
					style:    redSty,
				},
				&fs{
					start:    b.Set[2].Start() + 2*lengthOf(b.Set[2])/5,
					end:      b.Set[2].End() - 2*lengthOf(b.Set[2])/5,
					orient:   Backward,
					location: b.Set[2],
					style:    blueSty,
				},
			},
			segments: 5,
			twist:    Individual | Flat,
		},
		{
			orient: []Orientation{NotOriented, NotOriented, NotOriented},
			ends: []Feature{
				&fs{
					start:    b.Set[0].Start(),
					end:      b.Set[0].Start() + lengthOf(b.Set[0])/8,
					orient:   NotOriented,
					location: b.Set[0],
					style:    redSty,
				},
				&fs{
					start:    b.Set[1].Start(),
					end:      b.Set[1].Start() + lengthOf(b.Set[1])/8,
					orient:   NotOriented,
					location: b.Set[1],
					style:    redSty,
				},
				&fs{
					start:    b.Set[0].Start() + 2*lengthOf(b.Set[0])/5,
					end:      b.Set[0].End() - 2*lengthOf(b.Set[0])/5,
					orient:   NotOriented,
					location: b.Set[0],
					style:    blueSty,
				},
			},
			segments: 5,
			twist:    Individual | Flat,
		},
		{
			orient: []Orientation{NotOriented, NotOriented, NotOriented},
			ends: []Feature{
				&fs{
					start:    b.Set[0].Start(),
					end:      b.Set[0].Start() + lengthOf(b.Set[0])/8,
					orient:   NotOriented,
					location: b.Set[0],
					style:    redSty,
				},
				&fs{
					start:    b.Set[1].Start(),
					end:      b.Set[1].Start() + lengthOf(b.Set[1])/8,
					orient:   NotOriented,
					location: b.Set[1],
					style:    redSty,
				},
				&fs{
					start:    b.Set[0].Start() + 2*lengthOf(b.Set[0])/5,
					end:      b.Set[0].End() - 2*lengthOf(b.Set[0])/5,
					orient:   NotOriented,
					location: b.Set[0],
					style:    blueSty,
				},
				&fs{
					start:    b.Set[0].Start() + lengthOf(b.Set[0])/5,
					end:      b.Set[0].Start() + 2*lengthOf(b.Set[0])/7,
					orient:   Backward,
					location: b.Set[0],
					style:    blueSty,
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

		r, err := NewSail(test.ends, b, 70)
		if err != nil {
			t.Fatalf("unexpected error for NewSail: %v", err)
		}
		r.Bezier = &Bezier{Segments: test.segments}
		r.Twist = test.twist
		r.LineStyle = plotter.DefaultLineStyle
		r.Color = color.RGBA{R: 0xc4, G: 0x18, B: 0x80, A: 0xff}

		p.Add(r)
		p.HideAxes()
		p.Add(b)

		checkImage(t, fmt.Sprintf("sail-%d", i), p, *allPics)
	}
}
