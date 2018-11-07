package main

import "testing"

func TestSphereTransformation(t *testing.T) {
	T := Identity().Translate(2, 3, 4)
	s := MakeSphere(T)

	AssertMatrixEqual(T, s.Transform(), t)
}
