package arls
/*
Copyright (c) 2020 The Go Authors. All rights reserved.

Redistribution and use in source and binary forms, with or without
modification, are permitted provided that the following conditions are
met:

   * Redistributions of source code must retain the above copyright
notice, this list of conditions and the following disclaimer.
   * Redistributions in binary form must reproduce the above
copyright notice, this list of conditions and the following disclaimer
in the documentation and/or other materials provided with the
distribution.
   * Neither the name of Google Inc. nor the names of its
contributors may be used to endorse or promote products derived from
this software without specific prior written permission.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
"AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
(INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
*/

/*
The routines in this package solve the linear system of equation, 
    Ax = b
for any shape matrix. (Some routines add constraint equations.)
The system can be underdetermined, square, or over-determined.
That is, A(m,n) can be such that m < n, m = n, or m > n.
The right-hand-side column is b(n).

Our solvers automatically detect whether the system is 
ill-conditioned or not, and to what degree,
and produce an auto-regularized solution.
Any degree of singularity is handled.
There is no need to provide any ancillary information such as error estimates,
iteration limits, etc. And the only error mode is failure of the
SVD algorithm to converge, which is rare.

On the other hand, no software of this nature can be perfect, and occasionally
these routines may fail to please. But we are confident that these are the current
state-of-the-art for automatically solving difficult dense systems of linear equations.

Please search for "func Arls(" in this file for instructions for calling our primary
auto-regularizing routine.

Please search for "func Arlsvd(" in this file for instructions for calling the same
algorithm by providing the svd (Singular Value Decomposition) of A instead of A itself
in order to much more quickly solve systems of equations which use the same matrix but
different right-hand-side vectors.

Please search for "func Arlsnn(" in this file for instructions for calling our 
auto-regularizing routine that adds a non-negativity constraint.
*/

import "gonum.org/v1/gonum/mat"
import "math"

// SIMPLE MATH UTILITIES

func assumedErr() float64 { return 1.0E-9 }

func mini(a, b int) int { if a < b { return a } 
    return b }

func maxi(a, b int) int { if a > b { return a }  
    return b }

func minf(a, b float64) float64 { if a < b { return a } 
    return b }

func maxf(a, b float64) float64 { if a > b { return a } 
    return b }

func myabs(a float64) float64 { if a >= 0.0 { return a } 
    return -a }

func myrms(x *mat.VecDense) float64 { 
    return mat.Norm(x,2) / math.Sqrt(float64(x.Len())) }
    
func isMatZero(A *mat.Dense) bool {
    m, n := A.Dims()
    if m==0 || n == 0 { return true }
    for i:=0; i < m; i++ {
        for j:=0; j < n; j++ {
            if (A.At(i,j) == 0.0) { continue }
            return false
        }
    }
    return true
}     

func isVecZero(b *mat.VecDense) bool {
    m, _:= b.Dims()
    if m==0 { return true }
    for i:=0; i < m; i++ {
        if (b.AtVec(i) == 0.0) { continue }
        return false
    }
    return true
}  
   
// Return the (algebraic) minimum value in a vector
func vecMin(x *mat.VecDense) float64 {
    m, n := x.Dims()
    if n != 1 { panic("vecMin requires exactly one column.")	} 
    x1 := x.At(0,0)
    xmin := x1
    for i:=1; i<m; i++ {
        x1 = x.At(i,0)
        if x1 < xmin { xmin = x1 }
    }
    return xmin
} 
// SHAPE CHANGING UTILITIES

// Transpose a matrix
func trans(A *mat.Dense) *mat.Dense { 
    m, n := A.Dims()
    AT := mat.NewDense(n, m, nil)
    for i:=0; i < m; i++ {
        for j:=0; j < n; j++ {
            AT.Set(j,i,A.At(i,j))
        }
    }
    return AT
}

// Delete a column from a matrix
func deleteColumn(A *mat.Dense, jkill int) *mat.Dense {
    m, n := A.Dims()
    B := mat.NewDense(m, n-1, nil)
    jj := 0
    for j:=0; j < n; j++ {
        if j == jkill { continue }
        for i:=0; i< m; i++ { B.Set(i, jj, A.At(i,j)) }
        jj++
    }
    return B
} 

// Delete a row from a matrix
func deleteRow(A *mat.Dense, ikill int) *mat.Dense {
    m, n := A.Dims()
    if m<=1 { return mat.NewDense(1, n, nil) }
    B := mat.NewDense(m-1, n, nil)
    ii := 0
    for i:=0; i<m; i++ {
        if i == ikill { continue }
        for j:=0; j<n; j++ { B.Set(ii, j, A.At(i,j)) }  
        ii++
    }
    return B
} 

// Delete element from vector
func deleteElement(b *mat.VecDense, ikill int) *mat.VecDense {
    m, _ := b.Dims()
    if m<=1 { return mat.NewVecDense(1, nil) }
    bb := mat.NewVecDense(m-1, nil)
    ii := 0
    for i:=0; i<m; i++ {
        if i == ikill { continue }
        bb.SetVec(ii, b.AtVec(i))
        ii++
    }
    return bb
}

