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

func TestLabelsBlocks(t *testing.T) {
	rand.Seed(1)
	b, err := NewGappedBlocks(randomFeatures(3, 100000, 1000000, false, plotter.DefaultLineStyle),
		Arc{0, Complete * Clockwise},
		80, 100, 0.01,
	)
	if err != nil {
		t.Fatalf("unexpected error for NewGappedBlocks: %v", err)
	}
	font, err := vg.MakeFont("Helvetica", 10)
	if err != nil {
		t.Fatalf("unexpected error for vg.MakeFont: %v", err)
	}

	for i, test := range []struct {
		feats     []feat.Feature
		placement TextPlacement
		actions   []interface{}
	}{
		{
			feats: b.Set,
			actions: []interface{}{
				setColor{col: color.Gray16{Y: 0x0}},
				push{},
				rotate{angle: 3.830501210403824},
				fillString{font: "Helvetica", size: 10, x: -233.01155605316876, y: 84.52288965789225, str: "feature0"},
				pop{},
				setColor{col: color.Gray16{Y: 0x0}},
				push{},
				rotate{angle: 2.2333001189620183},
				fillString{font: "Helvetica", size: 10, x: 8.090496656616615, y: -108.74070751750412, str: "feature1"},
				pop{},
				setColor{col: color.Gray16{Y: 0x0}},
				push{},
				rotate{angle: -0.026404764646908996},
				fillString{font: "Helvetica", size: 10, x: 130.07341402421386, y: 261.77339238481545, str: "feature2"},
				pop{},
			},
		},
		{
			feats: b.Set[1:],
			actions: []interface{}{
				setColor{col: color.Gray16{Y: 0x0}},
				push{},
				rotate{angle: 2.2333001189620183},
				fillString{font: "Helvetica", size: 10, x: 8.090496656616615, y: -108.74070751750412, str: "feature1"},
				pop{},
				setColor{col: color.Gray16{Y: 0x0}},
				push{},
				rotate{angle: -0.026404764646908996},
				fillString{font: "Helvetica", size: 10, x: 130.07341402421386, y: 261.77339238481545, str: "feature2"},
				pop{},
			},
		},
		{
			feats:     b.Set,
			placement: Radial,
			actions: []interface{}{
				setColor{col: color.Gray16{Y: 0x0}},
				push{},
				rotate{angle: 5.401297537198721},
				fillString{font: "Helvetica", size: 10, x: 70.87542872039222, y: 209.9646810531688, str: "feature0"},
				pop{},
				setColor{col: color.Gray16{Y: 0x0}},
				push{},
				rotate{angle: 3.804096445756915},
				fillString{font: "Helvetica", size: 10, x: -122.38816845500412, y: -31.137371656616615, str: "feature1"},
				pop{},
				setColor{col: color.Gray16{Y: 0x0}},
				push{},
				rotate{angle: 1.5443915621479876},
				fillString{font: "Helvetica", size: 10, x: 248.12593144731545, y: -153.12028902421383, str: "feature2"},
				pop{},
			},
		},
		{
			feats: b.Set,
			placement: func(a Angle) (rot Angle, xalign, yalign float64) {
				return a, 0, -0.5
			},
			actions: []interface{}{
				setColor{col: color.Gray16{Y: 0x0}},
				push{},
				rotate{angle: 5.401297537198721},
				fillString{font: "Helvetica", size: 10, x: 89.22259668914222, y: 209.9646810531688, str: "feature0"},
				pop{}, setColor{col: color.Gray16{Y: 0x0}},
				push{},
				rotate{angle: 3.804096445756915},
				fillString{font: "Helvetica", size: 10, x: -104.04100048625412, y: -31.137371656616615, str: "feature1"},
				pop{},
				setColor{col: color.Gray16{Y: 0x0}},
				push{},
				rotate{angle: 1.5443915621479876},
				fillString{font: "Helvetica", size: 10, x: 266.47309941606545, y: -153.12028902421383, str: "feature2"},
				pop{},
			},
		},
		{
			feats:     b.Set,
			placement: Horizontal,
			actions: []interface{}{
				setColor{col: color.Gray16{Y: 0x0}},
				fillString{font: "Helvetica", size: 10, x: 215.74248908596053, y: 59.10442792299643, str: "feature0"},
				setColor{col: color.Gray16{Y: 0x0}},
				fillString{font: "Helvetica", size: 10, x: 32.95691195262061, y: 77.12617831374112, str: "feature1"},
				setColor{col: color.Gray16{Y: 0x0}},
				fillString{font: "Helvetica", size: 10, x: 137.54141500264052, y: 262.66014286751687, str: "feature2"},
			},
		},
	} {
		p, err := plot.New()
		if err != nil {
			t.Fatalf("unexpected error for plot.New: %v", err)
		}

		l, err := NewLabels(b, 110, NameLabels(test.feats)...)
		if err != nil {
			t.Fatalf("unexpected error for NewLabels: %v", err)
		}
		l.TextStyle = draw.TextStyle{Color: color.Gray16{0}, Font: font}
		l.Placement = test.placement
		p.Add(l)

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
			err := p.Save(vg.Length(300), vg.Length(300), fmt.Sprintf("labels-%d-%s.svg", i, failure(!ok)))
			if err != nil {
				t.Fatalf("unexpected error writing file: %v", err)
			}
		}
	}
}

