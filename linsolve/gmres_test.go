// Copyright ©2017 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package linsolve

import (
	"math"
	"math/rand"
	"testing"

	"github.com/gonum/floats"
)

func TestGMRES(t *testing.T) {
	rnd := rand.New(rand.NewSource(1))
	for _, tc := range []testCase{
		randomSPD(1, rnd),
		randomSPD(2, rnd),
		randomSPD(3, rnd),
		randomSPD(4, rnd),
		randomSPD(5, rnd),
		randomSPD(10, rnd),
		randomSPD(20, rnd),
		randomSPD(50, rnd),
		randomSPD(100, rnd),
		randomSPD(200, rnd),
		randomSPD(500, rnd),
		market("nos1", 1e-10),
		market("nos4", 1e-12),
		market("nos5", 1e-12),
		market("bcsstm20", 1e-11),
		market("bcsstm22", 1e-12),

		market("gre_216a", 1e-12),
		market("gre__115", 1e-12),
		market("gre__185", 1e-9),
		market("gre__343", 1e-12),
		market("hor__131", 1e-12),
		market("impcol_a", 1e-8),
		market("impcol_b", 1e-11),
		market("impcol_c", 1e-12),
		market("impcol_d", 1e-12),
		market("impcol_e", 1e-11),
		market("steam1", 1e-8),
		market("steam3", 1e-10),
		market("west0067", 1e-12),
		market("west0132", 1e-6),
		market("west0167", 1e-8),
		market("west0381", 1e-11),
		market("west0479", 1e-6),
		market("west0497", 1e-7),
	} {
		n := tc.n
		// Compute the right-hand side b so that the vector [1,1,...,1]
		// is the solution.
		want := make([]float64, n)
		for i := range want {
			want[i] = 1
		}
		sys := System{
			MatVec: tc.a,
			B:      make([]float64, n),
		}
		// B = A*want
		sys.MatVec(sys.B, want, false)

		// TODO(vladimir-ch): Add tests with non-default Restart.
		r, err := Iterative(sys, nil, &GMRES{}, Settings{
			MaxIterations: tc.iters,
			Tolerance:     1e-15,
		})
		if err != nil {
			t.Errorf("Case %v (n=%v): unexpected error %v", tc.name, n, err)
			continue
		}
		dist := floats.Distance(r.X, want, math.Inf(1))
		if dist > tc.tol {
			t.Errorf("Case %v (n=%v): unexpected solution, |want-got|=%v", tc.name, n, dist)
		}
	}
}
