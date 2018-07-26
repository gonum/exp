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

func TestLinks(t *testing.T) {
	const marks = 16

	rand.Seed(1)
	b, err := NewGappedBlocks(randomFeatures(3, 100000, 1000000, false, plotter.DefaultLineStyle),
		Arc{0, Complete * Clockwise},
		80, 100, 0.01,
	)
	if err != nil {
		t.Fatalf("unexpected error for NewGappedBlocks: %v", err)
	}

	for i, test := range []struct {
		ends    [2]Feature
		bezier  *Bezier
		actions []interface{}
	}{
		{
			ends: [2]Feature{b.Set[1], b.Set[1]},
			bezier: &Bezier{Segments: 5,
				Radius: LengthDist{Length: 2 * 70 / 3, Min: floatPtr(0.95), Max: floatPtr(1.05)},
				Crest:  &FactorDist{Factor: 2, Min: floatPtr(0.7), Max: floatPtr(1.4)},
			},
			actions: []interface{}{
				setColor{col: color.RGBA{R: 0x0, G: 0x0, B: 0x0, A: 0xfe}},
				setWidth{w: 1},
				setLineDash{dashes: []vg.Length(nil), offsets: 0},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 82.94890965035401, Y: 144.58508805005508}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 110.88673705784406, Y: 145.6911355360719}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 119.90877782857362, Y: 142.69011417010475}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 121.7205646001176, Y: 137.9725095392452}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 117.02660854266361, Y: 129.7038865923218}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 95.53039935901188, Y: 111.82488963989985}, Radius: 0, Start: 0, Angle: 0},
				}},
				setColor{col: color.RGBA{R: 0x0, G: 0x0, B: 0x0, A: 0xfe}},
				setWidth{w: 1},
				setLineDash{dashes: []vg.Length(nil), offsets: 0},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 101.00146627327814, Y: 105.08796540963792}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 124.33180539850362, Y: 123.75628728089062}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 133.53600800602896, Y: 126.9867250450579}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 137.6146477184416, Y: 124.41109687567388}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 138.6539379248489, Y: 114.71197579923606}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 131.82573178087796, Y: 85.6226896952053}, Radius: 0, Start: 0, Angle: 0},
				}},
				setColor{col: color.RGBA{R: 0x0, G: 0x0, B: 0x0, A: 0xfe}},
				setWidth{w: 1},
				setLineDash{dashes: []vg.Length(nil), offsets: 0},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 86.06673359411869, Y: 130.44050964674835}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 114.30424672198936, Y: 137.24853899180263}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 124.78415298682266, Y: 135.76155186537417}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 128.4370331400469, Y: 130.83626423005006}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 126.81793871625356, Y: 120.37595356159008}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 112.10639203319728, Y: 95.33045884892681}, Radius: 0, Start: 0, Angle: 0},
				}},
				setColor{col: color.RGBA{R: 0x0, G: 0x0, B: 0x0, A: 0xfe}},
				setWidth{w: 1},
				setLineDash{dashes: []vg.Length(nil), offsets: 0},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 112.49323426783418, Y: 95.05908517744862}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 128.70102082222166, Y: 119.19376569198872}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 132.80841664755707, Y: 126.77283659426749}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 131.84970657102377, Y: 127.53581018481371}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 125.3860716832876, Y: 121.83191324975331}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 105.50558933849673, Y: 100.62008706080948}, Radius: 0, Start: 0, Angle: 0},
				}},
				setColor{col: color.RGBA{R: 0x0, G: 0x0, B: 0x0, A: 0xfe}},
				setWidth{w: 1},
				setLineDash{dashes: []vg.Length(nil), offsets: 0},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 93.69240194313716, Y: 114.5307175366394}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 114.65191354102839, Y: 125.24672151270721}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 125.03898144677848, Y: 125.76552208275317}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 130.4340376495267, Y: 121.16418734043059}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 131.56088669456625, Y: 110.82539897138867}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 124.28670568334428, Y: 88.43745225327233}, Radius: 0, Start: 0, Angle: 0},
				}},
				setColor{col: color.RGBA{R: 0x0, G: 0x0, B: 0x0, A: 0xfe}},
				setWidth{w: 1},
				setLineDash{dashes: []vg.Length(nil), offsets: 0},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 83.8277514120339, Y: 138.93083370753428}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 108.59239399980017, Y: 144.0411074755302}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 116.44898471788231, Y: 146.05546682642733}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 116.34896987113308, Y: 146.6424637159879}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 108.26861079817078, Y: 145.941422662984}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 83.20898387137919, Y: 142.56244074919752}, Radius: 0, Start: 0, Angle: 0},
				}},
				setColor{col: color.RGBA{R: 0x0, G: 0x0, B: 0x0, A: 0xfe}},
				setWidth{w: 1},
				setLineDash{dashes: []vg.Length(nil), offsets: 0},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 82.89555477360024, Y: 145.068902858587}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 113.55690662695783, Y: 144.1072053579495}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 126.11507706265655, Y: 137.66374229988656}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 132.79547322189862, Y: 129.52152473716296}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 136.6616641233423, Y: 115.94660562925395}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 131.6153806631019, Y: 85.68807984234503}, Radius: 0, Start: 0, Angle: 0},
				}},
				setColor{col: color.RGBA{R: 0x0, G: 0x0, B: 0x0, A: 0xfe}},
				setWidth{w: 1},
				setLineDash{dashes: []vg.Length(nil), offsets: 0},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 134.33271834067673, Y: 84.89859559808822}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 139.52290627403312, Y: 111.54966962519853}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 138.1000982117343, Y: 121.14179149445158}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 133.6670957824232, Y: 123.58605646196415}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 124.79602912710473, Y: 119.66976215054652}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 105.02835689914588, Y: 101.05640854970244}, Radius: 0, Start: 0, Angle: 0},
				}},
			},
		},
		{
			ends: [2]Feature{b.Set[0], b.Set[1]},
			bezier: &Bezier{Segments: 5,
				Radius: LengthDist{Length: 2 * 70 / 3, Min: floatPtr(0.95), Max: floatPtr(1.05)},
				Crest:  &FactorDist{Factor: 2, Min: floatPtr(0.7), Max: floatPtr(1.4)},
			},
			actions: []interface{}{
				setColor{col: color.RGBA{R: 0x0, G: 0x0, B: 0x0, A: 0xfe}},
				setWidth{w: 1},
				setLineDash{dashes: []vg.Length(nil), offsets: 0},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 220.6937479573146, Y: 136.70086269651873}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 187.27574270379353, Y: 137.25961633395113}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 166.15903174864417, Y: 130.9097077221543}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 148.13498972152996, Y: 127.32746428700274}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 126.19498892990453, Y: 125.11993485232419}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 95.53039935901188, Y: 111.82488963989985}, Radius: 0, Start: 0, Angle: 0},
				}},
				setColor{col: color.RGBA{R: 0x0, G: 0x0, B: 0x0, A: 0xfe}},
				setWidth{w: 1},
				setLineDash{dashes: []vg.Length(nil), offsets: 0},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 217.42031968199183, Y: 126.32076982820193}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 186.00019989187024, Y: 133.32053939223022}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 170.19552498423474, Y: 129.88602813395457}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 158.86972436475185, Y: 124.50089395100164}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 146.22968107268173, Y: 114.41065499972504}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 131.82573178087796, Y: 85.6226896952053}, Radius: 0, Start: 0, Angle: 0},
				}},
				setColor{col: color.RGBA{R: 0x0, G: 0x0, B: 0x0, A: 0xfe}},
				setWidth{w: 1},
				setLineDash{dashes: []vg.Length(nil), offsets: 0},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 159.5939989492541, Y: 82.8603907326586}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 153.92941429324918, Y: 111.49485043162561}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 147.70416599618704, Y: 120.67802902176378}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 141.0425372642429, Y: 122.42734783244794}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 131.1086354461495, Y: 117.48750167624756}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 112.10639203319728, Y: 95.33045884892681}, Radius: 0, Start: 0, Angle: 0},
				}},
				setColor{col: color.RGBA{R: 0x0, G: 0x0, B: 0x0, A: 0xfe}},
				setWidth{w: 1},
				setLineDash{dashes: []vg.Length(nil), offsets: 0},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 222.21436576390906, Y: 146.1827849382893}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 188.2180668286317, Y: 142.7267415391919}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 168.89416473245714, Y: 134.14002758002692}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 152.88163388545502, Y: 127.88879157549177}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 132.85124984954496, Y: 121.11172995462557}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 105.50558933849673, Y: 100.62008706080948}, Radius: 0, Start: 0, Angle: 0},
				}},
				setColor{col: color.RGBA{R: 0x0, G: 0x0, B: 0x0, A: 0xfe}},
				setWidth{w: 1},
				setLineDash{dashes: []vg.Length(nil), offsets: 0},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 176.32705020754776, Y: 86.6800054815636}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 165.39233488090187, Y: 107.93519719859239}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 155.82445738512044, Y: 114.91667264848387}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 146.64756661804992, Y: 115.22658406817885}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 136.63048444192012, Y: 108.90650939232411}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 124.28670568334428, Y: 88.43745225327233}, Radius: 0, Start: 0, Angle: 0},
				}},
				setColor{col: color.RGBA{R: 0x0, G: 0x0, B: 0x0, A: 0xfe}},
				setWidth{w: 1},
				setLineDash{dashes: []vg.Length(nil), offsets: 0},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 219.24142964767108, Y: 131.39119689359507}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 186.84454583290164, Y: 134.3999975821199}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 161.85475786820874, Y: 131.6380158146877}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 139.86707538111085, Y: 133.44368623016248}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 115.66269170344745, Y: 140.24558718419001}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 83.20898387137919, Y: 142.56244074919752}, Radius: 0, Start: 0, Angle: 0},
				}},
				setColor{col: color.RGBA{R: 0x0, G: 0x0, B: 0x0, A: 0xfe}},
				setWidth{w: 1},
				setLineDash{dashes: []vg.Length(nil), offsets: 0},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 190.46661300037744, Y: 93.69067848308764}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 171.54730266177307, Y: 116.94666453897386}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 160.7093692836963, Y: 122.34880957245115}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 152.63976847345785, Y: 121.2515041480676}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 143.63785457384108, Y: 113.15153396205824}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 131.6153806631019, Y: 85.68807984234503}, Radius: 0, Start: 0, Angle: 0},
				}},
				setColor{col: color.RGBA{R: 0x0, G: 0x0, B: 0x0, A: 0xfe}},
				setWidth{w: 1},
				setLineDash{dashes: []vg.Length(nil), offsets: 0},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 185.07819022344938, Y: 90.54306720176712}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 167.85681139241447, Y: 113.63280001983611}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 154.46407158138777, Y: 120.55789476680084}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 142.3545736044188, Y: 122.1482951523736}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 127.62784877361577, Y: 118.91626906056871}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 105.02835689914588, Y: 101.05640854970244}, Radius: 0, Start: 0, Angle: 0},
				}},
			},
		},
	} {
		p, err := plot.New()
		if err != nil {
			t.Fatalf("unexpected error for plot.New: %v", err)
		}

		var m [2][]Feature
		rand.Seed(2)
		for j := range m {
			m[j] = randomFeatures(marks/2, test.ends[j].Start(), test.ends[j].End(), true, plotter.DefaultLineStyle)
		}
		mp := make([]Pair, marks/2)
		for j := range mp {
			m[0][j].(*fs).location = test.ends[0]
			m[1][j].(*fs).location = test.ends[1]
			mp[j] = fp{feats: [2]*fs{m[0][j].(*fs), m[1][j].(*fs)}, sty: plotter.DefaultLineStyle}
		}
		l, err := NewLinks(mp, [2]ArcOfer{b, b}, [2]vg.Length{70, 70})
		if err != nil {
			t.Fatalf("unexpected error for NewLinks: %v", err)
		}
		l.Bezier = test.bezier
		l.LineStyle = plotter.DefaultLineStyle
		p.Add(l)

		p.HideAxes()

		tc := &canvas{dpi: defaultDPI}
		p.Draw(draw.NewCanvas(tc, 300, 300))

		base.append(test.actions...)
		ok := reflect.DeepEqual(tc.actions, base.actions)
		if !ok {
			t.Errorf("unexpected actions for test %d:\ngot :%#v\nwant:%#v", i, tc.actions, base.actions)
		}
		p.Add(b)
		checkImage(t, fmt.Sprintf("links-%d", i), p, *allPics)
		if *pics && !ok || *allPics {
			err = p.Save(vg.Length(300), vg.Length(300), fmt.Sprintf("links-%d-%s.svg", i, failure(!ok)))
			if err != nil {
				t.Fatalf("unexpected error writing file: %v", err)
			}
		}
	}
}
