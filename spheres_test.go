package main

import "testing"

func TestSphereDefaultTransformation(t *testing.T) {
	s := MakeSphere()

	AssertMatrixEqual(Identity(), s.Transform(), t)
}
