// Copyright ©2013 The bíogo Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rings

import (
	"fmt"
	"image/color"
	"math"
	"reflect"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"

	"gopkg.in/check.v1"
)

func (s *S) TestHighlight(c *check.C) {
	p, err := plot.New()
	c.Assert(err, check.Equals, nil)

	h := NewHighlight(
		color.NRGBA{R: 0xf3, G: 0xf3, B: 0x15, A: 0xff},
		Arc{0, Complete / 2 * Clockwise},
		30, 120,
	)
	h.LineStyle = plotter.DefaultLineStyle
	p.Add(h)

	p.HideAxes()

	tc := &canvas{dpi: defaultDPI}
	p.Draw(draw.NewCanvas(tc, 300, 300))

	base.append(
		setColor{col: color.NRGBA{R: 0xf3, G: 0xf3, B: 0x15, A: 0xff}},
		fill{path: vg.Path{
			{Type: vg.MoveComp, Pos: vg.Point{X: 182.5, Y: 152.5}, Radius: 0, Start: 0, Angle: 0},
			{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 30, Start: 0, Angle: -math.Pi},
			{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 120, Start: -math.Pi, Angle: math.Pi},
			{Type: vg.CloseComp, Pos: vg.Point{X: 0, Y: 0}, Radius: 0, Start: 0, Angle: 0},
		}},
		setColor{col: color.Gray16{Y: 0x0}},
		setWidth{w: 1},
		setLineDash{dashes: []vg.Length(nil), offsets: 0},
		stroke{path: vg.Path{
			{Type: vg.MoveComp, Pos: vg.Point{X: 182.5, Y: 152.5}, Radius: 0, Start: 0, Angle: 0},
			{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 30, Start: 0, Angle: -math.Pi},
			{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 120, Start: -math.Pi, Angle: math.Pi},
			{Type: vg.CloseComp, Pos: vg.Point{X: 0, Y: 0}, Radius: 0, Start: 0, Angle: 0},
		}},
	)
	c.Check(tc.actions, check.DeepEquals, base.actions)
	if ok := reflect.DeepEqual(tc.actions, base.actions); *pics && !ok || *allPics {
		c.Assert(p.Save(vg.Length(300), vg.Length(300), fmt.Sprintf("highlight-%s.svg", failure(!ok))), check.Equals, nil)
	}
}
