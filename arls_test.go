package arls
import "fmt"
import "math"
import "testing"
import "gonum.org/v1/gonum/mat"

func FailMe(t *testing.T, s string) { fmt.Println(s); t.Error(s); t.Fail() }

//test utilities

func Zeros(m int) *mat.VecDense {
    x := mat.NewVecDense(m, nil)
    for i:=0; i<m; i++ { x.SetVec(i,0.) }
    return x
}

func Ones(m int) *mat.VecDense {
    x := mat.NewVecDense(m, nil)
    for i:=0; i<m; i++ { x.SetVec(i,1.0) }
    return x
}

func Iota(m int) *mat.VecDense {
    x := mat.NewVecDense(m, nil)
    for i:=0; i<m; i++ { x.SetVec(i,float64(i)) }
    return x
}

func Mzeros(m, n int) *mat.Dense {
    A := mat.NewDense(m, n, nil)
    for i:=0; i<m; i++ {
        for j:=0; j<n; j++ {
            A.Set(i,j,0.0) 
        }
    }             
    return A
}

func Mones(m, n int) *mat.Dense {
    A := mat.NewDense(m, n, nil)
    for i:=0; i<m; i++ {
        for j:=0; j<n; j++ {
            A.Set(i,j,1.0) 
        }
    }
    return A
}

func Eye(m, n int) *mat.Dense {
    A := mat.NewDense(m, n, nil)
    for i:=0; i<m; i++ {
        for j:=0; j<n; j++ {
            A.Set(i,j,0.0)
            if (i==j) { A.Set(i,j,1.0) }
        }
    }
    return A
}

func LowerTri(m, n int) *mat.Dense {
    A := mat.NewDense(m, n, nil)
    for i:=0; i<m; i++ {
        for j:=0; j<m; j++ { A.Set(i,j,1.0) }
    }
    return A
}

func Hilbert(m, n  int) *mat.Dense {
    A := mat.NewDense(m, n, nil)
    for i:=0; i<m; i++ {
        for j:=0; j<n; j++ {
            A.Set(i,j,1.0/(1.0 + float64(i+j)))
        }
    }
    return A
}

func Sum(x *mat.VecDense) float64 {
    m, _ := x.Dims()
    sum := 0.0
    for i:=0; i<m; i++ { sum += x.AtVec(i) }
    return sum
}

func Diffxy(x, y *mat.VecDense) float64 {
    m, _ := x.Dims()
    my, _ := y.Dims()
    if m!=my { return 1.0E9  }  
    sum := 0.0
    for i:=0; i<m; i++ { sum += x.AtVec(i) - y.AtVec(i) }
    return sum
}  

func DiffAB(A, B *mat.Dense) float64 {
    m, n := A.Dims()
    mb, nb := B.Dims()
    if m!=mb || n!=nb { return 1.0E9  }
    
    sum := 0.0
    for i:=0; i<m; i++ {
        for j:=0; j<n; j++ { 
            sum += A.At(i,j) - B.At(i,j) }
    }
    return sum
}     

func MyVecRandom(m int, err float64) *mat.VecDense {
    b := mat.NewVecDense(m,  nil)
    for i:=0; i<m; i++ {
        // be sure this stays the same as in python
        b.SetVec(i, err*Myabs( math.Sin(
                    float64(2*m) + 2.0*float64(i))))
        }
    return b
}

func MyMatRandom(m,n,ibias int) *mat.Dense {
    A := mat.NewDense(m, n, nil)
    for i:=0; i<m; i++ {
        for j:=0; j<n; j++ {
            // be sure this stays the same as in python
            A.Set(i,j,Myabs(math.Sin(
                float64(ibias) + float64(2*m+3*n) + 
                2.0*float64(i) + 2.5*float64(j))))
        }
    }
    return A
}

func MyVecPrint(x *mat.VecDense) {
    m, _ := x.Dims()
    fmt.Println("Vector") 
    for i:=0; i<m; i++ { fmt.Println(x.AtVec(i)) }
    fmt.Println(" ") 
}

