package arls

// Copyright 2021 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
import (
	"fmt"
	"math"
	"testing"

	"gonum.org/v1/gonum/mat"
)

func FailMe(t *testing.T, s string) { fmt.Println(s); t.Error(s); t.Fail() }

// Print routines for use in debugging

func MyVecPrint(x *mat.VecDense) {
	m, _ := x.Dims()
	fmt.Println("Vector")
	for i := 0; i < m; i++ {
		fmt.Println(x.AtVec(i))
	}
	fmt.Println(" ")
}

func MyMatPrint(A *mat.Dense) {
	m, n := A.Dims()
	fmt.Println("Matrix")
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			fmt.Printf(" %0.5f", A.At(i, j))
		}
		fmt.Printf("\n")
	}
	fmt.Println(" ")
}

func squareof(x float64) float64 {
	return x * x
}

func ones(m int) *mat.VecDense {
	x := mat.NewVecDense(m, nil)
	for i := 0; i < m; i++ {
		x.SetVec(i, 1)
	}
	return x
}

func Iota(m int) *mat.VecDense {
	x := mat.NewVecDense(m, nil)
	for i := 0; i < m; i++ {
		x.SetVec(i, float64(i))
	}
	return x
}

func Mones(m, n int) *mat.Dense {
	A := mat.NewDense(m, n, nil)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			A.Set(i, j, 1)
		}
	}
	return A
}

func eye(m, n int) *mat.Dense {
	A := mat.NewDense(m, n, nil)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			A.Set(i, j, 0.0)
			if i == j {
				A.Set(i, j, 1)
			}
		}
	}
	return A
}

func LowerTri(m, n int) *mat.Dense {
	A := mat.NewDense(m, n, nil)
	for i := 0; i < m; i++ {
		for j := 0; j < m; j++ {
			A.Set(i, j, 1)
		}
	}
	return A
}

func Hilbert(m, n int) *mat.Dense {
	A := mat.NewDense(m, n, nil)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			A.Set(i, j, 1/(1+float64(i+j)))
		}
	}
	return A
}

func Sum(x *mat.VecDense) float64 {
	m, _ := x.Dims()
	sum := 0.0
	for i := 0; i < m; i++ {
		sum += x.AtVec(i)
	}
	return sum
}

func Diffrms(x, y *mat.VecDense) float64 {
	mx := x.Len()
	my := y.Len()
	if mx != my {
		return 1.0E9
	}
	sum := 0.0
	for i := 0; i < mx; i++ {
		sum += squareof(x.AtVec(i) - y.AtVec(i))
	}
	sum = math.Sqrt(sum / float64(mx))
	return sum
}

func DiffABrms(A, B *mat.Dense) float64 {
	m, n := A.Dims()
	mb, nb := B.Dims()
	if m != mb || n != nb {
		return 1.0E9
	}
	sum := 0.0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			sum += squareof(A.At(i, j) - B.At(i, j))
		}
	}
	sum = math.Sqrt(sum / float64(m*n))
	return sum
}

func MyVecRandom(m int, err float64) *mat.VecDense {
	b := mat.NewVecDense(m, nil)
	for i := 0; i < m; i++ {
		// be sure this stays the same as in python
		b.SetVec(i, err*Myabs(math.Sin(
			float64(2*m)+2*float64(i))))
	}
	return b
}

func MyMatRandom(m, n, ibias int) *mat.Dense {
	A := mat.NewDense(m, n, nil)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			// be sure this stays the same as in python
			A.Set(i, j, Myabs(math.Sin(
				float64(ibias)+float64(2*m+3*n)+
					2*float64(i)+2.5*float64(j))))
		}
	}
	return A
}

func NormOfResidual(A *mat.Dense, b *mat.VecDense, x *mat.VecDense) float64 {
	m, _ := A.Dims()
	bb := mat.NewVecDense(m, nil)
	bb.MulVec(A, x)
	r := mat.NewVecDense(m, nil)
	r.SubVec(b, bb)
	res := mat.Norm(r, 2)
	// for debugging...
	//MyMatPrint(A)
	//MyVecPrint(b)
	//MyVecPrint(x)
	//fmt.Println("sum of x= ", Sum(x))
	//fmt.Println("Norm of Residual= ", res)
	return res
}