// Trim off last few rows of a matrix
func trimRowSize(A *mat.Dense, mnew int) *mat.Dense {
    _, n := A.Dims()
    B := mat.NewDense(mnew, n, nil)
    for i:=0; i<mnew; i++ {
        for j:=0; j<n; j++ { B.Set(i, j, A.At(i,j)) }  
    }
    return B
} 

// Trim off last few elements of a vector
// mnew must not be larger than length of x
func trimSize(x *mat.VecDense, mnew int) *mat.VecDense {
    y := mat.NewVecDense(mnew, nil)
    for i:=0; i<mnew; i++ { y.SetVec(i, x.AtVec(i)) }
    return y
} 
 

// Add a new row to A
func appendRow(A *mat.Dense, row mat.Vector) *mat.Dense {
    m, n := A.Dims()
    B := mat.NewDense(m+1, n, nil)
    for i:=0; i<m; i++ {
        for j:=0; j<n; j++ { B.Set(i, j, A.At(i,j)) }  
    }
        for j:=0; j<n; j++ { B.Set(m, j, row.AtVec(j)) } 
    return B 
}

// Add a new element to b
func appendElement(b *mat.VecDense, val float64) *mat.VecDense {
    m, _ := b.Dims()
    bb := mat.NewVecDense(m+1, nil)
    for i:=0; i<m; i++ { bb.SetVec(i, b.AtVec(i)) }
    bb.SetVec(m, val)
    return bb
}

// "USABLE RANK" DETERMINATION

// Determines a first-stage usable rank based on large rise in Picard Vector
func splita(g *mat.VecDense, mg int) int {
    if mg < 2 { return mg }
    // initialize
    w := decideWidth(mg)
    sensitivity := g.At(0,0)
    small := sensitivity
    local := sensitivity
    urank := 1
    //look for sensitivity explosion
    for i := 1; i<mg; i++ {
        sensitivity = g.At(i,0)
        if i >= w && sensitivity > 25.0 * small && sensitivity > local {
            break 
        }
        if sensitivity < small {
            small = small + 0.40 * (sensitivity - small)
        } else { 
            small = small + 0.10 * (sensitivity - small)
        }  
        local = local + 0.40 * (sensitivity - local)
        urank = i + 1 
    }  
    return urank
}

// A utility for splitb, below
func decideWidth(mg int) int {
    if (mg < 3)           { return 1  
    } else if (mg <= 8)   { return 2  // 2 to 4 spans of 2
    } else if (mg <= 12)  { return 3  // 3 to 4 spans of 3   
    } else if (mg <= 20)  { return 4  // 3 to 5 spans of 4
    } else if (mg <= 28)  { return 5  // 4 to 5 spans of 5
    } else if (mg <= 36)  { return 6  // 4 to 6 spans of 6
    } else if (mg <= 50)  { return 7  // 5 to 7 spans of 7
    } else if (mg <= 64)  { return 8  // 6 to 8 spans of 8
    } else if (mg <= 80)  { return 9  // 7 to 8 spans of 9
    } else if (mg <= 200) { return 10 // 8 to 20 spans of 10
    } else if (mg <= 300) { return 12 //16 to 24 spans of 12
    } else if (mg <= 400) { return 14 //21 to 28 spans of 14
    } else if (mg <=1000) { return 16 //25 to 60 spans of 16
    } else                { return 20 //50 to ?? spans of 20
    }
}

// A utility for splitb, below
func computeMovSums(g *mat.VecDense, mg, w int) *mat.VecDense {
    numsums := mg - w + 1
    sums := mat.NewVecDense(numsums , nil)
    s := 0.0
    for i := 0; i < numsums; i++ {
        s = 0.0
        for j := i; j < i + w; j++ {
            s = s + g.AtVec(j)
        }
        sums.SetVec(i, s)
    }
    return sums
}

// A utility for splitb, below
func decideMultiple(width int) float64 {
    if (width < 3)           { return 30.0
    } else if (width <= 10)  { return 20.0
    } else if (width <= 20)  { return 15.0   
    } else                   { return  7.0
    }
}

// Determine a usable rank based on small rise 
// in Picard Vector after low point
func splitb(g *mat.VecDense, mg int ) int {
    w := decideWidth(mg)
    if w < 2 { return mg } // splitb needs w>=2 to be reliable
    
    // magnify any divergence by squaring
    gg := mat.NewVecDense(mg, nil)
    for i:= 0; i < mg; i++ {  
        a := g.AtVec(i)
        gg.SetVec(i, a * a)
    }   
    
    // suppress dropouts
    var gmin float64
    for i:= 1; i < mg - 1; i++ {
        gmin = minf(gg.AtVec(i-1) , gg.AtVec(i+1)) 
        if gg.AtVec(i) < 0.2 * gmin {  
            gg.SetVec(i, 0.5 * gmin)
        }
    }
    
    // choose breakpoint as multiple of lowest moving average
    sums := computeMovSums(gg, mg, w)
    ilow := 0
    glow := sums.AtVec(0)
    ms := sums.Len()
    for i:= 1; i < ms; i++ {  
        a := sums.AtVec(i)
        if a < glow {
            glow = a
            ilow = i
        }
    }
    multi := decideMultiple(w)
    bad := multi * sums.AtVec(ilow)

    // look for unexpected rise
    ibad := 0
    for i:= ilow + 1; i < mg - w + 1; i++ {
        if sums.AtVec(i) > bad {
            ibad = i
            break
        }
    }   

    // decide 
    urank := mg    // full rank
    if ibad > 0 { urank = ibad + w - 1 }
    return urank
}

