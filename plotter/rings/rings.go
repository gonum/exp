// Copyright ©2018 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package rings implements a number of graphical representations of genomic features
// and feature associations using the idioms developed in the Circos distribution.
//
// The rings package borrows significantly from the ideas of Circos and shares some implementation
// details in order to run as a work-a-like. Circos is available from http://circos.ca/.
package rings

import (
	"image/color"

	"gonum.org/v1/plot/vg/draw"
)

// Twist is a flag type used to specify Ribbon and Sail twist behaviour. Specific interpretation
// of Twist flags is documented in the relevant types.
type Twist uint

const (
	None       Twist = 0         // None indicates no explicit twist.
	Flat       Twist = 1 << iota // Render feature connections without twist.
	Individual                   // Individual allows a feature or feature pair to define its ribbon twist.
	Twisted                      // Twisted specifies that feature connections render with twist.
	Invert                       // Invert inverts all twist behaviour.
)

// ColorFunc allows dynamic assignment of color to objects based on passed parameters.
type ColorFunc func(interface{}) color.Color

// LineStyleFunc allows dynamic assignment of line styles to objects based on passed parameters.
type LineStyleFunc func(interface{}) draw.LineStyle

// Feature is a Range whose coordinates are defined relative to a feature
// location/parent. Start and End return the coordinates of the feature
// relative to its parent which can be nil. In the latter case callers
// should make no assumptions whether coordinates of such features are
// comparable.
type Feature interface {
	// Start and End indicate the position of the feature within the
	// containing Parent's coordinate system.
	Start() float64
	End() float64

	// Name returns the name of the feature.
	Name() string

	// Parent returns the reference feature on which the feature is located.
	Parent() Feature
}

func lengthOf(f Feature) float64 {
	return f.End() - f.Start()
}

// Conformationer wraps the Conformation method.
type Conformationer interface {
	Conformation() Conformation
}

// Conformation describes whether a feature is linear or circular.
type Conformation int8

func (c Conformation) String() string {
	switch c {
	case UndefinedConformation:
		return "undefined"
	case Linear:
		return "linear"
	case Circular:
		return "circular"
	}
	panic("rings: illegal conformation")
}

const (
	UndefinedConformation Conformation = iota - 1
	Linear
	Circular
)

// Orienter wraps the Orientation method.
type Orienter interface {
	Orientation() Orientation
}

// Orientation describes whether a feature is oriented forwards or backwards.
type Orientation int8

func (o Orientation) String() string {
	switch o {
	case Backward:
		return "backward"
	case NotOriented:
		return "not oriented"
	case Forward:
		return "forward"
	}
	panic("rings: illegal orientation")
}

const (
	Backward Orientation = iota - 1
	NotOriented
	Forward
)

// Pair represents a pair of associated features.
type Pair interface {
	Features() [2]Feature
}

// TextStyler is a type that can define its text style. For the purposes of the rings package
// the lines of a LineStyler that returns a nil Color or a TextStyle with Font.Size of 0 are not rendered.
type TextStyler interface {
	TextStyle() draw.TextStyle
}

// LineStyler is a type that can define its drawing line style. For the purposes of the rings package
// the lines of a LineStyler that returns a nil Color or a LineStyle with width 0 are not rendered.
type LineStyler interface {
	LineStyle() draw.LineStyle
}

// FillColorer is a type that can define its fill color. For the purposes of the rings package
// a FillColoer that returns a nil Color is not rendered filled.
type FillColorer interface {
	FillColor() color.Color
}

// XYer is a type that returns its x and y coordinates.
type XYer interface {
	XY() (x, y float64)
}