func Myabs(a float64) float64 {
	if a >= 0.0 {
		return a
	}
	return -a
}

func MyVecRms(x *mat.VecDense) float64 {
	return mat.Norm(x, 2) / math.Sqrt(float64(x.Len()))
}

func MyMatRms(A *mat.Dense) float64 {
	m, n := A.Dims()
	sum := 0.0
	a := 0.0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			a = A.At(i, j)
			sum = sum + a*a
		}
	}
	return math.Sqrt(sum / float64(m*n))
}

func IsNear(x, y float64) bool {
	tol := 1.0E-8*(Myabs(x)+Myabs(y)) + 1.0E-12
	return (Myabs(x-y) < tol)
}

func IsAbout(x, y float64) bool {
	tol := 1.0E-3*(Myabs(x)+Myabs(y)) + 1.0E-12
	return (Myabs(x-y) < tol)
}

// end of test utilities

func TestDeletes(t *testing.T) {
	fmt.Println("TestDeletes")
	A := mat.NewDense(4, 4, nil)
	for i := 0; i < 4; i++ {
		A.Set(i, 2, 99.0)
	}
	B := deleteColumn(A, 2)
	_, n := B.Dims()
	if n != 3 {
		FailMe(t, "TestDeletes(1) failed!")
	}
	if !isMatZero(B) {
		FailMe(t, "TestDeletes(2 failed!")
	}

	A = mat.NewDense(4, 4, nil)
	for j := 0; j < 4; j++ {
		A.Set(2, j, 99)
	}
	B = deleteRow(A, 2)
	m, _ := B.Dims()
	if m != 3 {
		FailMe(t, "TestDeletes(3) failed!")
	}
	if !isMatZero(B) {
		FailMe(t, "TestDeletes(4 failed!")
	}

	b := mat.NewVecDense(4, nil)
	b.SetVec(2, 99.0)
	b = deleteElement(b, 2)
	m, _ = b.Dims()
	if m != 3 {
		FailMe(t, "TestDeletes(5) failed!")
	}
	if !isVecZero(b) {
		FailMe(t, "TestDeletes(6) failed!")
	}
}

//return number of non-zero elements in A and their sum
func getStats(A *mat.Dense) (int, float64) {
	m, n := A.Dims()
	k := 0
	sum := 0.0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			sum += A.At(i, j)
			if A.At(i, j) != 0 {
				k++
			}
		}
	}
	return k, sum
}

func TestAppends(t *testing.T) {
	fmt.Println("TestAppends")
	G := eye(4, 4)
	A := mat.NewDense(4, 4, nil)
	A = appendRow(A, G.RowView(2).(*mat.VecDense))
	m, _ := A.Dims()
	if m != 5 {
		FailMe(t, "TestAppends(1) failed!")
	}
	if A.At(4, 2) != 1 {
		FailMe(t, "TestAppends(2) failed!")
	}
	k, sum := getStats(A)
	if k != 1 || sum != 1 {
		FailMe(t, "TestAppends(3) failed!")
	}

	b := mat.NewVecDense(4, nil)
	b = appendElement(b, 7)
	m, _ = b.Dims()
	if m != 5 {
		FailMe(t, "TestAppends(4) failed!")
	}
	if b.AtVec(4) != 7 {
		FailMe(t, "TestAppends(5) failed!")
	}
	if isVecZero(b) {
		FailMe(t, "TestAppends(6) failed!")
	}
}

func TestWidth(t *testing.T) {
	fmt.Println("TestWidth")
	if decideWidth(2) != 1 {
		FailMe(t, "TestWidth(2) failed.")
	}
	if decideWidth(5) != 2 {
		FailMe(t, "TestWidth(5) failed.")
	}
	if decideWidth(11) != 3 {
		FailMe(t, "TestWidth(11) failed.")
	}
	if decideWidth(19) != 4 {
		FailMe(t, "TestWidth(19) failed.")
	}
	if decideWidth(27) != 5 {
		FailMe(t, "TestWidth(27) failed.")
	}
	if decideWidth(30) != 6 {
		FailMe(t, "TestWidth(30) failed.")
	}
	if decideWidth(49) != 7 {
		FailMe(t, "TestWidth(49) failed.")
	}
	if decideWidth(60) != 8 {
		FailMe(t, "TestWidth(60) failed.")
	}
	if decideWidth(79) != 9 {
		FailMe(t, "TestWidth(79) failed.")
	}
	if decideWidth(150) != 10 {
		FailMe(t, "TestWidth(150) failed.")
	}
	if decideWidth(250) != 12 {
		FailMe(t, "TestWidth(250) failed.")
	}
	if decideWidth(350) != 14 {
		FailMe(t, "TestWidth(350) failed.")
	}
	if decideWidth(450) != 16 {
		FailMe(t, "TestWidth(450) failed.")
	}
	if decideWidth(1100) != 20 {
		FailMe(t, "TestWidth(1100) failed.")
	}
}

