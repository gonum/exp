// Copyright ©2017 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package linsolve provides iterative algorithms for solving linear systems.
package linsolve

// TODO(vladimir-ch): Improve documentation. Write an introduction about
// iterative methods and that they can be more efficient than direct methods
// when we are solving large (sparse) systems, when the solution does not have
// to be known to machine precision. Write that the matrix is accessed only via
// matrix-vector products. Write that the documentation is written from the
// perspective of users who want to call Iterative and designers who want to
// implement Method (or direct users of Method?).
