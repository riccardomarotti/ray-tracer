package main

import "math"

type Matrix struct {
	rows, cols int
	values     []float64
}

func MakeMatrix(rows, cols int, data []float64) Matrix {
	return Matrix{rows, cols, data}
}

func (A Matrix) flatten(i, j int) int {
	return A.cols*i + j
}

func (A Matrix) At(i, j int) float64 {
	return A.values[A.flatten(i, j)]
}

func (A Matrix) Equals(B Matrix) bool {
	for i, Avalue := range A.values {
		areEqual := math.Abs(Avalue-B.values[i]) < Epsilon
		if areEqual == false {
			return false
		}
	}

	return true
}

func (A Matrix) Multiply(B Matrix) Matrix {
	values := make([]float64, len(A.values))

	for i := 0; i < A.rows; i++ {
		for j := 0; j < A.cols; j++ {
			var value float64
			for offset := 0; offset < A.rows; offset++ {
				value += A.At(i, offset) * B.At(offset, j)
			}
			values[A.flatten(i, j)] = value
		}
	}

	return MakeMatrix(A.rows, A.cols, values)
}

func (A Matrix) MultiplyByTuple(b Tuple) Tuple {
	result := make([]float64, 4)
	for row := 0; row < 4; row++ {
		var value float64
		value += A.At(row, 0) * b.x
		value += A.At(row, 1) * b.y
		value += A.At(row, 2) * b.z
		value += A.At(row, 3) * b.w

		result[row] = value
	}

	return Tuple{result[0], result[1], result[2], result[3]}
}

func MakeIdentityMatrix(size int) (identity Matrix) {
	identity = MakeMatrix(size, size, make([]float64, size*size))
	for i := 0; i < size; i++ {
		identity.values[identity.flatten(i, i)] = 1
	}
	return
}

func Identity() Matrix {
	return MakeIdentityMatrix(4)
}

func (A Matrix) T() (At Matrix) {
	At = MakeMatrix(A.rows, A.cols, make([]float64, len(A.values)))
	for i := 0; i < A.rows; i++ {
		for j := 0; j < A.cols; j++ {
			At.values[At.flatten(i, j)] = A.values[A.flatten(j, i)]
		}
	}

	return
}

func (A Matrix) determinant2x2() float64 {
	return A.values[A.flatten(0, 0)]*A.values[A.flatten(1, 1)] - A.values[A.flatten(0, 1)]*A.values[A.flatten(1, 0)]
}

func (A Matrix) Determinant() float64 {
	if A.rows == 2 {
		return A.determinant2x2()
	}

	var det float64
	row := 0
	for i := 0; i < A.cols; i++ {
		det += A.values[A.flatten(i, 0)] * A.Cofactor(i, row)
	}
	return det
}

func (A Matrix) Submatrix(row, col int) Matrix {
	resultMatrix := MakeMatrix((A.rows - 1), (A.rows - 1), make([]float64, 0))

	for i := 0; i < A.rows; i++ {
		for j := 0; j < A.cols; j++ {
			if i != row && j != col {
				resultMatrix.values = append(resultMatrix.values, A.At(i, j))
			}

		}
	}
	return resultMatrix
}

func (A Matrix) Minor(i, j int) float64 {
	return A.Submatrix(i, j).Determinant()
}

func (A Matrix) Cofactor(i, j int) float64 {
	m := A.Minor(i, j)

	if (i+j)%2 == 1 {
		m = -m
	}

	return m
}

func (A Matrix) IsInvertible() bool {
	return A.Determinant() != 0
}

func (A Matrix) Inverse() (Ai Matrix) {
	Ai = MakeMatrix(A.rows, A.cols, make([]float64, A.rows*A.cols))
	Adet := A.Determinant()

	for i := 0; i < A.rows; i++ {
		for j := 0; j < A.cols; j++ {
			Ai.values[Ai.flatten(i, j)] = A.Cofactor(i, j) / Adet
		}
	}

	Ai = Ai.T()
	return
}

func ViewTransform(from, to, up Tuple) Matrix {
	forward := to.Subtract(from).Normalize()
	left := forward.Cross(up.Normalize())
	trueUp := left.Cross(forward)

	orientation := MakeMatrix(4, 4, []float64{
		left.x, left.y, left.z, 0,
		trueUp.x, trueUp.y, trueUp.z, 0,
		-forward.x, -forward.y, -forward.z, 0,
		0, 0, 0, 1,
	})

	return orientation.Translate(-from.x, -from.y, -from.z)
}