func MyMatPrint(A *mat.Dense) {
    m, n := A.Dims()
    fmt.Println("Matrix") 
    for i:=0; i<m; i++ { 
        for j:=0; j<n; j++ { 
            fmt.Printf(" %0.5f", A.At(i,j)) }
        fmt.Printf("\n") 
    }    
    fmt.Println(" ") 
}

func NormOfResidual(A *mat.Dense, b *mat.VecDense, x *mat.VecDense) float64 {
    m, _ := A.Dims()
    bb := Zeros(m)
    bb.MulVec(A,x)  
    r := Zeros(m)
    r.SubVec(b,bb)
    res := mat.Norm(r,2)
    
    //MyMatPrint(A)
    //MyVecPrint(b)
    //MyVecPrint(x)
    //fmt.Println("sum of x= ", Sum(x))
    //fmt.Println("Norm of Residual= ", res)
    return res
}
    
func Myabs(a float64) float64 { if a >= 0.0 { return a } 
    return -a 
}

func MyVecRms(x *mat.VecDense) float64 { 
    return mat.Norm(x,2) / math.Sqrt(float64(x.Len()))
}

func MyMatRms(A *mat.Dense) float64 { 
    m, n := A.Dims()
    sum := 0.0
    a := 0.0
    for i:=0; i<m; i++ { 
        for j:=0; j<n; j++ { 
            a = A.At(i,j)
            sum = sum + a*a
        }
    }    
    return math.Sqrt(sum/float64(m*n))
}

func IsNear(x, y float64) bool {
    tol := 1.0E-8 * (Myabs(x)+Myabs(y)) + 1.0E-12
    return (Myabs(x-y) < tol)  
} 

func IsAbout(x, y float64) bool {
    tol := 1.0E-3 * (Myabs(x)+Myabs(y)) + 1.0E-12
    return (Myabs(x-y) < tol)  
} 

// end of test utilities

func TestZeros(t *testing.T) {
    fmt.Println("TestZeros") 
    A := MyMatRandom(4,4,0)
    if isMatZero(A)  { FailMe(t, "TestZeros(1) failed!") }   
    A = Mzeros(4,4)
    if !isMatZero(A) { FailMe(t, "TestZeros(2) failed!") } 
    b := Ones(4)
    if isVecZero(b)  { FailMe(t, "TestZeros(3) failed!") } 
    b = Zeros(4)
    if !isVecZero(b) { FailMe(t, "TestZeros(4) failed!") } 
}

func TestDeletes(t *testing.T) {
    fmt.Println("TestDeletes") 
    A := Mzeros(4,4)
    for i:=0; i<4; i++ { A.Set(i,2,99.0) }
    B := deleteColumn(A,2)
    m, n := B.Dims()
    if n!=3 { FailMe(t, "TestDeletes(1) failed!") }  
    if !isMatZero(B) { FailMe(t, "TestDeletes(2 failed!") }
    
    A = Mzeros(4,4)
    for j:=0; j<4; j++ { A.Set(2,j,99.0) }
    B = deleteRow(A,2)
    m, n = B.Dims()
    if m!=3 { FailMe(t, "TestDeletes(3) failed!") }   
    if !isMatZero(B) { FailMe(t, "TestDeletes(4 failed!") }
    
    b := Zeros(4)
    b.SetVec(2, 99.0)
    b = deleteElement(b,2)
    m, n = b.Dims()
        if m!=3 { FailMe(t, "TestDeletes(5) failed!")  } 
    if !isVecZero(b) { FailMe(t, "TestDeletes(6) failed!") }
}