func TestMultiple(t *testing.T) {
	fmt.Println("TestMultiple")
	if decideMultiple(2) != 30 {
		FailMe(t, "TestMultiple(2) failed.")
	}
	if decideMultiple(8) != 20 {
		FailMe(t, "TestMultiple(8) failed.")
	}
	if decideMultiple(15) != 15 {
		FailMe(t, "TestMultiple(15) failed.")
	}
	if decideMultiple(30) != 7 {
		FailMe(t, "TestMultiple(30) failed.")
	}
}

func TestSplita(t *testing.T) {
	fmt.Println("TestSplita")
	var g = mat.NewVecDense(4, []float64{1, 1, 0, 20})
	if splita(g, 1) != 1 {
		FailMe(t, "TestSplita(1) failed.")
	}
	if splita(g, 2) != 2 {
		FailMe(t, "TestSplita(2) failed.")
	}
	if splita(g, 3) != 3 {
		FailMe(t, "TestSplita(3) failed.")
	}
	if splita(g, 4) != 3 {
		FailMe(t, "TestSplita(4) failed.")
	}
}

func TestMovSums(t *testing.T) {
	fmt.Println("TestMovSums")
	var g = mat.NewVecDense(4, []float64{1, 2, 3, 4})
	var ans = mat.NewVecDense(3, []float64{3, 5, 7})
	sums := computeMovSums(g, 4, 2)
	for i := 0; i < 3; i++ {
		if sums.AtVec(i) != ans.AtVec(i) {
			FailMe(t, "TestMovSums(1) failed.")
		}
	}
}

func TestSplitb(t *testing.T) {
	fmt.Println("TestSplitb")
	var g = mat.NewVecDense(6, []float64{1, 0.1, 0.01, 0.1, 1, 10})
	var ans = mat.NewVecDense(6, []float64{1, 2, 3, 4, 4, 4})
	var r float64
	for i := 0; i < 6; i++ {
		r = float64(splitb(g, i+1))
		if r != ans.At(i, 0) {
			FailMe(t, "TestSplitb(1) failed.")
		}
	}
}

func TestColumn(t *testing.T) {
	fmt.Println("TestColumn")
	A := mat.NewDense(3, 1, []float64{1, 1, 1})
	b := mat.NewVecDense(3, []float64{2, 2, 2})

	x, _, _, _, _ := Arls(A, b)
	if !IsNear(x.At(0, 0), 2) {
		FailMe(t, "TestColumn(1) failed.")
	}

	x, _, _, _, _ = Arlsnn(A, b)
	if !IsNear(x.At(0, 0), 2) {
		FailMe(t, "TestColumn(2) failed.")
	}
}

func TestRow(t *testing.T) {
	fmt.Println("TestRow")
	A := mat.NewDense(1, 3, []float64{1, 1, 1})
	b := mat.NewVecDense(1, []float64{3})

	x, _, _, _, _ := Arls(A, b)
	if !IsNear(MyVecRms(x), 1) {
		FailMe(t, "TestRow(1) failed.")
	}

	x, _, _, _, _ = Arlsnn(A, b)
	if !IsNear(MyVecRms(x), 1) {
		FailMe(t, "TestRow(2) failed.")
	}
}