func TestLabelsArcs(t *testing.T) {
	a := Arc{Theta: -0.031415926535897934, Phi: -1.7009436868899361} // This is feature0 from the blocks test.
	h := NewHighlight(
		color.NRGBA{R: 243, G: 243, B: 21, A: 128},
		a,
		80, 100,
	)
	h.LineStyle = plotter.DefaultLineStyle

	font, err := vg.MakeFont("Helvetica", 10)
	if err != nil {
		t.Fatalf("unexpected error for vg.MakeFont: %v", err)
	}

	for i, test := range []struct {
		arc       Arcer
		label     Label
		placement TextPlacement
		actions   []interface{}
	}{
		{
			arc:   a,
			label: "Label",
			actions: []interface{}{
				setColor{col: color.Gray16{Y: 0x0}},
				push{},
				rotate{angle: -2.4526840967757626},
				fillString{font: "Helvetica", size: 10, x: -226.8982748031688, y: 84.52288965789216, str: "Label"},
				pop{},
			},
		},
		{
			arc:   h,
			label: "Label",
			actions: []interface{}{
				setColor{col: color.Gray16{Y: 0x0}},
				push{},
				rotate{angle: -2.4526840967757626},
				fillString{font: "Helvetica", size: 10, x: -226.8982748031688, y: 84.52288965789216, str: "Label"},
				pop{},
			},
		},
		{
			arc:       h,
			label:     "Label",
			placement: Radial,
			actions: []interface{}{
				setColor{col: color.Gray16{Y: 0x0}},
				push{},
				rotate{angle: -0.881887769980866},
				fillString{font: "Helvetica", size: 10, x: 76.98870997039216, y: 209.9646810531688, str: "Label"},
				pop{},
			},
		},
		{
			arc:   h,
			label: "Label",
			placement: func(a Angle) (rot Angle, xalign, yalign float64) {
				return a, 0, -0.5
			},
			actions: []interface{}{
				setColor{col: color.Gray16{Y: 0x0}},
				push{},
				rotate{angle: -0.881887769980866},
				fillString{font: "Helvetica", size: 10, x: 89.22259668914216, y: 209.9646810531688, str: "Label"},
				pop{},
			},
		},
		{
			arc:       h,
			label:     "Label",
			placement: Horizontal,
			actions: []interface{}{
				setColor{col: color.Gray16{Y: 0x0}},
				fillString{font: "Helvetica", size: 10, x: 217.96958781584493, y: 59.10442792299643, str: "Label"},
			},
		},
	} {
		p, err := plot.New()
		if err != nil {
			t.Fatalf("unexpected error for plot.New: %v", err)
		}

		l, err := NewLabels(test.arc, 110, test.label)
		if err != nil {
			t.Fatalf("unexpected error for NewLabels: %v", err)
		}
		l.TextStyle = draw.TextStyle{Color: color.Gray16{0}, Font: font}
		l.Placement = test.placement
		p.Add(l)

		p.HideAxes()

		tc := &canvas{dpi: defaultDPI}
		p.Draw(draw.NewCanvas(tc, 300, 300))

		base.append(test.actions...)
		ok := reflect.DeepEqual(tc.actions, base.actions)
		if !ok {
			t.Errorf("unexpected actions for test %d:\ngot :%#v\nwant:%#v", i, tc.actions, base.actions)
		}
		if *pics && !ok || *allPics {
			p.Add(h)
			err := p.Save(vg.Length(300), vg.Length(300), fmt.Sprintf("labelsarcs-%d-%s.svg", i, failure(!ok)))
			if err != nil {
				t.Fatalf("unexpected error writing file: %v", err)
			}
		}
	}
}