// ARLS() SOLVERS

// Computes a regularized solution to Ax=b, 
// given the usable rank and the Tikhonov lambda value.
func rmslambda(b *mat.VecDense, U *mat.Dense, S *mat.DiagDense, V *mat.Dense, 
               ur int, lamb float64) (*mat.VecDense, float64) {
    m, n1 := U.Dims()  
    n, n2 := V.Dims()
    mn := mini(m, n)
    si := 0.0
    ps := mat.NewDiagDense(mn ,nil)   // initally zero
    for i:=0; i < ur; i++ {
        si = S.At(i,i)
        if si > 0.0 { ps.SetDiag(i, 1.0 / (si + lamb * lamb / si) )  }
    }
    
    xx := mat.NewDense(n,1,nil)
    xx.Product(V, ps, U.T(), b)
    var x mat.Vector
    x = xx.ColView(0)
    
    Ax := mat.NewDense(m,1,nil)
    Ax.Product(U, S, V.T(), x) 
    var bb mat.Vector
    bb = Ax.ColView(0)
    
    res := mat.NewVecDense(m ,nil)
    res.SubVec(b, bb)
    rn := myrms(res)
    
    mn = n1 + n2 // avoid compiler complaint
    return x.(*mat.VecDense), rn
}

// Computes Tikhonov's lambda using b's estimated RMS error
func discrep(b *mat.VecDense, 
             U *mat.Dense, S *mat.DiagDense, V *mat.Dense, 
             ur int, mysigma float64) (float64) {
    lo := 0.0  // for minimum achievable residual
    hi := 0.33 * S.At(0,0)  // for ridiculously large residual
    lamb := 0.0
    var n, mx int
    n, mx = V.Dims()
    x := mat.NewVecDense(n,nil)
    check := 0.0
    // bisect until we get the residual we want...but quit eventually
    for k:= 0; k < 50; k++ {
        lamb = (lo + hi) * 0.5
        x, check = rmslambda(b, U, S, V, ur, lamb)
        if math.Abs(check - mysigma) < 1.0E-9 * mysigma {
            break  // close enough!
        }
        if check > mysigma {
            hi = lamb;
        } else {
            lo = lamb;
        }
    }
    mx = mx + int(x.AtVec(0)) // to avoid compiler complaint
    return lamb
}

/*----------------------------------------------------------
    Arlsvd solves the linear system of equation, Ax = b, for any shape matrix
    even if the system is very poorly conditioned. 
    The singular value decomposition of A must be provided to this routine.
    
    The purpose of this version is to allow the user to solve multiple 
    problems using the same matrix, A, without having to compute the singular
    value decomposition more than once. Since the cost of the SVD is the 
    main cost of calling Arls(), reusing the svd can greatly decrease the
    execution time.
    
    Other than that difference the call to Arlsvd (and the results) are
    identical to those for Arls, described below. Indeed, Arls is just
    a simple interface to this routine. See the (small) body of Arls() below
    for how to compute the SVD yourself and call this routine directly.
*/
func Arlsvd(svd mat.SVD, b *mat.VecDense) (
             x *mat.VecDense, nr, ur int, sigma, lambda float64) {
    //extract SVD components 
    var m, n, mb, nb, mn, mx int
    var U, V mat.Dense
    svd.UTo(&U)
    m, nb = U.Dims()
    mb, nb = b.Dims()
    if mb != m || nb != 1 { panic("Dimensions do not match.") }
    svd.VTo(&V)
    n, nb = V.Dims()
    mn = mini(m, n)  
    mx = maxi(m, n)
    sv := make([]float64, mn)
    svd.Values(sv)
    S := mat.NewDiagDense(mn, sv)
    
    //compute sensitivity vector
    Ut := trans(&U)
    Utb := mat.NewVecDense(mn, nil)
    Utb.MulVec(Ut,b)
    eps := S.At(0,0) * float64(mx) * 1.0E-14
    if eps == 0.0 { eps = 1.0E-14 }
    si := 0.0
    sense := 0.0
    nr = 0
    g := mat.NewVecDense(mn, nil)
    for i:=0; i < mn; i++ {
        si = S.At(i,i)
        if si < eps { break }    
        sense = Utb.AtVec(i) / si
        if sense < 0.0 { sense = -sense }
        g.SetVec(i,sense)
        nr = i + 1
    }   
    if nr < 1 { x := mat.NewVecDense(n, nil); return x, 0, 0, 0.0, 0.0 }
    
    //get usable rank
    ura := splita(g, nr)
    urb := splitb(g, ura)
    ur = mini(ura, urb)
    
    //solve
    sigma = 0.0
    lambda = 0.0
    check := 0.0
    if ur >= nr {
        x, check = rmslambda(b, &U, S, &V, ur, 0.0)
    } else {
        noise := Utb
        noise.SliceVec(ur, mn)
        sigma = myrms(noise)   //sigma := myrms(Utb[ur:mn])
        lambda = discrep(b, &U, S, &V, ur, sigma)
        x, check = rmslambda(b, &U, S, &V, ur, lambda)
    }
    eps = check // to avoid compiler complaint
    return
}