func TestArls(t *testing.T) {
	fmt.Println("TestArls")
	n := 3
	m := 3
	A := mat.NewDense(m, n, nil)
	b := ones(3)
	bb := b

	//TEST WITH ZERO MATRIX
	x, _, _, _, _ := Arls(A, b)
	if mat.Norm(x, 2) != 0 {
		FailMe(t, "TestArls(1) failed.")
	}

	//TEST WITH ZERO RIGHT HAND SIDE
	A = eye(3, 3)
	b = mat.NewVecDense(3, nil)
	var svd mat.SVD
	ok := svd.Factorize(A, mat.SVDThin)
	if !ok {
		FailMe(t, "TestArls(2) failed.")
	}

	x, _, _, _, _ = Arlsvd(svd, b)
	if mat.Norm(x, 2) != 0 {
		FailMe(t, "TestArls(3) failed.")
	}

	x, _, _, _, _ = Arls(A, b)
	if mat.Norm(x, 2) != 0 {
		FailMe(t, "TestArls(4) failed.")
	}

	//DESIRED SOLUTION FOR NEXT SEVERAL TESTS
	n = 6
	ans := mat.NewVecDense(n, nil)
	for i := 0; i < n; i++ {
		ans.SetVec(i, float64(n-2-i))
	}

	//OVERDETERMINED TEST WITH HILBERT(7,6)
	A = Hilbert(7, 6)
	m, _ = A.Dims()
	b = mat.NewVecDense(m, nil)
	bb = mat.NewVecDense(m, nil)
	b.MulVec(A, ans)
	b.AddVec(b, MyVecRandom(m, 0.002))
	x, _, _, _, _ = Arls(A, b)
	bb.MulVec(A, x)
	if Diffrms(x, ans) > 0.21 {
		FailMe(t, "TestArls(5A) failed.")
	}
	if Diffrms(b, bb) > 0.0008 {
		FailMe(t, "TestArls(5B) failed.")
	}

	//SQUARE TEST WITH HILBERT(6,6)
	A = Hilbert(6, 6)
	m, _ = A.Dims()
	b = mat.NewVecDense(m, nil)
	b.MulVec(A, ans)
	x, _, _, _, _ = Arls(A, b)
	bb = mat.NewVecDense(m, nil)
	bb.MulVec(A, x)
	if Diffrms(x, ans) > 2.0E-9 {
		FailMe(t, "TestArls(6A) failed.")
	}
	if Diffrms(b, bb) > 1.0E-12 {
		FailMe(t, "TestArls(6B) failed.")
	}

	//DUPLICATE ROW TEST WITH HILBERT(6,6)
	A = Hilbert(6, 6)
	for j := 0; j < 6; j++ {
		A.Set(3, j, A.At(2, j))
	}
	m, _ = A.Dims()
	b = mat.NewVecDense(m, nil)
	b.MulVec(A, ans)
	x, _, _, _, _ = Arls(A, b)
	bb = mat.NewVecDense(m, nil)
	bb.MulVec(A, x)
	if Diffrms(x, ans) > 0.01 {
		FailMe(t, "TestArls(7A) failed.")
	}
	if Diffrms(b, bb) > 1.0E-12 {
		FailMe(t, "TestArls(7B) failed.")
	}

	//DUPLICATE COLUMN TEST WITH HILBERT(6,6)
	A = Hilbert(6, 6)
	for i := 0; i < 6; i++ {
		A.Set(i, 3, A.At(i, 2))
	}
	m, _ = A.Dims()
	b = mat.NewVecDense(m, nil)
	b.MulVec(A, ans)
	x, _, _, _, _ = Arls(A, b)
	bb = mat.NewVecDense(m, nil)
	bb.MulVec(A, x)
	if Diffrms(x, ans) > 0.3 {
		FailMe(t, "TestArls(8A) failed.")
	}
	if Diffrms(b, bb) > 1.0E-12 {
		FailMe(t, "TestArls(8B) failed.")
	}

	//UNDERDETERMINED TEST WITH HILBERT(5,6)
	A = Hilbert(5, 6)
	m, _ = A.Dims()
	b = mat.NewVecDense(m, nil)
	b.MulVec(A, ans)
	x, _, _, _, _ = Arls(A, b)
	bb = mat.NewVecDense(m, nil)
	bb.MulVec(A, x)
	if Diffrms(x, ans) > 0.02 {
		FailMe(t, "TestArls(9A) failed.")
	}
	if Diffrms(b, bb) > 1.0E-12 {
		FailMe(t, "TestArls(9B) failed.")
	}

	//TEST LARGER, REAL-LIKE SYSTEM
	n = 15
	ans = ones(n)
	for i := 0; i < n; i++ {
		ans.SetVec(i, float64(n-1-i))
	}
	A = Hilbert(n, n)
	m, _ = A.Dims()
	b = mat.NewVecDense(m, nil)
	b.MulVec(A, ans)
	b.AddVec(b, MyVecRandom(m, 0.0001))
	x, _, _, _, _ = Arls(A, b)
	bb = mat.NewVecDense(m, nil)
	bb.MulVec(A, x)
	if Diffrms(x, ans) > 0.25 {
		FailMe(t, "TestArls(10A) failed.")
	}
	if Diffrms(b, bb) > 4.0E-5 {
		FailMe(t, "TestArls(10B) failed.")
	}

	//COMPARE TO PYTHON
	n = 12
	A = Hilbert(n, n)
	ans = ones(n)
	b = mat.NewVecDense(n, nil)
	b.MulVec(A, ans)
	for i := 0; i < n; i++ {
		b.SetVec(i, b.AtVec(i)+0.00001*math.Sin(float64(i+i)))
	}
	x, _, _, _, _ = Arls(A, b)
	var xp = mat.NewVecDense(12,
		[]float64{0.998635, 1.013942, 0.980540, 0.986143,
			1.000395, 1.011578, 1.016739, 1.015970,
			1.010249, 1.000679, 0.988241, 0.973741})
	if Diffrms(x, xp) > 0.00001 {
		FailMe(t, "TestArls(11) failed.")
	}
}

