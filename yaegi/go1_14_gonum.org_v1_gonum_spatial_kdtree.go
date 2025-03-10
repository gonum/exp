// Code generated by 'goexports gonum.org/v1/gonum/spatial/kdtree'. DO NOT EDIT.

// Copyright ©2019 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build go1.14 && !go1.15
// +build go1.14,!go1.15

package yaegi

import (
	"reflect"

	"gonum.org/v1/gonum/spatial/kdtree"
)

func init() {
	Symbols["gonum.org/v1/gonum/spatial/kdtree"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"MedianOfMedians": reflect.ValueOf(kdtree.MedianOfMedians),
		"MedianOfRandoms": reflect.ValueOf(kdtree.MedianOfRandoms),
		"New":             reflect.ValueOf(kdtree.New),
		"NewDistKeeper":   reflect.ValueOf(kdtree.NewDistKeeper),
		"NewNKeeper":      reflect.ValueOf(kdtree.NewNKeeper),
		"Partition":       reflect.ValueOf(kdtree.Partition),
		"Select":          reflect.ValueOf(kdtree.Select),

		// type definitions
		"Bounder":        reflect.ValueOf((*kdtree.Bounder)(nil)),
		"Bounding":       reflect.ValueOf((*kdtree.Bounding)(nil)),
		"Comparable":     reflect.ValueOf((*kdtree.Comparable)(nil)),
		"ComparableDist": reflect.ValueOf((*kdtree.ComparableDist)(nil)),
		"Dim":            reflect.ValueOf((*kdtree.Dim)(nil)),
		"DistKeeper":     reflect.ValueOf((*kdtree.DistKeeper)(nil)),
		"Extender":       reflect.ValueOf((*kdtree.Extender)(nil)),
		"Heap":           reflect.ValueOf((*kdtree.Heap)(nil)),
		"Interface":      reflect.ValueOf((*kdtree.Interface)(nil)),
		"Keeper":         reflect.ValueOf((*kdtree.Keeper)(nil)),
		"NKeeper":        reflect.ValueOf((*kdtree.NKeeper)(nil)),
		"Node":           reflect.ValueOf((*kdtree.Node)(nil)),
		"Operation":      reflect.ValueOf((*kdtree.Operation)(nil)),
		"Plane":          reflect.ValueOf((*kdtree.Plane)(nil)),
		"Point":          reflect.ValueOf((*kdtree.Point)(nil)),
		"Points":         reflect.ValueOf((*kdtree.Points)(nil)),
		"SortSlicer":     reflect.ValueOf((*kdtree.SortSlicer)(nil)),
		"Tree":           reflect.ValueOf((*kdtree.Tree)(nil)),

		// interface wrapper definitions
		"_Bounder":    reflect.ValueOf((*_gonum_org_v1_gonum_spatial_kdtree_Bounder)(nil)),
		"_Comparable": reflect.ValueOf((*_gonum_org_v1_gonum_spatial_kdtree_Comparable)(nil)),
		"_Extender":   reflect.ValueOf((*_gonum_org_v1_gonum_spatial_kdtree_Extender)(nil)),
		"_Interface":  reflect.ValueOf((*_gonum_org_v1_gonum_spatial_kdtree_Interface)(nil)),
		"_Keeper":     reflect.ValueOf((*_gonum_org_v1_gonum_spatial_kdtree_Keeper)(nil)),
		"_SortSlicer": reflect.ValueOf((*_gonum_org_v1_gonum_spatial_kdtree_SortSlicer)(nil)),
	}
}

// _gonum_org_v1_gonum_spatial_kdtree_Bounder is an interface wrapper for Bounder type
type _gonum_org_v1_gonum_spatial_kdtree_Bounder struct {
	WBounds func() *kdtree.Bounding
}

func (W _gonum_org_v1_gonum_spatial_kdtree_Bounder) Bounds() *kdtree.Bounding { return W.WBounds() }

// _gonum_org_v1_gonum_spatial_kdtree_Comparable is an interface wrapper for Comparable type
type _gonum_org_v1_gonum_spatial_kdtree_Comparable struct {
	WCompare  func(a0 kdtree.Comparable, a1 kdtree.Dim) float64
	WDims     func() int
	WDistance func(a0 kdtree.Comparable) float64
}

func (W _gonum_org_v1_gonum_spatial_kdtree_Comparable) Compare(a0 kdtree.Comparable, a1 kdtree.Dim) float64 {
	return W.WCompare(a0, a1)
}
func (W _gonum_org_v1_gonum_spatial_kdtree_Comparable) Dims() int { return W.WDims() }
func (W _gonum_org_v1_gonum_spatial_kdtree_Comparable) Distance(a0 kdtree.Comparable) float64 {
	return W.WDistance(a0)
}

