// Copyright ©2018 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rings

import (
	"fmt"
	"image/color"
	"math/rand/v2"
	"testing"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/palette"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg/draw"
)

func TestScores(t *testing.T) {
	rnd := rand.New(rand.NewPCG(1, 1))
	b, err := NewGappedBlocks(randomFeatures(rnd, 3, 100000, 1000000, false, plotter.DefaultLineStyle),
		Arc{0, Complete * Clockwise},
		80, 100, 0.01,
	)
	if err != nil {
		t.Fatalf("unexpected error for NewGappedBlocks: %v", err)
	}

	for i, test := range []struct {
		orient   Orientation
		scores   []Scorer
		renderer ScoreRenderer
	}{
		{
			scores:   makeScorers(b.Set[1].(*fs), 10, 5, func(i, j int) float64 { return float64(i * j) }),
			renderer: &Heat{Palette: palette.Radial(10, palette.Cyan, palette.Magenta, 1).Colors()},
		},
		{
			scores:   makeScorers(b.Set[1].(*fs), 10, 5, func(_, _ int) float64 { return rnd.NormFloat64() }),
			renderer: &Heat{Palette: palette.Radial(10, palette.Cyan, palette.Magenta, 1).Colors()},
		},
		{
			scores: makeScorers(b.Set[1].(*fs), 10, 1, func(v, _ int) float64 { return float64(v) }),
			renderer: &Trace{
				LineStyles: []draw.LineStyle{func() draw.LineStyle {
					sty := plotter.DefaultLineStyle
					sty.Color = color.Gray{0}
					return sty
				}()},
			},
		},
		{
			scores: makeScorers(b.Set[1].(*fs), 10, 1, func(v, _ int) float64 { return float64(v) }),
			renderer: &Trace{
				LineStyles: []draw.LineStyle{func() draw.LineStyle {
					sty := plotter.DefaultLineStyle
					sty.Color = color.Gray{0}
					return sty
				}()},
				Join: true,
			},
		},
		{
			orient: Forward,
			scores: []Scorer{
				&fs{
					start:  b.Set[1].Start(),
					end:    b.Set[1].Start() + lengthOf(b.Set[1])/3,
					name:   fmt.Sprintf("%s#0", b.Set[1].Name()),
					parent: b.Set[1],
					scores: []float64{1},
				},
				&fs{
					start:  b.Set[1].Start() + lengthOf(b.Set[1])/3,
					end:    b.Set[1].Start() + lengthOf(b.Set[1])/2,
					name:   fmt.Sprintf("%s#1", b.Set[1].Name()),
					parent: b.Set[1],
					scores: []float64{3},
				},
				&fs{
					start:  b.Set[1].Start() + lengthOf(b.Set[1])/2,
					end:    b.Set[1].End(),
					name:   fmt.Sprintf("%s#2", b.Set[1].Name()),
					parent: b.Set[1],
					scores: []float64{2},
				},
			},
			renderer: &Trace{
				LineStyles: []draw.LineStyle{func() draw.LineStyle {
					sty := plotter.DefaultLineStyle
					sty.Color = color.Gray{0}
					return sty
				}()},
				Join: true,
			},
		},
		{
			orient: Forward,
			scores: []Scorer{
				&fs{
					start:  b.Set[1].Start(),
					end:    b.Set[1].Start() + lengthOf(b.Set[1])/3,
					name:   fmt.Sprintf("%s#0", b.Set[1].Name()),
					parent: b.Set[1],
					scores: []float64{1},
				},
				&fs{
					start:  b.Set[1].Start() + lengthOf(b.Set[1])/3 + 1,
					end:    b.Set[1].Start() + lengthOf(b.Set[1])/2,
					name:   fmt.Sprintf("%s#1", b.Set[1].Name()),
					parent: b.Set[1],
					scores: []float64{3},
				},
				&fs{
					start:  b.Set[1].Start() + lengthOf(b.Set[1])/2,
					end:    b.Set[1].End(),
					name:   fmt.Sprintf("%s#2", b.Set[1].Name()),
					parent: b.Set[1],
					scores: []float64{2},
				},
			},
			renderer: &Trace{
				LineStyles: []draw.LineStyle{func() draw.LineStyle {
					sty := plotter.DefaultLineStyle
					sty.Color = color.Gray{0}
					return sty
				}()},
				Join: true,
			},
		},
		{
			orient: Forward,
			scores: []Scorer{
				&fs{
					start:  b.Set[1].Start(),
					end:    b.Set[1].Start() + lengthOf(b.Set[1])/3,
					name:   fmt.Sprintf("%s#0", b.Set[1].Name()),
					parent: b.Set[1],
					scores: []float64{1},
				},
				&fs{
					start:  b.Set[1].Start() + lengthOf(b.Set[1])/3,
					end:    b.Set[1].Start() + lengthOf(b.Set[1])/2,
					name:   fmt.Sprintf("%s#1", b.Set[1].Name()),
					parent: b.Set[1],
					scores: []float64{3},
				},
				&fs{
					start:  b.Set[1].Start() + lengthOf(b.Set[1])/2 + 1,
					end:    b.Set[1].End(),
					name:   fmt.Sprintf("%s#2", b.Set[1].Name()),
					parent: b.Set[1],
					scores: []float64{2},
				},
			},
			renderer: &Trace{
				LineStyles: []draw.LineStyle{func() draw.LineStyle {
					sty := plotter.DefaultLineStyle
					sty.Color = color.Gray{0}
					return sty
				}()},
				Join: true,
			},
		},
		{
			orient: Backward,
			scores: []Scorer{
				&fs{
					start:  b.Set[1].Start(),
					end:    b.Set[1].Start() + lengthOf(b.Set[1])/3,
					name:   fmt.Sprintf("%s#0", b.Set[1].Name()),
					parent: b.Set[1],
					scores: []float64{1},
				},
				&fs{
					start:  b.Set[1].Start() + lengthOf(b.Set[1])/3,
					end:    b.Set[1].Start() + lengthOf(b.Set[1])/2,
					name:   fmt.Sprintf("%s#1", b.Set[1].Name()),
					parent: b.Set[1],
					scores: []float64{3},
				},
				&fs{
					start:  b.Set[1].Start() + lengthOf(b.Set[1])/2,
					end:    b.Set[1].End(),
					name:   fmt.Sprintf("%s#2", b.Set[1].Name()),
					parent: b.Set[1],
					scores: []float64{2},
				},
			},
			renderer: &Trace{
				LineStyles: []draw.LineStyle{func() draw.LineStyle {
					sty := plotter.DefaultLineStyle
					sty.Color = color.Gray{0}
					return sty
				}()},
				Join: true,
			},
		},
		{
			orient: Backward,
			scores: []Scorer{
				&fs{
					start:  b.Set[1].Start(),
					end:    b.Set[1].Start() + lengthOf(b.Set[1])/3,
					name:   fmt.Sprintf("%s#0", b.Set[1].Name()),
					parent: b.Set[1],
					scores: []float64{1},
				},
				&fs{
					start:  b.Set[1].Start() + lengthOf(b.Set[1])/3 + 1,
					end:    b.Set[1].Start() + lengthOf(b.Set[1])/2,
					name:   fmt.Sprintf("%s#1", b.Set[1].Name()),
					parent: b.Set[1],
					scores: []float64{3},
				},
				&fs{
					start:  b.Set[1].Start() + lengthOf(b.Set[1])/2,
					end:    b.Set[1].End(),
					name:   fmt.Sprintf("%s#2", b.Set[1].Name()),
					parent: b.Set[1],
					scores: []float64{2},
				},
			},
			renderer: &Trace{
				LineStyles: []draw.LineStyle{func() draw.LineStyle {
					sty := plotter.DefaultLineStyle
					sty.Color = color.Gray{0}
					return sty
				}()},
				Join: true,
			},
		},
		{
			orient: Backward,
			scores: []Scorer{
				&fs{
					start:  b.Set[1].Start(),
					end:    b.Set[1].Start() + lengthOf(b.Set[1])/3,
					name:   fmt.Sprintf("%s#0", b.Set[1].Name()),
					parent: b.Set[1],
					scores: []float64{1},
				},
				&fs{
					start:  b.Set[1].Start() + lengthOf(b.Set[1])/3,
					end:    b.Set[1].Start() + lengthOf(b.Set[1])/2,
					name:   fmt.Sprintf("%s#1", b.Set[1].Name()),
					parent: b.Set[1],
					scores: []float64{3},
				},
				&fs{
					start:  b.Set[1].Start() + lengthOf(b.Set[1])/2 + 1,
					end:    b.Set[1].End(),
					name:   fmt.Sprintf("%s#2", b.Set[1].Name()),
					parent: b.Set[1],
					scores: []float64{2},
				},
			},
			renderer: &Trace{
				LineStyles: []draw.LineStyle{func() draw.LineStyle {
					sty := plotter.DefaultLineStyle
					sty.Color = color.Gray{0}
					return sty
				}()},
				Join: true,
			},
		},
		{
			scores: makeScorers(b.Set[1].(*fs), 10, 2, func(_, _ int) float64 { return rnd.NormFloat64() }),
			renderer: &Trace{
				LineStyles: func() []draw.LineStyle {
					sty := []draw.LineStyle{plotter.DefaultLineStyle, plotter.DefaultLineStyle}
					sty[0].Color = color.NRGBA{R: 0xff, A: 0xff}
					sty[1].Color = color.RGBA{G: 0xff, A: 0x80}
					return sty
				}(),
			},
		},
		{
			scores: makeScorers(b.Set[1].(*fs), 10, 2, func(_, _ int) float64 { return rnd.NormFloat64() }),
			renderer: &Trace{
				LineStyles: func() []draw.LineStyle {
					sty := []draw.LineStyle{plotter.DefaultLineStyle, plotter.DefaultLineStyle}
					sty[0].Color = color.NRGBA{R: 0xff, A: 0xff}
					sty[1].Color = color.RGBA{G: 0xff, A: 0x80}
					return sty
				}(),
				Join: true,
			},
		},
	} {
		t.Run(fmt.Sprintf("scores-%d", i), func(t *testing.T) {
			p := plot.New()
			b.Set[1].(*fs).orient = test.orient
			b.Base = NewGappedArcs(b.Base, b.Set, 0.01)
			r, err := NewScores(test.scores, b, 40, 75, test.renderer)
			if err != nil {
				t.Fatalf("unexpected error for NewScores: %v", err)
			}

			p.Add(r)
			p.HideAxes()
			p.Add(b)

			checkImage(t, p)
		})
	}
}
