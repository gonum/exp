// Code generated by 'goexports gonum.org/v1/gonum/graph/testgraph'. DO NOT EDIT.

// Copyright ©2019 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build go1.12,!go1.13

package yaegi

import (
	"reflect"

	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/testgraph"
)

func init() {
	Symbols["gonum.org/v1/gonum/graph/testgraph"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"AddArbitraryNodes":       reflect.ValueOf(testgraph.AddArbitraryNodes),
		"AddEdges":                reflect.ValueOf(testgraph.AddEdges),
		"AddLines":                reflect.ValueOf(testgraph.AddLines),
		"AddNodes":                reflect.ValueOf(testgraph.AddNodes),
		"AddWeightedEdges":        reflect.ValueOf(testgraph.AddWeightedEdges),
		"AddWeightedLines":        reflect.ValueOf(testgraph.AddWeightedLines),
		"AdjacencyMatrix":         reflect.ValueOf(testgraph.AdjacencyMatrix),
		"EdgeExistence":           reflect.ValueOf(testgraph.EdgeExistence),
		"LineExistence":           reflect.ValueOf(testgraph.LineExistence),
		"NewRandomNodes":          reflect.ValueOf(testgraph.NewRandomNodes),
		"NoLoopAddEdges":          reflect.ValueOf(testgraph.NoLoopAddEdges),
		"NoLoopAddWeightedEdges":  reflect.ValueOf(testgraph.NoLoopAddWeightedEdges),
		"NodeExistence":           reflect.ValueOf(testgraph.NodeExistence),
		"RemoveEdges":             reflect.ValueOf(testgraph.RemoveEdges),
		"RemoveLines":             reflect.ValueOf(testgraph.RemoveLines),
		"RemoveNodes":             reflect.ValueOf(testgraph.RemoveNodes),
		"ReturnAdjacentNodes":     reflect.ValueOf(testgraph.ReturnAdjacentNodes),
		"ReturnAllEdges":          reflect.ValueOf(testgraph.ReturnAllEdges),
		"ReturnAllLines":          reflect.ValueOf(testgraph.ReturnAllLines),
		"ReturnAllNodes":          reflect.ValueOf(testgraph.ReturnAllNodes),
		"ReturnAllWeightedEdges":  reflect.ValueOf(testgraph.ReturnAllWeightedEdges),
		"ReturnAllWeightedLines":  reflect.ValueOf(testgraph.ReturnAllWeightedLines),
		"ReturnEdgeSlice":         reflect.ValueOf(testgraph.ReturnEdgeSlice),
		"ReturnNodeSlice":         reflect.ValueOf(testgraph.ReturnNodeSlice),
		"ReturnWeightedEdgeSlice": reflect.ValueOf(testgraph.ReturnWeightedEdgeSlice),
		"Weight":                  reflect.ValueOf(testgraph.Weight),

		// type definitions
		"Builder":           reflect.ValueOf((*testgraph.Builder)(nil)),
		"Edge":              reflect.ValueOf((*testgraph.Edge)(nil)),
		"EdgeAdder":         reflect.ValueOf((*testgraph.EdgeAdder)(nil)),
		"EdgeRemover":       reflect.ValueOf((*testgraph.EdgeRemover)(nil)),
		"LineAdder":         reflect.ValueOf((*testgraph.LineAdder)(nil)),
		"LineRemover":       reflect.ValueOf((*testgraph.LineRemover)(nil)),
		"NodeAdder":         reflect.ValueOf((*testgraph.NodeAdder)(nil)),
		"NodeRemover":       reflect.ValueOf((*testgraph.NodeRemover)(nil)),
		"RandomNodes":       reflect.ValueOf((*testgraph.RandomNodes)(nil)),
		"WeightedEdgeAdder": reflect.ValueOf((*testgraph.WeightedEdgeAdder)(nil)),
		"WeightedLine":      reflect.ValueOf((*testgraph.WeightedLine)(nil)),
		"WeightedLineAdder": reflect.ValueOf((*testgraph.WeightedLineAdder)(nil)),

		// interface wrapper definitions
		"_Edge":              reflect.ValueOf((*_gonum_org_v1_gonum_graph_testgraph_Edge)(nil)),
		"_EdgeAdder":         reflect.ValueOf((*_gonum_org_v1_gonum_graph_testgraph_EdgeAdder)(nil)),
		"_EdgeRemover":       reflect.ValueOf((*_gonum_org_v1_gonum_graph_testgraph_EdgeRemover)(nil)),
		"_LineAdder":         reflect.ValueOf((*_gonum_org_v1_gonum_graph_testgraph_LineAdder)(nil)),
		"_LineRemover":       reflect.ValueOf((*_gonum_org_v1_gonum_graph_testgraph_LineRemover)(nil)),
		"_NodeAdder":         reflect.ValueOf((*_gonum_org_v1_gonum_graph_testgraph_NodeAdder)(nil)),
		"_NodeRemover":       reflect.ValueOf((*_gonum_org_v1_gonum_graph_testgraph_NodeRemover)(nil)),
		"_WeightedEdgeAdder": reflect.ValueOf((*_gonum_org_v1_gonum_graph_testgraph_WeightedEdgeAdder)(nil)),
		"_WeightedLine":      reflect.ValueOf((*_gonum_org_v1_gonum_graph_testgraph_WeightedLine)(nil)),
		"_WeightedLineAdder": reflect.ValueOf((*_gonum_org_v1_gonum_graph_testgraph_WeightedLineAdder)(nil)),
	}
}

