package main

import (
	"fmt"
	"testing"
)

func TestMatrixCreation(t *testing.T) {
	A := MakeMatrix(4, 4, []float64{1, 2, 3, 4, 5.5, 6.5, 7.5, 8.5, 9, 10, 11, 12, 13.5, 14.5, 15.5, 16.5})
	Assert(A.At(0, 0) == 1, "", t)
	Assert(A.At(0, 3) == 4, "", t)
	Assert(A.At(1, 0) == 5.5, "", t)
	Assert(A.At(1, 2) == 7.5, "", t)
	Assert(A.At(2, 2) == 11, "", t)
	Assert(A.At(3, 0) == 13.5, "", t)
	Assert(A.At(3, 2) == 15.5, "", t)

	A = MakeMatrix(2, 2, []float64{-3, 5, 1, -2})
	Assert(A.At(0, 0) == -3, "", t)
	Assert(A.At(0, 1) == 5, "", t)
	Assert(A.At(1, 0) == 1, "", t)
	Assert(A.At(1, 1) == -2, "", t)

	A = MakeMatrix(3, 3, []float64{-3, 5, 0, 1, -2, -7, 0, 1, 1})
	Assert(A.At(0, 0) == -3, "", t)
	Assert(A.At(1, 1) == -2, "", t)
	Assert(A.At(2, 2) == 1, "", t)
}

func TestEqualMatrices(t *testing.T) {
	A := MakeMatrix(4, 4, []float64{1, 2, 3, 4, 2, 3, 4, 5, 3, 4, 5, 6, 4, 5, 6, 7})
	B := MakeMatrix(4, 4, []float64{1, 2, 3, 4, 2, 3, 4, 5, 3, 4, 5, 6, 4, 5, 6, 7})

	Assert(A.Equals(B), "Matrices should be equal", t)
	Assert(B.Equals(A), "Matrices should be different equal", t)
}

func TestDifferentMatricestr(t *testing.T) {
	A := MakeMatrix(4, 4, []float64{0, 2, 3, 4, 2, 3, 4, 5, 3, 4, 5, 6, 4, 5, 6, 7})
	B := MakeMatrix(4, 4, []float64{1, 2, 3, 4, 2, 3, 4, 5, 3, 4, 5, 6, 4, 5, 6, 7})

	Assert(false == A.Equals(B), "Matrices should be different", t)
	Assert(false == B.Equals(A), "Matrices should be different", t)
}

func TestMatrixMultiplication(t *testing.T) {
	A := MakeMatrix(4, 4, []float64{1, 2, 3, 4, 2, 3, 4, 5, 3, 4, 5, 6, 4, 5, 6, 7})
	B := MakeMatrix(4, 4, []float64{0, 1, 2, 4, 1, 2, 4, 8, 2, 4, 8, 16, 4, 8, 16, 32})
	AB := MakeMatrix(4, 4, []float64{24, 49, 98, 196, 31, 64, 128, 256, 38, 79, 158, 316, 45, 94, 188, 376})

	actualProduct := A.Multiply(B)
	Assert(actualProduct.Equals(AB), fmt.Sprintf("Expected:\n%v\nBut was:\n%v", AB, actualProduct), t)
}

func TestMatrixTupleMultiplication(t *testing.T) {
	A := MakeMatrix(4, 4, []float64{1, 2, 3, 4, 2, 4, 4, 2, 8, 6, 4, 1, 0, 0, 0, 1})
	b := Tuple{1, 2, 3, 1}

	Ab := A.MultiplyByTuple(b)
	expectedTuple := Tuple{18, 24, 33, 1}

	AssertTupleEqual(expectedTuple, Ab, t)
}

func TestIdentityMatrix(t *testing.T) {
	A := MakeMatrix(4, 4, []float64{0, 1, 2, 4, 1, 2, 4, 8, 2, 4, 8, 16, 4, 8, 16, 32})

	AI := A.Multiply(MakeIdentityMatrix(4))

	Assert(A.Equals(AI), "", t)
}

func TestTransposeMatrix(t *testing.T) {
	A := MakeMatrix(4, 4, []float64{0, 9, 3, 0, 9, 8, 0, 8, 1, 8, 5, 3, 0, 0, 5, 8})
	At := MakeMatrix(4, 4, []float64{0, 9, 1, 0, 9, 8, 8, 0, 3, 0, 5, 5, 0, 8, 3, 8})

	actualTransposed := A.T()
	Assert(actualTransposed.Equals(At), fmt.Sprintf("Transposed matrix expected to be\n%v\nbut was\n%v", At, actualTransposed), t)
}

