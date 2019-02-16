// Copyright ©2019 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package barneshut

import (
	"fmt"
	"math"
)

const (
	ne = iota
	se
	sw
	nw
)

// Point2 is a 2D point.
type Point2 struct {
	X, Y float64
}

// Add returns the vector sum of p and q.
func (p Point2) Add(q Point2) Point2 {
	p.X += q.X
	p.Y += q.Y
	return p
}

// Sub returns the vector sum of p and -q.
func (p Point2) Sub(q Point2) Point2 {
	p.X -= q.X
	p.Y -= q.Y
	return p
}

// Scale returns the vector p scaled by f.
func (p Point2) Scale(f float64) Point2 {
	p.X *= f
	p.Y *= f
	return p
}

// Box2 is a 2D bounding box.
type Box2 struct {
	Min, Max Point2
}

// quadrant returns which quadrant of b that p should be placed in.
func (b Box2) quadrant(p Particle2) int {
	center := Point2{
		X: (b.Min.X + b.Max.X) / 2,
		Y: (b.Min.Y + b.Max.Y) / 2,
	}
	c := p.Coord2()
	if checkBounds && (c.X < b.Min.X || b.Max.X < c.X || c.Y < b.Min.Y || b.Max.Y < c.Y) {
		panic(fmt.Sprintf("p out of range %+v: %#v", b, p))
	}
	if c.X < center.X {
		if c.Y < center.Y {
			return nw
		} else {
			return sw
		}
	} else {
		if c.Y < center.Y {
			return ne
		} else {
			return se
		}
	}
}

// split returns a quadrant subdivision of b in the given direction.
func (b Box2) split(dir int) Box2 {
	halfX := (b.Max.X - b.Min.X) / 2
	halfY := (b.Max.Y - b.Min.Y) / 2
	switch dir {
	case ne:
		b.Min.X += halfX
		b.Max.Y -= halfY
	case se:
		b.Min.X += halfX
		b.Min.Y += halfY
	case sw:
		b.Max.X -= halfX
		b.Min.Y += halfY
	case nw:
		b.Max.X -= halfX
		b.Max.Y -= halfY
	}
	return b
}

// Particle2 is a particle in a plane.
type Particle2 interface {
	Coord2() Point2
	Mass() float64
}

// Gravity2 returns a vector force on m1 by m2, equalt to (m1⋅m2)/‖v‖²
// in the directions of v.
func Gravity2(m1, m2 float64, v Point2) Point2 {
	d2 := v.X*v.X + v.Y*v.Y
	if d2 == 0 {
		return Point2{}
	}
	return v.Scale((m1 * m2) / (d2 * math.Sqrt(d2)))
}

// Plane implements Barnes-Hut force approximation calculations.
type Plane struct {
	root tile

	Particles []Particle2
}

// NewPlane returns a new Plane.
func NewPlane(p []Particle2) *Plane {
	q := Plane{Particles: p}
	q.Reset()
	return &q
}

// Reset reconstructs the Barnes-Hut tree. Reset must be called if the
// Particles field or elements of Particles have been altered, unless
// ForceOn is called with theta=0 or no data structures have been
// previously built.
func (q *Plane) Reset() {
	// Note that Plane does not perform the normal Barnes-Hut
	// tree construction; Plane allows internal nodes, with the
	// exception of root, to be filled with particles and does
	// not relocate particles that have been put in the tree.

	q.root = tile{}
	if len(q.Particles) == 0 {
		return
	}

	q.root.bounds.Min = q.Particles[0].Coord2()
	q.root.bounds.Max = q.root.bounds.Min
	for _, e := range q.Particles[1:] {
		c := e.Coord2()
		if c.X < q.root.bounds.Min.X {
			q.root.bounds.Min.X = c.X
		}
		if c.X > q.root.bounds.Max.X {
			q.root.bounds.Max.X = c.X
		}
		if c.Y < q.root.bounds.Min.Y {
			q.root.bounds.Min.Y = c.Y
		}
		if c.Y > q.root.bounds.Max.Y {
			q.root.bounds.Max.Y = c.Y
		}
	}

	// TODO(kortschak): Partially parallelise this by
	// choosing the direction and using one of four
	// goroutines to work on each root quadrant.
	for _, e := range q.Particles {
		q.root.insert(e)
	}
	q.root.summarize()
}

// ForceOn returns a force vector on p given p's mass and the force calculation
// function, using the Barnes-Hut theta approximation parameter.
//
// When calculating the force component on a particle, m1 is the mass of the
// particle, m2 is the mass of the component mass center and v is the
// vector from the particle to the component mass center. The returned vector
// from force is the force vector component acting on p by the component
// mass center.
//
// It is safe to call ForceOn concurrently.
func (q *Plane) ForceOn(p Particle2, theta float64, force func(m1, m2 float64, v Point2) Point2) (vector Point2) {
	var empty tile
	if theta > 0 && q.root != empty {
		return q.root.forceOnMassAt(p.Coord2(), p.Mass(), theta, force)
	}

	// For the degenerate case, just iterate over the
	// slice of particles rather than walking the tree.
	var v Point2
	m := p.Mass()
	pv := p.Coord2()
	for _, e := range q.Particles {
		v = v.Add(force(m, e.Mass(), e.Coord2().Sub(pv)))
	}
	return v
}

// tile is a quad tree quadrant with Barnes-Hut extensions.
type tile struct {
	particle Particle2

	bounds Box2

	nodes [4]*tile

	center Point2
	mass   float64
}

// insert inserts p into the subtree rooted at t.
func (t *tile) insert(p Particle2) {
	dir := t.bounds.quadrant(p)
	if t.nodes[dir] != nil {
		t.nodes[dir].insert(p)
		return
	}
	t.nodes[dir] = &tile{
		particle: p,
		bounds:   t.bounds.split(dir),
		center:   p.Coord2(),
		mass:     p.Mass(),
	}
}

// summarize updates node masses and centers of mass.
func (t *tile) summarize() (center Point2, mass float64) {
	for _, d := range &t.nodes {
		if d == nil {
			continue
		}
		c, m := d.summarize()
		t.center.X += c.X * m
		t.center.Y += c.Y * m
		t.mass += m
	}
	t.center.X /= t.mass
	t.center.Y /= t.mass
	return t.center, t.mass
}

// forceOnMassAt returns a force vector on p given p's mass m and the force
// calculation function, using the Barnes-Hut theta approximation parameter.
func (t *tile) forceOnMassAt(p Point2, m, theta float64, force func(m1, m2 float64, v Point2) Point2) (vector Point2) {
	s := ((t.bounds.Max.X - t.bounds.Min.X) + (t.bounds.Max.Y - t.bounds.Min.Y)) / 2
	d := math.Hypot(p.X-t.center.X, p.Y-t.center.Y)
	if s/d < theta {
		return force(m, t.mass, t.center.Sub(p))
	}

	var v Point2
	if t.particle != nil {
		v = force(m, t.particle.Mass(), t.particle.Coord2().Sub(p))
	}
	for _, d := range &t.nodes {
		if d == nil {
			continue
		}
		v = v.Add(d.forceOnMassAt(p, m, theta, force))
	}
	return v
}
