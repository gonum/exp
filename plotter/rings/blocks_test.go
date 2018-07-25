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

	"gopkg.in/check.v1"
)

func (s *S) TestBlocks(c *check.C) {
	p, err := plot.New()
	c.Assert(err, check.Equals, nil)

	rand.Seed(1)
	b, err := NewGappedBlocks(randomFeatures(3, 100000, 1000000, false, plotter.DefaultLineStyle),
		Arc{0, Complete * Clockwise},
		80, 100, 0.01,
	)
	c.Assert(err, check.Equals, nil)
	b.LineStyle = plotter.DefaultLineStyle
	b.Color = color.RGBA{R: 0xc4, G: 0x18, B: 0x80, A: 0xff}
	p.Add(b)

	p.HideAxes()

	tc := &canvas{dpi: defaultDPI}
	p.Draw(draw.NewCanvas(tc, 300, 300))

	base.append(
		setColor{col: color.RGBA{R: 0xc4, G: 0x18, B: 0x80, A: 0xff}},
		fill{path: vg.Path{
			{Type: vg.MoveComp, Pos: vg.Point{X: 232.46052482925853, Y: 149.98713927374973}, Radius: 0, Start: 0, Angle: 0},
			{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 80, Start: 6.2517693806436885, Angle: -1.7009436868899361},
			{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 100, Start: 4.550825693753753, Angle: 1.7009436868899361},
			{Type: vg.CloseComp, Pos: vg.Point{X: 0, Y: 0}, Radius: 0, Start: 0, Angle: 0},
		}},
		setColor{col: color.Gray16{Y: 0x0}},
		setWidth{w: 1},
		setLineDash{dashes: []vg.Length(nil), offsets: 0},
		stroke{path: vg.Path{
			{Type: vg.MoveComp, Pos: vg.Point{X: 232.46052482925853, Y: 149.98713927374973}, Radius: 0, Start: 0, Angle: 0},
			{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 80, Start: 6.2517693806436885, Angle: -1.7009436868899361},
			{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 100, Start: 4.550825693753753, Angle: 1.7009436868899361},
			{Type: vg.CloseComp, Pos: vg.Point{X: 0, Y: 0}, Radius: 0, Start: 0, Angle: 0},
		}},
		setColor{col: color.RGBA{R: 0xc4, G: 0x18, B: 0x80, A: 0xff}},
		fill{path: vg.Path{
			{Type: vg.MoveComp, Pos: vg.Point{X: 134.69866343184597, Y: 74.50568984607081}, Radius: 0, Start: 0, Angle: 0},
			{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 80, Start: 4.487993840681956, Angle: -1.367794789850083},
			{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 100, Start: 3.1201990508318733, Angle: 1.367794789850083},
			{Type: vg.CloseComp, Pos: vg.Point{X: 0, Y: 0}, Radius: 0, Start: 0, Angle: 0},
		}},
		setColor{col: color.Gray16{Y: 0x0}},
		setWidth{w: 1},
		setLineDash{dashes: []vg.Length(nil), offsets: 0},
		stroke{path: vg.Path{
			{Type: vg.MoveComp, Pos: vg.Point{X: 134.69866343184597, Y: 74.50568984607081}, Radius: 0, Start: 0, Angle: 0},
			{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 80, Start: 4.487993840681956, Angle: -1.367794789850083},
			{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 100, Start: 3.1201990508318733, Angle: 1.367794789850083},
			{Type: vg.CloseComp, Pos: vg.Point{X: 0, Y: 0}, Radius: 0, Start: 0, Angle: 0},
		}},
		setColor{col: color.RGBA{R: 0xc4, G: 0x18, B: 0x80, A: 0xff}},
		fill{path: vg.Path{
			{Type: vg.MoveComp, Pos: vg.Point{X: 72.78358939002953, Y: 159.230072768001}, Radius: 0, Start: 0, Angle: 0},
			{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 80, Start: 3.057367197760077, Angle: -3.0259512712241787},
			{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 100, Start: 0.0314159265358982, Angle: 3.0259512712241787},
			{Type: vg.CloseComp, Pos: vg.Point{X: 0, Y: 0}, Radius: 0, Start: 0, Angle: 0},
		}},
		setColor{col: color.Gray16{Y: 0x0}},
		setWidth{w: 1},
		setLineDash{dashes: []vg.Length(nil), offsets: 0},
		stroke{path: vg.Path{
			{Type: vg.MoveComp, Pos: vg.Point{X: 72.78358939002953, Y: 159.230072768001}, Radius: 0, Start: 0, Angle: 0},
			{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 80, Start: 3.057367197760077, Angle: -3.0259512712241787},
			{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 100, Start: 0.0314159265358982, Angle: 3.0259512712241787},
			{Type: vg.CloseComp, Pos: vg.Point{X: 0, Y: 0}, Radius: 0, Start: 0, Angle: 0},
		}},
	)
	c.Check(tc.actions, check.DeepEquals, base.actions)
	if ok := reflect.DeepEqual(tc.actions, base.actions); *pics && !ok || *allPics {
		c.Assert(p.Save(vg.Length(300), vg.Length(300), fmt.Sprintf("blocks-%s.svg", failure(!ok))), check.Equals, nil)
	}
}