func TestTrims(t *testing.T) {
    fmt.Println("TestTrims") 
    A := Mzeros(4,4)
    A.Set(2,2,99.0)
    A.Set(3,3,99.0)
    A = trimRowSize(A, 2)
    m, _ := A.Dims()
    if m!=2 { FailMe(t, "TestTrims(1) failed!") } 
    if !isMatZero(A) { FailMe(t, "TestTrims(2) failed!") }
    
    b := Zeros(4)
    b.SetVec(2, 99.0)
    b.SetVec(3, 99.0)
    b = trimSize(b,2)
    m, _ = b.Dims()
    if m!=2 { FailMe(t, "TestTrims(3) failed!") } 
    if !isVecZero(b) { FailMe(t, "TestTrims(4) failed!") }
}

//return number of non-zero elements in A and their sum
func getStats(A *mat.Dense) (int, float64) {
    m, n := A.Dims()
    k := 0
    sum := 0.0
    for i:=0; i < m; i++ {
        for j:=0; j < n; j++ {
            sum += A.At(i,j)
            if (A.At(i,j) != 0.0) { k++ }
        }
    }
    return k, sum
}     

func TestAppends(t *testing.T) {
    fmt.Println("TestAppends") 
    G := Eye(4,4)
    A := Mzeros(4,4)
    A = appendRow(A,G.RowView(2)) 
    m, _ := A.Dims()
    if m!=5 { FailMe(t, "TestAppends(1) failed!") } 
    if A.At(4,2)!=1.0 { FailMe(t, "TestAppends(2) failed!") } 
    k, sum := getStats(A)
    if k!=1 || sum != 1.0 { FailMe(t, "TestAppends(3) failed!") }
    
    b := Zeros(4)
    b = appendElement(b,7.0) 
    m, _ = b.Dims()
    if m!=5 { FailMe(t, "TestAppends(4) failed!") } 
    if b.AtVec(4) != 7.0 { FailMe(t, "TestAppends(5) failed!") } 
    if isVecZero(b) { FailMe(t, "TestAppends(6) failed!") }
}

func TestTrans(t *testing.T) {
    fmt.Println("TestTrans") 
    A := MyMatRandom(4,4,0)
    m, n := A.Dims()
    AT := trans(A)
    for j:=0; j<n; j++ {
        for i:=0; i<m; i++ { 
            if !IsNear(A.At(i,j) , AT.At(j,i)) { 
                FailMe(t, "TestTrans(1) failed!")  }    
        }
    }    
}

func TestWidth(t *testing.T) {  
    fmt.Println("TestWidth")   
    if decideWidth(2)    != 1  { FailMe(t, "TestWidth(2) failed.") }
    if decideWidth(5)    != 2  { FailMe(t, "TestWidth(5) failed.") }
    if decideWidth(11)   != 3  { FailMe(t, "TestWidth(11) failed.") }
    if decideWidth(19)   != 4  { FailMe(t, "TestWidth(19) failed.") }
    if decideWidth(27)   != 5  { FailMe(t, "TestWidth(27) failed.") }
    if decideWidth(30)   != 6  { FailMe(t, "TestWidth(30) failed.") }
    if decideWidth(49)   != 7  { FailMe(t, "TestWidth(49) failed.") }
    if decideWidth(60)   != 8  { FailMe(t, "TestWidth(60) failed.") }
    if decideWidth(79)   != 9  { FailMe(t, "TestWidth(79) failed.") }
    if decideWidth(150)  != 10 { FailMe(t, "TestWidth(150) failed.") }
    if decideWidth(250)  != 12 { FailMe(t, "TestWidth(250) failed.") }
    if decideWidth(350)  != 14 { FailMe(t, "TestWidth(350) failed.") }
    if decideWidth(450)  != 16 { FailMe(t, "TestWidth(450) failed.") }
    if decideWidth(1100) != 20 { FailMe(t, "TestWidth(1100) failed.") }
}              

func TestMultiple(t *testing.T) {
    fmt.Println("TestMultiple")
    if decideMultiple(2)   != 30.0  { FailMe(t, "TestMultiple(2) failed.") }
    if decideMultiple(8)   != 20.0  { FailMe(t, "TestMultiple(8) failed.") }
    if decideMultiple(15)  != 15.0  { FailMe(t, "TestMultiple(15) failed.") }
    if decideMultiple(30)  != 7.0   { FailMe(t, "TestMultiple(30) failed.") }
}

