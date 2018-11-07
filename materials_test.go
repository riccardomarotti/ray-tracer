package main

import "testing"

func TestDefaultMaterial(t *testing.T) {
	m := MakeMaterial()

	Assert(Color{1, 1, 1} == m.color, "", t)
	AssertEqual(0.1, m.ambient, t)
	AssertEqual(0.9, m.diffuse, t)
	AssertEqual(0.9, m.specular, t)
	AssertEqual(200, m.shininess, t)
}