/*----------------------------------------------------------  
    Arls() solves the linear system of equation, Ax = b, for any shape matrix.
    The system can be underdetermined, square, or over-determined.
    That is, A(m,n) can be such that m < n, m = n, or m > n.
    Argument b(n) is the right-hand-side column.    
    This solver automatically detects whether the system is 
    ill-conditioned or not, and to what degree.
    
    Then...
     -- If the equations are consistent then the solution will usually be
        exact within round-off error.
     -- If the equations are inconsistent then the the solution will be
        by least-squares. That is, it solves ``min ||b - Ax||_2``.
     -- If the equations are inconsistent and diagnosable as ill-conditioned
        using the "Discrete Picard Condition" (see references) then the system
        will be automatically regularized. The residual of the regularized
        system will usually be larger than the residual of the least-squares 
        solution.
     -- If either A or b is all zeros then the solution will be all zeros.

    Parameters
    ----------
    A : *mat.Dense
        Coefficient matrix
    b : *mat.VecDense
        Columns of dependent variables.

    Returns
    -------
    x : *mat.VecDense
        The solution.
    nr : int
        The traditional numerical rank of A, for information only.
    ur : int
        The "Usable Rank".
        Note that "numerical rank" is an attribute of a matrix
        but the "usable rank" that Arls computes is an attribute
        of the whole problem, Ax=b.
    sigma : float64
        The estimated right-hand-side root-mean-square error.
    lambda : float64
        The estimated Tikhonov regularization parameter.

    Example
    --------
    Arls() will behave like any good least-squares solver when the system
    is well conditioned.
    Here is a tiny example of an ill-conditioned system as handled by Arls(),

       x + y = 2
       x + 1.01 y =3
       
    Then we have these arrays:
           A          b
      [ 1., 1.  ]    [2.]
      [ 1., 1.01]    [3.]

    Then standard solvers will return:
       x = [-98. , 100.]'      
    (Where "'" means to transpose this row to a column)   

    But Arls() will see the violation of the Picard Condition and return
       x = [1.122168 , 1.127793]'

    Notes:
    -----
    1. When the system is ill-conditioned, the process works best when the rows
       of A are scaled so that the elements of b have similar estimated errors.
       
    2. As with any linear equation solver, please check whether the solution 
       is reasonable. In particular, you should check the residual vector, A*x - b.
       
    3. Arls() neither needs nor accepts optional parameters such as iteration
       limits, error estimates, variable bounds, condition number limits, etc.
       It also does not return any error flags as there are no error states.
       As long as the SVD completes (and SVD failure is remarkably rare)
       then Arls() and Arlsvd() (below) will complete normally.
       In the case of Arlsnn() (below) the svd factorization is called 
       potentially up to n-1 times, so each of these factorizations will 
       need to complete successfully for Arlsnn() to complete successfully.
       But each successive call to the svd factorization will be for a further 
       reduced A matrix, so if the first factoriztion works the remainder of
       the smaller factorizations will probably work well also.
       
    4. Arls()'s main application is to find a reasonable solution even in the 
       midst of excessive inaccuracy, ill-conditioning, singularities,
       duplicated data, etc.
       The only commonly available routine that can handle difficult problems
       like Arls() is the method called LSMR() which is available for some
       languages. However, in our tests Arls() is much more consistent and 
       stable in its behavior than LSMR(). 
       On the other hand, when LSMR works, it often works very well, 
       and can produce more perfect answers than Arls(). 
       In contrast, Arls() tends to produce a slightly overly-smooth solution.
       (This trait is instrinsic to Arls()'s current algorithm.)
       
    5. In view of note 4, Arls() is not appropriate for situations
       where the requirements are more for high accuracy rather than
       robustness. So we assume, when appropriate in the coding, that none
       of the input data needs to be considered more accurate than 
       about 8 or 9 significant figures.
       
    6. There is one situation where Arls() is not likely to give a good result:
       when the system A*x = b is ill-conditioned and a majority of the 
       elements of the matrix A are zero.
       Elimination of less unimportant variables from the problem, 
       if possible, might help produce a useful solution.
       Multiplication of both sides by the transpose of A might help
       produce a useful answer if the system is over-determined.
       Or, you might consider using a Sparse solver.

    Resources
    ----------
    For a short presentation on the Picard Conditionm which is at the heart
    of this package's algorithms, please see http://www.rejtrix.net/ .
    For a complete description, see "The Discrete Picard Condition for Discrete
    Ill-posed Problems", Per Christian Hansen, 1990.
    See link.springer.com/article/10.1007/BF01933214
    
    Rondall E. Jones, Ph.D., University of New Mexico, 1985
    http://www.rejtrix.net/ 
    rejones7@msn.com
*/  
func Arls(A *mat.Dense, b *mat.VecDense) (
          x *mat.VecDense, nr, ur int, sigma, lambda float64) {
    _, n := A.Dims()
    if (isMatZero(A) || isVecZero(b)) { 
        return mat.NewVecDense(n, nil), 0, 0, 0., 0. }     
    var svd mat.SVD
    ok := svd.Factorize(A, mat.SVDThin)
    if !ok { panic("SVD failed to factorize A") }
    x, nr, ur, sigma, lambda = Arlsvd(svd, b)
    return
}


