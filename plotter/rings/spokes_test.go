// Copyright ©2018 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rings

import (
	"math/rand/v2"
	"testing"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
)

func TestSpokes(t *testing.T) {
	p := plot.New()
	rnd := rand.New(rand.NewPCG(1, 1))
	b, err := NewGappedBlocks(randomFeatures(rnd, 3, 100000, 1000000, false, plotter.DefaultLineStyle),
		Arc{0, Complete * Clockwise},
		80, 100, 0.01,
	)
	if err != nil {
		t.Fatalf("unexpected error for NewGappedBlocks: %v", err)
	}

	m := randomFeatures(rnd, 10, b.Set[1].Start(), b.Set[1].End(), true, plotter.DefaultLineStyle)
	for _, mf := range m {
		mf.(*fs).parent = b.Set[1]
	}
	ms, err := NewSpokes(m, b, 73, 78)
	if err != nil {
		t.Fatalf("unexpected error for NewSpokes: %v", err)
	}
	ms.LineStyle = plotter.DefaultLineStyle

	p.Add(ms)
	p.HideAxes()
	p.Add(b)

	checkImage(t, p)
}
