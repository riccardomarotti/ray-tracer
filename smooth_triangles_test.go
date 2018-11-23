package main

import "testing"

func TestConstructingASmoothTriangle(t *testing.T) {
	p1 := Point(0, 1, 0)
	p2 := Point(-1, 0, 0)
	p3 := Point(1, 0, 0)
	n1 := Vector(0, 1, 0)
	n2 := Vector(-1, 0, 0)
	n3 := Vector(1, 0, 0)

	tri := MakeSmoothTriangle(p1, p2, p3, n1, n2, n3)

	AssertTupleEqual(p1, tri.p1, t)
	AssertTupleEqual(p2, tri.p2, t)
	AssertTupleEqual(p3, tri.p3, t)
	AssertTupleEqual(n1, tri.n1, t)
	AssertTupleEqual(n2, tri.n2, t)
	AssertTupleEqual(n3, tri.n3, t)
}
