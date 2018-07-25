// Copyright ©2013 The bíogo Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rings

import (
	"fmt"
	"image/color"
	"math/rand"
	"reflect"

	"github.com/biogo/biogo/feat"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
	check "gopkg.in/check.v1"
)

func (s *S) TestSail(c *check.C) {
	rand.Seed(1)
	b, err := NewGappedBlocks(randomFeatures(3, 100000, 1000000, false, plotter.DefaultLineStyle),
		Arc{0, Complete * Clockwise},
		80, 100, 0.01,
	)
	c.Assert(err, check.Equals, nil)

	redSty := plotter.DefaultLineStyle
	redSty.Width *= 2
	redSty.Color = color.RGBA{R: 255, G: 0, B: 0, A: 255}
	blueSty := plotter.DefaultLineStyle
	blueSty.Width *= 2
	blueSty.Color = color.RGBA{R: 0, G: 0, B: 255, A: 255}

	for i, t := range []struct {
		orient   []feat.Orientation
		ends     []feat.Feature
		segments int
		twist    Twist
		actions  []interface{}
	}{
		{
			orient: []feat.Orientation{feat.NotOriented, feat.NotOriented, feat.NotOriented},
			ends: []feat.Feature{
				&fs{
					start:    b.Set[0].Start(),
					end:      b.Set[0].Start() + b.Set[0].Len()/8,
					orient:   feat.NotOriented,
					location: b.Set[0],
					style:    redSty,
				},
				&fs{
					start:    b.Set[1].Start() + b.Set[0].Len()/8,
					end:      b.Set[1].Start() + b.Set[0].Len()/4,
					orient:   feat.NotOriented,
					location: b.Set[1],
					style:    redSty,
				},
				&fs{
					start:    b.Set[2].Start() + 2*b.Set[2].Len()/5,
					end:      b.Set[2].End() - 2*b.Set[2].Len()/5,
					orient:   feat.Reverse,
					location: b.Set[2],
					style:    blueSty,
				},
			},
			segments: 5,
			twist:    Individual | Flat,
			actions: []interface{}{
				setColor{col: color.RGBA{R: 0xc4, G: 0x18, B: 0x80, A: 0xff}},
				fill{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 133.41119163435226, Y: 219.8469924731578}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 70, Start: 1.8469917648582488, Angle: -0.605200405420522},
					{Type: vg.LineComp, Pos: vg.Point{X: 165.28125595455933, Y: 192.66745599072038}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 153.86738894959836, Y: 167.42972044885494}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 140.87550574358931, Y: 143.03229497168846}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 126.30560633653224, Y: 119.47517955922092}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 110.15769072842713, Y: 96.7583742114524}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 70, Start: 4.062761091201875, Angle: 0.21261637474004047},
					{Type: vg.LineComp, Pos: vg.Point{X: 136.25616936425217, Y: 111.23375127848458}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 152.70267184177737, Y: 126.9621675160278}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 172.21314634858985, Y: 136.26381125713445}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 194.7875928846896, Y: 139.13868250180454}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 220.4260114500766, Y: 135.58678125003806}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 70, Start: 6.039153005903646, Angle: 0.21261637474004225},
					{Type: vg.LineComp, Pos: vg.Point{X: 196.51434156975887, Y: 153.78667769222616}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 174.6333559827128, Y: 162.48396766693642}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 156.822502464463, Y: 176.39311678866179}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 143.08178101500948, Y: 195.51412505740225}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 133.41119163435226, Y: 219.8469924731578}, Radius: 0, Start: 0, Angle: 0},
				}},
				setColor{col: color.Gray16{Y: 0x0}},
				setWidth{w: 1},
				setLineDash{dashes: []vg.Length(nil), offsets: 0},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 133.41119163435226, Y: 219.8469924731578}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.MoveComp, Pos: vg.Point{X: 175.11710675847226, Y: 218.74550159728483}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 165.28125595455933, Y: 192.66745599072038}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 153.86738894959836, Y: 167.42972044885494}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 140.87550574358931, Y: 143.03229497168846}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 126.30560633653224, Y: 119.47517955922092}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 110.15769072842713, Y: 96.7583742114524}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.MoveComp, Pos: vg.Point{X: 122.87363891601422, Y: 89.07856254450478}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 136.25616936425217, Y: 111.23375127848458}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 152.70267184177737, Y: 126.9621675160278}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 172.21314634858985, Y: 136.26381125713445}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 194.7875928846896, Y: 139.13868250180454}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 220.4260114500766, Y: 135.58678125003806}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.MoveComp, Pos: vg.Point{X: 222.4654592256012, Y: 150.301246864531}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 196.51434156975887, Y: 153.78667769222616}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 174.6333559827128, Y: 162.48396766693642}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 156.822502464463, Y: 176.39311678866179}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 143.08178101500948, Y: 195.51412505740225}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 133.41119163435226, Y: 219.8469924731578}, Radius: 0, Start: 0, Angle: 0},
				}},
				setColor{col: color.RGBA{R: 0x0, G: 0x0, B: 0xff, A: 0xff}},
				setWidth{w: 2},
				setLineDash{dashes: []vg.Length(nil), offsets: 0},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 133.41119163435226, Y: 219.8469924731578}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 70, Start: 1.8469917648582488, Angle: -0.605200405420522},
				}},
				setColor{col: color.RGBA{R: 0xff, G: 0x0, B: 0x0, A: 0xff}},
				setWidth{w: 2},
				setLineDash{dashes: []vg.Length(nil), offsets: 0},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 110.15769072842713, Y: 96.7583742114524}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 70, Start: 4.062761091201875, Angle: 0.21261637474004047},
				}},
				setColor{col: color.RGBA{R: 0xff, G: 0x0, B: 0x0, A: 0xff}},
				setWidth{w: 2},
				setLineDash{dashes: []vg.Length(nil), offsets: 0},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 220.4260114500766, Y: 135.58678125003806}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 70, Start: 6.039153005903646, Angle: 0.21261637474004225},
				}},
			},
		},
		{
			orient: []feat.Orientation{feat.Reverse, feat.NotOriented, feat.NotOriented},
			ends: []feat.Feature{
				&fs{
					start:    b.Set[0].Start(),
					end:      b.Set[0].Start() + b.Set[0].Len()/8,
					orient:   feat.NotOriented,
					location: b.Set[0],
					style:    redSty,
				},
				&fs{
					start:    b.Set[1].Start() + b.Set[0].Len()/8,
					end:      b.Set[1].Start() + b.Set[0].Len()/4,
					orient:   feat.NotOriented,
					location: b.Set[1],
					style:    redSty,
				},
				&fs{
					start:    b.Set[2].Start() + 2*b.Set[2].Len()/5,
					end:      b.Set[2].End() - 2*b.Set[2].Len()/5,
					orient:   feat.Reverse,
					location: b.Set[2],
					style:    blueSty,
				},
			},
			segments: 5,
			twist:    Individual | Flat,
			actions: []interface{}{
				setColor{col: color.RGBA{R: 0xc4, G: 0x18, B: 0x80, A: 0xff}},
				fill{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 133.41119163435226, Y: 219.8469924731578}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 70, Start: 1.8469917648582488, Angle: -0.605200405420522},
					{Type: vg.LineComp, Pos: vg.Point{X: 165.28125595455933, Y: 192.66745599072038}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 153.86738894959836, Y: 167.42972044885494}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 140.87550574358931, Y: 143.03229497168846}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 126.30560633653224, Y: 119.47517955922092}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 110.15769072842713, Y: 96.7583742114524}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 70, Start: 4.062761091201875, Angle: 0.21261637474004047},
					{Type: vg.LineComp, Pos: vg.Point{X: 133.08871718078365, Y: 109.14674438051988}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 140.03286310790335, Y: 118.61413992416902}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 143.70607669737328, Y: 117.48074917545219}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 144.10835794919348, Y: 105.74657213436939}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 141.23970686336392, Y: 83.41160880092065}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 70, Start: 4.550825693753751, Angle: 0.21261637474004225},
					{Type: vg.LineComp, Pos: vg.Point{X: 154.0226325850737, Y: 110.4522507778316}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 150.7317696788274, Y: 138.1083525275895}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 146.1995752182917, Y: 165.55951006006313}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 140.42604920346668, Y: 192.8057233752526}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 133.41119163435226, Y: 219.8469924731578}, Radius: 0, Start: 0, Angle: 0},
				}},
				setColor{col: color.Gray16{Y: 0x0}},
				setWidth{w: 1},
				setLineDash{dashes: []vg.Length(nil), offsets: 0},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 133.41119163435226, Y: 219.8469924731578}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.MoveComp, Pos: vg.Point{X: 175.11710675847226, Y: 218.74550159728483}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 165.28125595455933, Y: 192.66745599072038}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 153.86738894959836, Y: 167.42972044885494}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 140.87550574358931, Y: 143.03229497168846}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 126.30560633653224, Y: 119.47517955922092}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 110.15769072842713, Y: 96.7583742114524}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.MoveComp, Pos: vg.Point{X: 122.87363891601422, Y: 89.07856254450478}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 133.08871718078365, Y: 109.14674438051988}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 140.03286310790335, Y: 118.61413992416902}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 143.70607669737328, Y: 117.48074917545219}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 144.10835794919348, Y: 105.74657213436939}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 141.23970686336392, Y: 83.41160880092065}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.MoveComp, Pos: vg.Point{X: 156.0721639370306, Y: 82.59120481078952}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 154.0226325850737, Y: 110.4522507778316}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 150.7317696788274, Y: 138.1083525275895}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 146.1995752182917, Y: 165.55951006006313}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 140.42604920346668, Y: 192.8057233752526}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 133.41119163435226, Y: 219.8469924731578}, Radius: 0, Start: 0, Angle: 0},
				}},
				setColor{col: color.RGBA{R: 0x0, G: 0x0, B: 0xff, A: 0xff}},
				setWidth{w: 2},
				setLineDash{dashes: []vg.Length(nil), offsets: 0},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 133.41119163435226, Y: 219.8469924731578}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 70, Start: 1.8469917648582488, Angle: -0.605200405420522},
				}},
				setColor{col: color.RGBA{R: 0xff, G: 0x0, B: 0x0, A: 0xff}},
				setWidth{w: 2},
				setLineDash{dashes: []vg.Length(nil), offsets: 0},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 110.15769072842713, Y: 96.7583742114524}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 70, Start: 4.062761091201875, Angle: 0.21261637474004047},
				}},
				setColor{col: color.RGBA{R: 0xff, G: 0x0, B: 0x0, A: 0xff}},
				setWidth{w: 2},
				setLineDash{dashes: []vg.Length(nil), offsets: 0},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 141.23970686336392, Y: 83.41160880092065}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 70, Start: 4.550825693753751, Angle: 0.21261637474004225},
				}},
			},
		},
		{
			orient: []feat.Orientation{feat.NotOriented, feat.NotOriented, feat.NotOriented},
			ends: []feat.Feature{
				&fs{
					start:    b.Set[0].Start(),
					end:      b.Set[0].Start() + b.Set[0].Len()/8,
					orient:   feat.NotOriented,
					location: b.Set[0],
					style:    redSty,
				},
				&fs{
					start:    b.Set[1].Start(),
					end:      b.Set[1].Start() + b.Set[1].Len()/8,
					orient:   feat.NotOriented,
					location: b.Set[1],
					style:    redSty,
				},
				&fs{
					start:    b.Set[0].Start() + 2*b.Set[0].Len()/5,
					end:      b.Set[0].End() - 2*b.Set[0].Len()/5,
					orient:   feat.NotOriented,
					location: b.Set[0],
					style:    blueSty,
				},
			},
			segments: 5,
			twist:    Individual | Flat,
			actions: []interface{}{
				setColor{col: color.RGBA{R: 0xc4, G: 0x18, B: 0x80, A: 0xff}},
				fill{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 125.53976099018678, Y: 87.90011213218907}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 70, Start: 4.317022664193097, Angle: 0.17097117648885884},
					{Type: vg.LineComp, Pos: vg.Point{X: 143.91961502120444, Y: 106.39163397144404}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 152.44603297851438, Y: 118.20558293208987}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 162.503084374795, Y: 119.69682549724942}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 174.0907692100463, Y: 110.8653616669227}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 187.20908748426825, Y: 91.71119144110972}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 70, Start: 5.231195555127963, Angle: 0.34020396414151577},
					{Type: vg.LineComp, Pos: vg.Point{X: 189.1394495052649, Y: 122.56069345242653}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 182.44951692109703, Y: 133.33357248887017}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 185.43396638384303, Y: 139.09554680061996}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 198.09279789350288, Y: 139.8466163876759}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 220.4260114500766, Y: 135.58678125003806}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 70, Start: 6.039153005903646, Angle: 0.21261637474004225},
					{Type: vg.LineComp, Pos: vg.Point{X: 196.19948434399225, Y: 148.5088024785874}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 173.3739270796463, Y: 141.37246681238142}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 153.98878743256344, Y: 128.89223986591304}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 138.04406540274357, Y: 111.06812163918224}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 125.53976099018678, Y: 87.90011213218907}, Radius: 0, Start: 0, Angle: 0},
				}},
				setColor{col: color.Gray16{Y: 0x0}},
				setWidth{w: 1},
				setLineDash{dashes: []vg.Length(nil), offsets: 0},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 125.53976099018678, Y: 87.90011213218907}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.MoveComp, Pos: vg.Point{X: 136.92383050286517, Y: 84.25497861531197}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 143.91961502120444, Y: 106.39163397144404}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 152.44603297851438, Y: 118.20558293208987}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 162.503084374795, Y: 119.69682549724942}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 174.0907692100463, Y: 110.8653616669227}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 187.20908748426825, Y: 91.71119144110972}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.MoveComp, Pos: vg.Point{X: 205.50376413634663, Y: 106.77690969128909}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 189.1394495052649, Y: 122.56069345242653}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 182.44951692109703, Y: 133.33357248887017}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 185.43396638384303, Y: 139.09554680061996}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 198.09279789350288, Y: 139.8466163876759}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 220.4260114500766, Y: 135.58678125003806}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.MoveComp, Pos: vg.Point{X: 222.4654592256012, Y: 150.301246864531}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 196.19948434399225, Y: 148.5088024785874}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 173.3739270796463, Y: 141.37246681238142}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 153.98878743256344, Y: 128.89223986591304}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 138.04406540274357, Y: 111.06812163918224}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 125.53976099018678, Y: 87.90011213218907}, Radius: 0, Start: 0, Angle: 0},
				}},
				setColor{col: color.RGBA{R: 0xff, G: 0x0, B: 0x0, A: 0xff}},
				setWidth{w: 2},
				setLineDash{dashes: []vg.Length(nil), offsets: 0},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 125.53976099018678, Y: 87.90011213218907}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 70, Start: 4.317022664193097, Angle: 0.17097117648885884},
				}},
				setColor{col: color.RGBA{R: 0x0, G: 0x0, B: 0xff, A: 0xff}},
				setWidth{w: 2},
				setLineDash{dashes: []vg.Length(nil), offsets: 0},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 187.20908748426825, Y: 91.71119144110972}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 70, Start: 5.231195555127963, Angle: 0.34020396414151577},
				}},
				setColor{col: color.RGBA{R: 0xff, G: 0x0, B: 0x0, A: 0xff}},
				setWidth{w: 2},
				setLineDash{dashes: []vg.Length(nil), offsets: 0},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 220.4260114500766, Y: 135.58678125003806}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 70, Start: 6.039153005903646, Angle: 0.21261637474004225},
				}},
			},
		},
		{
			orient: []feat.Orientation{feat.NotOriented, feat.NotOriented, feat.NotOriented},
			ends: []feat.Feature{
				&fs{
					start:    b.Set[0].Start(),
					end:      b.Set[0].Start() + b.Set[0].Len()/8,
					orient:   feat.NotOriented,
					location: b.Set[0],
					style:    redSty,
				},
				&fs{
					start:    b.Set[1].Start(),
					end:      b.Set[1].Start() + b.Set[1].Len()/8,
					orient:   feat.NotOriented,
					location: b.Set[1],
					style:    redSty,
				},
				&fs{
					start:    b.Set[0].Start() + 2*b.Set[0].Len()/5,
					end:      b.Set[0].End() - 2*b.Set[0].Len()/5,
					orient:   feat.NotOriented,
					location: b.Set[0],
					style:    blueSty,
				},
				&fs{
					start:    b.Set[0].Start() + b.Set[0].Len()/5,
					end:      b.Set[0].Start() + 2*b.Set[0].Len()/7,
					orient:   feat.Reverse,
					location: b.Set[0],
					style:    blueSty,
				},
			},
			segments: 5,
			twist:    Individual | Flat,
			actions: []interface{}{
				setColor{col: color.RGBA{R: 0xc4, G: 0x18, B: 0x80, A: 0xff}},
				fill{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 125.53976099018678, Y: 87.90011213218907}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 70, Start: 4.317022664193097, Angle: 0.17097117648885884},
					{Type: vg.LineComp, Pos: vg.Point{X: 143.91961502120444, Y: 106.39163397144404}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 152.44603297851438, Y: 118.20558293208987}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 162.503084374795, Y: 119.69682549724942}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 174.0907692100463, Y: 110.8653616669227}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 187.20908748426825, Y: 91.71119144110972}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 70, Start: 5.231195555127963, Angle: 0.34020396414151577},
					{Type: vg.LineComp, Pos: vg.Point{X: 189.03130781590716, Y: 122.2205377799216}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 182.01695016366594, Y: 131.9729497988504}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 184.46069117962304, Y: 136.03414574807553}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 196.3625308637785, Y: 134.40412562759693}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 217.7224692161322, Y: 127.08288943741465}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 70, Start: 5.911590794441388, Angle: -0.14579626078796082},
					{Type: vg.LineComp, Pos: vg.Point{X: 194.15328928066876, Y: 129.66475902163717}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 185.2698017947617, Y: 137.32960937155113}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 186.687426327694, Y: 140.87156319292262}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 198.40616287946563, Y: 140.2906204857516}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 220.4260114500766, Y: 135.58678125003806}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 70, Start: 6.039153005903646, Angle: 0.21261637474004225},
					{Type: vg.LineComp, Pos: vg.Point{X: 196.19948434399225, Y: 148.5088024785874}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 173.3739270796463, Y: 141.37246681238142}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 153.98878743256344, Y: 128.89223986591304}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 138.04406540274357, Y: 111.06812163918224}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 125.53976099018678, Y: 87.90011213218907}, Radius: 0, Start: 0, Angle: 0},
				}},
				setColor{col: color.Gray16{Y: 0x0}},
				setWidth{w: 1},
				setLineDash{dashes: []vg.Length(nil), offsets: 0},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 125.53976099018678, Y: 87.90011213218907}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.MoveComp, Pos: vg.Point{X: 136.92383050286517, Y: 84.25497861531197}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 143.91961502120444, Y: 106.39163397144404}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 152.44603297851438, Y: 118.20558293208987}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 162.503084374795, Y: 119.69682549724942}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 174.0907692100463, Y: 110.8653616669227}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 187.20908748426825, Y: 91.71119144110972}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.MoveComp, Pos: vg.Point{X: 205.50376413634663, Y: 106.77690969128909}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 189.03130781590716, Y: 122.2205377799216}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 182.01695016366594, Y: 131.9729497988504}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 184.46069117962304, Y: 136.03414574807553}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 196.3625308637785, Y: 134.40412562759693}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 217.7224692161322, Y: 127.08288943741465}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.MoveComp, Pos: vg.Point{X: 213.33788878541515, Y: 117.87701214318068}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 194.15328928066876, Y: 129.66475902163717}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 185.2698017947617, Y: 137.32960937155113}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 186.687426327694, Y: 140.87156319292262}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 198.40616287946563, Y: 140.2906204857516}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 220.4260114500766, Y: 135.58678125003806}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.MoveComp, Pos: vg.Point{X: 222.4654592256012, Y: 150.301246864531}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 196.19948434399225, Y: 148.5088024785874}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 173.3739270796463, Y: 141.37246681238142}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 153.98878743256344, Y: 128.89223986591304}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 138.04406540274357, Y: 111.06812163918224}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 125.53976099018678, Y: 87.90011213218907}, Radius: 0, Start: 0, Angle: 0},
				}},
				setColor{col: color.RGBA{R: 0xff, G: 0x0, B: 0x0, A: 0xff}},
				setWidth{w: 2},
				setLineDash{dashes: []vg.Length(nil), offsets: 0},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 125.53976099018678, Y: 87.90011213218907}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 70, Start: 4.317022664193097, Angle: 0.17097117648885884},
				}},
				setColor{col: color.RGBA{R: 0x0, G: 0x0, B: 0xff, A: 0xff}},
				setWidth{w: 2},
				setLineDash{dashes: []vg.Length(nil), offsets: 0},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 187.20908748426825, Y: 91.71119144110972}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 70, Start: 5.231195555127963, Angle: 0.34020396414151577},
				}},
				setColor{col: color.RGBA{R: 0x0, G: 0x0, B: 0xff, A: 0xff}},
				setWidth{w: 2},
				setLineDash{dashes: []vg.Length(nil), offsets: 0},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 217.7224692161322, Y: 127.08288943741465}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 70, Start: 5.911590794441388, Angle: -0.14579626078796082},
				}},
				setColor{col: color.RGBA{R: 0xff, G: 0x0, B: 0x0, A: 0xff}},
				setWidth{w: 2},
				setLineDash{dashes: []vg.Length(nil), offsets: 0},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 220.4260114500766, Y: 135.58678125003806}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 70, Start: 6.039153005903646, Angle: 0.21261637474004225},
				}},
			},
		},
	} {
		p, err := plot.New()
		c.Assert(err, check.Equals, nil)

		for j, o := range t.orient {
			b.Set[j].(*fs).orient = o
		}
		b.Base = NewGappedArcs(b.Base, b.Set, 0.01)

		r, err := NewSail(t.ends, b, 70)
		c.Assert(err, check.Equals, nil)
		r.Bezier = &Bezier{Segments: t.segments}
		r.Twist = t.twist
		r.LineStyle = plotter.DefaultLineStyle
		r.Color = color.RGBA{R: 0xc4, G: 0x18, B: 0x80, A: 0xff}
		p.Add(r)

		p.HideAxes()

		tc := &canvas{dpi: defaultDPI}
		p.Draw(draw.NewCanvas(tc, 300, 300))

		base.append(t.actions...)
		c.Check(tc.actions, check.DeepEquals, base.actions, check.Commentf("Test %d", i))
		if ok := reflect.DeepEqual(tc.actions, base.actions); *pics && !ok || *allPics {
			p.Add(b)
			c.Assert(p.Save(vg.Length(300), vg.Length(300), fmt.Sprintf("sail-%d-%s.svg", i, failure(!ok))), check.Equals, nil)
		}
	}
}
