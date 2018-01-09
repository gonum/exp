package blas64

import (
	"testing"
)

// Single vector operations ---------------------------------------------------
type OneVectorCase struct {
	Name      string
	Vector    SparseVector
	Panic     bool
	Asum      float64
	Nrm2      float64
	Idmax     int
	ScalCases []ScalCase
}

type ScalCase struct {
	Alpha  float64
	Answer SparseVector // Values of the sparse vector
	Name   string
}

var OneVectorCases = []OneVectorCase{
	{
		Name:   "Dense All Positive",
		Vector: SparseVector{5, []int{0, 1, 2, 3, 4}, []float64{6, 5, 4, 2, 6}},
		Panic:  false,
		Asum:   23,
		Nrm2:   10.81665382639196787935766380241148783875388972153573863813135,
		Idmax:  0,
		ScalCases: []ScalCase{
			{
				Alpha:  0,
				Answer: SparseVector{5, []int{}, []float64{}},
			},
			{
				Alpha:  1,
				Answer: SparseVector{5, []int{0, 1, 2, 3, 4}, []float64{6, 5, 4, 2, 6}},
			},
			{
				Alpha:  -2,
				Answer: SparseVector{5, []int{0, 1, 2, 3, 4}, []float64{-12, -10, -8, -4, -12}},
			},
		},
	},
	{
		Name:   "Sparse All Positive",
		Vector: SparseVector{15, []int{1, 3, 4, 10, 14}, []float64{6, 5, 4, 2, 6}},
		Panic:  false,
		Asum:   23,
		Nrm2:   10.81665382639196787935766380241148783875388972153573863813135,
		Idmax:  1,
		ScalCases: []ScalCase{
			{
				Alpha:  0,
				Answer: SparseVector{15, []int{}, []float64{}},
			},
			{
				Alpha:  1,
				Answer: SparseVector{15, []int{1, 3, 4, 10, 14}, []float64{6, 5, 4, 2, 6}},
			},
			{
				Alpha:  -2,
				Answer: SparseVector{15, []int{1, 3, 4, 10, 14}, []float64{-12, -10, -8, -4, -12}},
			},
		},
	},
	{
		Name:   "Dense All negative",
		Vector: SparseVector{5, []int{0, 1, 2, 3, 4}, []float64{-6, -5, -4, -2, -6}},
		Panic:  false,
		Asum:   23,
		Nrm2:   10.81665382639196787935766380241148783875388972153573863813135,
		Idmax:  3,
		ScalCases: []ScalCase{
			{
				Alpha:  0,
				Answer: SparseVector{5, []int{}, []float64{}},
			},
			{
				Alpha:  1,
				Answer: SparseVector{5, []int{0, 1, 2, 3, 4}, []float64{-6, -5, -4, -2, -6}},
			},
			{
				Alpha:  -2,
				Answer: SparseVector{5, []int{0, 1, 2, 3, 4}, []float64{12, 10, 8, 4, 12}},
			},
		},
	},
	{
		Name:   "Sparse All Negative",
		Vector: SparseVector{15, []int{1, 3, 4, 10, 14}, []float64{-6, -5, -4, -2, -6}},
		Panic:  false,
		Asum:   23,
		Nrm2:   10.81665382639196787935766380241148783875388972153573863813135,
		Idmax:  10,
		ScalCases: []ScalCase{
			{
				Alpha:  0,
				Answer: SparseVector{15, []int{}, []float64{}},
			},
			{
				Alpha:  1,
				Answer: SparseVector{15, []int{1, 3, 4, 10, 14}, []float64{-6, -5, -4, -2, -6}},
			},
			{
				Alpha:  -2,
				Answer: SparseVector{15, []int{1, 3, 4, 10, 14}, []float64{12, 10, 8, 4, 12}},
			},
		},
	},
	{
		Name:   "Empty",
		Vector: SparseVector{0, []int{}, []float64{}},
		Panic:  false,
		Asum:   0,
		Nrm2:   0.0,
		Idmax:  0,
		ScalCases: []ScalCase{
			{
				Alpha:  0,
				Answer: SparseVector{0, []int{}, []float64{}},
			},
			{
				Alpha:  1,
				Answer: SparseVector{0, []int{}, []float64{}},
			},
			{
				Alpha:  -2,
				Answer: SparseVector{0, []int{}, []float64{}},
			},
		},
	},
}

// Two vectors operations -----------------------------------------------------