/*----------------------------------------------------------
    Arlsnn() solves Ax = b in the least squares sense, with the solution
    constrained to be non-negative. 
    The call to Arlsnn() and the parameters returned are 
    exactly the same as for Arls(). Please see above for details.
    This version actually deletes variables from the problem rather 
    than zeroing their column.

    Example
    -------
    Suppose we have:
           A          b
      [2., 2., 1.]  [3.9]
      [2., 1., 0.]  [3.0]
      [1., 1., 0.]  [2.0]
      
    Then any least-squares solver will produce
        x =  [1. ,1., -0.1]'
    (Where "'" meansto transpose this row to a column.)
    
    But Arlsnn() produces
        x =  [1.0400, 0.9200, 0.0000]'

    Arlsnn() tries to produce a small residual for the final solution,
    while being based toward making the fewest changes feasible
    to the problem. (That is, fewest columns of A set to zero.)
    Older solvers like the classic NNLS() focus only on minimizing the 
    residual, resulting in extra interference with the user's model.
    Arlsnn seeks a better balance.
*/
func Arlsnn(A *mat.Dense, b *mat.VecDense) (
            x *mat.VecDense, nr, ur int, sigma, lambda float64) {
    m, n := A.Dims()
    mb, nb := b.Dims()
    if mb != m || nb != 1 { panic("Dimensions are not right.") }
    
    // get initial solution and Tikhonov parameter
    var svd mat.SVD
    ok := svd.Factorize(A, mat.SVDThin)
    if !ok { panic("SVD failed to factorize A") }
    var xt *mat.VecDense
    xt, nr, ur, sigma, lambda = Arlsvd(svd, b)
    
    // see if unconstrained solution is already non-negative
    if vecMin(xt) >= 0.0 { return xt, nr, ur, sigma, lambda }
        
    C := matCopy(A)   
    // cols is a list of active column numbers
    cols := make([]int, n)  //var cols [n] int
    for i:=0; i<n; i++ { cols[i] = i }
    mc := m
    nc := n
    mrc := mini(mc, nc)
    t := 0.0
    
    // the approach here is to actually delete columns,
    // for SVD speed and stability, not just zero columns.
    for {
        // choose a column to zero
        p := -1
        worst := 0.0
        for j:=0; j<nc; j++ {
            t = xt.AtVec(j)
            if t < worst { p = j; worst = t }
        }
        if p < 0 { break }
   
        // remove column p and re-Factorize
        for i:=p; i<nc-1; i++ { cols[i] = cols[i+1] }
        C = deleteColumn(C, p)  
        ok := svd.Factorize(C, mat.SVDThin)
        if !ok { panic("SVD failed to factorize A") }
        
        mc, nc = C.Dims()
        mrc = mini(mc, nc)
        U := mat.NewDense(mc, mrc, nil)
        svd.UTo(U)
        V := mat.NewDense(nc, mrc, nil)
        svd.VTo(V)
        sv := make([]float64, mrc)
        svd.Values(sv)
        S := mat.NewDiagDense(mrc, sv)
        
        // compute new pseudoinverse
        si := 0.0
        pi := 0.0
        ps := mat.NewDiagDense(mrc, nil)
        for i:=0; i<mrc; i++ {
            si = S.At(i,i)
            if si > 0.0 {
                pi = 1.0 / (si + lambda*lambda / si)
            } else {
                pi = 0.0
            }    
            ps.SetDiag(i, pi)
        }
        // compute new solution
        xx := mat.NewDense(nc,1,nil)
        xx.Product(V, ps, U.T(), b)
        xt = (xx.ColView(0)).(*mat.VecDense)
        if nc < 2 { break }
    }
    
    // degenerate case: nc==1
    if xt.AtVec(0) < 0.0 { xt.SetVec(0, 0.0) }
    
    // rebuild full solution vector
    x = mat.NewVecDense(n, nil)
    for j:=0; j<nc; j++ { x.SetVec(cols[j], xt.AtVec(j)) }
    return
}

// UTILITIES FOR CONSTRAINED SOLVERS

func vecCopy(b *mat.VecDense) *mat.VecDense {
    m, _ := b.Dims()
    c := mat.NewVecDense(m, nil)
    c.CopyVec(b)
    return c
} 

func matCopy(A *mat.Dense) *mat.Dense {
    m, n := A.Dims()
    B := mat.NewDense(m, n, nil)
    B.Copy(A)
    return B
} 

// Exchange two rows of A, in place
func exchangeRowsOf(A *mat.Dense, i1,i2 int) {
    if i1==i2 { return }
    _, n := A.Dims()
    t := 0.0
    for j:=0; j<n; j++ {
        t = A.At(i1,j)
        A.Set(i1, j, A.At(i2,j))
        A.Set(i2, j, t)
    }
} 

// Multiply a row of A by r, in place
func scaleRow(A *mat.Dense, i int, r float64) {
    _, n := A.Dims()  
    for j:=0; j<n; j++ { A.Set(i, j, A.At(i,j) * r) }
} 

// Dot product two rows of A, returning a float64
func dotRows(A *mat.Dense, i1,i2 int) float64 {
    _, n := A.Dims()  
    t := 0.0
    for j:=0; j<n; j++ { t += A.At(i1,j) * A.At(i2,j) }
    return t
} 