func TestSplita(t *testing.T) {
    fmt.Println("TestSplita") 
    var g = mat.NewVecDense(4, []float64{1.0, 1.0, 0.0, 20.0})
    if splita(g, 1)  != 1 { FailMe(t, "TestSplita(1) failed.") }
    if splita(g, 2)  != 2 { FailMe(t, "TestSplita(2) failed.") }
    if splita(g, 3)  != 3 { FailMe(t, "TestSplita(3) failed.") }
    if splita(g, 4)  != 3 { FailMe(t, "TestSplita(4) failed.") }
    return
}  

func TestMovSums(t *testing.T) {
    fmt.Println("TestMovSums") 
    var g = mat.NewVecDense(4, []float64{1.0, 2.0, 3.0, 4.0} )
    var ans = mat.NewVecDense(3, []float64{3.0, 5.0, 7.0} )
    sums := computeMovSums(g, 4, 2)
    for i:=0; i<3; i++ {
        if sums.AtVec(i) != ans.AtVec(i) { FailMe(t, "TestMovSums(1) failed.") }
    }
    return
}

func TestSplitb(t *testing.T) {
    fmt.Println("TestSplitb") 
    var g = mat.NewVecDense(6, []float64{1.0, 0.1, 0.01, 0.1, 1.0, 10.0} )
    var ans = mat.NewVecDense(6, []float64{1, 2, 3, 4, 4, 4} )
    var r float64
    for i:=0; i<6; i++ {
        r = float64(splitb(g, i+1))
        if r != ans.At(i,0) { FailMe(t, "TestSplitb(1) failed.") }
    }
    return
}
  
func TestColumn(t *testing.T) {
    fmt.Println("TestColumn")
    A := mat.NewDense(3, 1, []float64 {1., 1., 1.}) 
    b := mat.NewVecDense(3, []float64 {2., 2., 2.})
    
    x, nr, ur, sigma, lambda := Arls(A, b)
    if ! IsNear(x.At(0,0) , 2.0) { FailMe(t, "TestColumn(1) failed.") }
    
    x, nr, ur, sigma, lambda = Arlsnn(A, b)
    if ! IsNear(x.At(0,0) , 2.0) { FailMe(t, "TestColumn(2) failed.") }

    ur = nr + ur + int(sigma) + int(lambda) // to avoid compiler complaint 
    return
}

func TestRow(t *testing.T) {
    fmt.Println("TestRow")
    A := mat.NewDense(1, 3, []float64 {1., 1., 1.})
    b := mat.NewVecDense(1, []float64 {3.})
    
    x, nr, ur, sigma, lambda := Arls(A, b)
    if ! IsNear(MyVecRms(x) , 1.0) { FailMe(t, "TestRow(1) failed.")  }
    
    x, nr, ur, sigma, lambda = Arlsnn(A, b)
    if ! IsNear(MyVecRms(x) , 1.0) { FailMe(t, "TestRow(2) failed.")  }
    
    ur = nr + ur + int(sigma) + int(lambda) // to avoid compiler complaint 
    return
}