// _gonum_org_v1_gonum_graph_testgraph_Edge is an interface wrapper for Edge type
type _gonum_org_v1_gonum_graph_testgraph_Edge struct {
	WFrom func() graph.Node
	WTo   func() graph.Node
}

func (W _gonum_org_v1_gonum_graph_testgraph_Edge) From() graph.Node { return W.WFrom() }
func (W _gonum_org_v1_gonum_graph_testgraph_Edge) To() graph.Node   { return W.WTo() }

// _gonum_org_v1_gonum_graph_testgraph_EdgeAdder is an interface wrapper for EdgeAdder type
type _gonum_org_v1_gonum_graph_testgraph_EdgeAdder struct {
	WEdge           func(uid int64, vid int64) graph.Edge
	WFrom           func(id int64) graph.Nodes
	WHasEdgeBetween func(xid int64, yid int64) bool
	WNewEdge        func(from graph.Node, to graph.Node) graph.Edge
	WNode           func(id int64) graph.Node
	WNodes          func() graph.Nodes
	WSetEdge        func(e graph.Edge)
}

func (W _gonum_org_v1_gonum_graph_testgraph_EdgeAdder) Edge(uid int64, vid int64) graph.Edge {
	return W.WEdge(uid, vid)
}
func (W _gonum_org_v1_gonum_graph_testgraph_EdgeAdder) From(id int64) graph.Nodes {
	return W.WFrom(id)
}
func (W _gonum_org_v1_gonum_graph_testgraph_EdgeAdder) HasEdgeBetween(xid int64, yid int64) bool {
	return W.WHasEdgeBetween(xid, yid)
}
func (W _gonum_org_v1_gonum_graph_testgraph_EdgeAdder) NewEdge(from graph.Node, to graph.Node) graph.Edge {
	return W.WNewEdge(from, to)
}
func (W _gonum_org_v1_gonum_graph_testgraph_EdgeAdder) Node(id int64) graph.Node { return W.WNode(id) }
func (W _gonum_org_v1_gonum_graph_testgraph_EdgeAdder) Nodes() graph.Nodes       { return W.WNodes() }
func (W _gonum_org_v1_gonum_graph_testgraph_EdgeAdder) SetEdge(e graph.Edge)     { W.WSetEdge(e) }

// _gonum_org_v1_gonum_graph_testgraph_EdgeRemover is an interface wrapper for EdgeRemover type
type _gonum_org_v1_gonum_graph_testgraph_EdgeRemover struct {
	WEdge           func(uid int64, vid int64) graph.Edge
	WFrom           func(id int64) graph.Nodes
	WHasEdgeBetween func(xid int64, yid int64) bool
	WNode           func(id int64) graph.Node
	WNodes          func() graph.Nodes
	WRemoveEdge     func(fid int64, tid int64)
}

