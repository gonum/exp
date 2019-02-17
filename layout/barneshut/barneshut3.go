// Copyright ©2019 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package barneshut

import (
	"fmt"
	"math"
)

const (
	lne = iota
	lse
	lsw
	lnw
	une
	use
	usw
	unw
)

// Point3 is a 3D point.
type Point3 struct {
	X, Y, Z float64
}

// Add returns the vector sum of p and q.
func (p Point3) Add(q Point3) Point3 {
	p.X += q.X
	p.Y += q.Y
	p.Z += q.Z
	return p
}

// Sub returns the vector sum of p and -q.
func (p Point3) Sub(q Point3) Point3 {
	p.X -= q.X
	p.Y -= q.Y
	p.Z -= q.Z
	return p
}

// Scale returns the vector p scaled by f.
func (p Point3) Scale(f float64) Point3 {
	p.X *= f
	p.Y *= f
	p.Z *= f
	return p
}

// Box3 is a 3D bounding box.
type Box3 struct {
	Min, Max Point3
}

// octant returns which octant of b that p should be placed in.
func (b Box3) octant(p Particle3) int {
	center := Point3{
		X: (b.Min.X + b.Max.X) / 2,
		Y: (b.Min.Y + b.Max.Y) / 2,
		Z: (b.Min.Z + b.Max.Z) / 2,
	}
	c := p.Coord3()
	if checkBounds && (c.X < b.Min.X || b.Max.X < c.X || c.Y < b.Min.Y || b.Max.Y < c.Y || c.Z < b.Min.Z || b.Max.Z < c.Z) {
		panic(fmt.Sprintf("p out of range %+v: %#v", b, p))
	}
	if c.X < center.X {
		if c.Y < center.Y {
			if c.Z < center.Z {
				return lnw
			} else {
				return unw
			}
		} else {
			if c.Z < center.Z {
				return lsw
			} else {
				return usw
			}
		}
	} else {
		if c.Y < center.Y {
			if c.Z < center.Z {
				return lne
			} else {
				return une
			}
		} else {
			if c.Z < center.Z {
				return lse
			} else {
				return use
			}
		}
	}
}

// split returns a octant subdivision of b in the given direction.
func (b Box3) split(dir int) Box3 {
	halfX := (b.Max.X - b.Min.X) / 2
	halfY := (b.Max.Y - b.Min.Y) / 2
	halfZ := (b.Max.Z - b.Min.Z) / 2
	switch dir {
	case lne:
		b.Min.X += halfX
		b.Max.Y -= halfY
		b.Max.Z -= halfZ
	case lse:
		b.Min.X += halfX
		b.Min.Y += halfY
		b.Max.Z -= halfZ
	case lsw:
		b.Max.X -= halfX
		b.Min.Y += halfY
		b.Max.Z -= halfZ
	case lnw:
		b.Max.X -= halfX
		b.Max.Y -= halfY
		b.Max.Z -= halfZ
	case une:
		b.Min.X += halfX
		b.Max.Y -= halfY
		b.Min.Z += halfZ
	case use:
		b.Min.X += halfX
		b.Min.Y += halfY
		b.Min.Z += halfZ
	case usw:
		b.Max.X -= halfX
		b.Min.Y += halfY
		b.Min.Z += halfZ
	case unw:
		b.Max.X -= halfX
		b.Max.Y -= halfY
		b.Min.Z += halfZ
	}
	return b
}

// Particle3 is a particle in a plane.
type Particle3 interface {
	Coord3() Point3
	Mass() float64
}

// Gravity3 returns a vector force on m1 by m2, equalt to (m1⋅m2)/‖v‖²
// in the directions of v.
func Gravity3(m1, m2 float64, v Point3) Point3 {
	d2 := v.X*v.X + v.Y*v.Y + v.Z*v.Z
	if d2 == 0 {
		return Point3{}
	}
	return v.Scale((m1 * m2) / (d2 * math.Sqrt(d2)))
}

// Volume implements Barnes-Hut force approximation calculations.
type Volume struct {
	root bucket

	Particles []Particle3
}

// NewVolume returns a new Volume.
func NewVolume(p []Particle3) *Volume {
	q := Volume{Particles: p}
	q.Reset()
	return &q
}