func TestDeterminantOf2x2Matrix(t *testing.T) {
	A := MakeMatrix(2, 2, []float64{1, 5, -3, 2})

	AssertEqual(17, A.Determinant(), t)
}

func TestSubmatrix3x3(t *testing.T) {
	A := MakeMatrix(3, 3, []float64{1, 5, 0, -3, 2, 7, 0, 6, -3})
	ExpectedSubA := MakeMatrix(2, 2, []float64{-3, 2, 0, 6})

	ActualSubA := A.Submatrix(0, 2)
	Assert(ActualSubA.Equals(ExpectedSubA), fmt.Sprintf("Expected submatrix\n%v\nbut was\n%v", ExpectedSubA, ActualSubA), t)
}

func TestSubmatrix4x4(t *testing.T) {
	A := MakeMatrix(4, 4, []float64{-6, 1, 1, 6, -8, 5, 8, 6, -1, 0, 8, 2, -7, 1, -1, 1})
	ExpectedSubA := MakeMatrix(3, 3, []float64{-6, 1, 6, -8, 8, 6, -7, -1, 1})

	ActualSubA := A.Submatrix(2, 1)
	Assert(ActualSubA.Equals(ExpectedSubA), fmt.Sprintf("Expected submatrix\n%v\nbut was\n%v", ExpectedSubA, ActualSubA), t)
}

func TestMinor3x3(t *testing.T) {
	A := MakeMatrix(3, 3, []float64{3, 5, 0, 2, -1, -7, 6, -1, 5})

	AssertEqual(25, A.Minor(1, 0), t)
}

func TestCofactor(t *testing.T) {
	A := MakeMatrix(3, 3, []float64{3, 5, 0, 2, -1, -7, 6, -1, 5})

	AssertEqual(-12, A.Cofactor(0, 0), t)
	AssertEqual(-25, A.Cofactor(1, 0), t)
}

func TestDeterminantOf3x3Matrix(t *testing.T) {
	A := MakeMatrix(3, 3, []float64{1, 2, 6, -5, 8, -4, 2, 6, 4})

	AssertEqual(-196, A.Determinant(), t)
}

func TestDeterminantOf4x4Matrix(t *testing.T) {
	A := MakeMatrix(4, 4, []float64{-2, -8, 3, 5, -3, 1, 7, 3, 1, 2, -9, 6, -6, 7, 7, -9})

	AssertEqual(-4071, A.Determinant(), t)
}

func TestInvertibleMatrix(t *testing.T) {
	A := MakeMatrix(4, 4, []float64{6, 4, 4, 4, 5, 5, 7, 6, 4, -8, 3, -7, 9, 1, 7, -6})
	Assert(A.IsInvertible(), "", t)

	A = MakeMatrix(4, 4, []float64{-4, 2, -2, -3, 9, 6, 2, 6, 0, -5, 1, -5, 0, 0, 0, 0})
	Assert(A.IsInvertible() == false, "", t)
}

func TestMatrixInversion(t *testing.T) {
	A := MakeMatrix(4, 4, []float64{-5, 2, 6, -8, 1, -5, 1, 8, 7, 7, -6, -7, 1, -3, 7, 4})

	expectedAinvert := MakeMatrix(4, 4,
		[]float64{
			0.21805, 0.45113, 0.24060, -0.04511,
			-0.80827, -1.45677, -0.44361, 0.52068,
			-0.07895, -0.22368, -0.05263, 0.19737,
			-0.52256, -0.81391, -0.30075, 0.30639,
		})
	actualInvert := A.Inverse()

	Assert(expectedAinvert.Equals(actualInvert), fmt.Sprintf("Expected invers matrix\n%v\n but was\n%v", expectedAinvert, actualInvert), t)

	B := MakeMatrix(4, 4, []float64{8, -5, 9, 2, 7, 5, 6, 1, -6, 0, 9, 6, -3, 0, -9, -4})

	expectedAinvert = MakeMatrix(4, 4,
		[]float64{
			-0.15385, -0.15385, -0.28205, -0.53846,
			-0.07692, 0.12308, 0.02564, 0.03077,
			0.35897, 0.35897, 0.43590, 0.92308,
			-0.69231, -0.69231, -0.76923, -1.92308,
		})
	actualInvert = B.Inverse()

	Assert(expectedAinvert.Equals(actualInvert), fmt.Sprintf("Expected invers matrix\n%v\n but was\n%v", expectedAinvert, actualInvert), t)

	C := A.Multiply(B)

	Assert(C.Multiply(B.Inverse()).Equals(A), "", t)
}