// Dot product row i of A with row j of E, returning a float64
// Caller must assure the matrices have identical column sizes
func dotRowsAB(A *mat.Dense, i int, E *mat.Dense, j int) float64 {
    _,n := A.Dims()  
    t := 0.0
    for k:=0; k<n; k++ { t += A.At(i,k) * E.At(j,k) }
    return t
} 

// Find the (first) row of Ex=f which has the highest ratio of f[i]
// to the norm of the row. 
func findMaxSense(A *mat.Dense, b *mat.VecDense) int {
    snmax := -1.0
    imax := 0  // default
    m,_:= A.Dims()
    s := 0.0
    for i:=0; i<m; i++ {
        rn := mat.Norm(A.RowView(i), 2)
        if rn > 0.0 {
            s = math.Abs(b.AtVec(i)) / rn
            if s > snmax {
                snmax = s
                imax = i
            }
        }
    }    
    return imax
}

// Find row of A with largest 2-norm
func findMaxRowNorm(A *mat.Dense, istart int) int {
    m, _:= A.Dims()
    rnmax := -1.0
    imax := istart
    rn := 0.0
    for i:=istart; i<m; i++ {
        rn = mat.Norm(A.RowView(i), 2)
        if rn > rnmax {
            rnmax = rn
            imax = i
        }
    }
    return imax
}

// Orthogonalize and order the rows of Ex=f for Arlseq
func prepeq(E *mat.Dense, f *mat.VecDense) (
              *mat.Dense, *mat.VecDense) {
    if (isMatZero(E)) { return E, f }
    
    EE := matCopy(E)
    ff := vecCopy(f)
    m, n := EE.Dims()
    t, rin, scale, d := 0.0, 0.0, 0.0, 0.0
    imax := 0
    for i:=0; i<m; i++ {
        
        // determine new best row and put it next
        if i == 0 { imax = findMaxSense(EE, ff) 
        } else    { imax = findMaxRowNorm(EE, i) }
        exchangeRowsOf(EE, i, imax) 
        t = ff.AtVec(i)
        ff.SetVec(i, ff.AtVec(imax))
        ff.SetVec(imax, t)

        // normalize
        rin = mat.Norm(EE.RowView(i), 2)
        if rin > 0.0 {
            scale = 1.0/rin
            scaleRow(EE, i, scale)
            ff.SetVec(i, scale * ff.AtVec(i))
        } else {
            ff.SetVec(i, 0.0)
        }

        // subtract projections onto EE[i,:]
        for k:=i+1; k<m; k++ {
            d = dotRows(EE, k, i)
            for j:=0; j<n; j++ {
                EE.Set(k, j, EE.At(k,j) - d*EE.At(i,j)) 
            }
            ff.SetVec(k, ff.AtVec(k) - d*ff.AtVec(i) )
        }
    }

    // reject ill-conditioned rows
    if m > 2 {
        g := mat.NewVecDense(m, nil)
        for k:=0; k<m; k++ { g.SetVec(k, math.Abs(ff.AtVec(k))) }
        m1 := splita(g, m)
        mm := splitb(g, m1)
        if mm < m {
            EE = trimRowSize(EE,mm)
            ff = trimSize(ff,mm)
        }    
    }
    return EE, ff
}

// Subtract from Ax=b its projection onto Ex=f.
// E should normally have been processed with prepeq() before calling arlspj.
// Caller must guarantee that A and E have identical 2nd dimensions.
func arlspj(A *mat.Dense, b *mat.VecDense, 
            E *mat.Dense, f *mat.VecDense, neglect float64) (
              *mat.Dense, *mat.VecDense) {
    AA := matCopy(A)
    bb := vecCopy(b)
    ma, na := AA.Dims()
    me, _ := E.Dims()
    i := 0
    d := 0.0
    for i<ma {
        for j:=0; j<me; j++ {
            d = dotRowsAB(AA, i, E, j)
            for k:=0; k<na; k++ {
                AA.Set(i, k, AA.At(i,k) - d * E.At(j,k)) 
            }    
            bb.SetVec(i, bb.AtVec(i) - d*f.AtVec(j) )
        }
        nm := mat.Norm(AA.RowView(i), 2)
        if nm < neglect {
            AA = deleteRow(AA, i)  //??
            bb = deleteElement(bb, i)
            ma, na = AA.Dims()
        } else {
            scaleRow(AA, i, 1.0/nm) 
            bb.SetVec(i, bb.AtVec(i)/nm )
            i += 1
        }
        if ma<2 { if isMatZero(AA) { break} }
    }
    return AA, bb 
}

// CONSTRAINED SOLVERS

