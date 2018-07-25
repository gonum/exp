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

	"github.com/biogo/biogo/feat"
)

func TestScoresAxis(t *testing.T) {
	rand.Seed(1)
	b, err := NewGappedBlocks(randomFeatures(3, 100000, 1000000, false, plotter.DefaultLineStyle),
		Arc{0, Complete * Clockwise},
		80, 100, 0.01,
	)
	if err != nil {
		t.Fatalf("unexpected error for NewGappedBlocks: %v", err)
	}
	font, err := vg.MakeFont("Helvetica", 5)
	if err != nil {
		t.Fatalf("unexpected error for vg.MakeFont: %v", err)
	}

	for i, test := range []struct {
		orient   feat.Orientation
		scores   []Scorer
		renderer ScoreRenderer
		actions  []interface{}
	}{
		{
			scores: makeScorers(b.Set[1].(*fs), 10, 1, func(v, _ int) float64 { return float64(v) }),
			renderer: &Trace{
				LineStyles: []draw.LineStyle{func() draw.LineStyle {
					sty := plotter.DefaultLineStyle
					sty.Color = color.Gray{0}
					return sty
				}()},
				Axis: func() *Axis {
					a, err := b.ArcOf(b.Set[1], nil)
					if err != nil {
						t.Fatalf("unexpected error for ArcOf: %v", err)
					}
					return &Axis{
						Angle:     a.Theta + a.Phi - Complete*0.01/2,
						Grid:      plotter.DefaultGridLineStyle,
						LineStyle: plotter.DefaultLineStyle,
						Label: AxisLabel{
							Text:      "Test",
							TextStyle: draw.TextStyle{Color: color.Gray16{0}, Font: font},
						},
						Tick: TickConfig{
							Marker:    plot.DefaultTicks{},
							LineStyle: plotter.DefaultLineStyle,
							Length:    -2,
							Label:     draw.TextStyle{Color: color.Gray16{0}, Font: font},
						},
					}
				}(),
			},
			actions: []interface{}{
				setColor{col: color.Gray{Y: 0x80}},
				setWidth{w: 0.25},
				setLineDash{dashes: []vg.Length(nil), offsets: 0},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 143.59933171592297, Y: 113.50284492303541}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 40, Start: 4.487993840681956, Angle: -1.367794789850083},
				}},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 140.1379607165597, Y: 98.33728461532695}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 55.55555555555556, Start: 4.487993840681956, Angle: -1.367794789850083},
				}},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 136.6765897171964, Y: 83.17172430761849}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 71.11111111111111, Start: 4.487993840681956, Angle: -1.367794789850083},
				}},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 142.73398896608217, Y: 109.7114548461083}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 43.888888888888886, Start: 4.487993840681956, Angle: -1.367794789850083},
				}},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 141.86864621624136, Y: 105.92006476918118}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 47.77777777777778, Start: 4.487993840681956, Angle: -1.367794789850083},
				}},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 141.00330346640052, Y: 102.12867469225407}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 51.666666666666664, Start: 4.487993840681956, Angle: -1.367794789850083},
				}},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 139.27261796671888, Y: 94.54589453839984}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 59.44444444444444, Start: 4.487993840681956, Angle: -1.367794789850083},
				}},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 138.40727521687808, Y: 90.75450446147272}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 63.33333333333333, Start: 4.487993840681956, Angle: -1.367794789850083},
				}},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 137.54193246703724, Y: 86.9631143845456}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 67.22222222222223, Start: 4.487993840681956, Angle: -1.367794789850083},
				}},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 135.8112469673556, Y: 79.38033423069139}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 75, Start: 4.487993840681956, Angle: -1.367794789850083},
				}},
				setColor{col: color.Gray16{Y: 0x0}},
				setWidth{w: 1},
				setLineDash{dashes: []vg.Length(nil), offsets: 0},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 112.55576396611607, Y: 154.6113994575564}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 77.60455743646764, Y: 156.45887398291822}, Radius: 0, Start: 0, Angle: 0},
				}},
				setColor{col: color.Gray16{Y: 0x0}},
				setWidth{w: 1},
				setLineDash{dashes: []vg.Length(nil), offsets: 0},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 112.55576396611607, Y: 154.6113994575564}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 112.66133393899389, Y: 156.60861125925058}, Radius: 0, Start: 0, Angle: 0},
				}},
				setColor{col: color.Gray16{Y: 0x0}},
				push{},
				rotate{angle: 1.517986797501079},
				fillString{font: "Helvetica", size: 5, x: 162.94672945174116, y: -106.58754296287375, str: "0"},
				pop{},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 97.02189439738343, Y: 155.43249924660608}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 97.12746437026125, Y: 157.42971104830028}, Radius: 0, Start: 0, Angle: 0},
				}},
				setColor{col: color.Gray16{Y: 0x0}},
				push{},
				rotate{angle: 1.517986797501079},
				fillString{font: "Helvetica", size: 5, x: 162.94672945174113, y: -91.03198740731821, str: "4"},
				pop{},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 81.4880248286508, Y: 156.25359903565578}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 81.59359480152861, Y: 158.25081083734997}, Radius: 0, Start: 0, Angle: 0},
				}},
				setColor{col: color.Gray16{Y: 0x0}},
				push{},
				rotate{angle: 1.517986797501079},
				fillString{font: "Helvetica", size: 5, x: 162.9467294517412, y: -75.47643185176261, str: "8"},
				pop{},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 108.67229657393293, Y: 154.8166744048188}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 108.72508156037182, Y: 155.8152803056659}, Radius: 0, Start: 0, Angle: 0},
				}},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 104.78882918174975, Y: 155.02194935208124}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 104.84161416818867, Y: 156.02055525292835}, Radius: 0, Start: 0, Angle: 0},
				}},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 100.9053617895666, Y: 155.22722429934367}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 100.9581467760055, Y: 156.22583020019076}, Radius: 0, Start: 0, Angle: 0},
				}},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 93.13842700520027, Y: 155.63777419386852}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 93.1912119916392, Y: 156.6363800947156}, Radius: 0, Start: 0, Angle: 0},
				}},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 89.25495961301712, Y: 155.84304914113093}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 89.30774459945604, Y: 156.84165504197804}, Radius: 0, Start: 0, Angle: 0},
				}},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 85.37149222083396, Y: 156.04832408839337}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 85.42427720727287, Y: 157.04692998924045}, Radius: 0, Start: 0, Angle: 0},
				}},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 77.60455743646764, Y: 156.45887398291822}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 77.65734242290655, Y: 157.4574798837653}, Radius: 0, Start: 0, Angle: 0},
				}},
				setColor{col: color.Gray16{Y: 0x0}},
				push{},
				rotate{angle: 1.517986797501079},
				fillString{font: "Helvetica", size: 5, x: 155.75214937361616, y: -89.08754296287375, str: "Test"},
				pop{},
				setColor{col: color.Gray{Y: 0x0}},
				setWidth{w: 1},
				setLineDash{dashes: []vg.Length(nil), offsets: 0},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 77.5170812433199, Y: 154.10059199332972}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 75, Start: 3.1202498067103024, Angle: 0.1367744033971654},
				}},
				setColor{col: color.Gray{Y: 0x0}},
				setWidth{w: 1},
				setLineDash{dashes: []vg.Length(nil), offsets: 0},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 81.86212109304435, Y: 144.30975049323513}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 71.11111111111111, Start: 3.257024210107468, Angle: 0.1367744033971654},
				}},
				setColor{col: color.Gray{Y: 0x0}},
				setWidth{w: 1},
				setLineDash{dashes: []vg.Length(nil), offsets: 0},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 87.4043997360034, Y: 135.7253170890239}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 67.22222222222223, Start: 3.393798613504633, Angle: 0.1367744033971655},
				}},
				setColor{col: color.Gray{Y: 0x0}},
				setWidth{w: 1},
				setLineDash{dashes: []vg.Length(nil), offsets: 0},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 93.89790524006592, Y: 128.4811407252797}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 63.33333333333333, Start: 3.5305730169017986, Angle: 0.1367744033971653},
				}},
				setColor{col: color.Gray{Y: 0x0}},
				setWidth{w: 1},
				setLineDash{dashes: []vg.Length(nil), offsets: 0},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 101.0838014425834, Y: 122.66685565997186}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 59.44444444444444, Start: 3.667347420298964, Angle: 0.1367744033971654},
				}},
				setColor{col: color.Gray{Y: 0x0}},
				setWidth{w: 1},
				setLineDash{dashes: []vg.Length(nil), offsets: 0},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 108.6978316852225, Y: 118.32706035456823}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 55.55555555555556, Start: 3.8041218236961294, Angle: 0.1367744033971654},
				}},
				setColor{col: color.Gray{Y: 0x0}},
				setWidth{w: 1},
				setLineDash{dashes: []vg.Length(nil), offsets: 0},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 116.47768348270975, Y: 115.46167988187636}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 51.666666666666664, Start: 3.940896227093295, Angle: 0.1367744033971654},
				}},
				setColor{col: color.Gray{Y: 0x0}},
				setWidth{w: 1},
				setLineDash{dashes: []vg.Length(nil), offsets: 0},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 124.17013120108297, Y: 114.02748275432367}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 47.77777777777778, Start: 4.07767063049046, Angle: 0.1367744033971654},
				}},
				setColor{col: color.Gray{Y: 0x0}},
				setWidth{w: 1},
				setLineDash{dashes: []vg.Length(nil), offsets: 0},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 131.53778143549647, Y: 113.94069553589927}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 43.888888888888886, Start: 4.214445033887626, Angle: 0.1367744033971654},
				}},
				setColor{col: color.Gray{Y: 0x0}},
				setWidth{w: 1},
				setLineDash{dashes: []vg.Length(nil), offsets: 0},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 138.36525744911535, Y: 115.0806326480496}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 40, Start: 4.351219437284791, Angle: 0.1367744033971654},
				}},
			},
		},
	} {
		p, err := plot.New()
		if err != nil {
			t.Fatalf("unexpected error for plot.New: %v", err)
		}

		b.Set[1].(*fs).orient = test.orient

		r, err := NewScores(test.scores, b, 40, 75, test.renderer)
		if err != nil {
			t.Fatalf("unexpected error for NewScores: %v", err)
		}
		p.Add(r)

		p.HideAxes()

		tc := &canvas{dpi: defaultDPI}
		p.Draw(draw.NewCanvas(tc, 300, 300))

		base.append(test.actions...)
		ok := reflect.DeepEqual(tc.actions, base.actions)
		if !ok {
			t.Errorf("unexpected actions for test %d:\ngot :%#v\nwant:%#v", i, tc.actions, base.actions)
		}
		if *pics && !ok || *allPics {
			p.Add(b)
			err := p.Save(vg.Length(300), vg.Length(300), fmt.Sprintf("axis-%d-%s.svg", i, failure(!ok)))
			if err != nil {
				t.Fatalf("unexpected error writing file: %v", err)
			}
		}
	}
}
