// Copyright ©2018 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rings

import (
	"fmt"
	"image/color"
	"testing"

	"golang.org/x/exp/rand"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/font"
	"gonum.org/v1/plot/font/liberation"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg/draw"
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
	cache := font.NewCache(liberation.Collection())
	fnt := cache.Lookup(font.Font{Typeface: "Liberation", Variant: "Sans"}, 5)

	for i, test := range []struct {
		orient   Orientation
		scores   []Scorer
		renderer ScoreRenderer
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
							TextStyle: draw.TextStyle{Color: color.Gray16{0}, Font: fnt.Font},
						},
						Tick: TickConfig{
							Marker:    plot.DefaultTicks{},
							LineStyle: plotter.DefaultLineStyle,
							Length:    -2,
							Label:     draw.TextStyle{Color: color.Gray16{0}, Font: fnt.Font},
						},
					}
				}(),
			},
		},
	} {
		t.Run(fmt.Sprintf("axis-%d", i), func(t *testing.T) {
			p := plot.New()
			b.Set[1].(*fs).orient = test.orient

			r, err := NewScores(test.scores, b, 40, 75, test.renderer)
			if err != nil {
				t.Fatalf("unexpected error for NewScores: %v", err)
			}

			p.Add(r)
			p.HideAxes()
			p.Add(b)

			checkImage(t, p, *regen)
		})
	}
}
