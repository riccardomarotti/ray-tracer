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