func TestArlsnn(t *testing.T) {
	fmt.Println("TestArlsnn")
	//as columns are removed arlsnn will deal with square & underdet
	A := Hilbert(7, 8)
	m, n := A.Dims()
	ans := ones(n)
	for i := 0; i < n; i++ {
		ans.SetVec(i, float64(n-2-i))
	}
	b := mat.NewVecDense(m, nil)
	b.MulVec(A, ans)
	b.AddVec(b, MyVecRandom(m, 0.000001))
	x, _, _, _, _ := Arlsnn(A, b)
	res := NormOfResidual(A, b, x)
	if res > 1.5 {
		FailMe(t, "TestArlsnn(1) failed.")
	}

	//TEST "IMPOSSIBLE" PROBLEM WITH ARLSNN
	A = eye(3, 3)
	b = mat.NewVecDense(3, []float64{-1., -1., -1.})
	x, _, _, _, _ = Arlsnn(A, b)
	res = NormOfResidual(A, b, x)
	if res > 1.8 {
		FailMe(t, "TestArlsnn(2) failed.")
	}
	if mat.Norm(x, 2) > 1.0E-8 {
		FailMe(t, "TestArlsnn(3) failed.")
	}

	//TEST COMPUTED ZERO PROBLEM
	A = mat.NewDense(2, 3, []float64{1, 1, 1, 0, 0, 0})
	b = mat.NewVecDense(2, []float64{0, 1})
	x, _, _, _, _ = Arlsnn(A, b)
	res = NormOfResidual(A, b, x)
	if res > 1.1 {
		FailMe(t, "Arls_test(4 failed: residual.")
	}
	if mat.Norm(x, 2) > 0.0000001 {
		FailMe(t, "TestArlsnn(4) failed: x not zero.")
	}
}

func TestVecMin(t *testing.T) {
	fmt.Println("TestVecMin")
	var g = mat.NewVecDense(6, []float64{1, 2, -1, 4, 5, 6})
	gmin := vecMin(g)
	if gmin != -1.0 {
		FailMe(t, "TestVecMin(1) failed.")
	}
}

func TestRowOps(t *testing.T) {
	fmt.Println("TestRowOps")
	A := mat.NewDense(3, 4,
		[]float64{1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0})
	exchangeRowsOf(A, 1, 2)
	B := mat.NewDense(3, 4,
		[]float64{1, 0, 0, 0, 0, 0, 1, 0, 0, 1, 0, 0})
	if DiffABrms(A, B) > 1.0e-9 {
		FailMe(t, "TestRowOps(1) failed!")
	}

	A = mat.NewDense(3, 3,
		[]float64{2, 2, 2, 1, 1, 1, 2, 2, 2})
	scaleRow(A, 1, 2.0)
	B = mat.NewDense(3, 3,
		[]float64{2, 2, 2, 2, 2, 2, 2, 2, 2})
	if DiffABrms(A, B) > 1.0e-9 {
		FailMe(t, "TestRowOps(2) failed!")
	}
}