func (W _gonum_org_v1_gonum_graph_testgraph_EdgeRemover) Edge(uid int64, vid int64) graph.Edge {
	return W.WEdge(uid, vid)
}
func (W _gonum_org_v1_gonum_graph_testgraph_EdgeRemover) From(id int64) graph.Nodes {
	return W.WFrom(id)
}
func (W _gonum_org_v1_gonum_graph_testgraph_EdgeRemover) HasEdgeBetween(xid int64, yid int64) bool {
	return W.WHasEdgeBetween(xid, yid)
}
func (W _gonum_org_v1_gonum_graph_testgraph_EdgeRemover) Node(id int64) graph.Node {
	return W.WNode(id)
}
func (W _gonum_org_v1_gonum_graph_testgraph_EdgeRemover) Nodes() graph.Nodes { return W.WNodes() }
func (W _gonum_org_v1_gonum_graph_testgraph_EdgeRemover) RemoveEdge(fid int64, tid int64) {
	W.WRemoveEdge(fid, tid)
}

// _gonum_org_v1_gonum_graph_testgraph_LineAdder is an interface wrapper for LineAdder type
type _gonum_org_v1_gonum_graph_testgraph_LineAdder struct {
	WFrom           func(id int64) graph.Nodes
	WHasEdgeBetween func(xid int64, yid int64) bool
	WLines          func(uid int64, vid int64) graph.Lines
	WNewLine        func(from graph.Node, to graph.Node) graph.Line
	WNode           func(id int64) graph.Node
	WNodes          func() graph.Nodes
	WSetLine        func(l graph.Line)
}

func (W _gonum_org_v1_gonum_graph_testgraph_LineAdder) From(id int64) graph.Nodes {
	return W.WFrom(id)
}
func (W _gonum_org_v1_gonum_graph_testgraph_LineAdder) HasEdgeBetween(xid int64, yid int64) bool {
	return W.WHasEdgeBetween(xid, yid)
}
func (W _gonum_org_v1_gonum_graph_testgraph_LineAdder) Lines(uid int64, vid int64) graph.Lines {
	return W.WLines(uid, vid)
}
func (W _gonum_org_v1_gonum_graph_testgraph_LineAdder) NewLine(from graph.Node, to graph.Node) graph.Line {
	return W.WNewLine(from, to)
}
func (W _gonum_org_v1_gonum_graph_testgraph_LineAdder) Node(id int64) graph.Node { return W.WNode(id) }
func (W _gonum_org_v1_gonum_graph_testgraph_LineAdder) Nodes() graph.Nodes       { return W.WNodes() }
func (W _gonum_org_v1_gonum_graph_testgraph_LineAdder) SetLine(l graph.Line)     { W.WSetLine(l) }

// _gonum_org_v1_gonum_graph_testgraph_LineRemover is an interface wrapper for LineRemover type
type _gonum_org_v1_gonum_graph_testgraph_LineRemover struct {
	WFrom           func(id int64) graph.Nodes
	WHasEdgeBetween func(xid int64, yid int64) bool
	WLines          func(uid int64, vid int64) graph.Lines
	WNode           func(id int64) graph.Node
	WNodes          func() graph.Nodes
	WRemoveLine     func(fid int64, tid int64, id int64)
}

func (W _gonum_org_v1_gonum_graph_testgraph_LineRemover) From(id int64) graph.Nodes {
	return W.WFrom(id)
}
func (W _gonum_org_v1_gonum_graph_testgraph_LineRemover) HasEdgeBetween(xid int64, yid int64) bool {
	return W.WHasEdgeBetween(xid, yid)
}
func (W _gonum_org_v1_gonum_graph_testgraph_LineRemover) Lines(uid int64, vid int64) graph.Lines {
	return W.WLines(uid, vid)
}
func (W _gonum_org_v1_gonum_graph_testgraph_LineRemover) Node(id int64) graph.Node {
	return W.WNode(id)
}
func (W _gonum_org_v1_gonum_graph_testgraph_LineRemover) Nodes() graph.Nodes { return W.WNodes() }
func (W _gonum_org_v1_gonum_graph_testgraph_LineRemover) RemoveLine(fid int64, tid int64, id int64) {
	W.WRemoveLine(fid, tid, id)
}