func TestArls(t *testing.T) {
    fmt.Println("TestArls") 
    n := 3
    m := 3
    A := mat.NewDense(3, 3, nil)
    b := Ones(3)
    bb := Zeros(3)
    xx := Zeros(3)
    resid := Zeros(3)
    
    //TEST WITH ZERO MATRIX 
    x, nr, ur, sigma, lambda := Arls(A,b)
    if mat.Norm(x,2) != 0.0 { FailMe(t, "TestArls(1) failed.") }
    
    //TEST WITH ZERO RIGHT HAND SIDE
    A = Eye(3,3)
    b = Zeros(3)
    var svd mat.SVD
    ok := svd.Factorize(A, mat.SVDThin)
    if !ok { FailMe(t, "TestArls(2) failed.") }
            
    x, nr, ur, sigma, lambda = Arlsvd(svd, b)
    if mat.Norm(x,2) != 0.0 { FailMe(t, "TestArls(3) failed.") }    
    
    x, nr, ur, sigma, lambda = Arls(A, b)
    if mat.Norm(x,2) != 0.0 { FailMe(t, "TestArls(4) failed.") }
    
    //DESIRED SOLUTION FOR NEXT SEVERAL TESTS
    n = 6
    xx = Ones(n)
    for i:=0; i<n; i++ { xx.SetVec(i,float64(n-2-i)) }
    
    //OVERDETERMINED TESTS WITH HILBERT(7,6)
    A = Hilbert(7,6)
    m, n = A.Dims()
    b = Zeros(m)
    bb = Zeros(m)
    resid = Zeros(m)
    b.MulVec(A,xx)
    b.AddVec(b,MyVecRandom(m, 0.002))

    x, nr, ur, sigma, lambda = Arls(A, b)
    bb.MulVec(A,x)
    resid.SubVec(b,bb)
    if !IsNear(MyVecRms(resid), 0.31871896) { FailMe(t, "TestArls(5) failed.") }

    //SQUARE TESTS WITH HILBERT(6,6)
    A = Hilbert(6,6)
    m, n = A.Dims()
    b = Zeros(m)
    bb = Zeros(m)
    resid = Zeros(m)
    b.MulVec(A,xx)
    x, nr, ur, sigma, lambda = Arls(A, b)
    bb.MulVec(A,x)
    resid.SubVec(b,bb)
    if !IsNear(MyVecRms(resid), 0.0) { FailMe(t, "TestArls(6) failed.") }

    //DUPLICATE ROW TEST WITH HILBERT(6,6)
    A = Hilbert(6,6)
    for j:=0; j<6; j++ { A.Set(3,j, A.At(2,j)) }
    m, n = A.Dims()
    b = Zeros(m)
    bb = Zeros(m)
    resid = Zeros(m)
    b.MulVec(A,xx)
    x, nr, ur, sigma, lambda = Arls(A, b)
    bb.MulVec(A,x)
    resid.SubVec(b,bb)
    if !IsNear(MyVecRms(resid), 0.0) { FailMe(t, "TestArls(7) failed.") }

    //DUPLICATE COLUMN TEST WITH HILBERT(6,6)
    A = Hilbert(6,6)
    for i:=0; i<6; i++ { A.Set(i,3, A.At(i,2)) }
    m, n = A.Dims()
    b = Zeros(m)
    bb = Zeros(m)
    resid = Zeros(m)
    b.MulVec(A,xx)
    x, nr, ur, sigma, lambda = Arls(A, b)
    bb.MulVec(A,x)
    resid.SubVec(b,bb)
    if !IsNear(MyVecRms(resid), 0.0) { FailMe(t, "TestArls(8) failed.") }

    //UNDERDETERMINED TESTS WITH HILBERT(5,6)
    A = Hilbert(5,6)
    m, n = A.Dims()
    b = Zeros(m)
    bb = Zeros(m)
    resid = Zeros(m)
    b.MulVec(A,xx)
    x, nr, ur, sigma, lambda = Arls(A, b)
    bb.MulVec(A,x)
    resid.SubVec(b,bb) 
    if !IsNear(MyVecRms(resid), 0.0) { FailMe(t, "TestArls(9) failed.") }  
    nr = 2*nr + int(ur) + int(sigma) + int(lambda) // to avoid dianostic
}