func TestFindMax(t *testing.T) {
	fmt.Println("TestFindMax")
	A := Mones(5, 4)
	b := Iota(5)
	ib := findMaxSense(A, b)
	if ib != 4 {
		FailMe(t, "TestFindMax(1) failed!")
	}

	A = MyMatRandom(10, 11, 0)
	scaleRow(A, 7, 3.0)
	ib = findMaxRowNorm(A, 0)
	if ib != 7 {
		FailMe(t, "TestFindMax(2) failed!")
	}
}

func TestPrepeq(t *testing.T) {
	fmt.Println("TestPrepeq")
	E := mat.NewDense(3, 3, []float64{1, 0, 0, 1, 1, 1, 1, 1, 0})
	x := mat.NewVecDense(3, []float64{4, 0, 1})
	f := mat.NewVecDense(3, nil)
	f.MulVec(E, x) //[]float64{4.,5.,4.}
	EE, ff := prepeq(E, f)
	y := mat.NewVecDense(3, nil)
	y.MulVec(EE.T(), ff)
	if Diffrms(x, y) > 1.0e-9 {
		FailMe(t, "TestPrepeq(1) failed!")
	}
}

func TestArlspj(t *testing.T) {
	fmt.Println("TestArlspj")
	A := eye(5, 5)
	b := ones(5)
	x := ones(5)
	E := mat.NewDense(5, 5, []float64{
		0, 0, 0, 0, 0,
		1, 0, 0, 0, 0,
		1, 0, 0, 0, 0,
		1, 1, 1, 1, 0.01,
		1, 1, 1, 1, 0.01})
	f := mat.NewVecDense(5, nil)
	f.MulVec(E, x)

	AA, bb := arlspj(A, b, E, f, 1.0E-9)
	m, _ := AA.Dims()
	r := mat.NewVecDense(m, nil)
	r.MulVec(AA, x)
	r.SubVec(r, bb)
	if !IsNear(MyVecRms(r), 0) {
		FailMe(t, "TestArlspj(1) failed.")
	}
}

func TestArlseq(t *testing.T) {
	fmt.Println("TestArlseq")
	A := eye(3, 3)
	b := ones(3)
	E := mat.NewDense(3, 3, nil)
	f := mat.NewVecDense(3, nil)
	ans := ones(3)
	x, _, _, _, _ := Arlseq(A, b, E, f)
	if Diffrms(x, ans) > 1.0e-9 {
		FailMe(t, "TestArlseq(1) failed!")
	}

	E.Set(0, 0, 1)
	f.SetVec(0, 2)
	ans.SetVec(0, 2)
	x, _, _, _, _ = Arlseq(A, b, E, f)
	if Diffrms(x, ans) > 1.0e-9 {
		FailMe(t, "TestArlseq(2) failed!")
	}

	E.Set(1, 1, 1)
	f.SetVec(1, 2)
	ans.SetVec(1, 2)
	x, _, _, _, _ = Arlseq(A, b, E, f)
	if Diffrms(x, ans) > 1.0e-9 {
		FailMe(t, "TestArlseq(3) failed!")
	}

	E.Set(2, 2, 1)
	f.SetVec(2, 2)
	ans.SetVec(2, 2)
	x, _, _, _, _ = Arlseq(A, b, E, f)
	if Diffrms(x, ans) > 1.0e-9 {
		FailMe(t, "TestArlseq(4) failed!")
	}

	E = MyMatRandom(3, 3, 3)
	ans = ones(3)
	f.MulVec(E, ans)
	x, _, _, _, _ = Arlseq(A, b, E, f)
	if Diffrms(x, ans) > 1.0e-9 {
		FailMe(t, "TestArlseq(5) failed!")
	}
}

