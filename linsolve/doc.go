// Copyright ©2017 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Package linsolve provides iterative methods for solving linear systems.

Background

A system of linear equations can be written as

 A * x = b,

where A is a given n×n non-singular matrix, b is a given n-vector (the
right-hand side), and x is an unknown n-vector.

Direct methods such as the LU or QR decomposition can find (in the absence of
roundoff errors) the exact solution after a finite number of steps. For a
general matrix A they require O(n^3) arithmetic operations which becomes
infeasible for large problems due to excessive memory and time cost.

Iterative methods, in contrast, generally do not compute the exact solution x.
Starting from an initial estimate x_0, they instead compute a sequence x_i of
increasingly accurate approximations to x. When the error of the approximation
x_i becomes smaller that a prescribed threshold,  x_{i+1} is not computed and
the iterative process is stopped.

This package provides iterative methods that access the matrix A only only via
the matrix-vector multiplication. The multiplication requires O(n^2) arithmetic
operations for a general n×n matrix but can be much cheaper if the matrix is
sparse (has only a few non-zeros per row) or has some other special properties
or structure. This is often the case for matrices that arise in practical
applications. Furthermore, if an iterative method can find a good approximation
with a number of matrix-vector multiplications that is much lower than n or
independent of n, it can outperform a direct method for the same problem.

This means that iterative methods are most often useful in the following
situations:

 - the system matrix A is sparse, blocked or has other special structure,
 - the problem size is sufficiently large that a dense factorization of A is too
   costly in terms of computer time and/or storage,
 - computing the product of A (or A^T, if necessary) with a vector can be done
   more efficiently than if A were first formed explicitly as a matrix,
 - it suffices to find only an approximation to the solution.

Using linsolve

The two most important elements of the API for solving linear systems are the
MulVecToer interface and the Iterative function.

The MulVecToer interface represents the system matrix A in this package. It
separates the method implementations from the details of any particular matrix
storage format. All matrix types provided by gonum.org/v1/gonum/mat and
github.com/james-bowman/sparse packages implement this interface.

Note that methods in this package have only limited means for checking whether
the provided MulVecToer represents a matrix that satisfies all assumptions made
by the chosen Method, for example whether the matrix is actually symmetric
positive definite.

The Iterative function is the entry point to the functionality provided by this
package. It takes as parameters the matrix A (as a MulVecToer), the right-hand
side vector b, the iterative method and settings that control the iterative
process and provide a way for reusing memory.

Choosing an iterative method

The choice of an iterative method is typically guided by the properties of the
matrix A including symmetry, definiteness, sparsity, conditioning and block
structure. This is done to assure convergence of the algorithm for the given
problem and efficiency of the computation. The methods of choice are well
understood for symmetric problems in which case their convergence rate (at least
in exact arithmetics) is determined by spectral properties (eigenvalues) of A.
The situation is not as clear for problems with non-symmetric matrices and it
can be difficult to pick the best method for these problems. Any general advice
for non-symmetric problem is often complemented with a trial-and-error approach.

Preconditioning

Preconditioning is a family of techniques that attempt to transform the linear
system into one that has the same solution but more favorable spectrum. The
transformation matrix is called a preconditioner. A good preconditioner improves
the convergence rate of the iterative method, sufficiently to overcome the extra
cost of constructing and applying the preconditioner. Without a preconditioner
the iterative method may even fail to converge. In linsolve a preconditioner is
specified by Settings.PreconSolve.

Implementing Method interface

This package allows external implementations of iterative solvers by means of
the Method interface. It uses a reverse-communication style of API to
"outsource" operations such as matrix-vector multiplication, preconditioner
solve or convergence checks to the caller. The caller performs the commanded
operation and passes the result again to Method. The matrix A and the right-hand
side b are not directly available to Methods which encourages their cleaner
implementation. See the documentation for Method, Operation, and Context for
more information.

References

Further details about computational practice and mathematical theory of
iterative methods can be found in the following references:

 - Barrett, Richard et al. (1994). Templates for the Solution of Linear Systems:
   Building Blocks for Iterative Methods (2nd ed.). Philadelphia, PA: SIAM.
   Retrieved from http://www.netlib.org/templates/templates.pdf
 - Saad, Yousef (2003). Iterative methods for sparse linear systems (2nd ed.).
   Philadelphia, PA: SIAM. Retrieved from
   http://www-users.cs.umn.edu/~saad/IterMethBook_2ndEd.pdf
 - Greenbaum, A. (1997). Iterative methods for solving linear systems.
   Philadelphia, PA: SIAM.
*/
package linsolve