// _gonum_org_v1_gonum_graph_testgraph_NodeAdder is an interface wrapper for NodeAdder type
type _gonum_org_v1_gonum_graph_testgraph_NodeAdder struct {
	WAddNode        func(a0 graph.Node)
	WEdge           func(uid int64, vid int64) graph.Edge
	WFrom           func(id int64) graph.Nodes
	WHasEdgeBetween func(xid int64, yid int64) bool
	WNewNode        func() graph.Node
	WNode           func(id int64) graph.Node
	WNodes          func() graph.Nodes
}

func (W _gonum_org_v1_gonum_graph_testgraph_NodeAdder) AddNode(a0 graph.Node) { W.WAddNode(a0) }
func (W _gonum_org_v1_gonum_graph_testgraph_NodeAdder) Edge(uid int64, vid int64) graph.Edge {
	return W.WEdge(uid, vid)
}
func (W _gonum_org_v1_gonum_graph_testgraph_NodeAdder) From(id int64) graph.Nodes {
	return W.WFrom(id)
}
func (W _gonum_org_v1_gonum_graph_testgraph_NodeAdder) HasEdgeBetween(xid int64, yid int64) bool {
	return W.WHasEdgeBetween(xid, yid)
}
func (W _gonum_org_v1_gonum_graph_testgraph_NodeAdder) NewNode() graph.Node      { return W.WNewNode() }
func (W _gonum_org_v1_gonum_graph_testgraph_NodeAdder) Node(id int64) graph.Node { return W.WNode(id) }
func (W _gonum_org_v1_gonum_graph_testgraph_NodeAdder) Nodes() graph.Nodes       { return W.WNodes() }

// _gonum_org_v1_gonum_graph_testgraph_NodeRemover is an interface wrapper for NodeRemover type
type _gonum_org_v1_gonum_graph_testgraph_NodeRemover struct {
	WEdge           func(uid int64, vid int64) graph.Edge
	WFrom           func(id int64) graph.Nodes
	WHasEdgeBetween func(xid int64, yid int64) bool
	WNode           func(id int64) graph.Node
	WNodes          func() graph.Nodes
	WRemoveNode     func(id int64)
}

func (W _gonum_org_v1_gonum_graph_testgraph_NodeRemover) Edge(uid int64, vid int64) graph.Edge {
	return W.WEdge(uid, vid)
}
func (W _gonum_org_v1_gonum_graph_testgraph_NodeRemover) From(id int64) graph.Nodes {
	return W.WFrom(id)
}
func (W _gonum_org_v1_gonum_graph_testgraph_NodeRemover) HasEdgeBetween(xid int64, yid int64) bool {
	return W.WHasEdgeBetween(xid, yid)
}
func (W _gonum_org_v1_gonum_graph_testgraph_NodeRemover) Node(id int64) graph.Node {
	return W.WNode(id)
}
func (W _gonum_org_v1_gonum_graph_testgraph_NodeRemover) Nodes() graph.Nodes  { return W.WNodes() }
func (W _gonum_org_v1_gonum_graph_testgraph_NodeRemover) RemoveNode(id int64) { W.WRemoveNode(id) }

// _gonum_org_v1_gonum_graph_testgraph_WeightedEdgeAdder is an interface wrapper for WeightedEdgeAdder type
type _gonum_org_v1_gonum_graph_testgraph_WeightedEdgeAdder struct {
	WEdge            func(uid int64, vid int64) graph.Edge
	WFrom            func(id int64) graph.Nodes
	WHasEdgeBetween  func(xid int64, yid int64) bool
	WNewWeightedEdge func(from graph.Node, to graph.Node, weight float64) graph.WeightedEdge
	WNode            func(id int64) graph.Node
	WNodes           func() graph.Nodes
	WSetWeightedEdge func(e graph.WeightedEdge)
}

