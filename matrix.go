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
func (m Matrix) At(i, j int) float64 {
	return m.values[m.cols*i+j]
}