func TestArlsall(t *testing.T) {
	fmt.Println("TestArlall")
	A := Hilbert(7, 6)
	m, n := A.Dims()
	ans := mat.NewVecDense(6, []float64{4, 3, 2, 1, 0, -1})
	b := mat.NewVecDense(m, nil)
	b.MulVec(A, ans)
	b.AddVec(b, MyVecRandom(m, 0.000001))

	E := mat.NewDense(2, n, nil)
	f := mat.NewVecDense(2, nil)
	for j := 0; j < n; j++ {
		E.Set(0, j, 1.0)
	}
	f.SetVec(0, Sum(ans)) // sum must be exact
	E.Set(1, 0, 1)
	f.SetVec(1, 5) // x[1] must be exact

	G := eye(n, n)
	h := mat.NewVecDense(n, nil) // require all x[i] non-neg

	Z := mat.NewDense(1, n, nil) // zero matrix for dummy
	z := mat.NewVecDense(1, nil) // zero vector for dummy
	res := 0.0

	// solve with with (A,0,0)
	x, _, _, _, _ := Arls(A, b)
	//fmt.Println("after A 0 O")
	res = NormOfResidual(A, b, x)
	if res > 0.000001 {
		FailMe(t, "TestArlsAll(1) failed.")
	}

	// Solve with (A,E,0)
	x, _, _, _, _ = Arlsall(A, b, E, f, Z, z)
	//fmt.Println("after A E O")
	res = NormOfResidual(A, b, x)
	if res > 0.27 {
		FailMe(t, "TestArlsAll(2) failed.")
	}
	if math.Abs(Sum(ans)-Sum(x)) > 0.00001 {
		FailMe(t, "TestArlsAll(2A) failed.")
	}
	if math.Abs(x.AtVec(0)-5.0) > 0.000001 {
		FailMe(t, "TestArlsAll(2B) failed.")
	}

	// Solve with (A,0,G)
	x, _, _, _, _ = Arlsall(A, b, Z, z, G, h)
	//fmt.Println("after  A 0 G")
	res = NormOfResidual(A, b, x)
	if res > 0.002 {
		FailMe(t, "TestArlsAll(3A) failed.")
	}
	if vecMin(x) < -1.0E-9 {
		FailMe(t, "TestArlsAll(3B) failed.")
	}

	// Solve with (A,E,G)
	x, _, _, _, _ = Arlsall(A, b, E, f, G, h)
	//fmt.Println("after A E G")
	res = NormOfResidual(A, b, x)
	if res > 0.30 {
		FailMe(t, "TestArlsAll(4A) failed.")
	}
	if math.Abs(Sum(ans)-Sum(x)) > 0.00001 {
		FailMe(t, "TestArlsAll(4B) failed: sum(ans).")
	}
	if math.Abs(x.AtVec(0)-5.0) > 0.000001 {
		FailMe(t, "TestArlsAll(4C) failed.")
	}
	if vecMin(x) < -1.0E-9 {
		FailMe(t, "TestArlsAll(4D) failed.")
	}

	// Solve with Arlsgt
	x, _, _, _, _ = Arlsgt(A, b, G, h)
	//fmt.Println("after Arlsgt")
	res = NormOfResidual(A, b, x)
	if res > 0.002 {
		FailMe(t, "TestArlsAll(5A) failed: residual.")
	}
	if vecMin(x) < -1.0E-9 {
		FailMe(t, "TestArlsAll(5B) failed.")
	}
}

func TestArlsgt(t *testing.T) {
	fmt.Println("TestArlsgt")
	A := Hilbert(7, 6)
	m, _ := A.Dims()
	ans := mat.NewVecDense(6, []float64{4, 2, 3, 1, 0, 1})
	b := mat.NewVecDense(m, nil)
	b.MulVec(A, ans)
	b.AddVec(b, MyVecRandom(m, 0.000001))

	G := eye(5, 6)
	for i := 0; i < 5; i++ {
		G.Set(i, i+1, -1)
	} // solution must decrease
	h := mat.NewVecDense(5, nil)

	// Solve with Arlsgt
	x, _, _, _, _ := Arlsgt(A, b, G, h)
	for i := 0; i < 5; i++ {
		if x.AtVec(i+1) > x.AtVec(i)+0.00000001 {
			FailMe(t, "TestArlsgt(1) failed.")
		}
	}
}
