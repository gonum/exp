// Copyright ©2013 The bíogo Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rings

import (
	"fmt"
	"image/color"
	"math/rand"
	"reflect"
	"testing"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
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
		actions  []interface{}
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
			actions: []interface{}{
				setColor{col: color.RGBA{R: 0xc4, G: 0x18, B: 0x80, A: 0xff}},
				fill{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 222.4654592256012, Y: 150.301246864531}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 70, Start: 6.2517693806436885, Angle: -0.3401785862023008},
					{Type: vg.LineComp, Pos: vg.Point{X: 191.53092479391694, Y: 135.5344727962992}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 165.1342669001769, Y: 140.55553442288462}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 138.53249553491207, Y: 142.1460743171709}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 111.72561069812248, Y: 140.306092479158}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 84.71361238980812, Y: 135.03558890884597}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 70, Start: 3.393747857626204, Angle: 1.0942459830557514},
					{Type: vg.LineComp, Pos: vg.Point{X: 145.32986989085776, Y: 108.7352361883809}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 158.08705245712767, Y: 127.57999179983727}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 175.19537820167486, Y: 140.7892454496811}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 196.65484712449938, Y: 148.36299713791232}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 222.4654592256012, Y: 150.301246864531}, Radius: 0, Start: 0, Angle: 0},
				}},
				setColor{col: color.RGBA{R: 0xff, G: 0x0, B: 0xff, A: 0xfe}},
				setWidth{w: 1},
				setLineDash{dashes: []vg.Length(nil), offsets: 0},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 222.4654592256012, Y: 150.301246864531}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.MoveComp, Pos: vg.Point{X: 217.7224692161322, Y: 127.08288943741465}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 191.53092479391694, Y: 135.5344727962992}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 165.1342669001769, Y: 140.55553442288462}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 138.53249553491207, Y: 142.1460743171709}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 111.72561069812248, Y: 140.306092479158}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 84.71361238980812, Y: 135.03558890884597}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.MoveComp, Pos: vg.Point{X: 136.92383050286517, Y: 84.25497861531197}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 145.32986989085776, Y: 108.7352361883809}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 158.08705245712767, Y: 127.57999179983727}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 175.19537820167486, Y: 140.7892454496811}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 196.65484712449938, Y: 148.36299713791232}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 222.4654592256012, Y: 150.301246864531}, Radius: 0, Start: 0, Angle: 0},
				}},
				setColor{col: color.RGBA{R: 0xff, G: 0x0, B: 0x0, A: 0xff}},
				setWidth{w: 2},
				setLineDash{dashes: []vg.Length(nil), offsets: 0},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 222.4654592256012, Y: 150.301246864531}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 70, Start: 6.2517693806436885, Angle: -0.3401785862023008},
				}},
				setColor{col: color.RGBA{R: 0x0, G: 0x0, B: 0xff, A: 0xff}},
				setWidth{w: 2},
				setLineDash{dashes: []vg.Length(nil), offsets: 0},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 84.71361238980812, Y: 135.03558890884597}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 70, Start: 3.393747857626204, Angle: 1.0942459830557514},
				}},
				setColor{col: color.RGBA{R: 0xc4, G: 0x18, B: 0x80, A: 0xff}},
				fill{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 82.74814071627586, Y: 158.38881367200094}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 70, Start: 3.057367197760076, Angle: -1.0086461940848572},
					{Type: vg.LineComp, Pos: vg.Point{X: 134.46313551225836, Y: 193.39542462921173}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 151.18289877836887, Y: 179.33719016058748}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 170.46367310317905, Y: 172.48189130151792}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 192.3054584866889, Y: 172.82952805200304}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 216.70825492889844, Y: 180.38010041204285}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 70, Start: 0.4096487325905134, Angle: -0.3782328060546156},
					{Type: vg.LineComp, Pos: vg.Point{X: 194.4878195330358, Y: 154.14275455358018}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 166.52726783582057, Y: 154.233761316289}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 138.5838041339555, Y: 154.97177342359538}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 110.65742842744058, Y: 156.35679087549937}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 82.74814071627586, Y: 158.38881367200094}, Radius: 0, Start: 0, Angle: 0},
				}},
				setColor{col: color.RGBA{R: 0xff, G: 0x0, B: 0xff, A: 0xfe}},
				setWidth{w: 1},
				setLineDash{dashes: []vg.Length(nil), offsets: 0},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 82.74814071627586, Y: 158.38881367200094}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.MoveComp, Pos: vg.Point{X: 120.30438330484755, Y: 214.65659470739064}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 134.46313551225836, Y: 193.39542462921173}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 151.18289877836887, Y: 179.33719016058748}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 170.46367310317905, Y: 172.48189130151792}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 192.3054584866889, Y: 172.82952805200304}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 216.70825492889844, Y: 180.38010041204285}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.MoveComp, Pos: vg.Point{X: 222.4654592256012, Y: 154.69875313546896}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 194.4878195330358, Y: 154.14275455358018}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 166.52726783582057, Y: 154.233761316289}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 138.5838041339555, Y: 154.97177342359538}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 110.65742842744058, Y: 156.35679087549937}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 82.74814071627586, Y: 158.38881367200094}, Radius: 0, Start: 0, Angle: 0},
				}},
				setColor{col: color.RGBA{R: 0xff, G: 0x0, B: 0x0, A: 0xff}},
				setWidth{w: 2},
				setLineDash{dashes: []vg.Length(nil), offsets: 0},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 82.74814071627586, Y: 158.38881367200094}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 70, Start: 3.057367197760076, Angle: -1.0086461940848572},
				}},
				setColor{col: color.RGBA{R: 0x0, G: 0x0, B: 0xff, A: 0xff}},
				setWidth{w: 2},
				setLineDash{dashes: []vg.Length(nil), offsets: 0},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 216.70825492889844, Y: 180.38010041204285}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 70, Start: 0.4096487325905134, Angle: -0.3782328060546156}},
				}},
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
			actions: []interface{}{
				setColor{col: color.RGBA{R: 0xc4, G: 0x18, B: 0x80, A: 0xff}},
				fill{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 222.4654592256012, Y: 150.301246864531}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 70, Start: 6.2517693806436885, Angle: -0.3401785862023008},
					{Type: vg.LineComp, Pos: vg.Point{X: 192.90504382281705, Y: 133.77306323362805}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 170.6307430157773, Y: 133.50989617220003}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 150.89956679501304, Y: 126.29338825313054}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 133.7115151605242, Y: 112.1235394764196}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 119.0665881123108, Y: 91.00034984206721}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 70, Start: 4.214445033887625, Angle: -1.0942459830557514},
					{Type: vg.LineComp, Pos: vg.Point{X: 110.50887014975963, Y: 153.3704101695125}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 138.50024010275996, Y: 152.6872771642238}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 166.49012826640032, Y: 151.94803894496397}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 194.47853464068075, Y: 151.15269551173304}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 222.4654592256012, Y: 150.301246864531}, Radius: 0, Start: 0, Angle: 0},
				}},
				setColor{col: color.RGBA{R: 0xff, G: 0x0, B: 0xff, A: 0xfe}},
				setWidth{w: 1},
				setLineDash{dashes: []vg.Length(nil), offsets: 0},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 222.4654592256012, Y: 150.301246864531}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.MoveComp, Pos: vg.Point{X: 217.7224692161322, Y: 127.08288943741465}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 192.90504382281705, Y: 133.77306323362805}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 170.6307430157773, Y: 133.50989617220003}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 150.89956679501304, Y: 126.29338825313054}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 133.7115151605242, Y: 112.1235394764196}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 119.0665881123108, Y: 91.00034984206721}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.MoveComp, Pos: vg.Point{X: 82.51601840739936, Y: 153.99743796083007}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 110.50887014975963, Y: 153.3704101695125}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 138.50024010275996, Y: 152.6872771642238}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 166.49012826640032, Y: 151.94803894496397}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 194.47853464068075, Y: 151.15269551173304}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 222.4654592256012, Y: 150.301246864531}, Radius: 0, Start: 0, Angle: 0},
				}},
				setColor{col: color.RGBA{R: 0xff, G: 0x0, B: 0x0, A: 0xff}},
				setWidth{w: 2},
				setLineDash{dashes: []vg.Length(nil), offsets: 0},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 222.4654592256012, Y: 150.301246864531}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 70, Start: 6.2517693806436885, Angle: -0.3401785862023008},
				}},
				setColor{col: color.RGBA{R: 0x0, G: 0x0, B: 0xff, A: 0xff}},
				setWidth{w: 2},
				setLineDash{dashes: []vg.Length(nil), offsets: 0},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 119.0665881123108, Y: 91.00034984206721}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 70, Start: 4.214445033887625, Angle: -1.0942459830557514},
				}},
				setColor{col: color.RGBA{R: 0xc4, G: 0x18, B: 0x80, A: 0xff}},
				fill{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 82.74814071627586, Y: 158.38881367200094}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 70, Start: 3.057367197760076, Angle: -1.0086461940848572},
					{Type: vg.LineComp, Pos: vg.Point{X: 134.46313551225836, Y: 193.39542462921173}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 151.18289877836887, Y: 179.33719016058748}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 170.46367310317905, Y: 172.48189130151792}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 192.3054584866889, Y: 172.82952805200304}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 216.70825492889844, Y: 180.38010041204285}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 70, Start: 0.4096487325905134, Angle: -0.3782328060546156},
					{Type: vg.LineComp, Pos: vg.Point{X: 194.4878195330358, Y: 154.14275455358018}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 166.52726783582057, Y: 154.233761316289}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 138.5838041339555, Y: 154.97177342359538}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 110.65742842744058, Y: 156.35679087549937}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 82.74814071627586, Y: 158.38881367200094}, Radius: 0, Start: 0, Angle: 0},
				}},
				setColor{col: color.RGBA{R: 0xff, G: 0x0, B: 0xff, A: 0xfe}},
				setWidth{w: 1},
				setLineDash{dashes: []vg.Length(nil), offsets: 0},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 82.74814071627586, Y: 158.38881367200094}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.MoveComp, Pos: vg.Point{X: 120.30438330484755, Y: 214.65659470739064}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 134.46313551225836, Y: 193.39542462921173}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 151.18289877836887, Y: 179.33719016058748}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 170.46367310317905, Y: 172.48189130151792}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 192.3054584866889, Y: 172.82952805200304}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 216.70825492889844, Y: 180.38010041204285}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.MoveComp, Pos: vg.Point{X: 222.4654592256012, Y: 154.69875313546896}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 194.4878195330358, Y: 154.14275455358018}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 166.52726783582057, Y: 154.233761316289}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 138.5838041339555, Y: 154.97177342359538}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 110.65742842744058, Y: 156.35679087549937}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 82.74814071627586, Y: 158.38881367200094}, Radius: 0, Start: 0, Angle: 0},
				}},
				setColor{col: color.RGBA{R: 0xff, G: 0x0, B: 0x0, A: 0xff}},
				setWidth{w: 2},
				setLineDash{dashes: []vg.Length(nil), offsets: 0},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 82.74814071627586, Y: 158.38881367200094}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 70, Start: 3.057367197760076, Angle: -1.0086461940848572},
				}},
				setColor{col: color.RGBA{R: 0x0, G: 0x0, B: 0xff, A: 0xff}},
				setWidth{w: 2},
				setLineDash{dashes: []vg.Length(nil), offsets: 0},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 216.70825492889844, Y: 180.38010041204285}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 70, Start: 0.4096487325905134, Angle: -0.3782328060546156},
				}},
			},
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
			actions: []interface{}{
				setColor{col: color.RGBA{R: 0xc4, G: 0x18, B: 0x80, A: 0xff}},
				fill{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 222.4654592256012, Y: 150.301246864531}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 70, Start: 6.2517693806436885, Angle: -0.3401785862023008},
					{Type: vg.LineComp, Pos: vg.Point{X: 191.53092479391694, Y: 135.5344727962992}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 165.1342669001769, Y: 140.55553442288462}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 138.53249553491207, Y: 142.1460743171709}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 111.72561069812248, Y: 140.306092479158}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 84.71361238980812, Y: 135.03558890884597}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 70, Start: 3.393747857626204, Angle: 1.0942459830557514},
					{Type: vg.LineComp, Pos: vg.Point{X: 145.32986989085776, Y: 108.7352361883809}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 158.08705245712767, Y: 127.57999179983727}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 175.19537820167486, Y: 140.7892454496811}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 196.65484712449938, Y: 148.36299713791232}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 222.4654592256012, Y: 150.301246864531}, Radius: 0, Start: 0, Angle: 0},
				}},
				setColor{col: color.RGBA{R: 0xff, G: 0x0, B: 0xff, A: 0xfe}},
				setWidth{w: 1},
				setLineDash{dashes: []vg.Length(nil), offsets: 0},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 222.4654592256012, Y: 150.301246864531}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.MoveComp, Pos: vg.Point{X: 217.7224692161322, Y: 127.08288943741465}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 191.53092479391694, Y: 135.5344727962992}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 165.1342669001769, Y: 140.55553442288462}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 138.53249553491207, Y: 142.1460743171709}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 111.72561069812248, Y: 140.306092479158}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 84.71361238980812, Y: 135.03558890884597}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.MoveComp, Pos: vg.Point{X: 136.92383050286517, Y: 84.25497861531197}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 145.32986989085776, Y: 108.7352361883809}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 158.08705245712767, Y: 127.57999179983727}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 175.19537820167486, Y: 140.7892454496811}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 196.65484712449938, Y: 148.36299713791232}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 222.4654592256012, Y: 150.301246864531}, Radius: 0, Start: 0, Angle: 0},
				}},
				setColor{col: color.RGBA{R: 0xff, G: 0x0, B: 0x0, A: 0xff}},
				setWidth{w: 2},
				setLineDash{dashes: []vg.Length(nil), offsets: 0},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 222.4654592256012, Y: 150.301246864531}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 70, Start: 6.2517693806436885, Angle: -0.3401785862023008},
				}},
				setColor{col: color.RGBA{R: 0x0, G: 0x0, B: 0xff, A: 0xff}},
				setWidth{w: 2},
				setLineDash{dashes: []vg.Length(nil), offsets: 0},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 84.71361238980812, Y: 135.03558890884597}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 70, Start: 3.393747857626204, Angle: 1.0942459830557514},
				}},
				setColor{col: color.RGBA{R: 0xc4, G: 0x18, B: 0x80, A: 0xff}},
				fill{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 222.4654592256012, Y: 154.69875313546896}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 70, Start: 0.031415926535897754, Angle: 1.0086461940848572},
					{Type: vg.LineComp, Pos: vg.Point{X: 172.67038374470474, Y: 192.38633670499263}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 155.23186577061887, Y: 179.23025328639585}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 135.61611389965154, Y: 173.40224680492543}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 113.82312813180272, Y: 174.90231726058138}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 89.85290846707245, Y: 183.73046465336367}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 70, Start: 2.6791343917054604, Angle: 0.3782328060546156},
					{Type: vg.LineComp, Pos: vg.Point{X: 110.65742842744058, Y: 156.35679087549937}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 138.5838041339555, Y: 154.97177342359538}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 166.52726783582057, Y: 154.233761316289}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 194.4878195330358, Y: 154.14275455358018}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 222.4654592256012, Y: 154.69875313546896}, Radius: 0, Start: 0, Angle: 0},
				}},
				setColor{col: color.RGBA{R: 0xff, G: 0x0, B: 0xff, A: 0xfe}},
				setWidth{w: 1},
				setLineDash{dashes: []vg.Length(nil), offsets: 0},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 222.4654592256012, Y: 154.69875313546896}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.MoveComp, Pos: vg.Point{X: 187.93166782190912, Y: 212.87049706071576}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 172.67038374470474, Y: 192.38633670499263}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 155.23186577061887, Y: 179.23025328639585}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 135.61611389965154, Y: 173.40224680492543}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 113.82312813180272, Y: 174.90231726058138}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 89.85290846707245, Y: 183.73046465336367}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.MoveComp, Pos: vg.Point{X: 82.74814071627586, Y: 158.38881367200094}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 110.65742842744058, Y: 156.35679087549937}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 138.5838041339555, Y: 154.97177342359538}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 166.52726783582057, Y: 154.233761316289}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 194.4878195330358, Y: 154.14275455358018}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 222.4654592256012, Y: 154.69875313546896}, Radius: 0, Start: 0, Angle: 0},
				}},
				setColor{col: color.RGBA{R: 0xff, G: 0x0, B: 0x0, A: 0xff}},
				setWidth{w: 2},
				setLineDash{dashes: []vg.Length(nil), offsets: 0},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 222.4654592256012, Y: 154.69875313546896}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 70, Start: 0.031415926535897754, Angle: 1.0086461940848572},
				}},
				setColor{col: color.RGBA{R: 0x0, G: 0x0, B: 0xff, A: 0xff}},
				setWidth{w: 2},
				setLineDash{dashes: []vg.Length(nil), offsets: 0},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 89.85290846707245, Y: 183.73046465336367}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 70, Start: 2.6791343917054604, Angle: 0.3782328060546156},
				}},
			},
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
			actions: []interface{}{
				setColor{col: color.RGBA{R: 0xc4, G: 0x18, B: 0x80, A: 0xff}},
				fill{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 222.4654592256012, Y: 150.301246864531}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 70, Start: 6.2517693806436885, Angle: -0.3401785862023008},
					{Type: vg.LineComp, Pos: vg.Point{X: 191.53092479391694, Y: 135.5344727962992}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 165.1342669001769, Y: 140.55553442288462}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 138.53249553491207, Y: 142.1460743171709}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 111.72561069812248, Y: 140.306092479158}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 84.71361238980812, Y: 135.03558890884597}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 70, Start: 3.393747857626204, Angle: 1.0942459830557514},
					{Type: vg.LineComp, Pos: vg.Point{X: 145.32986989085776, Y: 108.7352361883809}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 158.08705245712767, Y: 127.57999179983727}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 175.19537820167486, Y: 140.7892454496811}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 196.65484712449938, Y: 148.36299713791232}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 222.4654592256012, Y: 150.301246864531}, Radius: 0, Start: 0, Angle: 0},
				}},
				setColor{col: color.RGBA{R: 0xff, G: 0x0, B: 0xff, A: 0xfe}},
				setWidth{w: 1},
				setLineDash{dashes: []vg.Length(nil), offsets: 0},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 222.4654592256012, Y: 150.301246864531}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.MoveComp, Pos: vg.Point{X: 217.7224692161322, Y: 127.08288943741465}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 191.53092479391694, Y: 135.5344727962992}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 165.1342669001769, Y: 140.55553442288462}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 138.53249553491207, Y: 142.1460743171709}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 111.72561069812248, Y: 140.306092479158}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 84.71361238980812, Y: 135.03558890884597}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.MoveComp, Pos: vg.Point{X: 136.92383050286517, Y: 84.25497861531197}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 145.32986989085776, Y: 108.7352361883809}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 158.08705245712767, Y: 127.57999179983727}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 175.19537820167486, Y: 140.7892454496811}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 196.65484712449938, Y: 148.36299713791232}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 222.4654592256012, Y: 150.301246864531}, Radius: 0, Start: 0, Angle: 0},
				}},
				setColor{col: color.RGBA{R: 0xff, G: 0x0, B: 0x0, A: 0xff}},
				setWidth{w: 2},
				setLineDash{dashes: []vg.Length(nil), offsets: 0},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 222.4654592256012, Y: 150.301246864531}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 70, Start: 6.2517693806436885, Angle: -0.3401785862023008},
				}},
				setColor{col: color.RGBA{R: 0x0, G: 0x0, B: 0xff, A: 0xff}},
				setWidth{w: 2},
				setLineDash{dashes: []vg.Length(nil), offsets: 0},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 84.71361238980812, Y: 135.03558890884597}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 70, Start: 3.393747857626204, Angle: 1.0942459830557514},
				}},
				setColor{col: color.RGBA{R: 0xc4, G: 0x18, B: 0x80, A: 0xff}},
				fill{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 82.74814071627586, Y: 158.38881367200094}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 70, Start: 3.057367197760076, Angle: -1.0086461940848572},
					{Type: vg.LineComp, Pos: vg.Point{X: 134.6934236841265, Y: 192.36817073814876}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 152.1040514658413, Y: 175.22817459633566}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 172.53626664999203, Y: 163.23660628195134}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 195.99006923657868, Y: 156.39346579499576}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 222.4654592256012, Y: 154.69875313546896}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 70, Start: 0.031415926535897754, Angle: 0.3782328060546156},
					{Type: vg.LineComp, Pos: vg.Point{X: 190.80320878314603, Y: 170.57881681058745}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 164.45467428900756, Y: 163.47904633585557}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 137.66265144648307, Y: 159.0807889878472}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 110.42714025557248, Y: 157.3840447665623}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 82.74814071627586, Y: 158.38881367200094}, Radius: 0, Start: 0, Angle: 0},
				}},
				setColor{col: color.RGBA{R: 0xff, G: 0x0, B: 0xff, A: 0xfe}},
				setWidth{w: 1},
				setLineDash{dashes: []vg.Length(nil), offsets: 0},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 82.74814071627586, Y: 158.38881367200094}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.MoveComp, Pos: vg.Point{X: 120.30438330484755, Y: 214.65659470739064}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 134.6934236841265, Y: 192.36817073814876}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 152.1040514658413, Y: 175.22817459633566}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 172.53626664999203, Y: 163.23660628195134}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 195.99006923657868, Y: 156.39346579499576}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 222.4654592256012, Y: 154.69875313546896}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.MoveComp, Pos: vg.Point{X: 216.70825492889844, Y: 180.38010041204285}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 190.80320878314603, Y: 170.57881681058745}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 164.45467428900756, Y: 163.47904633585557}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 137.66265144648307, Y: 159.0807889878472}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 110.42714025557248, Y: 157.3840447665623}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 82.74814071627586, Y: 158.38881367200094}, Radius: 0, Start: 0, Angle: 0},
				}},
				setColor{col: color.RGBA{R: 0xff, G: 0x0, B: 0x0, A: 0xff}},
				setWidth{w: 2},
				setLineDash{dashes: []vg.Length(nil), offsets: 0},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 82.74814071627586, Y: 158.38881367200094}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 70, Start: 3.057367197760076, Angle: -1.0086461940848572},
				}},
				setColor{col: color.RGBA{R: 0x0, G: 0x0, B: 0xff, A: 0xff}},
				setWidth{w: 2},
				setLineDash{dashes: []vg.Length(nil), offsets: 0},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 222.4654592256012, Y: 154.69875313546896}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 70, Start: 0.031415926535897754, Angle: 0.3782328060546156},
				}},
			},
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

		tc := &canvas{dpi: defaultDPI}
		p.Draw(draw.NewCanvas(tc, 300, 300))

		base.append(test.actions...)
		ok := reflect.DeepEqual(tc.actions, base.actions)
		if !ok {
			t.Errorf("unexpected actions for test %d:\ngot :%#v\nwant:%#v", i, tc.actions, base.actions)
		}
		p.Add(b)
		checkImage(t, fmt.Sprintf("ribbons-%d", i), p, *allPics)
		if *pics && !ok || *allPics {
			err = p.Save(vg.Length(300), vg.Length(300), fmt.Sprintf("ribbons-%d-%s.svg", i, failure(!ok)))
			if err != nil {
				t.Fatalf("unexpected error writing file: %v", err)
			}
		}
	}
}
