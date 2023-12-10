// Copyright ©2021 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package posit

import (
	"fmt"
	"math"
)

type Posit []byte

func (p Posit) regime() (k int) {
	// Regime lead bit.
	const rbit = 0b0100_0000
	// Indicates if leading Regime bit is a one.
	oneRun := rbit&p[0] != 0
outer:
	for i := range p {
		for j := 0; j < 8; j++ {
			// Ignore sign bit.
			if i == 0 && j == 0 {
				j++
			}
			// If current bit is part of the run, add to r.
			if (oneRun && p[i]&(1<<(7-j)) != 0) || (!oneRun && p[i]&(1<<(7-j)) == 0) {
				k++
			} else {
				break outer
			}
		}
	}
	if !oneRun {
		k *= -1
	} else {
		k--
	}
	return k
}

// sign returns true if posit negative
func (p Posit) sign() bool { return p[0]&(1<<7) != 0 }

func (p Posit) bits() int { return 8 * len(p) }

func (p Posit) String() string { return fmt.Sprintf("%08b", []byte(p)) }

// es represents max amount of exponent bits that can be present in the posit.
//
// Values taken from http://www.johngustafson.net/pdfs/BeatingFloatingPoint.pdf
// (table 3) to match or exceed IEEE float dynamic range.
func (p Posit) es() (es int) {
	switch p.bits() {
	case 16:
		es = 1
	case 32:
		es = 3
	case 64:
		es = 4
	case 128:
		es = 7
	case 256:
		es = 10
	default:
		es = p.bits() / 16 // 8 bit posit has es == 0.
	}
	return es
}

// exp returns the exponent part of the posit (2**e).
func (p Posit) exp() (exp int) {
	// bits in front of exp.
	flen := p.regimeLen() + 2 // sign and opposite bits included
	es := p.es()
	// Check if exp bits present for quick return.
	if flen >= p.bits() || es == 0 {
		return 0
	}

	expcount := 0
outer:
	for i := flen / 8; i < len(p); i++ {
		for j := flen % 8; j < 8; j++ {
			if expcount == es {
				break outer
			}
			exp <<= 1
			exp |= (int(p[i]) & (1 << (7 - j))) >> (7 - j)
			expcount++
		}
	}
	return exp
}

// returns regime length for a given posit in number of bits.
func (p Posit) regimeLen() int {
	r := p.regime()
	if r < 0 {
		return -r
	}
	return r + 1
}

// useed defines the midway point of accuracy from 1 to +inf and
// conversely from 1 to 0. It depends on es.
func (p Posit) useed() int { return 1 << (1 << (p.es())) }

// fraction returns the numerator of the fraction part.
func (p Posit) fraction() (frac int) {
	// bits in front of fraction.
	flen := p.regimeLen() + p.es() + 2 // sign and opposite bits included
	// Check if exp bits present for quick return.
	if flen >= p.bits() {
		return 0
	}

	for i := flen / 8; i < len(p); i++ {
		for j := flen % 8; j < 8; j++ {
			frac <<= 1
			frac |= (int(p[i]) & (1 << (7 - j))) >> (7 - j)
		}
	}
	return frac
}

func (p Posit) ToFloat64() float64 {
	reg := float64(p.regime())
	useed := float64(p.useed())
	exp := 1 << p.exp()
	return math.Pow(useed, reg) * float64(exp) * (1 + float64(p.fraction())/useed)
}

// Format implements fmt.Formatter.
func (p Posit) Format(fs fmt.State, c rune) {
	switch c {
	case 'v', 'f':
		fmt.Fprintf(fs, "%T{%f}", p, p.ToFloat64())
	default:
		fmt.Fprintf(fs, "%%!%c(%T=%[2]v)", c, p)
		return
	}
}