/*----------------------------------------------------------    
    Arlseq() solves the double linear system of equations
       Ax = b  (least squares)
       Ex == f  (exact)

    Both Ax=b and Ex=f system can be underdetermined, square,
    or over-determined. Arguments b and f must be single columns.

    Ex=f is treated as a set of equality constraints.
    These constraints are usually few in number and well behaved.
    But clearly the caller can easily provide equations in Ex=f that
    are impossible to satisfy as a group. For example, there could be
    one equation requiring x[0]=0, and another requiring x[0]=1.
    And, the solver must deal with there being redundant or other
    pathological situations within the E matrix.
    So the solution process will either solve each equation in Ex=f exactly
    (within roundoff) or if that is impossible, arlseq() will discard
    one or more equations until the remaining equations are solvable
    exactly (within roundoff).
    We will refer below to the solution of this reduced system as "xe".

    After Ex=f is processed as above, the rows of Ax=b will have their
    projections onto every row of Ex=f subtracted from them.
    We will call this reduced set of equations A'x = b'.
    (Thus, the rows of A' will all be orthogonal to the rows of E.)
    This reduced problem A'x = b', will then be solved with arls().
    We will refer to the solution of this system as "xt".

    The final solution will be x = xe + xt.

    Parameters
    ----------
    A : (m, n)  array_like "Coefficient" matrix, type float.
    b : (m)     array_like column of dependent variables, type float.
    E : (me, n) array_like "Coefficient" matrix, type float.
    f : (me)    array_like column of dependent variables, type float.

    Returns
    -------
    x : (n) array_like column, type float.
    nr : int
        The numerical rank of the matrix, A, after its projection onto the rows
        of E are subtracted.
    ur : int
        The "usable" rank of the "reduced" problem, Ax=b, after its projection
        onto the rows of Ex=f are subtracted.
        Note that "numerical rank" is an attribute of a matrix
        but the "usable rank" that arls computes is an attribute
        of the problem, Ax=b.
    sigma : float
        The estimated right-hand-side root-mean-square error.
    lambda : float
        The estimated Tikhonov regularization.

    Examples
    --------
    Here is a tiny example of a problem which has an "unknown" amount
    of error in the right hand side, but for which the user knows that the
    correct SUM of the unknowns must be 3:

         x + 2 y = 5.3   (Least Squares)
       2 x + 3 y = 7.8
           x + y = 3     ( Exact )

    Then we have these arrays:
            A          b
        [1., 2.]    [5.3]
        [2., 3.]    [7.8]

           E          f
        [1., 1.]    [3.0]

    Without using the equality constraint we are given here,
    standard solvers will return [x,y] = [-.3 , 2.8]'.
    (Where "'" means to transpose this row to a column.)
    Even arls() will return the same [x,y] = [-.3 , 2.8]'.
    The residual for this solution is [0.0 , 0.0]' (within roundoff).
    But of course x + y = 2.5, not the 3.0 we really want.

    Arlsnn() could help here by disallowing presumably unacceptable
    negative values, producing [x,y] = [0. , 2.6]'.
    The residual for this solution is [-0.1 , 0.]' which is of course
    an increase from zero, but this is natural since we have forced
    the solution away from being the "exact" result, for good reason.
    Note that x + y = 2.6, which is a little better.

    If we solve with arlseq(A,b,E,f) then we get [x,y] = [1.004, 1.996]'.
    This answer is close to the "correct" answer of [x,y] = [1.0 , 2.0]'
    if the right hand side had been the correct [5.,8.]' instead of [5.3,7.8]'.
    The residual for this solution is [-0.3 , 0.2]' which is yet larger.
    Again, when adding constraints to the problem the residual
    typically increases, but the solution becomes more acceptable.
    Note that x + y = 3 exactly.

    Notes:
    -----
    See arls() above for notes and references.
*/
func Arlseq(A *mat.Dense, b *mat.VecDense, 
            E *mat.Dense, f *mat.VecDense) (
            x *mat.VecDense, nr, ur int, sigma, lambda float64) {
    ma, na := A.Dims()
    mb, nb := b.Dims()
    me, ne := E.Dims()
    mf, nf := f.Dims()
    if ( ma < 1 || ma != mb || nb != 1  || me != mf || nf != 1 || na != ne) { 
        panic("Dimensions do not match.") }
        
    if (isMatZero(E)) { return Arls(A, b) }

    imax := findMaxRowNorm(A,0)
    rnmax := mat.Norm(A.RowView(imax), 2)
    neglect := rnmax * assumedErr()  // see Note 5. for arls()
  
    EE, ff := prepeq(E, f)
    AA, bb := arlspj(A, b, EE, ff, neglect)
    EEt := trans(EE)
    xe := mat.NewVecDense(ne, nil)
    xe.MulVec(EEt,ff)      
    xt, nr, ur, sigma, lambdah := Arls(AA, bb)
    xt.AddVec(xt,xe)
    return xt, nr, ur, sigma, lambdah
}

//  Find the most violated inequality constraint
func get_worst(GG *mat.Dense, hh, x *mat.VecDense) int {
    p := -1
    m, _ := GG.Dims()
    if (m<1) { return p}
    
    rhs := mat.NewVecDense(m, nil)
    rhs.MulVec(GG,x)
    worst := 0.0
    diff := 0.0
    for i:=0; i<m; i++ {
        if (rhs.AtVec(i) < hh.AtVec(i)) {
            diff = hh.AtVec(i) - rhs.AtVec(i)
            if p < 0 || diff > worst {
                p = i
                worst = diff
            }
        }
    }
    return p
}    

