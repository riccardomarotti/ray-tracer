package main

import (
	"math"
)

// Translate creates a translation matrix
func (A Matrix) Translate(x, y, z float64) Matrix {
	T := MakeIdentityMatrix(A.rows)

	T.values[T.flatten(0, 3)] = x
	T.values[T.flatten(1, 3)] = y
	T.values[T.flatten(2, 3)] = z
	return A.Multiply(T)
}

// Scale creates a scaling matrix
func (A Matrix) Scale(x, y, z float64) Matrix {
	S := MakeIdentityMatrix(A.rows)
	S.values[S.flatten(0, 0)] = x
	S.values[S.flatten(1, 1)] = y
	S.values[S.flatten(2, 2)] = z

	return A.Multiply(S)
}

// RotateX creates a matrix rotation around X axis of the given radians
func (A Matrix) RotateX(r float64) Matrix {
	Rx := MakeIdentityMatrix(A.rows)
	cosR := math.Cos(r)
	sinR := math.Sin(r)

	Rx.values[Rx.flatten(1, 1)] = cosR
	Rx.values[Rx.flatten(2, 2)] = cosR
	Rx.values[Rx.flatten(1, 2)] = -sinR
	Rx.values[Rx.flatten(2, 1)] = sinR

	return A.Multiply(Rx)
}

// RotateY creates a matrix rotation around Y axis of the given radians
func (A Matrix) RotateY(r float64) Matrix {
	Ry := MakeIdentityMatrix(A.rows)
	cosR := math.Cos(r)
	sinR := math.Sin(r)

	Ry.values[Ry.flatten(0, 0)] = cosR
	Ry.values[Ry.flatten(2, 2)] = cosR
	Ry.values[Ry.flatten(2, 0)] = -sinR
	Ry.values[Ry.flatten(0, 2)] = sinR

	return A.Multiply(Ry)
}

// RotateZ creates a matrix rotation around Z axis of the given radians
func (A Matrix) RotateZ(r float64) Matrix {
	Rz := MakeIdentityMatrix(A.rows)
	cosR := math.Cos(r)
	sinR := math.Sin(r)

	Rz.values[Rz.flatten(0, 0)] = cosR
	Rz.values[Rz.flatten(1, 1)] = cosR
	Rz.values[Rz.flatten(0, 1)] = -sinR
	Rz.values[Rz.flatten(1, 0)] = sinR

	return A.Multiply(Rz)
}

// Shear creates a shearing transform
// xy is shearing of x in proportion of y
// xz is shearing of x in proportion of z
// yx is shearing of y in proportion of x
// yz is shearing of z in proportion of z
// zx is shearing of z in proportion of x
// zy is shearing of z in proportion of y
func (A Matrix) Shear(xy, xz, yx, yz, zx, zy float64) Matrix {
	S := MakeIdentityMatrix(A.rows)
	S.values[S.flatten(0, 1)] = xy
	S.values[S.flatten(0, 2)] = xz
	S.values[S.flatten(1, 0)] = yx
	S.values[S.flatten(2, 0)] = zx
	S.values[S.flatten(1, 2)] = yz
	S.values[S.flatten(2, 1)] = zy
	return A.Multiply(S)
}
