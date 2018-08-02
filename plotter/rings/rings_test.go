// Copyright ©2018 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rings

import (
	"bytes"
	"flag"
	"fmt"
	"image"
	"image/color"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"golang.org/x/exp/rand"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
)

var regen = flag.Bool("regen", false, "Uses the current state to regenerate the test data.")

// fs is a Feature implementation for testing.
type fs struct {
	start, end float64
	name       string
	parent     Feature
	orient     Orientation
	style      draw.LineStyle
	scores     []float64
}

func (f *fs) Start() float64            { return f.start }
func (f *fs) End() float64              { return f.end }
func (f *fs) Name() string              { return f.name }
func (f *fs) Parent() Feature           { return f.parent }
func (f *fs) Orientation() Orientation  { return f.orient }
func (f *fs) LineStyle() draw.LineStyle { return f.style }
func (f *fs) Scores() []float64         { return f.scores }

type fp struct {
	feats [2]*fs
	sty   draw.LineStyle
}

func (p fp) Features() [2]Feature { return [2]Feature{p.feats[0], p.feats[1]} }
func (p fp) LineStyle() draw.LineStyle {
	var col color.RGBA
	for _, f := range p.feats {
		r, g, b, a := f.style.Color.RGBA()
		col.R += byte(r / 2)
		col.G += byte(g / 2)
		col.B += byte(b / 2)
		col.A += byte(a / 2)
	}
	p.sty.Color = col
	return p.sty
}

func randomFeatures(n int, min, max float64, single bool, sty draw.LineStyle) []Feature {
	data := make([]Feature, n)
	for i := range data {
		// Intn is used here to avoid drastic random
		// sequence changes at this stage.
		start := float64(rand.Intn(int(max-min))) + min
		var end float64
		if !single {
			end = float64(rand.Intn(int(max - start)))
		}
		data[i] = &fs{
			start: start,
			end:   start + end,
			name:  fmt.Sprintf("feature%v", i),
			style: sty,
		}
	}
	return data
}

func floatPtr(f float64) *float64 { return &f }

// makeScorers returns n Scorers each with m scores.
func makeScorers(f *fs, n, m int, fn func(i, j int) float64) []Scorer {
	s := make([]Scorer, n)
	for i := 0; i < n; i++ {
		cs := &fs{
			start:  f.Start() + float64(i)*(lengthOf(f)/float64(n)),
			end:    f.Start() + float64(i+1)*(lengthOf(f)/float64(n)),
			name:   fmt.Sprintf("%s#%d", f.Name(), i),
			parent: f,
			scores: make([]float64, m),
		}
		for j := range cs.scores {
			cs.scores[j] = fn(i, j)
		}
		s[i] = cs
	}
	return s
}

// checkImage compares the plot in p to the image in testdata/name_golden.png.
// If regen is true the plot in p is first saved to testdata/name_golden.png.
func checkImage(t *testing.T, p *plot.Plot, regen bool) {
	name := filepath.FromSlash(t.Name())
	path := filepath.Join("testdata", name+"_golden.png")
	w, err := p.WriterTo(vg.Length(300), vg.Length(300), "png")
	var buf bytes.Buffer
	_, err = w.WriteTo(&buf)
	if err != nil {
		t.Fatalf("unexpected error writing plot: %v", err)
	}
	got := buf.Bytes()
	if regen {
		err = os.Mkdir(filepath.Dir(path), 0775)
		if err != nil && !os.IsExist(err) {
			t.Fatalf("failed to created testdata subdir: %v", err)
			return
		}
		err = ioutil.WriteFile(path, got, 0664)
		if err != nil {
			t.Fatalf("unexpected error writing golden file: %v", err)
		}
		// Fallthrough rather than returning just
		// to confirm we have written correctly.
	}
	gold, err := ioutil.ReadFile(path)
	if err != nil {
		t.Fatalf("unexpected error reading golden file: %v", err)
	}
	ok, err := equalImage(got, gold)
	if err != nil {
		t.Errorf("failed to compare image for %s: %v", path, err)
	}
	if !ok {
		t.Errorf("image mismatch for %q", name)
		// TODO(kortschak): Add image diffing.
		err = ioutil.WriteFile(filepath.Join("testdata", name+"_failed.png"), got, 0664)
		if err != nil {
			t.Fatalf("unexpected error writing failed file: %v", err)
		}
	}
}

// TODO(kortschak): Use cmpimg when rings lives in plot.
func equalImage(raw1, raw2 []byte) (bool, error) {
	v1, _, err := image.Decode(bytes.NewReader(raw1))
	if err != nil {
		return false, err
	}
	v2, _, err := image.Decode(bytes.NewReader(raw2))
	if err != nil {
		return false, err
	}
	return reflect.DeepEqual(v1, v2), nil
}
