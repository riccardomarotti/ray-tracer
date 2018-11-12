package main

import "testing"

func TestPlaneImplementsObjectInterface(t *testing.T) {
	p := Plane{}

	p.Transform()
	p.NormalAt(Point(0, 0, 0))
	p.Material()
}
func TestTheNormalOfAPlaneIsConstantEverywhere(t *testing.T) {
	p := Plane{}
	n1 := p.NormalAt(Point(0, 0, 0))
	n2 := p.NormalAt(Point(10, 0, -10))
	n3 := p.NormalAt(Point(-5, 0, 150))

	AssertTupleEqual(Vector(0, 1, 0), n1, t)
	AssertTupleEqual(Vector(0, 1, 0), n2, t)
	AssertTupleEqual(Vector(0, 1, 0), n3, t)
}