type TwoSparseVectorsCase struct {
	Name      string
	Vector1   SparseVector
	Vector2   SparseVector
	Panic     bool
	Dot       float64
	CopyCases []CopyCase
	ScalCases []ScalCase
}

type CopyCase struct {
	Name  string
	Input SparseVector
	Panic bool
}

var TwoSparseVectorsCases = []TwoSparseVectorsCase{
	{
		Name:    "Dense All Positive",
		Vector1: SparseVector{5, []int{0, 1, 2, 3, 4}, []float64{6, 5, 4, 2, 6}},
		Vector2: SparseVector{5, []int{0, 1, 2, 3, 4}, []float64{6, 5, 4, 2, 6}},
		Panic:   false,
		Dot:     117,
		ScalCases: []ScalCase{
			{
				Name:   "Zero alpha",
				Alpha:  0,
				Answer: SparseVector{5, []int{0, 1, 2, 3, 4}, []float64{6, 5, 4, 2, 6}},
			},
			{
				Name:   "Positive alpha",
				Alpha:  1,
				Answer: SparseVector{5, []int{0, 1, 2, 3, 4}, []float64{12, 10, 8, 4, 12}},
			},
			{
				Name:   "Negative alpha",
				Alpha:  -2,
				Answer: SparseVector{5, []int{0, 1, 2, 3, 4}, []float64{-6, -5, -4, -2, -6}},
			},
		},
		CopyCases: []CopyCase{
			{
				Name:  "Input does not have the same length",
				Panic: true,
				Input: SparseVector{1, []int{0}, []float64{0}},
			},
			{
				Name:  "Same length but Ind and Data not matching",
				Panic: false,
				Input: SparseVector{5, []int{0, 1, 1}, []float64{0}},
			},
			{
				Name:  "Size, Ind and Data Matching",
				Panic: false,
				Input: SparseVector{5, []int{0, 1, 1, 1, 1}, []float64{0, 3, 4, 5, 1}},
			},
		},
	},
}

// Mixed vectors operations ---------------------------------------------------

type MixedVectorsCase struct {
	Name   string
	Sparse SparseVector
	Dense  DenseVector
	Panic  bool
}

var MixedVectorsCases = []MixedVectorsCase{
	{
		Name:   "Dense to dense",
		Sparse: SparseVector{5, []int{0, 1, 2, 3, 4}, []float64{0, 4, 6, 3, 0}},
		Dense:  DenseVector{[]float64{0, 4, 6, 3, 0}},
		Panic:  false,
	},
}

// Test single vector operations ----------------------------------------------

func TestNrm2(t *testing.T) {
	for _, c := range OneVectorCases {
		n := Nrm2(c.Vector)
		if n != c.Nrm2 {
			t.Errorf("nrm2: mismatch %v: expected %v, found %v", c.Name, c.Nrm2, n)
		}
	}
}

func TestAsum(t *testing.T) {
	for _, c := range OneVectorCases {
		n := Asum(c.Vector)
		if n != c.Asum {
			t.Errorf("asum: mismatch %v: expected %v, found %v", c.Name, c.Asum, n)
		}
	}
}

func TestIdmax(t *testing.T) {
	for _, c := range OneVectorCases {
		n := Idmax(c.Vector)
		if n != c.Idmax {
			t.Errorf("idmax: mismatch %v: expected %v, found %v", c.Name, c.Idmax, n)
		}
	}
}

func TestAscal(t *testing.T) {
	for _, c := range OneVectorCases {
		for _, scal := range c.ScalCases {
			scaled := Scal(scal.Alpha, c.Vector)
			// Size
			if len(scaled.Ind) != len(scaled.Data) {
				t.Errorf("ascal: mismatch of data and indices list size in %v: data %v, indices %v", c.Name, len(scaled.Data), len(scaled.Ind))
			}
			if scaled.Size != scal.Answer.Size {
				t.Errorf("ascal: mismatch of size in %v: expected %v, found %v", c.Name, scal.Answer.Size, scaled.Size)
			}
			if len(scaled.Ind) != len(scal.Answer.Ind) {
				t.Errorf("ascal: mismatch of indices list size in %v: expected %v, found %v", c.Name, len(scal.Answer.Ind), len(scaled.Ind))
			}
			// Indices (sanity check)
			if scal.Answer.Size != 0 {
				for i, ind := range scaled.Ind {
					if ind != scal.Answer.Ind[i] {
						t.Errorf("ascal: mismatch of indices in %v: expected %v, found %v", c.Name, ind, scal.Answer.Ind[i])
					}
				}
			}
			// Values
			if scal.Answer.Size != 0 {
				for i, val := range scaled.Data {
					if val != scal.Answer.Data[i] {
						t.Errorf("ascal: mismatch of data in %v: expected %v, found %v", c.Name, scal.Answer.Data[i], val)
					}
				}
			}
		}
	}
}