func TestArlsnn(t *testing.T) {
    fmt.Println("TestArlsnn") 
    //as columns are removed arlsnn will deal with square & underdet
    A := Hilbert(7,8)
    m, n := A.Dims()
    x := Ones(n)
    for i:=0; i<n; i++ { x.SetVec(i,float64(n-2-i)) }
    b := Zeros(m)
    b.MulVec(A,x)
    b.AddVec(b,MyVecRandom(m, 0.000001))      
    xx, nr, ur, sigma, lambda := Arlsnn(A, b) 
    res := NormOfResidual(A,b,xx)  
    if res > 1.5 { FailMe(t, "TestArlsnn(1) failed.") }

    //TEST "IMPOSSIBLE" PROBLEM WITH ARLSNN
    A = Eye(3,3)
    b = mat.NewVecDense(3,[]float64{ -1., -1., -1.})
    xx, nr, ur, sigma, lambda = Arlsnn(A, b)
    res = NormOfResidual(A,b,xx)
    if res > 1.8 { FailMe(t, "TestArlsnn(2) failed.") }
    if mat.Norm(xx,2) > 1.0E-8 { FailMe(t, "TestArlsnn(3) failed.") }
        
    //TEST COMPUTED ZERO PROBLEM
    A = mat.NewDense(2, 3, []float64 { 1., 1., 1., 0., 0., 0.})
    b = mat.NewVecDense(2, []float64 { 0., 1. }) 
    xx, nr, ur, sigma, lambda = Arlsnn(A,b)
    res = NormOfResidual(A,b,xx)
    if res > 1.1 { FailMe(t, "Arls_test(4 failed: residual.") }  
    if mat.Norm(xx,2) > 0.0000001 { FailMe(t, "TestArlsnn(4) failed: x not zero.") }    

    ur = nr + ur + int(sigma) + int(lambda) // to avoid compiler complaint      
    return
}

func TestVecMin(t *testing.T) {
    fmt.Println("TestVecMin") 
    var g = mat.NewVecDense(6, []float64{1.0, 2.0, -1.0, 4.0, 5.0, 6.0} )
    gmin := vecMin(g)
    if gmin != -1.0 { FailMe(t, "TestVecMin(1) failed.") }
    return
}

func TestRowOps(t *testing.T) {
    fmt.Println("TestRowOps") 
    A := mat.NewDense(3, 4, 
        []float64{1.,0.,0.,0.,  0.,1.,0.,0.,  0.,0.,1.,0.} )
    exchangeRowsOf(A,1,2)        
    B := mat.NewDense(3, 4, 
        []float64{1.,0.,0.,0.,  0.,0.,1.,0.,  0.,1.,0.,0.} )
    if DiffAB(A,B) > 1.0e-9 { FailMe(t, "TestRowOps(1) failed!") }
    
    A = mat.NewDense(3, 3, 
        []float64{2.,2.,2.,  1.,1.,1.,  2.,2.,2.} )
    scaleRow(A,1,2.0) 
    B = mat.NewDense(3, 3, 
        []float64{2.,2.,2.,  2.,2.,2.,  2.,2.,2.} )
    if DiffAB(A,B) > 1.0e-9 { FailMe(t, "TestRowOps(2) failed!") }  
                
    A = mat.NewDense(3, 3, []float64{1.,2.,3.,  4.,5.,6.,  7.,8.,9.} )  
    B = A              
    r := dotRows(A,0,2)
    if !IsNear(r, 50.0) { FailMe(t, "TestRowOps(3) failed!") }     
    r = dotRowsAB(A,0,B,2)
    if !IsNear(r, 50.0) { FailMe(t, "TestRowOps(4) failed!") }          
}

func TestFindMax(t *testing.T) {
    fmt.Println("TestFindMax") 
    A := Mones(5,4)
    b := Iota(5)
    ib := findMaxSense(A,b)
    if ib != 4 { FailMe(t, "TestFindMax(1) failed!") }   
    
    A = MyMatRandom(10,11,0)
    scaleRow(A,7,3.0) 
    ib = findMaxRowNorm(A,0)
    if ib != 7 { FailMe(t, "TestFindMax(2) failed!") } 
}