func (W _gonum_org_v1_gonum_graph_testgraph_WeightedEdgeAdder) Edge(uid int64, vid int64) graph.Edge {
	return W.WEdge(uid, vid)
}
func (W _gonum_org_v1_gonum_graph_testgraph_WeightedEdgeAdder) From(id int64) graph.Nodes {
	return W.WFrom(id)
}
func (W _gonum_org_v1_gonum_graph_testgraph_WeightedEdgeAdder) HasEdgeBetween(xid int64, yid int64) bool {
	return W.WHasEdgeBetween(xid, yid)
}
func (W _gonum_org_v1_gonum_graph_testgraph_WeightedEdgeAdder) NewWeightedEdge(from graph.Node, to graph.Node, weight float64) graph.WeightedEdge {
	return W.WNewWeightedEdge(from, to, weight)
}
func (W _gonum_org_v1_gonum_graph_testgraph_WeightedEdgeAdder) Node(id int64) graph.Node {
	return W.WNode(id)
}
func (W _gonum_org_v1_gonum_graph_testgraph_WeightedEdgeAdder) Nodes() graph.Nodes {
	return W.WNodes()
}
func (W _gonum_org_v1_gonum_graph_testgraph_WeightedEdgeAdder) SetWeightedEdge(e graph.WeightedEdge) {
	W.WSetWeightedEdge(e)
}

// _gonum_org_v1_gonum_graph_testgraph_WeightedLine is an interface wrapper for WeightedLine type
type _gonum_org_v1_gonum_graph_testgraph_WeightedLine struct {
	WFrom   func() graph.Node
	WID     func() int64
	WTo     func() graph.Node
	WWeight func() float64
}

func (W _gonum_org_v1_gonum_graph_testgraph_WeightedLine) From() graph.Node { return W.WFrom() }
func (W _gonum_org_v1_gonum_graph_testgraph_WeightedLine) ID() int64        { return W.WID() }
func (W _gonum_org_v1_gonum_graph_testgraph_WeightedLine) To() graph.Node   { return W.WTo() }
func (W _gonum_org_v1_gonum_graph_testgraph_WeightedLine) Weight() float64  { return W.WWeight() }

// _gonum_org_v1_gonum_graph_testgraph_WeightedLineAdder is an interface wrapper for WeightedLineAdder type
type _gonum_org_v1_gonum_graph_testgraph_WeightedLineAdder struct {
	WFrom            func(id int64) graph.Nodes
	WHasEdgeBetween  func(xid int64, yid int64) bool
	WLines           func(uid int64, vid int64) graph.Lines
	WNewWeightedLine func(from graph.Node, to graph.Node, weight float64) graph.WeightedLine
	WNode            func(id int64) graph.Node
	WNodes           func() graph.Nodes
	WSetWeightedLine func(l graph.WeightedLine)
}

func (W _gonum_org_v1_gonum_graph_testgraph_WeightedLineAdder) From(id int64) graph.Nodes {
	return W.WFrom(id)
}
func (W _gonum_org_v1_gonum_graph_testgraph_WeightedLineAdder) HasEdgeBetween(xid int64, yid int64) bool {
	return W.WHasEdgeBetween(xid, yid)
}
func (W _gonum_org_v1_gonum_graph_testgraph_WeightedLineAdder) Lines(uid int64, vid int64) graph.Lines {
	return W.WLines(uid, vid)
}
func (W _gonum_org_v1_gonum_graph_testgraph_WeightedLineAdder) NewWeightedLine(from graph.Node, to graph.Node, weight float64) graph.WeightedLine {
	return W.WNewWeightedLine(from, to, weight)
}
func (W _gonum_org_v1_gonum_graph_testgraph_WeightedLineAdder) Node(id int64) graph.Node {
	return W.WNode(id)
}
func (W _gonum_org_v1_gonum_graph_testgraph_WeightedLineAdder) Nodes() graph.Nodes {
	return W.WNodes()
}
func (W _gonum_org_v1_gonum_graph_testgraph_WeightedLineAdder) SetWeightedLine(l graph.WeightedLine) {
	W.WSetWeightedLine(l)
}
