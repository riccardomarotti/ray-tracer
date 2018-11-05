package main

// Translation creates a translation matrix
func Translation(x, y, z float64) Matrix {
	T := MakeIdentityMatrix(4)

	T.values[T.flatten(0, 3)] = x
	T.values[T.flatten(1, 3)] = y
	T.values[T.flatten(2, 3)] = z
	return T
}

// Scaling creates a scaling matrix
func Scaling(x, y, z float64) Matrix {
	S := MakeIdentityMatrix(4)
	S.values[S.flatten(0, 0)] = x
	S.values[S.flatten(1, 1)] = y
	S.values[S.flatten(2, 2)] = z

	return S
}