func TestPrepeq(t *testing.T) {
    fmt.Println("TestPrepeq") 
    E := mat.NewDense(3, 3, []float64{1.,0.,0.,  1.,1.,1.,  1.,1.,0.} ) 
    x := mat.NewVecDense(3, []float64{4.,0.,1.} )
    f := Zeros(3)
    f.MulVec(E,x)     //[]float64{4.,5.,4.}
    EE, ff := prepeq(E,f)    
    if DiffAB(EE, Eye(3,3)) > 1.0e-9 { FailMe(t, "TestPrepeq(1) failed!") } 
        
    y := Zeros(3)
    y.MulVec(trans(EE),ff)
    if Diffxy(x,y) > 1.0e-9 { FailMe(t, "TestPrepeq(2) failed!") } 
}

func TestArlspj(t *testing.T) {
    fmt.Println("TestArlspj") 
    A := Eye(5,5)
    b := Ones(5)
    x := Ones(5)
    E := mat.NewDense(5, 5, []float64{
         0.,0.,0.,0.,0., 
         1.,0.,0.,0.,0., 
         1.,0.,0.,0.,0.,
         1.,1.,1.,1.,0.01,
         1.,1.,1.,1.,0.01 } ) 
    f := Zeros(5)
    f.MulVec(E,x) 

    AA, bb := arlspj(A, b, E, f, 1.0E-9)
    m, _ := AA.Dims()   
    r := Zeros(m)
    r.MulVec(AA,x)
    r.SubVec(r,bb) 
    if !IsNear(MyVecRms(r), 0.0) { FailMe(t, "TestArlspj(1) failed.") }  
}

func TestArlseq(t *testing.T) {
    fmt.Println("TestArlseq") 
    A := Eye(3,3)
    b := Ones(3)
    E := Mzeros(3,3)
    f := Zeros(3)
    ans := Ones(3)
    x, nr, ur, sigma, lambda := Arlseq(A, b, E, f)
    if Diffxy(x,ans) > 1.0e-9 { FailMe(t, "TestArlseq(1) failed!") } 
        
    E.Set(0,0, 1.0)
    f.SetVec(0, 2.0)
    ans.SetVec(0, 2.0)
    x, nr, ur, sigma, lambda = Arlseq(A, b, E, f)
    if Diffxy(x,ans) > 1.0e-9 { FailMe(t, "TestArlseq(2) failed!") } 
        
    E.Set(1,1, 1.0)
    f.SetVec(1, 2.0)
    ans.SetVec(1, 2.0)
    x, nr, ur, sigma, lambda = Arlseq(A, b, E, f)
    if Diffxy(x,ans) > 1.0e-9 { FailMe(t, "TestArlseq(3) failed!") }    
            
    E.Set(2,2, 1.0)
    f.SetVec(2, 2.0)
    ans.SetVec(2, 2.0)
    x, nr, ur, sigma, lambda = Arlseq(A, b, E, f)
    if Diffxy(x,ans) > 1.0e-9 { FailMe(t, "TestArlseq(4) failed!") }   
        
    E = MyMatRandom(3,3,3)
    ans = Ones(3)
    f.MulVec(E,ans)       
    x, nr, ur, sigma, lambda = Arlseq(A, b, E, f)
    if Diffxy(x,ans) > 1.0e-9 { FailMe(t, "TestArlseq(5) failed!") } 
        
    ur = nr + ur + int(sigma) + int(lambda) // to avoid diagnostic        
    return
}  

