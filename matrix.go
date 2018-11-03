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

// At returns the matrix value at i X j
func (A Matrix) At(i, j int) float64 {
	return A.values[A.cols*i+j]
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