// Test 2 sparse vectors cases ------------------------------------------------

func TestDot(t *testing.T) {
	for _, c := range TwoSparseVectorsCases {
		dotProduct := Dot(c.Vector1, c.Vector2)
		dotProductsym := Dot(c.Vector2, c.Vector1)
		if dotProduct != dotProductsym {
			t.Errorf("dot: the result is not symmetrical in the inputs in %v: %d and %d", c.Name, dotProduct, dotProductsym)
		}
		if dotProduct != c.Dot {
			t.Error("dot: mismatch of dot product value in %v: expected %v, got %v", c.Name, c.Dot, dotProduct)
		}

	}
}

func TestSwap(t *testing.T) {
	for _, c := range TwoSparseVectorsCases {
		Vec1Copy := SparseVector{Size: c.Vector1.Size, Ind: c.Vector1.Ind, Data: c.Vector1.Data}
		Vec2Copy := SparseVector{Size: c.Vector2.Size, Ind: c.Vector2.Ind, Data: c.Vector2.Data}
		Swap(&c.Vector1, &c.Vector2)
		if !isEqualSparse(c.Vector2, Vec1Copy) {
			t.Error("swap: the vectors are different in %v", c.Name)
		}
		if !isEqualSparse(c.Vector1, Vec2Copy) {
			t.Error("swap: the vectors are different in %v", c.Name)
		}
	}
}

func TestCopy(t *testing.T) {
	for _, c := range TwoSparseVectorsCases {
		for _, cop := range c.CopyCases {
			// Check if funcion panics
			func() {
				if cop.Panic { // If we expect the function to panic
					defer func() {
						if r := recover(); r == nil {
							t.Errorf("copy: %v in %v should have panicked", cop.Name, c.Name)
						}
					}()
					Copy(&c.Vector1, &cop.Input)
				} else {
					Copy(&c.Vector1, &cop.Input)
					if !isEqualSparse(cop.Input, c.Vector1) {
						t.Error("copy: %v: %v: expected vectors to be equal", c.Name, cop.Name)
					}
				}
			}()
		}
	}
}

func TestAxpy(t *testing.T) {
	for _, c := range TwoSparseVectorsCases {
		for _, cal := range c.ScalCases {
			// Create a deep copy of y
			var output = SparseVector{Size: c.Vector2.Size,
				Ind:  make([]int, len(c.Vector2.Ind)),
				Data: make([]float64, len(c.Vector2.Data))}
			copy(output.Ind, c.Vector2.Ind)
			copy(output.Data, c.Vector2.Data)
			// Test
			Axpy(&c.Vector1, &output, cal.Alpha)
			if !isEqualSparse(output, cal.Answer) {
				t.Errorf("axpy: %s: %s: output does not match the expected vector", c.Name, cal.Name)
			}
		}
	}
}

// Test 1 sparse and 1 dense vectors cases ------------------------------------

func testFromDense(t *testing.T) {
	for _, c := range MixedVectorsCases {
		output := Gather(c.Dense)
		if !isEqualSparse(output, c.Sparse) {
			t.Errorf("fromDense: %s: output does not match the expected vector", c.Name)
		}
	}
}

func testScatter(t *testing.T) {
	for _, c := range MixedVectorsCases {
		output := Scatter(c.Sparse)
		if !isEqualDense(output, c.Dense) {
		}
	}
}

// Helper functions -----------------------------------------------------------

func isEqualSparse(x, y SparseVector) bool {
	if x.Size != y.Size {
		return false
	}
	for idx, i := range x.Ind {
		if i != y.Ind[idx] {
			return false
		}
	}
	for idx, d := range x.Data {
		if d != y.Data[idx] {
			return false
		}
	}

	return true
}

func isEqualDense(x, y DenseVector) bool {
	if len(x.Data) != len(y.Data) {
		return false
	}
	for i, dx := range x.Data {
		if dx != y.Data[i] {
			return false
		}
	}

	return true
}