// Reset reconstructs the Barnes-Hut tree. Reset must be called if the
// Particles field or elements of Particles have been altered, unless
// ForceOn is called with theta=0 or no data structures have been
// previously built.
func (q *Volume) Reset() {
	if len(q.Particles) == 0 {
		q.root = bucket{}
		return
	}

	q.root = bucket{
		particle: q.Particles[0],
		center:   q.Particles[0].Coord3(),
		mass:     q.Particles[0].Mass(),
	}
	q.root.bounds.Min = q.root.center
	q.root.bounds.Max = q.root.center
	for _, e := range q.Particles[1:] {
		c := e.Coord3()
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
		if c.Z < q.root.bounds.Min.Z {
			q.root.bounds.Min.Z = c.Z
		}
		if c.Z > q.root.bounds.Max.Z {
			q.root.bounds.Max.Z = c.Z
		}
	}

	// TODO(kortschak): Partially parallelise this by
	// choosing the direction and using one of eight
	// goroutines to work on each root octant.
	for _, e := range q.Particles[1:] {
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
func (q *Volume) ForceOn(p Particle3, theta float64, force func(m1, m2 float64, v Point3) Point3) (vector Point3) {
	var empty bucket
	if theta > 0 && q.root != empty {
		return q.root.forceOnMassAt(p.Coord3(), p.Mass(), theta, force)
	}

	// For the degenerate case, just iterate over the
	// slice of particles rather than walking the tree.
	var v Point3
	m := p.Mass()
	pv := p.Coord3()
	for _, e := range q.Particles {
		v = v.Add(force(m, e.Mass(), e.Coord3().Sub(pv)))
	}
	return v
}

// bucket is a oct tree octant with Barnes-Hut extensions.
type bucket struct {
	particle Particle3

	bounds Box3

	nodes [8]*bucket

	center Point3
	mass   float64
}

// insert inserts p into the subtree rooted at b.
func (b *bucket) insert(p Particle3) {
	if b.particle == nil {
		for _, q := range b.nodes {
			if q != nil {
				dir := b.bounds.octant(p)
				if b.nodes[dir] == nil {
					b.nodes[dir] = &bucket{bounds: b.bounds.split(dir)}
				}
				b.nodes[dir].insert(p)
				return
			}
		}
		b.particle = p
		b.center = p.Coord3()
		b.mass = p.Mass()
		return
	}
	dir := b.bounds.octant(p)
	if b.nodes[dir] == nil {
		b.nodes[dir] = &bucket{bounds: b.bounds.split(dir)}
	}
	b.nodes[dir].insert(p)
	dir = b.bounds.octant(b.particle)
	if b.nodes[dir] == nil {
		b.nodes[dir] = &bucket{bounds: b.bounds.split(dir)}
	}
	b.nodes[dir].insert(b.particle)
	b.particle = nil
	b.center = Point3{}
	b.mass = 0
}

// summarize updates node masses and centers of mass.
func (b *bucket) summarize() (center Point3, mass float64) {
	for _, d := range &b.nodes {
		if d == nil {
			continue
		}
		c, m := d.summarize()
		b.center.X += c.X * m
		b.center.Y += c.Y * m
		b.center.Z += c.Z * m
		b.mass += m
	}
	b.center.X /= b.mass
	b.center.Y /= b.mass
	b.center.Z /= b.mass
	return b.center, b.mass
}

// forceOnMassAt returns a force vector on p given p's mass m and the force
// calculation function, using the Barnes-Hut theta approximation parameter.
func (b *bucket) forceOnMassAt(p Point3, m, theta float64, force func(m1, m2 float64, v Point3) Point3) (vector Point3) {
	s := ((b.bounds.Max.X - b.bounds.Min.X) + (b.bounds.Max.Y - b.bounds.Min.Y) + (b.bounds.Max.Z - b.bounds.Min.Z)) / 3
	d := math.Hypot(math.Hypot(p.X-b.center.X, p.Y-b.center.Y), p.Z-b.center.Z)
	if s/d < theta || b.particle != nil {
		return force(m, b.mass, b.center.Sub(p))
	}

	var v Point3
	for _, d := range &b.nodes {
		if d == nil {
			continue
		}
		v = v.Add(d.forceOnMassAt(p, m, theta, force))
	}
	return v
}