func (s *S) TestBlocksScale(c *check.C) {
	rand.Seed(1)
	b, err := NewGappedBlocks(randomFeatures(3, 100000, 1000000, false, plotter.DefaultLineStyle),
		Arc{0, Complete * Clockwise},
		80, 100, 0.01,
	)
	c.Assert(err, check.Equals, nil)
	font, err := vg.MakeFont("Helvetica", 5)
	c.Assert(err, check.Equals, nil)

	for i, t := range []struct {
		feats   []feat.Feature
		grid    draw.LineStyle
		inner   vg.Length
		outer   vg.Length
		actions []interface{}
	}{
		{
			feats: b.Set,
			actions: []interface{}{
				setColor{col: color.Gray16{Y: 0x0}},
				setWidth{w: 1},
				setLineDash{dashes: []vg.Length(nil), offsets: 0},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 262.4457216402305, Y: 149.04481650140588}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 110, Start: 6.2517693806436885, Angle: -1.7009436868899357},
				}},
				setColor{col: color.Gray16{Y: 0x0}},
				setWidth{w: 1},
				setLineDash{dashes: []vg.Length(nil), offsets: 0},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 262.3290023868793, Y: 146.36891243718895}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 265.3243388156124, Y: 146.20170095820322}, Radius: 0, Start: 0, Angle: 0},
				}},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 245.50327522917544, Y: 93.75997278987488}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 248.03972818997113, Y: 92.15797204778056}, Radius: 0, Start: 0, Angle: 0},
				}},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 205.2283815234374, Y: 55.96131458364049}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 206.66642829225842, Y: 53.32844134501249}, Radius: 0, Start: 0, Angle: 0},
				}},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 151.6589388127897, Y: 42.503215428452776}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 151.63600078041122, Y: 39.50330312195604}, Radius: 0, Start: 0, Angle: 0},
				}},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 260.67012415777026, Y: 132.51940342000364}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 262.14517130537627, Y: 132.24694073936735}, Radius: 0, Start: 0, Angle: 0},
				}},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 257.27193490354546, Y: 118.99117046855808}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 258.7006431067756, Y: 118.53423188403842}, Radius: 0, Start: 0, Angle: 0},
				}},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 252.18907547888924, Y: 106.00173949313452}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 253.54847196269228, Y: 105.36767230440455}, Radius: 0, Start: 0, Angle: 0},
				}},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 237.3220378313814, Y: 82.4627106597371}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 238.4787019836275, Y: 81.50765671418806}, Radius: 0, Start: 0, Angle: 0},
				}},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 227.77691269913558, Y: 72.29160633395828}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 228.80341605412377, Y: 71.19785551123954}, Radius: 0, Start: 0, Angle: 0},
				}},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 217.02137974667383, Y: 63.410205098532714}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 217.90121674321938, Y: 62.19534425896725}, Radius: 0, Start: 0, Angle: 0},
				}},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 192.5875424011144, Y: 50.06470850220197}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 193.13419070658415, Y: 48.667863618141084}, Radius: 0, Start: 0, Angle: 0},
				}},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 179.30211952693816, Y: 45.8152007600721}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 179.66760297503276, Y: 44.360408043163986}, Radius: 0, Start: 0, Angle: 0},
				}},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 165.58573457009211, Y: 43.28112090503265}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 165.7641764051388, Y: 41.791772553737644}, Radius: 0, Start: 0, Angle: 0},
				}},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 137.74566681731258, Y: 43.49399258603107}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 137.5444713648214, Y: 42.00754703038605}, Radius: 0, Start: 0, Angle: 0},
				}},
				setColor{col: color.Gray16{Y: 0x0}},
				push{},
				rotate{angle: 4.656622921172369},
				fillString{font: "Helvetica", size: 5, x: -169.10513713195704, y: 260.82918038819696, str: "800000"},
				pop{},
				setColor{col: color.Gray16{Y: 0x0}},
				push{},
				rotate{angle: 4.149064136880077},
				fillString{font: "Helvetica", size: 5, x: -218.7136817198258, y: 164.56748322669478, str: "840000"},
				pop{},
				setColor{col: color.Gray16{Y: 0x0}},
				push{},
				rotate{angle: 3.641505352587785},
				fillString{font: "Helvetica", size: 5, x: -215.28071886824114, y: 56.32914989418752, str: "880000"},
				pop{},
				setColor{col: color.Gray16{Y: 0x0}},
				push{},
				rotate{angle: 3.1339465682954932},
				fillString{font: "Helvetica", size: 5, x: -159.67181075726256, y: -36.59539678344792, str: "920000"},
				pop{}, setColor{col: color.Gray16{Y: 0x0}},
				setWidth{w: 1},
				setLineDash{dashes: []vg.Length(nil), offsets: 0},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 128.0231622187882, Y: 45.25782353834737}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 110, Start: 4.487993840681956, Angle: -1.367794789850083},
				}},
				setColor{col: color.Gray16{Y: 0x0}},
				setWidth{w: 1},
				setLineDash{dashes: []vg.Length(nil), offsets: 0},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 117.07926421412347, Y: 48.358886714289795}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 116.11324414723593, Y: 45.51867453377044}, Radius: 0, Start: 0, Angle: 0},
				}},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 70.9273349616073, Y: 78.70365646907574}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 68.70262591510568, Y: 76.69102891823235}, Radius: 0, Start: 0, Angle: 0},
				}},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 45.34254168939712, Y: 127.65489729521317}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 42.420065553653416, Y: 126.97730358508262}, Radius: 0, Start: 0, Angle: 0},
				}},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 104.18503453050306, Y: 53.67862522874255}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 103.52619409228265, Y: 52.33106102731631}, Radius: 0, Start: 0, Angle: 0},
				}},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 92.06768065037733, Y: 60.58735245884155}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 91.24360356833702, Y: 59.33399817418939}, Radius: 0, Start: 0, Angle: 0},
				}},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 80.92204239077553, Y: 68.97398019483948}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 79.94597933246793, Y: 67.8349890156782}, Radius: 0, Start: 0, Angle: 0},
				}},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 62.24426729206817, Y: 89.61993389511773}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 61.01350730059637, Y: 88.76247844823297}, Radius: 0, Start: 0, Angle: 0},
				}},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 55.01245792699254, Y: 101.54728524834422}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 53.68308235326971, Y: 100.85247550173074}, Radius: 0, Start: 0, Angle: 0},
				}},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 49.34819004399671, Y: 114.29392583893788}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 47.94157445368758, Y: 113.77293391855977}, Radius: 0, Start: 0, Angle: 0},
				}},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 43.05992129759416, Y: 141.4153631718844}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 41.567556588015904, Y: 141.2642090333192}, Radius: 0, Start: 0, Angle: 0},
				}},
				setColor{col: color.Gray16{Y: 0x0}},
				push{},
				rotate{angle: 2.8137443446786836},
				fillString{font: "Helvetica", size: 5, x: -103.61371759920122, y: -76.41731046714297, str: "740000"},
				pop{},
				setColor{col: color.Gray16{Y: 0x0}},
				push{}, rotate{angle: 2.3061855603863917},
				fillString{font: "Helvetica", size: 5, x: 2.438524206376691, y: -98.33178158899622, str: "780000"},
				pop{},
				setColor{col: color.Gray16{Y: 0x0}}, push{}, rotate{angle: 1.7986267760940993},
				fillString{font: "Helvetica", size: 5, x: 105.77257147908588, y: -65.93738838923346, str: "820000"},
				pop{},
				setColor{col: color.Gray16{Y: 0x0}}, setWidth{w: 1}, setLineDash{dashes: []vg.Length(nil), offsets: 0},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 42.88993541129061, Y: 161.75385005600137}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 110, Start: 3.057367197760077, Angle: -3.0259512712241787},
				}},
				setColor{col: color.Gray16{Y: 0x0}},
				setWidth{w: 1},
				setLineDash{dashes: []vg.Length(nil), offsets: 0},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 50.14593244350226, Y: 192.79447672622013}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 47.354457873779594, Y: 193.8934170005716}, Radius: 0, Start: 0, Angle: 0},
				}},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 160.53873174123612, Y: 262.2058740086074}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 160.75796987963346, Y: 265.19785239066033}, Radius: 0, Start: 0, Angle: 0},
				}},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 259.63444739908493, Y: 177.44414120575627}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 262.5562959645145, Y: 178.12443596591328}, Radius: 0, Start: 0, Angle: 0},
				}},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 93.94804733691842, Y: 245.62179572656584}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 93.1496116187855, Y: 246.89163839556446}, Radius: 0, Start: 0, Angle: 0},
				}},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 224.00072122194018, Y: 236.09214595129373}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 224.9757310567848, Y: 237.23203885062958}, Radius: 0, Start: 0, Angle: 0},
				}},
				setColor{col: color.Gray16{Y: 0x0}},
				push{},
				rotate{angle: 1.1957523765353044},
				fillString{font: "Helvetica", size: 5, x: 185.94512762064295, y: 31.029002094671917, str: "325000.00"},
				pop{},
				setColor{col: color.Gray16{Y: 0x0}},
				push{},
				rotate{angle: -0.07314458419542547},
				fillString{font: "Helvetica", size: 5, x: 129.13000209936703, y: 280.30300189893075, str: "425000.00"},
				pop{},
				setColor{col: color.Gray16{Y: 0x0}},
				push{},
				rotate{angle: -1.3420415449261553},
				fillString{font: "Helvetica", size: 5, x: -125.76327872114877, y: 300.1751144933594, str: "525000.00"},
				pop{},
			},
		},
		{
			feats: b.Set,
			grid:  plotter.DefaultGridLineStyle,
			inner: b.Inner - 5,
			outer: b.Outer + 5,
			actions: []interface{}{
				setColor{col: color.Gray{Y: 0x80}},
				setWidth{w: 0.25},
				setLineDash{dashes: []vg.Length(nil), offsets: 0},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 227.38341071832681, Y: 148.31971302535612}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 257.33677500565756, Y: 146.64759823549855}, Radius: 0, Start: 0, Angle: 0},
				}},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 215.91132401989233, Y: 112.44998144764196}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 241.2758536278493, Y: 96.42997402669874}, Radius: 0, Start: 0, Angle: 0},
				}},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 188.45116922052551, Y: 86.67816903430034}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 202.8316369087357, Y: 60.34943664802046}, Radius: 0, Start: 0, Angle: 0},
				}},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 151.92654919053842, Y: 77.50219233758145}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 151.6971688667538, Y: 47.50306927261401}, Radius: 0, Start: 0, Angle: 0},
				}},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 226.25235738029792, Y: 138.8768659681843}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 255.7533003324171, Y: 133.42761235545802}, Radius: 0, Start: 0, Angle: 0},
				}},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 223.9354101615083, Y: 129.65307077401687}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 252.50957422611157, Y: 120.51429908362363}, Radius: 0, Start: 0, Angle: 0},
				}},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 220.46982419015177, Y: 120.79664056350082}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 247.65775386621246, Y: 108.11529678890113}, Radius: 0, Start: 0, Angle: 0},
				}},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 210.3332076123055, Y: 104.74730272254803}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 233.4664906572277, Y: 85.64622381156722}, Radius: 0, Start: 0, Angle: 0},
				}},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 203.82516774941064, Y: 97.81245886406248}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 224.35523484917488, Y: 75.93744240968746}, Radius: 0, Start: 0, Angle: 0},
				}},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 196.4918498272776, Y: 91.75695802172686}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 214.08858975818868, Y: 67.45974123041759}, Radius: 0, Start: 0, Angle: 0},
				}},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 179.8324152734871, Y: 82.65775579695588}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 190.7653813828819, Y: 54.72085811573824}, Radius: 0, Start: 0, Angle: 0},
				}},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 170.77417240473056, Y: 79.76036415459461}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 178.08384136662278, Y: 50.66450981643246}, Radius: 0, Start: 0, Angle: 0},
				}},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 161.42209175233552, Y: 78.03258243524954}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 164.99092845326973, Y: 48.245615409349355}, Radius: 0, Start: 0, Angle: 0},
				}},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 142.4402273754404, Y: 78.17772221774847}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 138.41631832561654, Y: 48.44881110484785}, Radius: 0, Start: 0, Angle: 0},
				}},
				setColor{col: color.Gray16{Y: 0x0}},
				setWidth{w: 1},
				setLineDash{dashes: []vg.Length(nil), offsets: 0},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 262.4457216402305, Y: 149.04481650140588}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 110, Start: 6.2517693806436885, Angle: -1.7009436868899357}}},
				setColor{col: color.Gray16{Y: 0x0}},
				setWidth{w: 1},
				setLineDash{dashes: []vg.Length(nil), offsets: 0},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 262.3290023868793, Y: 146.36891243718895}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 265.3243388156124, Y: 146.20170095820322}, Radius: 0, Start: 0, Angle: 0},
				}},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 245.50327522917544, Y: 93.75997278987488}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 248.03972818997113, Y: 92.15797204778056}, Radius: 0, Start: 0, Angle: 0},
				}},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 205.2283815234374, Y: 55.96131458364049}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 206.66642829225842, Y: 53.32844134501249}, Radius: 0, Start: 0, Angle: 0},
				}},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 151.6589388127897, Y: 42.503215428452776}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 151.63600078041122, Y: 39.50330312195604}, Radius: 0, Start: 0, Angle: 0},
				}},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 260.67012415777026, Y: 132.51940342000364}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 262.14517130537627, Y: 132.24694073936735}, Radius: 0, Start: 0, Angle: 0},
				}},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 257.27193490354546, Y: 118.99117046855808}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 258.7006431067756, Y: 118.53423188403842}, Radius: 0, Start: 0, Angle: 0},
				}},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 252.18907547888924, Y: 106.00173949313452}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 253.54847196269228, Y: 105.36767230440455}, Radius: 0, Start: 0, Angle: 0},
				}},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 237.3220378313814, Y: 82.4627106597371}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 238.4787019836275, Y: 81.50765671418806}, Radius: 0, Start: 0, Angle: 0},
				}},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 227.77691269913558, Y: 72.29160633395828}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 228.80341605412377, Y: 71.19785551123954}, Radius: 0, Start: 0, Angle: 0},
				}},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 217.02137974667383, Y: 63.410205098532714}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 217.90121674321938, Y: 62.19534425896725}, Radius: 0, Start: 0, Angle: 0},
				}},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 192.5875424011144, Y: 50.06470850220197}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 193.13419070658415, Y: 48.667863618141084}, Radius: 0, Start: 0, Angle: 0},
				}},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 179.30211952693816, Y: 45.8152007600721}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 179.66760297503276, Y: 44.360408043163986}, Radius: 0, Start: 0, Angle: 0},
				}},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 165.58573457009211, Y: 43.28112090503265}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 165.7641764051388, Y: 41.791772553737644}, Radius: 0, Start: 0, Angle: 0},
				}},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 137.74566681731258, Y: 43.49399258603107}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 137.5444713648214, Y: 42.00754703038605}, Radius: 0, Start: 0, Angle: 0},
				}},
				setColor{col: color.Gray16{Y: 0x0}},
				push{},
				rotate{angle: 4.656622921172369},
				fillString{font: "Helvetica", size: 5, x: -169.10513713195704, y: 260.82918038819696, str: "800000"},
				pop{},
				setColor{col: color.Gray16{Y: 0x0}},
				push{},
				rotate{angle: 4.149064136880077},
				fillString{font: "Helvetica", size: 5, x: -218.7136817198258, y: 164.56748322669478, str: "840000"},
				pop{},
				setColor{col: color.Gray16{Y: 0x0}},
				push{},
				rotate{angle: 3.641505352587785},
				fillString{font: "Helvetica", size: 5, x: -215.28071886824114, y: 56.32914989418752, str: "880000"},
				pop{},
				setColor{col: color.Gray16{Y: 0x0}},
				push{},
				rotate{angle: 3.1339465682954932},
				fillString{font: "Helvetica", size: 5, x: -159.67181075726256, y: -36.59539678344792, str: "920000"},
				pop{},
				setColor{col: color.Gray{Y: 0x80}},
				setWidth{w: 0.25},
				setLineDash{dashes: []vg.Length(nil), offsets: 0},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 128.34949832781146, Y: 81.49469548701578}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 118.68929765893604, Y: 53.09257368182209}, Radius: 0, Start: 0, Angle: 0},
				}},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 96.88227383745954, Y: 102.18431122891528}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 74.63518337244334, Y: 82.05803572048139}, Radius: 0, Start: 0, Angle: 0},
				}},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 79.43809660640713, Y: 135.56015724673625}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 50.21333524896998, Y: 128.78422014543077}, Radius: 0, Start: 0, Angle: 0},
				}},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 119.55797808897937, Y: 85.1217899286881}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 106.38116932457112, Y: 58.17050590016335}, Radius: 0, Start: 0, Angle: 0},
				}},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 111.29614589798454, Y: 89.83228576739197}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 94.81460425717836, Y: 64.76520007434874}, Radius: 0, Start: 0, Angle: 0},
				}},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 103.69684708461969, Y: 95.55044104193601}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 84.17558591846756, Y: 72.77061745871042}, Radius: 0, Start: 0, Angle: 0},
				}},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 90.96200042641013, Y: 109.62722765576208}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 66.34680059697416, Y: 92.47811871806692}, Radius: 0, Start: 0, Angle: 0},
				}},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 86.03122131385855, Y: 117.7595126693256}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 59.44370983940196, Y: 103.86331773705585}, Radius: 0, Start: 0, Angle: 0},
				}},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 82.1692204845432, Y: 126.45040398109401}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 54.03690867836049, Y: 116.03056557353162}, Radius: 0, Start: 0, Angle: 0},
				}},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 77.88176452108692, Y: 144.94229307173936}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 48.034470329521696, Y: 141.9192103004351}, Radius: 0, Start: 0, Angle: 0},
				}},
				setColor{col: color.Gray16{Y: 0x0}},
				setWidth{w: 1},
				setLineDash{dashes: []vg.Length(nil), offsets: 0},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 128.0231622187882, Y: 45.25782353834737}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 110, Start: 4.487993840681956, Angle: -1.367794789850083}}},
				setColor{col: color.Gray16{Y: 0x0}},
				setWidth{w: 1},
				setLineDash{dashes: []vg.Length(nil), offsets: 0},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 117.07926421412347, Y: 48.358886714289795}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 116.11324414723593, Y: 45.51867453377044}, Radius: 0, Start: 0, Angle: 0},
				}},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 70.9273349616073, Y: 78.70365646907574}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 68.70262591510568, Y: 76.69102891823235}, Radius: 0, Start: 0, Angle: 0},
				}},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 45.34254168939712, Y: 127.65489729521317}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 42.420065553653416, Y: 126.97730358508262}, Radius: 0, Start: 0, Angle: 0},
				}},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 104.18503453050306, Y: 53.67862522874255}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 103.52619409228265, Y: 52.33106102731631}, Radius: 0, Start: 0, Angle: 0},
				}},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 92.06768065037733, Y: 60.58735245884155}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 91.24360356833702, Y: 59.33399817418939}, Radius: 0, Start: 0, Angle: 0},
				}},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 80.92204239077553, Y: 68.97398019483948}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 79.94597933246793, Y: 67.8349890156782}, Radius: 0, Start: 0, Angle: 0},
				}},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 62.24426729206817, Y: 89.61993389511773}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 61.01350730059637, Y: 88.76247844823297}, Radius: 0, Start: 0, Angle: 0},
				}},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 55.01245792699254, Y: 101.54728524834422}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 53.68308235326971, Y: 100.85247550173074}, Radius: 0, Start: 0, Angle: 0},
				}},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 49.34819004399671, Y: 114.29392583893788}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 47.94157445368758, Y: 113.77293391855977}, Radius: 0, Start: 0, Angle: 0},
				}},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 43.05992129759416, Y: 141.4153631718844}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 41.567556588015904, Y: 141.2642090333192}, Radius: 0, Start: 0, Angle: 0},
				}},
				setColor{col: color.Gray16{Y: 0x0}},
				push{},
				rotate{angle: 2.8137443446786836},
				fillString{font: "Helvetica", size: 5, x: -103.61371759920122, y: -76.41731046714297, str: "740000"},
				pop{},
				setColor{col: color.Gray16{Y: 0x0}},
				push{},
				rotate{angle: 2.3061855603863917},
				fillString{font: "Helvetica", size: 5, x: 2.438524206376691, y: -98.33178158899622, str: "780000"},
				pop{},
				setColor{col: color.Gray16{Y: 0x0}},
				push{},
				rotate{angle: 1.7986267760940993},
				fillString{font: "Helvetica", size: 5, x: 105.77257147908588, y: -65.93738838923346, str: "820000"},
				pop{},
				setColor{col: color.Gray{Y: 0x80}},
				setWidth{w: 0.25},
				setLineDash{dashes: []vg.Length(nil), offsets: 0},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 82.71313575693335, Y: 179.97350685878644}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 54.798390059706705, Y: 190.96290960230104}, Radius: 0, Start: 0, Angle: 0},
				}},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 157.98095345993372, Y: 227.29945955132325}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 160.1733348439072, Y: 257.21924337185254}, Radius: 0, Start: 0, Angle: 0},
				}},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 225.54621413573972, Y: 169.50736900392474}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 254.7646997900356, Y: 176.31031660549462}, Radius: 0, Start: 0, Angle: 0},
				}},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 112.57821409335347, Y: 215.99213344993126}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 96.60949973069486, Y: 241.38898682990373}, Radius: 0, Start: 0, Angle: 0},
				}},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 201.25049174223193, Y: 209.49464496679118}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 220.7506884391247, Y: 232.29250295350766}, Radius: 0, Start: 0, Angle: 0},
				}},
				setColor{col: color.Gray16{Y: 0x0}},
				setWidth{w: 1},
				setLineDash{dashes: []vg.Length(nil), offsets: 0},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 42.88993541129061, Y: 161.75385005600137}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 110, Start: 3.057367197760077, Angle: -3.0259512712241787}}},
				setColor{col: color.Gray16{Y: 0x0}},
				setWidth{w: 1},
				setLineDash{dashes: []vg.Length(nil), offsets: 0},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 50.14593244350226, Y: 192.79447672622013}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 47.354457873779594, Y: 193.8934170005716}, Radius: 0, Start: 0, Angle: 0},
				}},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 160.53873174123612, Y: 262.2058740086074}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 160.75796987963346, Y: 265.19785239066033}, Radius: 0, Start: 0, Angle: 0},
				}},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 259.63444739908493, Y: 177.44414120575627}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 262.5562959645145, Y: 178.12443596591328}, Radius: 0, Start: 0, Angle: 0},
				}},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 93.94804733691842, Y: 245.62179572656584}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 93.1496116187855, Y: 246.89163839556446}, Radius: 0, Start: 0, Angle: 0},
				}},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 224.00072122194018, Y: 236.09214595129373}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 224.9757310567848, Y: 237.23203885062958}, Radius: 0, Start: 0, Angle: 0},
				}},
				setColor{col: color.Gray16{Y: 0x0}},
				push{},
				rotate{angle: 1.1957523765353044},
				fillString{font: "Helvetica", size: 5, x: 185.94512762064295, y: 31.029002094671917, str: "325000.00"},
				pop{},
				setColor{col: color.Gray16{Y: 0x0}},
				push{},
				rotate{angle: -0.07314458419542547},
				fillString{font: "Helvetica", size: 5, x: 129.13000209936703, y: 280.30300189893075, str: "425000.00"},
				pop{},
				setColor{col: color.Gray16{Y: 0x0}},
				push{},
				rotate{angle: -1.3420415449261553},
				fillString{font: "Helvetica", size: 5, x: -125.76327872114877, y: 300.1751144933594, str: "525000.00"},
				pop{},
			},
		},
	} {
		p, err := plot.New()
		c.Assert(err, check.Equals, nil)

		s, err := NewScale(t.feats, b, 110)
		c.Assert(err, check.Equals, nil)
		s.LineStyle = plotter.DefaultLineStyle
		s.Tick.Length = 3
		s.Tick.LineStyle = plotter.DefaultLineStyle
		s.Tick.Label = draw.TextStyle{Color: color.Gray16{0}, Font: font}
		s.Grid.LineStyle = t.grid
		s.Grid.Inner = t.inner
		s.Grid.Outer = t.outer
		p.Add(s)

		p.HideAxes()

		tc := &canvas{dpi: defaultDPI}
		p.Draw(draw.NewCanvas(tc, 300, 300))

		base.append(t.actions...)
		c.Check(tc.actions, check.DeepEquals, base.actions, check.Commentf("Test %d", i))
		if ok := reflect.DeepEqual(tc.actions, base.actions); *pics && !ok || *allPics {
			p.Add(b)
			c.Assert(p.Save(vg.Length(300), vg.Length(300), fmt.Sprintf("scale-%d-%s.svg", i, failure(!ok))), check.Equals, nil)
		}
	}
}
