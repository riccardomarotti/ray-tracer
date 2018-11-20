package main

import "testing"

func TestConstructingTheTriangle(t *testing.T) {
	p1 := Point(0, 1, 0)
	p2 := Point(-1, 0, 0)
	p3 := Point(1, 0, 0)

	triangle := MakeTriangle(p1, p2, p3)

	AssertTupleEqual(p1, triangle.p1, t)
	AssertTupleEqual(p2, triangle.p2, t)
	AssertTupleEqual(p3, triangle.p3, t)
	AssertTupleEqual(Vector(-1, -1, 0), triangle.e1, t)
	AssertTupleEqual(Vector(1, -1, 0), triangle.e2, t)
	AssertTupleEqual(Vector(0, 0, -1), triangle.normal, t)
}
