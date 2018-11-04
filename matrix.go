package main

// Matrix represents a matrix
type Matrix struct {
	rows, cols int
	values     []float64
}

// MakeMatrix creates a matrix of rows X cols with data values
func MakeMatrix(rows, cols int, data []float64) Matrix {
	return Matrix{rows, cols, data}
}

func (A Matrix) flatten(i, j int) int {
	return A.cols*i + j
}

// At returns the matrix value at i X j
func (A Matrix) At(i, j int) float64 {
	return A.values[A.flatten(i, j)]
}

// Equals checks for matrix equality
func (A Matrix) Equals(B Matrix) bool {
	for i, Avalue := range A.values {
		if Avalue != B.values[i] {
			return false
		}
	}

	return true
}

// Multiply calulates the product of two matrices
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

// MultiplyByTuple multiplies a matrix by a tuple
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

// MakeIdentityMatrix creates an identiy matrix of the given size
func MakeIdentityMatrix(size int) (identity Matrix) {
	identity = MakeMatrix(size, size, make([]float64, size*size))
	for i := 0; i < size; i++ {
		identity.values[identity.flatten(i, i)] = 1
	}
	return
}
