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
	"gonum.org/v1/plot/palette"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"

	"github.com/biogo/biogo/feat"
)

func TestScores(t *testing.T) {
	rand.Seed(1)
	b, err := NewGappedBlocks(randomFeatures(3, 100000, 1000000, false, plotter.DefaultLineStyle),
		Arc{0, Complete * Clockwise},
		80, 100, 0.01,
	)
	if err != nil {
		t.Fatalf("unexpected error for NewGappedBlocks: %v", err)
	}

	for i, test := range []struct {
		orient   feat.Orientation
		scores   []Scorer
		renderer ScoreRenderer
		actions  []interface{}
	}{
		{
			scores:   makeScorers(b.Set[1].(*fs), 10, 5, func(i, j int) float64 { return float64(i * j) }),
			renderer: &Heat{Palette: palette.Radial(10, palette.Cyan, palette.Magenta, 1).Colors()},
			actions: []interface{}{
				setColor{col: color.NRGBA{R: 0x7f, G: 0xff, B: 0xff, A: 0xff}},
				fill{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 143.59933171592297, Y: 113.50284492303541}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 40, Start: 4.487993840681956, Angle: -0.1367744033971654},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 47, Start: 4.351219437284791, Angle: 0.1367744033971654},
					{Type: vg.CloseComp, Pos: vg.Point{X: 0, Y: 0}, Radius: 0, Start: 0, Angle: 0},
				}},
				setColor{col: color.NRGBA{R: 0x7f, G: 0xff, B: 0xff, A: 0xff}},
				fill{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 142.0417147662095, Y: 106.67834278456661}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 47, Start: 4.487993840681956, Angle: -0.1367744033971654},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 54, Start: 4.351219437284791, Angle: 0.1367744033971654},
					{Type: vg.CloseComp, Pos: vg.Point{X: 0, Y: 0}, Radius: 0, Start: 0, Angle: 0},
				}},
				setColor{col: color.NRGBA{R: 0x7f, G: 0xff, B: 0xff, A: 0xff}},
				fill{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 140.48409781649605, Y: 99.8538406460978}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 54, Start: 4.487993840681956, Angle: -0.1367744033971654},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 61, Start: 4.351219437284791, Angle: 0.1367744033971654},
					{Type: vg.CloseComp, Pos: vg.Point{X: 0, Y: 0}, Radius: 0, Start: 0, Angle: 0},
				}},
				setColor{col: color.NRGBA{R: 0x7f, G: 0xff, B: 0xff, A: 0xff}},
				fill{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 138.92648086678255, Y: 93.02933850762899}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 61, Start: 4.487993840681956, Angle: -0.1367744033971654},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 68, Start: 4.351219437284791, Angle: 0.1367744033971654},
					{Type: vg.CloseComp, Pos: vg.Point{X: 0, Y: 0}, Radius: 0, Start: 0, Angle: 0},
				}},
				setColor{col: color.NRGBA{R: 0x7f, G: 0xff, B: 0xff, A: 0xff}},
				fill{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 137.36886391706906, Y: 86.20483636916019}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 68, Start: 4.487993840681956, Angle: -0.1367744033971654},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 75, Start: 4.351219437284791, Angle: 0.1367744033971654},
					{Type: vg.CloseComp, Pos: vg.Point{X: 0, Y: 0}, Radius: 0, Start: 0, Angle: 0},
				}},
				setColor{col: color.NRGBA{R: 0x7f, G: 0xff, B: 0xff, A: 0xff}},
				fill{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 138.36525744911535, Y: 115.0806326480496}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 40, Start: 4.351219437284791, Angle: -0.1367744033971654},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 47, Start: 4.214445033887626, Angle: 0.1367744033971654},
					{Type: vg.CloseComp, Pos: vg.Point{X: 0, Y: 0}, Radius: 0, Start: 0, Angle: 0},
				}},
				setColor{col: color.NRGBA{R: 0x7f, G: 0xff, B: 0xff, A: 0xff}},
				fill{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 135.89167750271054, Y: 108.53224336145828}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 47, Start: 4.351219437284791, Angle: -0.1367744033971654},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 54, Start: 4.214445033887626, Angle: 0.1367744033971654},
					{Type: vg.CloseComp, Pos: vg.Point{X: 0, Y: 0}, Radius: 0, Start: 0, Angle: 0},
				}},
				setColor{col: color.NRGBA{R: 0x99, G: 0xff, B: 0xff, A: 0xff}},
				fill{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 133.41809755630572, Y: 101.98385407486697}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 54, Start: 4.351219437284791, Angle: -0.1367744033971654},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 61, Start: 4.214445033887626, Angle: 0.1367744033971654},
					{Type: vg.CloseComp, Pos: vg.Point{X: 0, Y: 0}, Radius: 0, Start: 0, Angle: 0},
				}},
				setColor{col: color.NRGBA{R: 0x99, G: 0xff, B: 0xff, A: 0xff}},
				fill{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 130.9445176099009, Y: 95.43546478827565}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 61, Start: 4.351219437284791, Angle: -0.1367744033971654},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 68, Start: 4.214445033887626, Angle: 0.1367744033971654},
					{Type: vg.CloseComp, Pos: vg.Point{X: 0, Y: 0}, Radius: 0, Start: 0, Angle: 0},
				}},
				setColor{col: color.NRGBA{R: 0x99, G: 0xff, B: 0xff, A: 0xff}},
				fill{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 128.4709376634961, Y: 88.88707550168434}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 68, Start: 4.351219437284791, Angle: -0.1367744033971654},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 75, Start: 4.214445033887626, Angle: 0.1367744033971654},
					{Type: vg.CloseComp, Pos: vg.Point{X: 0, Y: 0}, Radius: 0, Start: 0, Angle: 0},
				}},
				setColor{col: color.NRGBA{R: 0x7f, G: 0xff, B: 0xff, A: 0xff}},
				fill{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 133.39519320703477, Y: 117.35734276689553}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 40, Start: 4.214445033887626, Angle: -0.1367744033971654},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 47, Start: 4.07767063049046, Angle: 0.1367744033971654},
					{Type: vg.CloseComp, Pos: vg.Point{X: 0, Y: 0}, Radius: 0, Start: 0, Angle: 0},
				}},
				setColor{col: color.NRGBA{R: 0x99, G: 0xff, B: 0xff, A: 0xff}},
				fill{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 130.05185201826586, Y: 111.20737775110226}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 47, Start: 4.214445033887626, Angle: -0.1367744033971654},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 54, Start: 4.07767063049046, Angle: 0.1367744033971654},
					{Type: vg.CloseComp, Pos: vg.Point{X: 0, Y: 0}, Radius: 0, Start: 0, Angle: 0},
				}},
				setColor{col: color.NRGBA{R: 0x99, G: 0xff, B: 0xff, A: 0xff}},
				fill{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 126.70851082949693, Y: 105.05741273530897}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 54, Start: 4.214445033887626, Angle: -0.1367744033971654},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 61, Start: 4.07767063049046, Angle: 0.1367744033971654},
					{Type: vg.CloseComp, Pos: vg.Point{X: 0, Y: 0}, Radius: 0, Start: 0, Angle: 0},
				}},
				setColor{col: color.NRGBA{R: 0xb3, G: 0xff, B: 0xff, A: 0xff}},
				fill{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 123.36516964072803, Y: 98.90744771951569}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 61, Start: 4.214445033887626, Angle: -0.1367744033971654},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 68, Start: 4.07767063049046, Angle: 0.1367744033971654},
					{Type: vg.CloseComp, Pos: vg.Point{X: 0, Y: 0}, Radius: 0, Start: 0, Angle: 0},
				}},
				setColor{col: color.NRGBA{R: 0xb3, G: 0xff, B: 0xff, A: 0xff}},
				fill{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 120.0218284519591, Y: 92.75748270372242}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 68, Start: 4.214445033887626, Angle: -0.1367744033971654},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 75, Start: 4.07767063049046, Angle: 0.1367744033971654},
					{Type: vg.CloseComp, Pos: vg.Point{X: 0, Y: 0}, Radius: 0, Start: 0, Angle: 0},
				}},
				setColor{col: color.NRGBA{R: 0x7f, G: 0xff, B: 0xff, A: 0xff}},
				fill{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 128.7819703078834, Y: 120.29045067803841}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 40, Start: 4.07767063049046, Angle: -0.1367744033971654},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 47, Start: 3.940896227093295, Angle: 0.1367744033971654},
					{Type: vg.CloseComp, Pos: vg.Point{X: 0, Y: 0}, Radius: 0, Start: 0, Angle: 0},
				}},
				setColor{col: color.NRGBA{R: 0x99, G: 0xff, B: 0xff, A: 0xff}},
				fill{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 124.631315111763, Y: 114.65377954669515}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 47, Start: 4.07767063049046, Angle: -0.1367744033971654},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 54, Start: 3.940896227093295, Angle: 0.1367744033971654},
					{Type: vg.CloseComp, Pos: vg.Point{X: 0, Y: 0}, Radius: 0, Start: 0, Angle: 0},
				}},
				setColor{col: color.NRGBA{R: 0xb3, G: 0xff, B: 0xff, A: 0xff}},
				fill{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 120.4806599156426, Y: 109.01710841535187}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 54, Start: 4.07767063049046, Angle: -0.1367744033971654},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 61, Start: 3.940896227093295, Angle: 0.1367744033971654},
					{Type: vg.CloseComp, Pos: vg.Point{X: 0, Y: 0}, Radius: 0, Start: 0, Angle: 0},
				}},
				setColor{col: color.NRGBA{R: 0xb3, G: 0xff, B: 0xff, A: 0xff}},
				fill{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 116.33000471952221, Y: 103.3804372840086}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 61, Start: 4.07767063049046, Angle: -0.1367744033971654},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 68, Start: 3.940896227093295, Angle: 0.1367744033971654},
					{Type: vg.CloseComp, Pos: vg.Point{X: 0, Y: 0}, Radius: 0, Start: 0, Angle: 0},
				}},
				setColor{col: color.NRGBA{R: 0xcc, G: 0xff, B: 0xff, A: 0xff}},
				fill{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 112.1793495234018, Y: 97.74376615266532}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 68, Start: 4.07767063049046, Angle: -0.1367744033971654},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 75, Start: 3.940896227093295, Angle: 0.1367744033971654},
					{Type: vg.CloseComp, Pos: vg.Point{X: 0, Y: 0}, Radius: 0, Start: 0, Angle: 0},
				}},
				setColor{col: color.NRGBA{R: 0x7f, G: 0xff, B: 0xff, A: 0xff}},
				fill{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 124.61175495435593, Y: 123.82517152145267}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 40, Start: 3.940896227093295, Angle: -0.1367744033971654},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 47, Start: 3.8041218236961294, Angle: 0.1367744033971654},
					{Type: vg.CloseComp, Pos: vg.Point{X: 0, Y: 0}, Radius: 0, Start: 0, Angle: 0},
				}},
				setColor{col: color.NRGBA{R: 0x99, G: 0xff, B: 0xff, A: 0xff}},
				fill{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 119.73131207136822, Y: 118.80707653770689}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 47, Start: 3.940896227093295, Angle: -0.1367744033971654},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 54, Start: 3.8041218236961294, Angle: 0.1367744033971654},
					{Type: vg.CloseComp, Pos: vg.Point{X: 0, Y: 0}, Radius: 0, Start: 0, Angle: 0},
				}},
				setColor{col: color.NRGBA{R: 0xb3, G: 0xff, B: 0xff, A: 0xff}},
				fill{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 114.8508691883805, Y: 113.78898155396111}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 54, Start: 3.940896227093295, Angle: -0.1367744033971654},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 61, Start: 3.8041218236961294, Angle: 0.1367744033971654},
					{Type: vg.CloseComp, Pos: vg.Point{X: 0, Y: 0}, Radius: 0, Start: 0, Angle: 0},
				}},
				setColor{col: color.NRGBA{R: 0xcc, G: 0xff, B: 0xff, A: 0xff}},
				fill{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 109.97042630539279, Y: 108.77088657021532}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 61, Start: 3.940896227093295, Angle: -0.1367744033971654},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 68, Start: 3.8041218236961294, Angle: 0.1367744033971654},
					{Type: vg.CloseComp, Pos: vg.Point{X: 0, Y: 0}, Radius: 0, Start: 0, Angle: 0},
				}},
				setColor{col: color.NRGBA{R: 0xe6, G: 0xff, B: 0xff, A: 0xff}},
				fill{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 105.08998342240508, Y: 103.75279158646954}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 68, Start: 3.940896227093295, Angle: -0.1367744033971654},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 75, Start: 3.8041218236961294, Angle: 0.1367744033971654},
					{Type: vg.CloseComp, Pos: vg.Point{X: 0, Y: 0}, Radius: 0, Start: 0, Angle: 0},
				}},
				setColor{col: color.NRGBA{R: 0x7f, G: 0xff, B: 0xff, A: 0xff}},
				fill{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 120.96243881336021, Y: 127.89548345528912}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 40, Start: 3.8041218236961294, Angle: -0.1367744033971654},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 47, Start: 3.667347420298964, Angle: 0.1367744033971654},
					{Type: vg.CloseComp, Pos: vg.Point{X: 0, Y: 0}, Radius: 0, Start: 0, Angle: 0},
				}},
				setColor{col: color.NRGBA{R: 0x99, G: 0xff, B: 0xff, A: 0xff}},
				fill{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 115.44336560569823, Y: 123.58969305996473}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 47, Start: 3.8041218236961294, Angle: -0.1367744033971654},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 54, Start: 3.667347420298964, Angle: 0.1367744033971654},
					{Type: vg.CloseComp, Pos: vg.Point{X: 0, Y: 0}, Radius: 0, Start: 0, Angle: 0},
				}},
				setColor{col: color.NRGBA{R: 0xcc, G: 0xff, B: 0xff, A: 0xff}},
				fill{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 109.92429239803627, Y: 119.28390266464032}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 54, Start: 3.8041218236961294, Angle: -0.1367744033971654},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 61, Start: 3.667347420298964, Angle: 0.1367744033971654},
					{Type: vg.CloseComp, Pos: vg.Point{X: 0, Y: 0}, Radius: 0, Start: 0, Angle: 0},
				}},
				setColor{col: color.NRGBA{R: 0xe6, G: 0xff, B: 0xff, A: 0xff}},
				fill{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 104.40521919037431, Y: 114.97811226931591}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 61, Start: 3.8041218236961294, Angle: -0.1367744033971654},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 68, Start: 3.667347420298964, Angle: 0.1367744033971654},
					{Type: vg.CloseComp, Pos: vg.Point{X: 0, Y: 0}, Radius: 0, Start: 0, Angle: 0},
				}},
				setColor{col: color.NRGBA{R: 0xff, G: 0xe6, B: 0xff, A: 0xff}},
				fill{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 98.88614598271235, Y: 110.6723218739915}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 68, Start: 3.8041218236961294, Angle: -0.1367744033971654},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 75, Start: 3.667347420298964, Angle: 0.1367744033971654},
					{Type: vg.CloseComp, Pos: vg.Point{X: 0, Y: 0}, Radius: 0, Start: 0, Angle: 0},
				}},
				setColor{col: color.NRGBA{R: 0x7f, G: 0xff, B: 0xff, A: 0xff}},
				fill{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 117.90218414828041, Y: 132.425360817925}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 40, Start: 3.667347420298964, Angle: -0.1367744033971653},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 47, Start: 3.5305730169017986, Angle: 0.1367744033971653},
					{Type: vg.CloseComp, Pos: vg.Point{X: 0, Y: 0}, Radius: 0, Start: 0, Angle: 0},
				}},
				setColor{col: color.NRGBA{R: 0xb3, G: 0xff, B: 0xff, A: 0xff}},
				fill{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 111.8475663742295, Y: 128.91229896106185}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 47, Start: 3.667347420298964, Angle: -0.1367744033971653},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 54, Start: 3.5305730169017986, Angle: 0.1367744033971653},
					{Type: vg.CloseComp, Pos: vg.Point{X: 0, Y: 0}, Radius: 0, Start: 0, Angle: 0},
				}},
				setColor{col: color.NRGBA{R: 0xcc, G: 0xff, B: 0xff, A: 0xff}},
				fill{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 105.79294860017856, Y: 125.39923710419873}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 54, Start: 3.667347420298964, Angle: -0.1367744033971653},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 61, Start: 3.5305730169017986, Angle: 0.1367744033971653},
					{Type: vg.CloseComp, Pos: vg.Point{X: 0, Y: 0}, Radius: 0, Start: 0, Angle: 0},
				}},
				setColor{col: color.NRGBA{R: 0xff, G: 0xe6, B: 0xff, A: 0xff}},
				fill{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 99.73833082612765, Y: 121.88617524733561}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 61, Start: 3.667347420298964, Angle: -0.1367744033971653},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 68, Start: 3.5305730169017986, Angle: 0.1367744033971653},
					{Type: vg.CloseComp, Pos: vg.Point{X: 0, Y: 0}, Radius: 0, Start: 0, Angle: 0},
				}},
				setColor{col: color.NRGBA{R: 0xff, G: 0xcc, B: 0xff, A: 0xff}},
				fill{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 93.68371305207671, Y: 118.37311339047248}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 68, Start: 3.667347420298964, Angle: -0.1367744033971653},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 75, Start: 3.5305730169017986, Angle: 0.1367744033971653},
					{Type: vg.CloseComp, Pos: vg.Point{X: 0, Y: 0}, Radius: 0, Start: 0, Angle: 0},
				}},
				setColor{col: color.NRGBA{R: 0x7f, G: 0xff, B: 0xff, A: 0xff}},
				fill{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 115.48815067793637, Y: 137.3301941422819}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 40, Start: 3.5305730169017986, Angle: -0.1367744033971655},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 47, Start: 3.393798613504633, Angle: 0.1367744033971655},
					{Type: vg.CloseComp, Pos: vg.Point{X: 0, Y: 0}, Radius: 0, Start: 0, Angle: 0},
				}},
				setColor{col: color.NRGBA{R: 0xb3, G: 0xff, B: 0xff, A: 0xff}},
				fill{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 109.01107704657524, Y: 134.67547811718123}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 47, Start: 3.5305730169017986, Angle: -0.1367744033971655},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 54, Start: 3.393798613504633, Angle: 0.1367744033971655},
					{Type: vg.CloseComp, Pos: vg.Point{X: 0, Y: 0}, Radius: 0, Start: 0, Angle: 0},
				}},
				setColor{col: color.NRGBA{R: 0xe6, G: 0xff, B: 0xff, A: 0xff}},
				fill{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 102.5340034152141, Y: 132.02076209208056}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 54, Start: 3.5305730169017986, Angle: -0.1367744033971655},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 61, Start: 3.393798613504633, Angle: 0.1367744033971655},
					{Type: vg.CloseComp, Pos: vg.Point{X: 0, Y: 0}, Radius: 0, Start: 0, Angle: 0},
				}},
				setColor{col: color.NRGBA{R: 0xff, G: 0xe6, B: 0xff, A: 0xff}},
				fill{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 96.05692978385296, Y: 129.3660460669799}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 61, Start: 3.5305730169017986, Angle: -0.1367744033971655},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 68, Start: 3.393798613504633, Angle: 0.1367744033971655},
					{Type: vg.CloseComp, Pos: vg.Point{X: 0, Y: 0}, Radius: 0, Start: 0, Angle: 0},
				}},
				setColor{col: color.NRGBA{R: 0xff, G: 0xb3, B: 0xff, A: 0xff}},
				fill{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 89.57985615249183, Y: 126.71133004187924}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 68, Start: 3.5305730169017986, Angle: -0.1367744033971655},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 75, Start: 3.393798613504633, Angle: 0.1367744033971655},
					{Type: vg.CloseComp, Pos: vg.Point{X: 0, Y: 0}, Radius: 0, Start: 0, Angle: 0},
				}},
				setColor{col: color.NRGBA{R: 0x7f, G: 0xff, B: 0xff, A: 0xff}},
				fill{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 113.76542794208467, Y: 142.51837049925388}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 40, Start: 3.393798613504633, Angle: -0.1367744033971654},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 47, Start: 3.257024210107468, Angle: 0.1367744033971654},
					{Type: vg.CloseComp, Pos: vg.Point{X: 0, Y: 0}, Radius: 0, Start: 0, Angle: 0},
				}},
				setColor{col: color.NRGBA{R: 0xb3, G: 0xff, B: 0xff, A: 0xff}},
				fill{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 106.9868778319495, Y: 140.7715853366233}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 47, Start: 3.393798613504633, Angle: -0.1367744033971654},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 54, Start: 3.257024210107468, Angle: 0.1367744033971654},
					{Type: vg.CloseComp, Pos: vg.Point{X: 0, Y: 0}, Radius: 0, Start: 0, Angle: 0},
				}},
				setColor{col: color.NRGBA{R: 0xe6, G: 0xff, B: 0xff, A: 0xff}},
				fill{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 100.2083277218143, Y: 139.02480017399276}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 54, Start: 3.393798613504633, Angle: -0.1367744033971654},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 61, Start: 3.257024210107468, Angle: 0.1367744033971654},
					{Type: vg.CloseComp, Pos: vg.Point{X: 0, Y: 0}, Radius: 0, Start: 0, Angle: 0},
				}},
				setColor{col: color.NRGBA{R: 0xff, G: 0xcc, B: 0xff, A: 0xff}},
				fill{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 93.42977761167913, Y: 137.27801501136219}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 61, Start: 3.393798613504633, Angle: -0.1367744033971654},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 68, Start: 3.257024210107468, Angle: 0.1367744033971654},
					{Type: vg.CloseComp, Pos: vg.Point{X: 0, Y: 0}, Radius: 0, Start: 0, Angle: 0},
				}},
				setColor{col: color.NRGBA{R: 0xff, G: 0x99, B: 0xff, A: 0xff}},
				fill{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 86.65122750154394, Y: 135.5312298487316}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 68, Start: 3.393798613504633, Angle: -0.1367744033971654},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 75, Start: 3.257024210107468, Angle: 0.1367744033971654},
					{Type: vg.CloseComp, Pos: vg.Point{X: 0, Y: 0}, Radius: 0, Start: 0, Angle: 0},
				}},
				setColor{col: color.NRGBA{R: 0x7f, G: 0xff, B: 0xff, A: 0xff}},
				fill{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 112.76619311483745, Y: 147.89298465244477}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 40, Start: 3.257024210107468, Angle: -0.1367744033971654},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 47, Start: 3.1202498067103024, Angle: 0.1367744033971654},
					{Type: vg.CloseComp, Pos: vg.Point{X: 0, Y: 0}, Radius: 0, Start: 0, Angle: 0},
				}},
				setColor{col: color.NRGBA{R: 0xb3, G: 0xff, B: 0xff, A: 0xff}},
				fill{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 105.812776909934, Y: 147.0867569666226}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 47, Start: 3.257024210107468, Angle: -0.1367744033971654},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 54, Start: 3.1202498067103024, Angle: 0.1367744033971654},
					{Type: vg.CloseComp, Pos: vg.Point{X: 0, Y: 0}, Radius: 0, Start: 0, Angle: 0},
				}},
				setColor{col: color.NRGBA{R: 0xff, G: 0xe6, B: 0xff, A: 0xff}},
				fill{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 98.85936070503055, Y: 146.28052928080044}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 54, Start: 3.257024210107468, Angle: -0.1367744033971654},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 61, Start: 3.1202498067103024, Angle: 0.1367744033971654},
					{Type: vg.CloseComp, Pos: vg.Point{X: 0, Y: 0}, Radius: 0, Start: 0, Angle: 0},
				}},
				setColor{col: color.NRGBA{R: 0xff, G: 0xb3, B: 0xff, A: 0xff}},
				fill{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 91.9059445001271, Y: 145.47430159497827}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 61, Start: 3.257024210107468, Angle: -0.1367744033971654},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 68, Start: 3.1202498067103024, Angle: 0.1367744033971654},
					{Type: vg.CloseComp, Pos: vg.Point{X: 0, Y: 0}, Radius: 0, Start: 0, Angle: 0},
				}},
				setColor{col: color.NRGBA{R: 0xff, G: 0x7f, B: 0xff, A: 0xff}},
				fill{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 84.95252829522367, Y: 144.6680739091561}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 68, Start: 3.257024210107468, Angle: -0.1367744033971654},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 75, Start: 3.1202498067103024, Angle: 0.1367744033971654},
					{Type: vg.CloseComp, Pos: vg.Point{X: 0, Y: 0}, Radius: 0, Start: 0, Angle: 0},
				}},
			},
		},
		{
			scores:   makeScorers(b.Set[1].(*fs), 10, 5, func(_, _ int) float64 { return rand.NormFloat64() }),
			renderer: &Heat{Palette: palette.Radial(10, palette.Cyan, palette.Magenta, 1).Colors()},
			actions: []interface{}{
				setColor{col: color.NRGBA{R: 0xff, G: 0xcc, B: 0xff, A: 0xff}},
				fill{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 143.59933171592297, Y: 113.50284492303541}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 40, Start: 4.487993840681956, Angle: -0.1367744033971654},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 47, Start: 4.351219437284791, Angle: 0.1367744033971654},
					{Type: vg.CloseComp, Pos: vg.Point{X: 0, Y: 0}, Radius: 0, Start: 0, Angle: 0},
				}},
				setColor{col: color.NRGBA{R: 0xff, G: 0xcc, B: 0xff, A: 0xff}},
				fill{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 142.0417147662095, Y: 106.67834278456661}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 47, Start: 4.487993840681956, Angle: -0.1367744033971654},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 54, Start: 4.351219437284791, Angle: 0.1367744033971654},
					{Type: vg.CloseComp, Pos: vg.Point{X: 0, Y: 0}, Radius: 0, Start: 0, Angle: 0},
				}},
				setColor{col: color.NRGBA{R: 0xff, G: 0xe6, B: 0xff, A: 0xff}},
				fill{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 140.48409781649605, Y: 99.8538406460978}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 54, Start: 4.487993840681956, Angle: -0.1367744033971654},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 61, Start: 4.351219437284791, Angle: 0.1367744033971654},
					{Type: vg.CloseComp, Pos: vg.Point{X: 0, Y: 0}, Radius: 0, Start: 0, Angle: 0},
				}},
				setColor{col: color.NRGBA{R: 0xff, G: 0xb3, B: 0xff, A: 0xff}},
				fill{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 138.92648086678255, Y: 93.02933850762899}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 61, Start: 4.487993840681956, Angle: -0.1367744033971654},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 68, Start: 4.351219437284791, Angle: 0.1367744033971654},
					{Type: vg.CloseComp, Pos: vg.Point{X: 0, Y: 0}, Radius: 0, Start: 0, Angle: 0},
				}},
				setColor{col: color.NRGBA{R: 0xcc, G: 0xff, B: 0xff, A: 0xff}},
				fill{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 137.36886391706906, Y: 86.20483636916019}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 68, Start: 4.487993840681956, Angle: -0.1367744033971654},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 75, Start: 4.351219437284791, Angle: 0.1367744033971654},
					{Type: vg.CloseComp, Pos: vg.Point{X: 0, Y: 0}, Radius: 0, Start: 0, Angle: 0},
				}},
				setColor{col: color.NRGBA{R: 0xff, G: 0xcc, B: 0xff, A: 0xff}},
				fill{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 138.36525744911535, Y: 115.0806326480496}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 40, Start: 4.351219437284791, Angle: -0.1367744033971654},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 47, Start: 4.214445033887626, Angle: 0.1367744033971654},
					{Type: vg.CloseComp, Pos: vg.Point{X: 0, Y: 0}, Radius: 0, Start: 0, Angle: 0},
				}},
				setColor{col: color.NRGBA{R: 0xff, G: 0x99, B: 0xff, A: 0xff}},
				fill{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 135.89167750271054, Y: 108.53224336145828}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 47, Start: 4.351219437284791, Angle: -0.1367744033971654},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 54, Start: 4.214445033887626, Angle: 0.1367744033971654},
					{Type: vg.CloseComp, Pos: vg.Point{X: 0, Y: 0}, Radius: 0, Start: 0, Angle: 0},
				}},
				setColor{col: color.NRGBA{R: 0xff, G: 0xb3, B: 0xff, A: 0xff}},
				fill{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 133.41809755630572, Y: 101.98385407486697}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 54, Start: 4.351219437284791, Angle: -0.1367744033971654},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 61, Start: 4.214445033887626, Angle: 0.1367744033971654},
					{Type: vg.CloseComp, Pos: vg.Point{X: 0, Y: 0}, Radius: 0, Start: 0, Angle: 0},
				}},
				setColor{col: color.NRGBA{R: 0xff, G: 0x99, B: 0xff, A: 0xff}},
				fill{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 130.9445176099009, Y: 95.43546478827565}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 61, Start: 4.351219437284791, Angle: -0.1367744033971654},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 68, Start: 4.214445033887626, Angle: 0.1367744033971654},
					{Type: vg.CloseComp, Pos: vg.Point{X: 0, Y: 0}, Radius: 0, Start: 0, Angle: 0},
				}},
				setColor{col: color.NRGBA{R: 0xff, G: 0xcc, B: 0xff, A: 0xff}},
				fill{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 128.4709376634961, Y: 88.88707550168434}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 68, Start: 4.351219437284791, Angle: -0.1367744033971654},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 75, Start: 4.214445033887626, Angle: 0.1367744033971654},
					{Type: vg.CloseComp, Pos: vg.Point{X: 0, Y: 0}, Radius: 0, Start: 0, Angle: 0},
				}},
				setColor{col: color.NRGBA{R: 0xff, G: 0xcc, B: 0xff, A: 0xff}},
				fill{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 133.39519320703477, Y: 117.35734276689553}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 40, Start: 4.214445033887626, Angle: -0.1367744033971654},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 47, Start: 4.07767063049046, Angle: 0.1367744033971654},
					{Type: vg.CloseComp, Pos: vg.Point{X: 0, Y: 0}, Radius: 0, Start: 0, Angle: 0},
				}},
				setColor{col: color.NRGBA{R: 0xb3, G: 0xff, B: 0xff, A: 0xff}},
				fill{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 130.05185201826586, Y: 111.20737775110226}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 47, Start: 4.214445033887626, Angle: -0.1367744033971654},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 54, Start: 4.07767063049046, Angle: 0.1367744033971654},
					{Type: vg.CloseComp, Pos: vg.Point{X: 0, Y: 0}, Radius: 0, Start: 0, Angle: 0},
				}},
				setColor{col: color.NRGBA{R: 0xff, G: 0xcc, B: 0xff, A: 0xff}},
				fill{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 126.70851082949693, Y: 105.05741273530897}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 54, Start: 4.214445033887626, Angle: -0.1367744033971654},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 61, Start: 4.07767063049046, Angle: 0.1367744033971654},
					{Type: vg.CloseComp, Pos: vg.Point{X: 0, Y: 0}, Radius: 0, Start: 0, Angle: 0},
				}},
				setColor{col: color.NRGBA{R: 0xff, G: 0xcc, B: 0xff, A: 0xff}},
				fill{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 123.36516964072803, Y: 98.90744771951569}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 61, Start: 4.214445033887626, Angle: -0.1367744033971654},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 68, Start: 4.07767063049046, Angle: 0.1367744033971654},
					{Type: vg.CloseComp, Pos: vg.Point{X: 0, Y: 0}, Radius: 0, Start: 0, Angle: 0},
				}},
				setColor{col: color.NRGBA{R: 0xff, G: 0xb3, B: 0xff, A: 0xff}},
				fill{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 120.0218284519591, Y: 92.75748270372242}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 68, Start: 4.214445033887626, Angle: -0.1367744033971654},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 75, Start: 4.07767063049046, Angle: 0.1367744033971654},
					{Type: vg.CloseComp, Pos: vg.Point{X: 0, Y: 0}, Radius: 0, Start: 0, Angle: 0},
				}},
				setColor{col: color.NRGBA{R: 0x99, G: 0xff, B: 0xff, A: 0xff}},
				fill{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 128.7819703078834, Y: 120.29045067803841}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 40, Start: 4.07767063049046, Angle: -0.1367744033971654},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 47, Start: 3.940896227093295, Angle: 0.1367744033971654},
					{Type: vg.CloseComp, Pos: vg.Point{X: 0, Y: 0}, Radius: 0, Start: 0, Angle: 0},
				}},
				setColor{col: color.NRGBA{R: 0xe6, G: 0xff, B: 0xff, A: 0xff}},
				fill{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 124.631315111763, Y: 114.65377954669515}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 47, Start: 4.07767063049046, Angle: -0.1367744033971654},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 54, Start: 3.940896227093295, Angle: 0.1367744033971654},
					{Type: vg.CloseComp, Pos: vg.Point{X: 0, Y: 0}, Radius: 0, Start: 0, Angle: 0},
				}},
				setColor{col: color.NRGBA{R: 0xff, G: 0x7f, B: 0xff, A: 0xff}},
				fill{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 120.4806599156426, Y: 109.01710841535187}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 54, Start: 4.07767063049046, Angle: -0.1367744033971654},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 61, Start: 3.940896227093295, Angle: 0.1367744033971654},
					{Type: vg.CloseComp, Pos: vg.Point{X: 0, Y: 0}, Radius: 0, Start: 0, Angle: 0},
				}},
				setColor{col: color.NRGBA{R: 0xff, G: 0xb3, B: 0xff, A: 0xff}},
				fill{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 116.33000471952221, Y: 103.3804372840086}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 61, Start: 4.07767063049046, Angle: -0.1367744033971654},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 68, Start: 3.940896227093295, Angle: 0.1367744033971654},
					{Type: vg.CloseComp, Pos: vg.Point{X: 0, Y: 0}, Radius: 0, Start: 0, Angle: 0},
				}},
				setColor{col: color.NRGBA{R: 0xcc, G: 0xff, B: 0xff, A: 0xff}},
				fill{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 112.1793495234018, Y: 97.74376615266532}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 68, Start: 4.07767063049046, Angle: -0.1367744033971654},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 75, Start: 3.940896227093295, Angle: 0.1367744033971654},
					{Type: vg.CloseComp, Pos: vg.Point{X: 0, Y: 0}, Radius: 0, Start: 0, Angle: 0},
				}},
				setColor{col: color.NRGBA{R: 0xff, G: 0xb3, B: 0xff, A: 0xff}},
				fill{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 124.61175495435593, Y: 123.82517152145267}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 40, Start: 3.940896227093295, Angle: -0.1367744033971654},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 47, Start: 3.8041218236961294, Angle: 0.1367744033971654},
					{Type: vg.CloseComp, Pos: vg.Point{X: 0, Y: 0}, Radius: 0, Start: 0, Angle: 0},
				}},
				setColor{col: color.NRGBA{R: 0xcc, G: 0xff, B: 0xff, A: 0xff}},
				fill{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 119.73131207136822, Y: 118.80707653770689}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 47, Start: 3.940896227093295, Angle: -0.1367744033971654},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 54, Start: 3.8041218236961294, Angle: 0.1367744033971654},
					{Type: vg.CloseComp, Pos: vg.Point{X: 0, Y: 0}, Radius: 0, Start: 0, Angle: 0},
				}},
				setColor{col: color.NRGBA{R: 0xb3, G: 0xff, B: 0xff, A: 0xff}},
				fill{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 114.8508691883805, Y: 113.78898155396111}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 54, Start: 3.940896227093295, Angle: -0.1367744033971654},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 61, Start: 3.8041218236961294, Angle: 0.1367744033971654},
					{Type: vg.CloseComp, Pos: vg.Point{X: 0, Y: 0}, Radius: 0, Start: 0, Angle: 0},
				}},
				setColor{col: color.NRGBA{R: 0x7f, G: 0xff, B: 0xff, A: 0xff}},
				fill{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 109.97042630539279, Y: 108.77088657021532}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 61, Start: 3.940896227093295, Angle: -0.1367744033971654},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 68, Start: 3.8041218236961294, Angle: 0.1367744033971654},
					{Type: vg.CloseComp, Pos: vg.Point{X: 0, Y: 0}, Radius: 0, Start: 0, Angle: 0},
				}},
				setColor{col: color.NRGBA{R: 0xff, G: 0xe6, B: 0xff, A: 0xff}},
				fill{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 105.08998342240508, Y: 103.75279158646954}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 68, Start: 3.940896227093295, Angle: -0.1367744033971654},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 75, Start: 3.8041218236961294, Angle: 0.1367744033971654},
					{Type: vg.CloseComp, Pos: vg.Point{X: 0, Y: 0}, Radius: 0, Start: 0, Angle: 0},
				}},
				setColor{col: color.NRGBA{R: 0xff, G: 0xcc, B: 0xff, A: 0xff}},
				fill{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 120.96243881336021, Y: 127.89548345528912}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 40, Start: 3.8041218236961294, Angle: -0.1367744033971654},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 47, Start: 3.667347420298964, Angle: 0.1367744033971654},
					{Type: vg.CloseComp, Pos: vg.Point{X: 0, Y: 0}, Radius: 0, Start: 0, Angle: 0},
				}},
				setColor{col: color.NRGBA{R: 0xcc, G: 0xff, B: 0xff, A: 0xff}},
				fill{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 115.44336560569823, Y: 123.58969305996473}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 47, Start: 3.8041218236961294, Angle: -0.1367744033971654},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 54, Start: 3.667347420298964, Angle: 0.1367744033971654},
					{Type: vg.CloseComp, Pos: vg.Point{X: 0, Y: 0}, Radius: 0, Start: 0, Angle: 0},
				}},
				setColor{col: color.NRGBA{R: 0xff, G: 0xe6, B: 0xff, A: 0xff}},
				fill{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 109.92429239803627, Y: 119.28390266464032}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 54, Start: 3.8041218236961294, Angle: -0.1367744033971654},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 61, Start: 3.667347420298964, Angle: 0.1367744033971654},
					{Type: vg.CloseComp, Pos: vg.Point{X: 0, Y: 0}, Radius: 0, Start: 0, Angle: 0},
				}},
				setColor{col: color.NRGBA{R: 0xff, G: 0xe6, B: 0xff, A: 0xff}},
				fill{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 104.40521919037431, Y: 114.97811226931591}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 61, Start: 3.8041218236961294, Angle: -0.1367744033971654},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 68, Start: 3.667347420298964, Angle: 0.1367744033971654},
					{Type: vg.CloseComp, Pos: vg.Point{X: 0, Y: 0}, Radius: 0, Start: 0, Angle: 0},
				}},
				setColor{col: color.NRGBA{R: 0xb3, G: 0xff, B: 0xff, A: 0xff}},
				fill{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 98.88614598271235, Y: 110.6723218739915}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 68, Start: 3.8041218236961294, Angle: -0.1367744033971654},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 75, Start: 3.667347420298964, Angle: 0.1367744033971654},
					{Type: vg.CloseComp, Pos: vg.Point{X: 0, Y: 0}, Radius: 0, Start: 0, Angle: 0},
				}},
				setColor{col: color.NRGBA{R: 0xff, G: 0xe6, B: 0xff, A: 0xff}},
				fill{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 117.90218414828041, Y: 132.425360817925}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 40, Start: 3.667347420298964, Angle: -0.1367744033971653},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 47, Start: 3.5305730169017986, Angle: 0.1367744033971653},
					{Type: vg.CloseComp, Pos: vg.Point{X: 0, Y: 0}, Radius: 0, Start: 0, Angle: 0},
				}},
				setColor{col: color.NRGBA{R: 0x99, G: 0xff, B: 0xff, A: 0xff}},
				fill{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 111.8475663742295, Y: 128.91229896106185}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 47, Start: 3.667347420298964, Angle: -0.1367744033971653},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 54, Start: 3.5305730169017986, Angle: 0.1367744033971653},
					{Type: vg.CloseComp, Pos: vg.Point{X: 0, Y: 0}, Radius: 0, Start: 0, Angle: 0},
				}},
				setColor{col: color.NRGBA{R: 0xff, G: 0xcc, B: 0xff, A: 0xff}},
				fill{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 105.79294860017856, Y: 125.39923710419873}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 54, Start: 3.667347420298964, Angle: -0.1367744033971653},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 61, Start: 3.5305730169017986, Angle: 0.1367744033971653},
					{Type: vg.CloseComp, Pos: vg.Point{X: 0, Y: 0}, Radius: 0, Start: 0, Angle: 0},
				}},
				setColor{col: color.NRGBA{R: 0xff, G: 0xcc, B: 0xff, A: 0xff}},
				fill{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 99.73833082612765, Y: 121.88617524733561}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 61, Start: 3.667347420298964, Angle: -0.1367744033971653},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 68, Start: 3.5305730169017986, Angle: 0.1367744033971653},
					{Type: vg.CloseComp, Pos: vg.Point{X: 0, Y: 0}, Radius: 0, Start: 0, Angle: 0},
				}},
				setColor{col: color.NRGBA{R: 0xb3, G: 0xff, B: 0xff, A: 0xff}},
				fill{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 93.68371305207671, Y: 118.37311339047248}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 68, Start: 3.667347420298964, Angle: -0.1367744033971653},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 75, Start: 3.5305730169017986, Angle: 0.1367744033971653},
					{Type: vg.CloseComp, Pos: vg.Point{X: 0, Y: 0}, Radius: 0, Start: 0, Angle: 0},
				}},
				setColor{col: color.NRGBA{R: 0xcc, G: 0xff, B: 0xff, A: 0xff}},
				fill{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 115.48815067793637, Y: 137.3301941422819}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 40, Start: 3.5305730169017986, Angle: -0.1367744033971655},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 47, Start: 3.393798613504633, Angle: 0.1367744033971655},
					{Type: vg.CloseComp, Pos: vg.Point{X: 0, Y: 0}, Radius: 0, Start: 0, Angle: 0},
				}},
				setColor{col: color.NRGBA{R: 0xff, G: 0xcc, B: 0xff, A: 0xff}},
				fill{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 109.01107704657524, Y: 134.67547811718123}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 47, Start: 3.5305730169017986, Angle: -0.1367744033971655},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 54, Start: 3.393798613504633, Angle: 0.1367744033971655},
					{Type: vg.CloseComp, Pos: vg.Point{X: 0, Y: 0}, Radius: 0, Start: 0, Angle: 0},
				}},
				setColor{col: color.NRGBA{R: 0xff, G: 0x7f, B: 0xff, A: 0xff}},
				fill{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 102.5340034152141, Y: 132.02076209208056}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 54, Start: 3.5305730169017986, Angle: -0.1367744033971655},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 61, Start: 3.393798613504633, Angle: 0.1367744033971655},
					{Type: vg.CloseComp, Pos: vg.Point{X: 0, Y: 0}, Radius: 0, Start: 0, Angle: 0},
				}},
				setColor{col: color.NRGBA{R: 0xb3, G: 0xff, B: 0xff, A: 0xff}},
				fill{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 96.05692978385296, Y: 129.3660460669799}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 61, Start: 3.5305730169017986, Angle: -0.1367744033971655},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 68, Start: 3.393798613504633, Angle: 0.1367744033971655},
					{Type: vg.CloseComp, Pos: vg.Point{X: 0, Y: 0}, Radius: 0, Start: 0, Angle: 0},
				}},
				setColor{col: color.NRGBA{R: 0xff, G: 0xcc, B: 0xff, A: 0xff}},
				fill{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 89.57985615249183, Y: 126.71133004187924}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 68, Start: 3.5305730169017986, Angle: -0.1367744033971655},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 75, Start: 3.393798613504633, Angle: 0.1367744033971655},
					{Type: vg.CloseComp, Pos: vg.Point{X: 0, Y: 0}, Radius: 0, Start: 0, Angle: 0},
				}},
				setColor{col: color.NRGBA{R: 0xff, G: 0xb3, B: 0xff, A: 0xff}},
				fill{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 113.76542794208467, Y: 142.51837049925388}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 40, Start: 3.393798613504633, Angle: -0.1367744033971654},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 47, Start: 3.257024210107468, Angle: 0.1367744033971654},
					{Type: vg.CloseComp, Pos: vg.Point{X: 0, Y: 0}, Radius: 0, Start: 0, Angle: 0},
				}},
				setColor{col: color.NRGBA{R: 0xb3, G: 0xff, B: 0xff, A: 0xff}},
				fill{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 106.9868778319495, Y: 140.7715853366233}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 47, Start: 3.393798613504633, Angle: -0.1367744033971654},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 54, Start: 3.257024210107468, Angle: 0.1367744033971654},
					{Type: vg.CloseComp, Pos: vg.Point{X: 0, Y: 0}, Radius: 0, Start: 0, Angle: 0},
				}},
				setColor{col: color.NRGBA{R: 0xff, G: 0xb3, B: 0xff, A: 0xff}},
				fill{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 100.2083277218143, Y: 139.02480017399276}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 54, Start: 3.393798613504633, Angle: -0.1367744033971654},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 61, Start: 3.257024210107468, Angle: 0.1367744033971654},
					{Type: vg.CloseComp, Pos: vg.Point{X: 0, Y: 0}, Radius: 0, Start: 0, Angle: 0},
				}},
				setColor{col: color.NRGBA{R: 0xe6, G: 0xff, B: 0xff, A: 0xff}},
				fill{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 93.42977761167913, Y: 137.27801501136219}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 61, Start: 3.393798613504633, Angle: -0.1367744033971654},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 68, Start: 3.257024210107468, Angle: 0.1367744033971654},
					{Type: vg.CloseComp, Pos: vg.Point{X: 0, Y: 0}, Radius: 0, Start: 0, Angle: 0},
				}},
				setColor{col: color.NRGBA{R: 0xff, G: 0xcc, B: 0xff, A: 0xff}},
				fill{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 86.65122750154394, Y: 135.5312298487316}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 68, Start: 3.393798613504633, Angle: -0.1367744033971654},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 75, Start: 3.257024210107468, Angle: 0.1367744033971654},
					{Type: vg.CloseComp, Pos: vg.Point{X: 0, Y: 0}, Radius: 0, Start: 0, Angle: 0},
				}},
				setColor{col: color.NRGBA{R: 0xe6, G: 0xff, B: 0xff, A: 0xff}},
				fill{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 112.76619311483745, Y: 147.89298465244477}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 40, Start: 3.257024210107468, Angle: -0.1367744033971654},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 47, Start: 3.1202498067103024, Angle: 0.1367744033971654},
					{Type: vg.CloseComp, Pos: vg.Point{X: 0, Y: 0}, Radius: 0, Start: 0, Angle: 0},
				}},
				setColor{col: color.NRGBA{R: 0xff, G: 0xe6, B: 0xff, A: 0xff}},
				fill{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 105.812776909934, Y: 147.0867569666226}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 47, Start: 3.257024210107468, Angle: -0.1367744033971654},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 54, Start: 3.1202498067103024, Angle: 0.1367744033971654},
					{Type: vg.CloseComp, Pos: vg.Point{X: 0, Y: 0}, Radius: 0, Start: 0, Angle: 0},
				}},
				setColor{col: color.NRGBA{R: 0xff, G: 0xe6, B: 0xff, A: 0xff}},
				fill{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 98.85936070503055, Y: 146.28052928080044}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 54, Start: 3.257024210107468, Angle: -0.1367744033971654},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 61, Start: 3.1202498067103024, Angle: 0.1367744033971654},
					{Type: vg.CloseComp, Pos: vg.Point{X: 0, Y: 0}, Radius: 0, Start: 0, Angle: 0},
				}},
				setColor{col: color.NRGBA{R: 0xff, G: 0xe6, B: 0xff, A: 0xff}},
				fill{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 91.9059445001271, Y: 145.47430159497827}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 61, Start: 3.257024210107468, Angle: -0.1367744033971654},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 68, Start: 3.1202498067103024, Angle: 0.1367744033971654},
					{Type: vg.CloseComp, Pos: vg.Point{X: 0, Y: 0}, Radius: 0, Start: 0, Angle: 0},
				}},
				setColor{col: color.NRGBA{R: 0xff, G: 0xcc, B: 0xff, A: 0xff}},
				fill{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 84.95252829522367, Y: 144.6680739091561}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 68, Start: 3.257024210107468, Angle: -0.1367744033971654},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 75, Start: 3.1202498067103024, Angle: 0.1367744033971654},
					{Type: vg.CloseComp, Pos: vg.Point{X: 0, Y: 0}, Radius: 0, Start: 0, Angle: 0},
				}},
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
			},
			actions: []interface{}{
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
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 40, Start: 4.351219437284791, Angle: 0.1367744033971654}},
				}},
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
			actions: []interface{}{
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
					{Type: vg.MoveComp, Pos: vg.Point{X: 77.99911209032022, Y: 143.86184622333394}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 81.86212109304435, Y: 144.30975049323513}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 71.11111111111111, Start: 3.257024210107468, Angle: 0.1367744033971654},
				}},
				setColor{col: color.Gray{Y: 0x0}},
				setWidth{w: 1},
				setLineDash{dashes: []vg.Length(nil), offsets: 0},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 83.63853856370608, Y: 134.75488088756248}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 87.4043997360034, Y: 135.7253170890239}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 67.22222222222223, Start: 3.393798613504633, Angle: 0.1367744033971655},
				}},
				setColor{col: color.Gray{Y: 0x0}},
				setWidth{w: 1},
				setLineDash{dashes: []vg.Length(nil), offsets: 0},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 90.29953100042084, Y: 127.00629848911265}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 93.89790524006592, Y: 128.4811407252797}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 63.33333333333333, Start: 3.5305730169017986, Angle: 0.1367744033971653},
				}},
				setColor{col: color.Gray{Y: 0x0}},
				setWidth{w: 1},
				setLineDash{dashes: []vg.Length(nil), offsets: 0},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 97.720124901444, Y: 120.71515462838124}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 101.0838014425834, Y: 122.66685565997186}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 59.44444444444444, Start: 3.667347420298964, Angle: 0.1367744033971654},
				}},
				setColor{col: color.Gray{Y: 0x0}},
				setWidth{w: 1},
				setLineDash{dashes: []vg.Length(nil), offsets: 0},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 105.63167990318809, Y: 115.934954579388}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 108.6978316852225, Y: 118.32706035456823}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 55.55555555555556, Start: 3.8041218236961294, Angle: 0.1367744033971654},
				}},
				setColor{col: color.Gray{Y: 0x0}},
				setWidth{w: 1},
				setLineDash{dashes: []vg.Length(nil), offsets: 0},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 113.76632632549435, Y: 112.67384933535092}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 116.47768348270975, Y: 115.46167988187636}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 51.666666666666664, Start: 3.940896227093295, Angle: 0.1367744033971654},
				}},
				setColor{col: color.Gray{Y: 0x0}},
				setWidth{w: 1},
				setLineDash{dashes: []vg.Length(nil), offsets: 0},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 121.86421164768274, Y: 110.89599879246629}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 124.17013120108297, Y: 114.02748275432367}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 47.77777777777778, Start: 4.07767063049046, Angle: 0.1367744033971654},
				}},
				setColor{col: color.Gray{Y: 0x0}},
				setWidth{w: 1},
				setLineDash{dashes: []vg.Length(nil), offsets: 0},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 129.6803696639582, Y: 110.524048304903}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 131.53778143549647, Y: 113.94069553589927}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 43.888888888888886, Start: 4.214445033887626, Angle: 0.1367744033971654},
				}},
				setColor{col: color.Gray{Y: 0x0}},
				setWidth{w: 1},
				setLineDash{dashes: []vg.Length(nil), offsets: 0},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 136.99104636777935, Y: 111.44263859994332}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 138.36525744911535, Y: 115.0806326480496}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 40, Start: 4.351219437284791, Angle: 0.1367744033971654},
				}},
			},
		},
		{
			orient: feat.Forward,
			scores: []Scorer{
				&fs{
					start:    b.Set[1].Start(),
					end:      b.Set[1].Start() + b.Set[1].Len()/3,
					name:     fmt.Sprintf("%s#0", b.Set[1].Name()),
					location: b.Set[1],
					scores:   []float64{1},
				},
				&fs{
					start:    b.Set[1].Start() + b.Set[1].Len()/3,
					end:      b.Set[1].Start() + b.Set[1].Len()/2,
					name:     fmt.Sprintf("%s#1", b.Set[1].Name()),
					location: b.Set[1],
					scores:   []float64{3},
				},
				&fs{
					start:    b.Set[1].Start() + b.Set[1].Len()/2,
					end:      b.Set[1].End(),
					name:     fmt.Sprintf("%s#2", b.Set[1].Name()),
					location: b.Set[1],
					scores:   []float64{2},
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
			actions: []interface{}{
				setColor{col: color.Gray{Y: 0x0}},
				setWidth{w: 1},
				setLineDash{dashes: []vg.Length(nil), offsets: 0},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 95.01315797750661, Y: 153.73003832496755}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 57.5, Start: 3.1201990508318733, Angle: 0.6838973949250415},
				}},
				setColor{col: color.Gray{Y: 0x0}},
				setWidth{w: 1},
				setLineDash{dashes: []vg.Length(nil), offsets: 0},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 107.16385821666168, Y: 117.13215799343922}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 93.36590202173261, Y: 106.36803216535552}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 75, Start: 3.8040964457569153, Angle: 0.22797002796488292},
				}},
				setColor{col: color.Gray{Y: 0x0}},
				setWidth{w: 1},
				setLineDash{dashes: []vg.Length(nil), offsets: 0},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 105.32171772278139, Y: 94.19725837174482}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 127.33824945215007, Y: 121.40520446493056}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 40, Start: 4.032066473721798, Angle: 0.4559273669601586},
				}},
			},
		},
		{
			orient: feat.Forward,
			scores: []Scorer{
				&fs{
					start:    b.Set[1].Start(),
					end:      b.Set[1].Start() + b.Set[1].Len()/3,
					name:     fmt.Sprintf("%s#0", b.Set[1].Name()),
					location: b.Set[1],
					scores:   []float64{1},
				},
				&fs{
					start:    b.Set[1].Start() + b.Set[1].Len()/3 + 1,
					end:      b.Set[1].Start() + b.Set[1].Len()/2,
					name:     fmt.Sprintf("%s#1", b.Set[1].Name()),
					location: b.Set[1],
					scores:   []float64{3},
				},
				&fs{
					start:    b.Set[1].Start() + b.Set[1].Len()/2,
					end:      b.Set[1].End(),
					name:     fmt.Sprintf("%s#2", b.Set[1].Name()),
					location: b.Set[1],
					scores:   []float64{2},
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
			actions: []interface{}{
				setColor{col: color.Gray{Y: 0x0}},
				setWidth{w: 1},
				setLineDash{dashes: []vg.Length(nil), offsets: 0},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 95.01315797750661, Y: 153.73003832496755}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 57.5, Start: 3.1201990508318733, Angle: 0.6838973949250415},
				}},
				setColor{col: color.Gray{Y: 0x0}},
				setWidth{w: 1},
				setLineDash{dashes: []vg.Length(nil), offsets: 0},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 107.16385821666168, Y: 117.13215799343922}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 93.36590202173261, Y: 106.36803216535552}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 75, Start: 3.8040964457569153, Angle: 0.2279573389952756},
				}},
				setColor{col: color.Gray{Y: 0x0}},
				setWidth{w: 1},
				setLineDash{dashes: []vg.Length(nil), offsets: 0},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 127.33824945215007, Y: 121.40520446493056}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 40, Start: 4.032066473721798, Angle: 0.4559273669601586},
				}},
			},
		},
		{
			orient: feat.Forward,
			scores: []Scorer{
				&fs{
					start:    b.Set[1].Start(),
					end:      b.Set[1].Start() + b.Set[1].Len()/3,
					name:     fmt.Sprintf("%s#0", b.Set[1].Name()),
					location: b.Set[1],
					scores:   []float64{1},
				},
				&fs{
					start:    b.Set[1].Start() + b.Set[1].Len()/3,
					end:      b.Set[1].Start() + b.Set[1].Len()/2,
					name:     fmt.Sprintf("%s#1", b.Set[1].Name()),
					location: b.Set[1],
					scores:   []float64{3},
				},
				&fs{
					start:    b.Set[1].Start() + b.Set[1].Len()/2 + 1,
					end:      b.Set[1].End(),
					name:     fmt.Sprintf("%s#2", b.Set[1].Name()),
					location: b.Set[1],
					scores:   []float64{2},
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
			actions: []interface{}{
				setColor{col: color.Gray{Y: 0x0}},
				setWidth{w: 1},
				setLineDash{dashes: []vg.Length(nil), offsets: 0},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 95.01315797750661, Y: 153.73003832496755}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 57.5, Start: 3.1201990508318733, Angle: 0.6838847059554342},
				}},
				setColor{col: color.Gray{Y: 0x0}},
				setWidth{w: 1},
				setLineDash{dashes: []vg.Length(nil), offsets: 0},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 93.36590202173261, Y: 106.36803216535552}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 75, Start: 3.8040964457569153, Angle: 0.22797002796488292},
				}},
				setColor{col: color.Gray{Y: 0x0}},
				setWidth{w: 1},
				setLineDash{dashes: []vg.Length(nil), offsets: 0},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 105.32171772278139, Y: 94.19725837174482}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 127.33824945215007, Y: 121.40520446493056}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 40, Start: 4.032066473721798, Angle: 0.4559273669601586},
				}},
			},
		},
		{
			orient: feat.Reverse,
			scores: []Scorer{
				&fs{
					start:    b.Set[1].Start(),
					end:      b.Set[1].Start() + b.Set[1].Len()/3,
					name:     fmt.Sprintf("%s#0", b.Set[1].Name()),
					location: b.Set[1],
					scores:   []float64{1},
				},
				&fs{
					start:    b.Set[1].Start() + b.Set[1].Len()/3,
					end:      b.Set[1].Start() + b.Set[1].Len()/2,
					name:     fmt.Sprintf("%s#1", b.Set[1].Name()),
					location: b.Set[1],
					scores:   []float64{3},
				},
				&fs{
					start:    b.Set[1].Start() + b.Set[1].Len()/2,
					end:      b.Set[1].End(),
					name:     fmt.Sprintf("%s#2", b.Set[1].Name()),
					location: b.Set[1],
					scores:   []float64{2},
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
			actions: []interface{}{
				setColor{col: color.Gray{Y: 0x0}},
				setWidth{w: 1},
				setLineDash{dashes: []vg.Length(nil), offsets: 0},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 112.50915337565678, Y: 153.35567883476003}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 40, Start: 3.1201990508318733, Angle: 0.4559273669601586},
				}},
				setColor{col: color.Gray{Y: 0x0}},
				setWidth{w: 1},
				setLineDash{dashes: []vg.Length(nil), offsets: 0},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 116.21734318635691, Y: 135.66049838791542}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 84.47001847441922, Y: 120.92593447734143}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 75, Start: 3.576126417792032, Angle: 0.22797002796488292},
				}},
				setColor{col: color.Gray{Y: 0x0}},
				setWidth{w: 1},
				setLineDash{dashes: []vg.Length(nil), offsets: 0},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 93.36590202173261, Y: 106.36803216535554}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 107.16385821666167, Y: 117.13215799343925}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 57.5, Start: 3.804096445756915, Angle: 0.6838973949250415},
				}},
			},
		},
		{
			orient: feat.Reverse,
			scores: []Scorer{
				&fs{
					start:    b.Set[1].Start(),
					end:      b.Set[1].Start() + b.Set[1].Len()/3,
					name:     fmt.Sprintf("%s#0", b.Set[1].Name()),
					location: b.Set[1],
					scores:   []float64{1},
				},
				&fs{
					start:    b.Set[1].Start() + b.Set[1].Len()/3 + 1,
					end:      b.Set[1].Start() + b.Set[1].Len()/2,
					name:     fmt.Sprintf("%s#1", b.Set[1].Name()),
					location: b.Set[1],
					scores:   []float64{3},
				},
				&fs{
					start:    b.Set[1].Start() + b.Set[1].Len()/2,
					end:      b.Set[1].End(),
					name:     fmt.Sprintf("%s#2", b.Set[1].Name()),
					location: b.Set[1],
					scores:   []float64{2},
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
			actions: []interface{}{
				setColor{col: color.Gray{Y: 0x0}},
				setWidth{w: 1},
				setLineDash{dashes: []vg.Length(nil), offsets: 0},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 112.50915337565678, Y: 153.35567883476003}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 40, Start: 3.1201990508318733, Angle: 0.4559273669601586},
				}},
				setColor{col: color.Gray{Y: 0x0}},
				setWidth{w: 1},
				setLineDash{dashes: []vg.Length(nil), offsets: 0},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 84.47041912225374, Y: 120.92507124951537}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 75, Start: 3.576139106761639, Angle: 0.2279573389952756},
				}},
				setColor{col: color.Gray{Y: 0x0}},
				setWidth{w: 1},
				setLineDash{dashes: []vg.Length(nil), offsets: 0},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 93.36590202173261, Y: 106.36803216535554}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 107.16385821666167, Y: 117.13215799343925}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 57.5, Start: 3.804096445756915, Angle: 0.6838973949250415}},
				}},
		},
		{
			orient: feat.Reverse,
			scores: []Scorer{
				&fs{
					start:    b.Set[1].Start(),
					end:      b.Set[1].Start() + b.Set[1].Len()/3,
					name:     fmt.Sprintf("%s#0", b.Set[1].Name()),
					location: b.Set[1],
					scores:   []float64{1},
				},
				&fs{
					start:    b.Set[1].Start() + b.Set[1].Len()/3,
					end:      b.Set[1].Start() + b.Set[1].Len()/2,
					name:     fmt.Sprintf("%s#1", b.Set[1].Name()),
					location: b.Set[1],
					scores:   []float64{3},
				},
				&fs{
					start:    b.Set[1].Start() + b.Set[1].Len()/2 + 1,
					end:      b.Set[1].End(),
					name:     fmt.Sprintf("%s#2", b.Set[1].Name()),
					location: b.Set[1],
					scores:   []float64{2},
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
			actions: []interface{}{
				setColor{col: color.Gray{Y: 0x0}},
				setWidth{w: 1},
				setLineDash{dashes: []vg.Length(nil), offsets: 0},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 112.50915337565678, Y: 153.35567883476003}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 40, Start: 3.1201990508318733, Angle: 0.4559273669601586},
				}},
				setColor{col: color.Gray{Y: 0x0}},
				setWidth{w: 1},
				setLineDash{dashes: []vg.Length(nil), offsets: 0},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 116.21734318635691, Y: 135.66049838791542}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 84.47001847441922, Y: 120.92593447734143}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 75, Start: 3.576126417792032, Angle: 0.22797002796488292},
				}},
				setColor{col: color.Gray{Y: 0x0}},
				setWidth{w: 1},
				setLineDash{dashes: []vg.Length(nil), offsets: 0},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 107.16430700178374, Y: 117.13158272736135}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 57.5, Start: 3.8041091347265223, Angle: 0.6838847059554342},
				}},
			},
		},
		{
			scores: makeScorers(b.Set[1].(*fs), 10, 2, func(_, _ int) float64 { return rand.NormFloat64() }),
			renderer: &Trace{
				LineStyles: func() []draw.LineStyle {
					sty := []draw.LineStyle{plotter.DefaultLineStyle, plotter.DefaultLineStyle}
					sty[0].Color = color.NRGBA{R: 0xff, A: 0xff}
					sty[1].Color = color.RGBA{G: 0xff, A: 0x80}
					return sty
				}(),
			},
			actions: []interface{}{
				setColor{col: color.NRGBA{R: 0xff, G: 0x0, B: 0x0, A: 0xff}},
				setWidth{w: 1},
				setLineDash{dashes: []vg.Length(nil), offsets: 0},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 106.42203131113078, Y: 153.48358438128594}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 46.088465332743745, Start: 3.1202498067103024, Angle: 0.1367744033971654},
				}},
				setColor{col: color.RGBA{R: 0x0, G: 0xff, B: 0x0, A: 0x80}},
				setWidth{w: 1},
				setLineDash{dashes: []vg.Length(nil), offsets: 0},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 100.66325238128137, Y: 153.60651178394392}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 51.84855612275754, Start: 3.1202498067103024, Angle: 0.1367744033971654},
				}},
				setColor{col: color.NRGBA{R: 0xff, G: 0x0, B: 0x0, A: 0xff}},
				setWidth{w: 1},
				setLineDash{dashes: []vg.Length(nil), offsets: 0},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 98.00316042721153, Y: 146.18125543493176}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 54.861936315635305, Start: 3.257024210107468, Angle: 0.1367744033971654},
				}},
				setColor{col: color.RGBA{R: 0x0, G: 0xff, B: 0x0, A: 0x80}},
				setWidth{w: 1},
				setLineDash{dashes: []vg.Length(nil), offsets: 0},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 83.47726274439339, Y: 144.4970213076765}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 69.48514896153198, Start: 3.257024210107468, Angle: 0.1367744033971654},
				}},
				setColor{col: color.NRGBA{R: 0xff, G: 0x0, B: 0x0, A: 0xff}},
				setWidth{w: 1},
				setLineDash{dashes: []vg.Length(nil), offsets: 0},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 79.87267739140876, Y: 133.78444468610104}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 75, Start: 3.393798613504633, Angle: 0.1367744033971655},
				}},
				setColor{col: color.RGBA{R: 0x0, G: 0xff, B: 0x0, A: 0x80}},
				setWidth{w: 1},
				setLineDash{dashes: []vg.Length(nil), offsets: 0},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 100.00010911484128, Y: 138.9711436888544}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 54.215020944763964, Start: 3.393798613504633, Angle: 0.1367744033971655},
				}},
				setColor{col: color.NRGBA{R: 0xff, G: 0x0, B: 0x0, A: 0xff}},
				setWidth{w: 1},
				setLineDash{dashes: []vg.Length(nil), offsets: 0},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 102.13189097188692, Y: 131.85595088134733}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 54.43457698082379, Start: 3.5305730169017986, Angle: 0.1367744033971653},
				}},
				setColor{col: color.RGBA{R: 0x0, G: 0xff, B: 0x0, A: 0x80}},
				setWidth{w: 1},
				setLineDash{dashes: []vg.Length(nil), offsets: 0},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 98.13649424314028, Y: 130.21838362086683}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 58.752541958990804, Start: 3.5305730169017986, Angle: 0.1367744033971653},
				}},
				setColor{col: color.NRGBA{R: 0xff, G: 0x0, B: 0x0, A: 0xff}},
				setWidth{w: 1},
				setLineDash{dashes: []vg.Length(nil), offsets: 0},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 106.16550357279438, Y: 125.61540412129611}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 53.569273419787514, Start: 3.667347420298964, Angle: 0.1367744033971654},
				}},
				setColor{col: color.RGBA{R: 0x0, G: 0xff, B: 0x0, A: 0x80}},
				setWidth{w: 1},
				setLineDash{dashes: []vg.Length(nil), offsets: 0},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 103.77523046391377, Y: 124.2284995139924}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 56.3327693804862, Start: 3.667347420298964, Angle: 0.1367744033971654},
				}},
				setColor{col: color.NRGBA{R: 0xff, G: 0x0, B: 0x0, A: 0xff}},
				setWidth{w: 1},
				setLineDash{dashes: []vg.Length(nil), offsets: 0},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 95.9460872748634, Y: 108.37858956249192}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 71.72896139996323, Start: 3.8041218236961294, Angle: 0.1367744033971654},
				}},
				setColor{col: color.RGBA{R: 0x0, G: 0xff, B: 0x0, A: 0x80}},
				setWidth{w: 1},
				setLineDash{dashes: []vg.Length(nil), offsets: 0},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 110.57729514584946, Y: 119.79335286648781}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 53.171777749143374, Start: 3.8041218236961294, Angle: 0.1367744033971654},
				}},
				setColor{col: color.NRGBA{R: 0xff, G: 0x0, B: 0x0, A: 0xff}},
				setWidth{w: 1},
				setLineDash{dashes: []vg.Length(nil), offsets: 0},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 118.3748140107319, Y: 117.4123186690577}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 48.94561982428966, Start: 3.940896227093295, Angle: 0.1367744033971654},
				}},
				setColor{col: color.RGBA{R: 0x0, G: 0xff, B: 0x0, A: 0x80}},
				setWidth{w: 1},
				setLineDash{dashes: []vg.Length(nil), offsets: 0},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 107.90197344763634, Y: 106.64409333474813}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 63.96677378497078, Start: 3.940896227093295, Angle: 0.1367744033971654},
				}},
				setColor{col: color.NRGBA{R: 0xff, G: 0x0, B: 0x0, A: 0xff}},
				setWidth{w: 1},
				setLineDash{dashes: []vg.Length(nil), offsets: 0},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 123.44289448732908, Y: 113.03988103932193}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 49.004248480773136, Start: 4.07767063049046, Angle: 0.1367744033971654},
				}},
				setColor{col: color.RGBA{R: 0x0, G: 0xff, B: 0x0, A: 0x80}},
				setWidth{w: 1},
				setLineDash{dashes: []vg.Length(nil), offsets: 0},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 125.6123529213532, Y: 115.98604812569974}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 45.34549863993757, Start: 4.07767063049046, Angle: 0.1367744033971654},
				}},
				setColor{col: color.NRGBA{R: 0xff, G: 0x0, B: 0x0, A: 0xff}},
				setWidth{w: 1},
				setLineDash{dashes: []vg.Length(nil), offsets: 0},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 119.5588498864459, Y: 91.9058490174767}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 68.96934466918282, Start: 4.214445033887626, Angle: 0.1367744033971654},
				}},
				setColor{col: color.RGBA{R: 0x0, G: 0xff, B: 0x0, A: 0x80}},
				setWidth{w: 1},
				setLineDash{dashes: []vg.Length(nil), offsets: 0},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 125.64884009622823, Y: 103.10818363486462}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 56.21864736922418, Start: 4.214445033887626, Angle: 0.1367744033971654},
				}},
				setColor{col: color.NRGBA{R: 0xff, G: 0x0, B: 0x0, A: 0xff}},
				setWidth{w: 1},
				setLineDash{dashes: []vg.Length(nil), offsets: 0},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 138.36525744911535, Y: 115.0806326480496}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 40, Start: 4.351219437284791, Angle: 0.1367744033971654},
				}},
				setColor{col: color.RGBA{R: 0x0, G: 0xff, B: 0x0, A: 0x80}},
				setWidth{w: 1},
				setLineDash{dashes: []vg.Length(nil), offsets: 0},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 129.29604494210653, Y: 91.07140905087653}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 65.66502353858921, Start: 4.351219437284791, Angle: 0.1367744033971654},
				}},
			},
		},
		{
			scores: makeScorers(b.Set[1].(*fs), 10, 2, func(_, _ int) float64 { return rand.NormFloat64() }),
			renderer: &Trace{
				LineStyles: func() []draw.LineStyle {
					sty := []draw.LineStyle{plotter.DefaultLineStyle, plotter.DefaultLineStyle}
					sty[0].Color = color.NRGBA{R: 0xff, A: 0xff}
					sty[1].Color = color.RGBA{G: 0xff, A: 0x80}
					return sty
				}(),
				Join: true,
			},
			actions: []interface{}{
				setColor{col: color.NRGBA{R: 0xff, G: 0x0, B: 0x0, A: 0xff}},
				setWidth{w: 1},
				setLineDash{dashes: []vg.Length(nil), offsets: 0},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 90.53810547581148, Y: 153.82264406229348}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 61.97600955484456, Start: 3.1202498067103024, Angle: 0.1367744033971654},
				}},
				setColor{col: color.RGBA{R: 0x0, G: 0xff, B: 0x0, A: 0x80}},
				setWidth{w: 1},
				setLineDash{dashes: []vg.Length(nil), offsets: 0},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 101.58451940234829, Y: 153.58684633690493}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 50.927079235411604, Start: 3.1202498067103024, Angle: 0.1367744033971654},
				}},
				setColor{col: color.NRGBA{R: 0xff, G: 0x0, B: 0x0, A: 0xff}},
				setWidth{w: 1},
				setLineDash{dashes: []vg.Length(nil), offsets: 0},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 90.93643012087043, Y: 145.36188932001502}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 110.07770693050352, Y: 147.58126276914487}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 42.70649745905707, Start: 3.257024210107468, Angle: 0.1367744033971654},
				}},
				setColor{col: color.RGBA{R: 0x0, G: 0xff, B: 0x0, A: 0x80}},
				setWidth{w: 1},
				setLineDash{dashes: []vg.Length(nil), offsets: 0},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 101.91183171086959, Y: 146.63445410890742}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 89.00466214778763, Y: 145.13790663631121}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 63.92071923611516, Start: 3.257024210107468, Angle: 0.1367744033971654},
				}},
				setColor{col: color.NRGBA{R: 0xff, G: 0x0, B: 0x0, A: 0xff}},
				setWidth{w: 1},
				setLineDash{dashes: []vg.Length(nil), offsets: 0},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 111.1445524207744, Y: 141.84298912722844}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 87.82359451519339, Y: 135.83334066466236}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 66.78933267996709, Start: 3.393798613504633, Angle: 0.1367744033971655},
				}},
				setColor{col: color.RGBA{R: 0x0, G: 0xff, B: 0x0, A: 0x80}},
				setWidth{w: 1},
				setLineDash{dashes: []vg.Length(nil), offsets: 0},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 90.60145736887307, Y: 136.5491765790971}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 93.54274517016367, Y: 137.3071259645621}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 60.88334188041047, Start: 3.393798613504633, Angle: 0.1367744033971655},
				}},
				setColor{col: color.NRGBA{R: 0xff, G: 0x0, B: 0x0, A: 0xff}},
				setWidth{w: 1},
				setLineDash{dashes: []vg.Length(nil), offsets: 0},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 90.70008206319696, Y: 127.17046974695882}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 108.55396890509792, Y: 134.4881261774861}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 47.494012755212246, Start: 3.5305730169017986, Angle: 0.1367744033971653},
				}},
				setColor{col: color.RGBA{R: 0x0, G: 0xff, B: 0x0, A: 0x80}},
				setWidth{w: 1},
				setLineDash{dashes: []vg.Length(nil), offsets: 0},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 96.16487310246404, Y: 129.4102880926274}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 115.48815067793637, Y: 137.3301941422819}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 40, Start: 3.5305730169017986, Angle: 0.1367744033971653},
				}},
				setColor{col: color.NRGBA{R: 0xff, G: 0x0, B: 0x0, A: 0xff}},
				setWidth{w: 1},
				setLineDash{dashes: []vg.Length(nil), offsets: 0},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 111.42027231589864, Y: 128.66437076575613}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 107.45743046568425, Y: 126.36501693892414}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 52.075622030431774, Start: 3.667347420298964, Angle: 0.1367744033971654},
				}},
				setColor{col: color.RGBA{R: 0x0, G: 0xff, B: 0x0, A: 0x80}},
				setWidth{w: 1},
				setLineDash{dashes: []vg.Length(nil), offsets: 0},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 117.90218414828041, Y: 132.425360817925}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 100.85844725334994, Y: 122.53609873426245}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 59.70498596556162, Start: 3.667347420298964, Angle: 0.1367744033971654},
				}},
				setColor{col: color.NRGBA{R: 0xff, G: 0x0, B: 0x0, A: 0xff}},
				setWidth{w: 1},
				setLineDash{dashes: []vg.Length(nil), offsets: 0},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 111.44154709707327, Y: 120.46761240440328}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 98.37240845370792, Y: 110.27152158835042}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 68.65158815035078, Start: 3.8041218236961294, Angle: 0.1367744033971654},
				}},
				setColor{col: color.RGBA{R: 0x0, G: 0xff, B: 0x0, A: 0x80}},
				setWidth{w: 1},
				setLineDash{dashes: []vg.Length(nil), offsets: 0},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 105.42625879909076, Y: 115.77469212521521}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 112.51850005161984, Y: 121.30781392256131}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 50.70969148409288, Start: 3.8041218236961294, Angle: 0.1367744033971654},
				}},
				setColor{col: color.NRGBA{R: 0xff, G: 0x0, B: 0x0, A: 0xff}},
				setWidth{w: 1},
				setLineDash{dashes: []vg.Length(nil), offsets: 0},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 104.63569217225957, Y: 103.28568712522048}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 111.68001463366127, Y: 110.52869356028557}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 58.54794419588548, Start: 3.940896227093295, Angle: 0.1367744033971654},
				}},
				setColor{col: color.RGBA{R: 0x0, G: 0xff, B: 0x0, A: 0x80}},
				setWidth{w: 1},
				setLineDash{dashes: []vg.Length(nil), offsets: 0},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 117.14489244256518, Y: 116.14770736233962}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 112.45431574879328, Y: 111.32483371291328}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 57.43736715690048, Start: 3.940896227093295, Angle: 0.1367744033971654},
				}},
				setColor{col: color.NRGBA{R: 0xff, G: 0x0, B: 0x0, A: 0xff}},
				setWidth{w: 1},
				setLineDash{dashes: []vg.Length(nil), offsets: 0},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 117.78395302874007, Y: 105.35492759307931}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 108.02869432728139, Y: 92.10709502132204}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 75, Start: 4.07767063049046, Angle: 0.1367744033971654},
				}},
				setColor{col: color.RGBA{R: 0x0, G: 0xff, B: 0x0, A: 0x80}},
				setWidth{w: 1},
				setLineDash{dashes: []vg.Length(nil), offsets: 0},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 118.44247050839081, Y: 106.24920724090495}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 117.12965586238778, Y: 104.46637921345365}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 59.65140375782332, Start: 4.07767063049046, Angle: 0.1367744033971654},
				}},
				setColor{col: color.NRGBA{R: 0xff, G: 0x0, B: 0x0, A: 0xff}},
				setWidth{w: 1},
				setLineDash{dashes: []vg.Length(nil), offsets: 0},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 116.67848726319019, Y: 86.60751768792913}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 131.98108700494163, Y: 114.75614139962155}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 42.960733845502894, Start: 4.214445033887626, Angle: 0.1367744033971654},
				}},
				setColor{col: color.RGBA{R: 0x0, G: 0xff, B: 0x0, A: 0x80}},
				setWidth{w: 1},
				setLineDash{dashes: []vg.Length(nil), offsets: 0},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 124.00928640694063, Y: 100.09227910663239}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 127.22778100678856, Y: 106.01259297082186}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 52.91279679942573, Start: 4.214445033887626, Angle: 0.1367744033971654},
				}},
				setColor{col: color.NRGBA{R: 0xff, G: 0x0, B: 0x0, A: 0xff}},
				setWidth{w: 1},
				setLineDash{dashes: []vg.Length(nil), offsets: 0},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 137.3190271824185, Y: 112.31091296314398}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 133.63616264540116, Y: 102.5611449343306}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 53.38289618417768, Start: 4.351219437284791, Angle: 0.1367744033971654},
				}},
				setColor{col: color.RGBA{R: 0x0, G: 0xff, B: 0x0, A: 0x80}},
				setWidth{w: 1},
				setLineDash{dashes: []vg.Length(nil), offsets: 0},
				stroke{path: vg.Path{
					{Type: vg.MoveComp, Pos: vg.Point{X: 133.8022809898211, Y: 103.0009154735796}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.LineComp, Pos: vg.Point{X: 137.13376676648508, Y: 111.82046699041888}, Radius: 0, Start: 0, Angle: 0},
					{Type: vg.ArcComp, Pos: vg.Point{X: 152.5, Y: 152.5}, Radius: 43.48500350309722, Start: 4.351219437284791, Angle: 0.1367744033971654},
				}},
			},
		},
	} {
		p, err := plot.New()
		if err != nil {
			t.Fatalf("unexpected error for plot.New: %v", err)
		}

		b.Set[1].(*fs).orient = test.orient
		b.Base = NewGappedArcs(b.Base, b.Set, 0.01)
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
			err = p.Save(vg.Length(300), vg.Length(300), fmt.Sprintf("scores-%d-%s.svg", i, failure(!ok)))
			if err != nil {
				t.Fatalf("unexpected error writing file: %v", err)
			}
		}
	}
}
