package main

import (
	"math"
)

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

// RotationX creates a matrix rotation around X axis of the given radians
func RotationX(r float64) Matrix {
	Rx := MakeIdentityMatrix(4)
	cosR := math.Cos(r)
	sinR := math.Sin(r)

	Rx.values[Rx.flatten(1, 1)] = cosR
	Rx.values[Rx.flatten(2, 2)] = cosR
	Rx.values[Rx.flatten(1, 2)] = -sinR
	Rx.values[Rx.flatten(2, 1)] = sinR

	return Rx
}

// RotationY creates a matrix rotation around Y axis of the given radians
func RotationY(r float64) Matrix {
	Ry := MakeIdentityMatrix(4)
	cosR := math.Cos(r)
	sinR := math.Sin(r)

	Ry.values[Ry.flatten(0, 0)] = cosR
	Ry.values[Ry.flatten(2, 2)] = cosR
	Ry.values[Ry.flatten(2, 0)] = -sinR
	Ry.values[Ry.flatten(0, 2)] = sinR

	return Ry
}

// RotationZ creates a matrix rotation around Z axis of the given radians
func RotationZ(r float64) Matrix {
	Rz := MakeIdentityMatrix(4)
	cosR := math.Cos(r)
	sinR := math.Sin(r)

	Rz.values[Rz.flatten(0, 0)] = cosR
	Rz.values[Rz.flatten(1, 1)] = cosR
	Rz.values[Rz.flatten(0, 1)] = -sinR
	Rz.values[Rz.flatten(1, 0)] = sinR

	return Rz
}
