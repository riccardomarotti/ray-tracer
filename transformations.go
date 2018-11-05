package main

// Translation create a translation matrix
func Translation(x, y, z float64) Matrix {
	T := MakeIdentityMatrix(4)

	T.values[T.flatten(0, 3)] = x
	T.values[T.flatten(1, 3)] = y
	T.values[T.flatten(2, 3)] = z
	return T
}