func TestArlsall(t *testing.T) { 
    fmt.Println("TestArlall") 
    A := Hilbert(7,6)     
    m, n := A.Dims()
    x := mat.NewVecDense(6, []float64{ 4., 3., 2., 1., 0., -1. })    
    b := Zeros(m)
    b.MulVec(A,x)
    b.AddVec(b,MyVecRandom(m, 0.000001))        
        
    E := Mzeros(2, n)
    f := Zeros(2)  
    for j:=0; j<n; j++ { E.Set(0,j,1.0) } //first row all 1's
    f.SetVec(0, Sum(x)) // sum must be exact
    E.Set(1,0,1.0)
    f.SetVec(1,5.0)  // x[1] must be exact 

    G := Eye(n,n)
    h := Zeros(n)     // require all x[i] non-neg
    
    Z := Mzeros(1,n)   // zero matrix for dummy
    z := Zeros(1)      // zero vector for dummy
    res := 0.0
    
    // solve with with (A,0,0) 
    xx, nr, ur, sigma, lambda := Arls(A,b)
    //fmt.Println("after A 0 O")
    res = NormOfResidual(A,b,xx)
    if res > 0.000001 { FailMe(t, "TestArlsAll(1) failed.") }
    
    // Solve with (A,E,0) 
    xx, nr, ur, sigma, lambda = Arlsall(A,b,E,f,Z,z)
    //fmt.Println("after A E O")
    res = NormOfResidual(A,b,xx)
    if res > 0.27 { 
        FailMe(t, "TestArlsAll(2) failed.") }
    if math.Abs(Sum(x) - Sum(xx)) > 0.00001 {
        FailMe(t, "TestArlsAll(3) failed.") }
    if math.Abs(xx.AtVec(0) - 5.0) > 0.000001 { 
        FailMe(t, "TestArlsAll(4) failed.") }

    // Solve with (A,0,G) 
    xx, nr, ur, sigma, lambda = Arlsall(A,b,Z,z,G,h)
    //fmt.Println("after  A 0 G")
    res = NormOfResidual(A,b,xx)
    
    if res > 0.002 { FailMe(t, "TestArlsAll(5) failed.") }
    if vecMin(xx) < -1.0E-9 { FailMe(t, "TestArlsAll(6) failed.") }
    
    // Solve with (A,E,G) 
    xx, nr, ur, sigma, lambda = Arlsall(A,b,E,f,G,h)
    //fmt.Println("after A E G")
    res = NormOfResidual(A,b,xx)
    if res > 0.29 { FailMe(t, "TestArlsAll(7) failed.") }
    if math.Abs(Sum(x) - Sum(xx)) > 0.00001 {
        FailMe(t, "TestArlsAll(8) failed: sum(x).") }
    if math.Abs(xx.AtVec(0) - 5.0) > 0.000001 { 
        FailMe(t, "TestArlsAll(9) failed.") }
    if vecMin(xx) < -1.0E-9 { FailMe(t, "TestArlsAll(10) failed.") }
        
    // Solve with Arlsgt 
    xx, nr, ur, sigma, lambda = Arlsgt(A,b,G,h)
    //fmt.Println("after Arlsgt")
    res = NormOfResidual(A,b,xx)
    if res > 0.002 { FailMe(t, "TestArlsAll(11) failed: residual.") }
    if vecMin(xx) < -1.0E-9 {  FailMe(t, "TestArlsAll(12) failed.") }
    
    ur = nr + ur + int(sigma) + int(lambda) + int(xx.AtVec(0)) //junk
    ur += int(G.At(0,0) + h.AtVec(0) + Z.At(0,0) + z.AtVec(0)) //junk
}

func TestArlsgt(t *testing.T) { 
    fmt.Println("TestArlsgt") 
    A := Hilbert(7,6)     
    m, _ := A.Dims()
    x := mat.NewVecDense(6, []float64{ 4., 2., 3., 1., 0., 1. })    
    b := Zeros(m)
    b.MulVec(A,x)
    b.AddVec(b,MyVecRandom(m, 0.000001))        
     
    G := Eye(5,6)
    for i:=0; i<5; i++ { G.Set(i,i+1,-1.0) } // solution must decrease
    h := Zeros(5)  
                
    // Solve with Arlsgt 
    xx, nr, ur, sigma, lambda := Arlsgt(A,b,G,h)
    for i:=0; i<5; i++ { 
        if xx.AtVec(i+1) > xx.AtVec(i) + 0.00000001 {
            FailMe(t, "TestArlsgt(1) failed.") }
    }
    ur = nr + ur + int(sigma) + int(lambda) + int(xx.AtVec(0)) //junk
    ur += int(G.At(0,0) + h.AtVec(0)) //junk
    //FailMe(t, "THIS IS JUST A TEST OF t.Fail()")
}    