// _gonum_org_v1_gonum_spatial_kdtree_Extender is an interface wrapper for Extender type
type _gonum_org_v1_gonum_spatial_kdtree_Extender struct {
	WCompare  func(a0 kdtree.Comparable, a1 kdtree.Dim) float64
	WDims     func() int
	WDistance func(a0 kdtree.Comparable) float64
	WExtend   func(a0 *kdtree.Bounding) *kdtree.Bounding
}

func (W _gonum_org_v1_gonum_spatial_kdtree_Extender) Compare(a0 kdtree.Comparable, a1 kdtree.Dim) float64 {
	return W.WCompare(a0, a1)
}
func (W _gonum_org_v1_gonum_spatial_kdtree_Extender) Dims() int { return W.WDims() }
func (W _gonum_org_v1_gonum_spatial_kdtree_Extender) Distance(a0 kdtree.Comparable) float64 {
	return W.WDistance(a0)
}
func (W _gonum_org_v1_gonum_spatial_kdtree_Extender) Extend(a0 *kdtree.Bounding) *kdtree.Bounding {
	return W.WExtend(a0)
}

// _gonum_org_v1_gonum_spatial_kdtree_Interface is an interface wrapper for Interface type
type _gonum_org_v1_gonum_spatial_kdtree_Interface struct {
	WIndex func(i int) kdtree.Comparable
	WLen   func() int
	WPivot func(a0 kdtree.Dim) int
	WSlice func(start int, end int) kdtree.Interface
}

func (W _gonum_org_v1_gonum_spatial_kdtree_Interface) Index(i int) kdtree.Comparable {
	return W.WIndex(i)
}
func (W _gonum_org_v1_gonum_spatial_kdtree_Interface) Len() int                { return W.WLen() }
func (W _gonum_org_v1_gonum_spatial_kdtree_Interface) Pivot(a0 kdtree.Dim) int { return W.WPivot(a0) }
func (W _gonum_org_v1_gonum_spatial_kdtree_Interface) Slice(start int, end int) kdtree.Interface {
	return W.WSlice(start, end)
}

// _gonum_org_v1_gonum_spatial_kdtree_Keeper is an interface wrapper for Keeper type
type _gonum_org_v1_gonum_spatial_kdtree_Keeper struct {
	WKeep func(a0 kdtree.ComparableDist)
	WLen  func() int
	WLess func(i int, j int) bool
	WMax  func() kdtree.ComparableDist
	WPop  func() interface{}
	WPush func(x interface{})
	WSwap func(i int, j int)
}

func (W _gonum_org_v1_gonum_spatial_kdtree_Keeper) Keep(a0 kdtree.ComparableDist) { W.WKeep(a0) }
func (W _gonum_org_v1_gonum_spatial_kdtree_Keeper) Len() int                      { return W.WLen() }
func (W _gonum_org_v1_gonum_spatial_kdtree_Keeper) Less(i int, j int) bool        { return W.WLess(i, j) }
func (W _gonum_org_v1_gonum_spatial_kdtree_Keeper) Max() kdtree.ComparableDist    { return W.WMax() }
func (W _gonum_org_v1_gonum_spatial_kdtree_Keeper) Pop() interface{}              { return W.WPop() }
func (W _gonum_org_v1_gonum_spatial_kdtree_Keeper) Push(x interface{})            { W.WPush(x) }
func (W _gonum_org_v1_gonum_spatial_kdtree_Keeper) Swap(i int, j int)             { W.WSwap(i, j) }

// _gonum_org_v1_gonum_spatial_kdtree_SortSlicer is an interface wrapper for SortSlicer type
type _gonum_org_v1_gonum_spatial_kdtree_SortSlicer struct {
	WLen   func() int
	WLess  func(i int, j int) bool
	WSlice func(start int, end int) kdtree.SortSlicer
	WSwap  func(i int, j int)
}

func (W _gonum_org_v1_gonum_spatial_kdtree_SortSlicer) Len() int               { return W.WLen() }
func (W _gonum_org_v1_gonum_spatial_kdtree_SortSlicer) Less(i int, j int) bool { return W.WLess(i, j) }
func (W _gonum_org_v1_gonum_spatial_kdtree_SortSlicer) Slice(start int, end int) kdtree.SortSlicer {
	return W.WSlice(start, end)
}
func (W _gonum_org_v1_gonum_spatial_kdtree_SortSlicer) Swap(i int, j int) { W.WSwap(i, j) }