/*----------------------------------------------------------
    Arlsall() solves the triple linear system of equations
       Ax =  b  (least squares)
       Ex == f  (exact)
       Gx >= h  ("greater than" inequality constraints)  
        
    Each of the three systems an be underdetermined, square, or
    over-determined. However, generally E should have very few rows
    compared to A. Arguments b, f, and h must be single columns.
    
    Arlsall() uses Arlseq(), above, as the core solver, and iteratively selects
    rows of Gx>=h to addto Ex==f, choosing first whatever remaining equation
    in Gx>=h most violates its requirement.

    Note that "less than" equations can be included by negating
    both sides of the equation, thus turning it into a "greater than".

    If either A or b is all zeros then the solution will be all zeros.

    Parameters
    ----------
    A : (m, n)  array_like "Coefficient" matrix, type float.
    b : (m)     array_like column of dependent variables, type float.
    E : (me, n) array_like "Coefficient" matrix, type float.
    f : (me)    array_like column of dependent variables, type float.
    G : (mg, n) array_like "Coefficient" matrix, type float.
    b : (mg)    array_like column of dependent variables, type float.

    Returns
    -------
    x : (n) array_like column, type float.
    nr : int
        The numerical rank of the matrix, A.
    ur : int
        The usable rank of the problem, Ax=b.
        Note that "numerical rank" is an attribute of a matrix
        but the "usable rank" that arls computes is an attribute
        of the problem, Ax=b.
    sigma : float
        The estimated right-hand-side root-mean-square error.
    lambda : float
        The estimated Tikhonov regularization.

    Example
    -------
   Let A and b be:
         A           b   
     [1.,1.,1.]    [5.9]
     [0.,1.,1.]    [5.0]
     [1.,0.,1.]    [3.9]

    Then any least-squares solver would produce x = [0.9, 2., 3.]'.
    (Where "'" means to transpose this row to a column.)
    The residual for this solution is zero within roundoff.

    But if we happen to know that all the answers should be at least 1.0
    then we can add inequalites to insure that:
        x[0] >= 1
        x[1] >= 1
        x[2] >= 1
        
    This can be expressed in the matrix equation Gx>=h where we have these arrays"
           G       h
        [1,0,0]   [1]
        [0,1,0]   [1]
        [0,0,1]   [1]

    Let's let E and f be zeros at this point.
    Then arlsall(A,b,E,f,G,h) produces x = [1., 2.013, 2.872]'.
    The residual vector and its norm are then:
       res = [-0.015, -0.115, 0.028]'
       norm(res) = 0.119
    So the price of adding this constraint is that the residual is no
    longer zero. This is normal behavior.

    Let's say that we have discovered that x[2] should be exactly 3.0. 
    We can add that constraint using the Ex==f system:
          E         f
      [0.,0.,1.]   [3.]    
      
    Calling arlsall(A,b,E,f,G,h) produces x = [1., 1.9, 3.0]'.
    The residual vector and its norm are then:
       res = [0.0, -0.1, 0.1]'
       norm(res) = 0.141
    So again, as we add constraints to force the solution to what we know
    it must be, the residual will usually increase steadily from what the
    least-squares equations left alone will produce.
    But it would be a mistake to accept an answer that did not meet
    the facts that we know.
*/
func Arlsall(A *mat.Dense, b *mat.VecDense, 
             E *mat.Dense, f *mat.VecDense, 
             G *mat.Dense, h *mat.VecDense)  (
             x *mat.VecDense, nr, ur int, sigma, lambda float64) {
    ma, na := A.Dims()
    mb, nb := b.Dims()
    me, ne := E.Dims()
    mf, nf := f.Dims()
    mg, ng := G.Dims()
    mh, nh := h.Dims()
    if ( ma != mb || nb != 1 ||
         me != mf || nf != 1 ||
         mg != mh || nh != 1 ||
         ne != na || ng != na ) { 
         panic("Dimensions do not match.") }    
                        
    EE := matCopy(E)
    ff := vecCopy(f)
    GG := matCopy(G)
    hh := vecCopy(h)
    
    // get initial solution... it might actually be ok
    x, nr, ur, sigma, lambdah := Arlseq(A, b, EE, ff)
    if (isMatZero(G)) { return x, nr, ur, sigma, lambdah }
    
    // while inequality constraints are not fully satisfied:
    for (true) {
        p := get_worst(GG, hh, x)
        if (p < 0) { break }
        // move row p of GGx=hh to new last row of Ex>=f
        if (me == 0) {
            EE = mat.NewDense(1,na,nil)
            for j:=0; j<ne; j++ { EE.Set(0,j,GG.At(p,j)) }
            ff = mat.NewVecDense(1,nil)
            ff.SetVec(0, hh.AtVec(p))
        } else {             
            EE = appendRow(EE,GG.RowView(p)) 
            ff = appendElement(ff,hh.AtVec(p))
        }
        me, _ = EE.Dims()
        GG = deleteRow(GG,p)
        hh = deleteElement(hh,p) 
        
        // re-solve modified system
        x, nr, ur, sigma, lambdah = Arlseq(A, b, EE, ff)
    }
    return x, nr, ur, sigma, lambdah
}

/*----------------------------------------------------------   
    Arlsgt() is the same as Arlsall except that equality constraints
    are not used.
*/
func Arlsgt(A *mat.Dense, b *mat.VecDense, G *mat.Dense, h *mat.VecDense)  (
            x *mat.VecDense, nr, ur int, sigma, lambda float64) {
    _, n := A.Dims()        
    E := mat.NewDense(1, n, nil)
    f := mat.NewVecDense(1, nil)
    return Arlsall(A, b, E, f, G, h)
}

