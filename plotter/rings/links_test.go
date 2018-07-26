// Copyright ©2018 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rings

import (
	"fmt"
	"math/rand"
	"testing"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
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
		ends   [2]Feature
		bezier *Bezier
	}{
		{
			ends: [2]Feature{b.Set[1], b.Set[1]},
			bezier: &Bezier{Segments: 5,
				Radius: LengthDist{Length: 2 * 70 / 3, Min: floatPtr(0.95), Max: floatPtr(1.05)},
				Crest:  &FactorDist{Factor: 2, Min: floatPtr(0.7), Max: floatPtr(1.4)},
			},
		},
		{
			ends: [2]Feature{b.Set[0], b.Set[1]},
			bezier: &Bezier{Segments: 5,
				Radius: LengthDist{Length: 2 * 70 / 3, Min: floatPtr(0.95), Max: floatPtr(1.05)},
				Crest:  &FactorDist{Factor: 2, Min: floatPtr(0.7), Max: floatPtr(1.4)},
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
			m[0][j].(*fs).parent = test.ends[0]
			m[1][j].(*fs).parent = test.ends[1]
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
		p.Add(b)

		checkImage(t, fmt.Sprintf("links-%d", i), p, *regen)
	}
}
