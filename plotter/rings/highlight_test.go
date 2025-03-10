// Copyright ©2018 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rings

import (
	"image/color"
	"testing"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
)

func TestHighlight(t *testing.T) {
	p := plot.New()
	h := NewHighlight(
		color.NRGBA{R: 0xf3, G: 0xf3, B: 0x15, A: 0xff},
		Arc{0, Complete / 2 * Clockwise},
		30, 120,
	)
	h.LineStyle = plotter.DefaultLineStyle

	p.Add(h)
	p.HideAxes()

	checkImage(t, p, *regen)
}