func TestLabelSpokes(t *testing.T) {
	p, err := plot.New()
	if err != nil {
		t.Fatalf("unexpected error for plot.New: %v", err)
	}

	rand.Seed(1)
	b, err := NewGappedBlocks(randomFeatures(3, 100000, 1000000, false, plotter.DefaultLineStyle),
		Arc{0, Complete * Clockwise},
		80, 100, 0.01,
	)
	if err != nil {
		t.Fatalf("unexpected error for NewGappedBlocks: %v", err)
	}

	m := randomFeatures(10, b.Set[1].Start(), b.Set[1].End(), true, plotter.DefaultLineStyle)
	for _, mf := range m {
		mf.(*fs).location = b.Set[1]
	}
	ms, err := NewSpokes(m, b, 73, 78)
	if err != nil {
		t.Fatalf("unexpected error for NewSpokes: %v", err)
	}
	ms.LineStyle = plotter.DefaultLineStyle

	font, err := vg.MakeFont("Helvetica", 10)
	if err != nil {
		t.Fatalf("unexpected error for vg.MakeFont: %v", err)
	}

	l, err := NewLabels(ms, 125, NameLabels([]feat.Feature{m[1], m[5], m[9]})...)
	if err != nil {
		t.Fatalf("unexpected error for NewLabels: %v", err)
	}
	l.TextStyle = draw.TextStyle{Color: color.Gray16{0}, Font: font}
	l.Placement = Radial
	p.Add(l)

	p.HideAxes()

	tc := &canvas{dpi: defaultDPI}
	p.Draw(draw.NewCanvas(tc, 300, 300))

	base.append(
		setColor{col: color.Gray16{Y: 0x0}},
		push{},
		rotate{angle: 4.2234542023088135},
		fillString{font: "Helvetica", size: 10, x: -99.60637951470011, y: 58.305353499178814, str: "feature1"},
		pop{},
		setColor{col: color.Gray16{Y: 0x0}},
		push{},
		rotate{angle: 3.909820940524999},
		fillString{font: "Helvetica", size: 10, x: -108.98294710980127, y: -8.402510595005197, str: "feature5"},
		pop{},
		setColor{col: color.Gray16{Y: 0x0}},
		push{},
		rotate{angle: 4.113415457874245},
		fillString{font: "Helvetica", size: 10, x: -105.27790643446565, y: 35.27356020294578, str: "feature9"},
		pop{})
	ok := reflect.DeepEqual(tc.actions, base.actions)
	if !ok {
		t.Errorf("unexpected actions:\ngot :%#v\nwant:%#v", tc.actions, base.actions)
	}
	if *pics && !ok || *allPics {
		err := p.Save(vg.Length(300), vg.Length(300), fmt.Sprintf("labelspokes-%s.svg", failure(!ok)))
		if err != nil {
			t.Fatalf("unexpected error writing file: %v", err)
		}
	}
}
