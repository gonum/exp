// Copyright ©2013 The bíogo Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rings

import (
	"fmt"
	"image/color"
	"math/rand"
	"reflect"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"

	"gopkg.in/check.v1"
)

func (s *S) TestSpokes(c *check.C) {
	p, err := plot.New()
	c.Assert(err, check.Equals, nil)

	rand.Seed(1)
	b, err := NewGappedBlocks(randomFeatures(3, 100000, 1000000, false, plotter.DefaultLineStyle),
		Arc{0, Complete * Clockwise},
		80, 100, 0.01,
	)
	c.Assert(err, check.Equals, nil)

	m := randomFeatures(10, b.Set[1].Start(), b.Set[1].End(), true, plotter.DefaultLineStyle)
	for _, mf := range m {
		mf.(*fs).location = b.Set[1]
	}
	ms, err := NewSpokes(m, b, 73, 78)
	c.Assert(err, check.Equals, nil)
	ms.LineStyle = plotter.DefaultLineStyle
	p.Add(ms)

	p.HideAxes()

	tc := &canvas{dpi: defaultDPI}
	p.Draw(draw.NewCanvas(tc, 300, 300))

	base.append(
		setColor{col: color.Gray16{Y: 0x0}},
		setWidth{w: 1},
		setLineDash{dashes: []vg.Length(nil), offsets: 0},
		stroke{path: vg.Path{
			{Type: vg.MoveComp, Pos: vg.Point{X: 88.01347691001654, Y: 118.28759961994731}, Radius: 0, Start: 0, Angle: 0},
			{Type: vg.LineComp, Pos: vg.Point{X: 83.596591766867, Y: 115.94428452542314}, Radius: 0, Start: 0, Angle: 0},
		}},
		setColor{col: color.Gray16{Y: 0x0}},
		setWidth{w: 1},
		setLineDash{dashes: []vg.Length(nil), offsets: 0},
		stroke{path: vg.Path{
			{Type: vg.MoveComp, Pos: vg.Point{X: 118.21294090448178, Y: 88.05314143745684}, Radius: 0, Start: 0, Angle: 0},
			{Type: vg.LineComp, Pos: vg.Point{X: 115.8645121993093, Y: 83.6389730427621}, Radius: 0, Start: 0, Angle: 0},
		}},
		setColor{col: color.Gray16{Y: 0x0}},
		setWidth{w: 1},
		setLineDash{dashes: []vg.Length(nil), offsets: 0},
		stroke{path: vg.Path{
			{Type: vg.MoveComp, Pos: vg.Point{X: 80.97439251045306, Y: 137.90248400407566}, Radius: 0, Start: 0, Angle: 0},
			{Type: vg.LineComp, Pos: vg.Point{X: 76.07537829884025, Y: 136.90265414134112}, Radius: 0, Start: 0, Angle: 0},
		}},
		setColor{col: color.Gray16{Y: 0x0}},
		setWidth{w: 1},
		setLineDash{dashes: []vg.Length(nil), offsets: 0},
		stroke{path: vg.Path{
			{Type: vg.MoveComp, Pos: vg.Point{X: 81.8483994975755, Y: 134.13287321201787}, Radius: 0, Start: 0, Angle: 0},
			{Type: vg.LineComp, Pos: vg.Point{X: 77.00924877823135, Y: 132.87485082927938}, Radius: 0, Start: 0, Angle: 0},
		}},
		setColor{col: color.Gray16{Y: 0x0}},
		setWidth{w: 1},
		setLineDash{dashes: []vg.Length(nil), offsets: 0},
		stroke{path: vg.Path{
			{Type: vg.MoveComp, Pos: vg.Point{X: 124.72319085921325, Y: 84.99112003627741}, Radius: 0, Start: 0, Angle: 0},
			{Type: vg.LineComp, Pos: vg.Point{X: 122.82066968518674, Y: 80.3672241483512}, Radius: 0, Start: 0, Angle: 0},
		}},
		setColor{col: color.Gray16{Y: 0x0}},
		setWidth{w: 1},
		setLineDash{dashes: []vg.Length(nil), offsets: 0},
		stroke{path: vg.Path{
			{Type: vg.MoveComp, Pos: vg.Point{X: 100.00256872966926, Y: 101.77505830445045}, Radius: 0, Start: 0, Angle: 0},
			{Type: vg.LineComp, Pos: vg.Point{X: 96.40685425909868, Y: 98.3007472294128}, Radius: 0, Start: 0, Angle: 0},
		}},
		setColor{col: color.Gray16{Y: 0x0}},
		setWidth{w: 1},
		setLineDash{dashes: []vg.Length(nil), offsets: 0},
		stroke{path: vg.Path{
			{Type: vg.MoveComp, Pos: vg.Point{X: 89.97379980140269, Y: 114.82408874725218}, Radius: 0, Start: 0, Angle: 0},
			{Type: vg.LineComp, Pos: vg.Point{X: 85.69118334944397, Y: 112.24354688062562}, Radius: 0, Start: 0, Angle: 0},
		}},
		setColor{col: color.Gray16{Y: 0x0}},
		setWidth{w: 1},
		setLineDash{dashes: []vg.Length(nil), offsets: 0},
		stroke{path: vg.Path{
			{Type: vg.MoveComp, Pos: vg.Point{X: 95.34966934773948, Y: 107.08150479884551}, Radius: 0, Start: 0, Angle: 0},
			{Type: vg.LineComp, Pos: vg.Point{X: 91.43526313868054, Y: 103.97064896315}, Radius: 0, Start: 0, Angle: 0},
		}},
		setColor{col: color.Gray16{Y: 0x0}},
		setWidth{w: 1},
		setLineDash{dashes: []vg.Length(nil), offsets: 0},
		stroke{path: vg.Path{
			{Type: vg.MoveComp, Pos: vg.Point{X: 102.38779948295821, Y: 99.41735350098115}, Radius: 0, Start: 0, Angle: 0},
			{Type: vg.LineComp, Pos: vg.Point{X: 98.95545698179097, Y: 95.7815557955689}, Radius: 0, Start: 0, Angle: 0},
		}},
		setColor{col: color.Gray16{Y: 0x0}},
		setWidth{w: 1},
		setLineDash{dashes: []vg.Length(nil), offsets: 0},
		stroke{path: vg.Path{
			{Type: vg.MoveComp, Pos: vg.Point{X: 111.3429659019641, Y: 92.20822158657857}, Radius: 0, Start: 0, Angle: 0},
			{Type: vg.LineComp, Pos: vg.Point{X: 108.52399096374245, Y: 88.0786477226456}, Radius: 0, Start: 0, Angle: 0},
		}},
	)
	c.Check(tc.actions, check.DeepEquals, base.actions)
	if ok := reflect.DeepEqual(tc.actions, base.actions); *pics && !ok || *allPics {
		p.Add(b)
		c.Assert(p.Save(vg.Length(300), vg.Length(300), fmt.Sprintf("spokes-%s.svg", failure(!ok))), check.Equals, nil)
	}
}
