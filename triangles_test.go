package main

import "testing"

func TestConstructingTheTriangle(t *testing.T) {
	p1 := Point(0, 1, 0)
	p2 := Point(-1, 0, 0)
	p3 := Point(1, 0, 0)

	triangle := MakeTriangle(p1, p2, p3, Identity())

	AssertTupleEqual(p1, triangle.p1, t)
	AssertTupleEqual(p2, triangle.p2, t)
	AssertTupleEqual(p3, triangle.p3, t)
	AssertTupleEqual(Vector(-1, -1, 0), triangle.e1, t)
	AssertTupleEqual(Vector(1, -1, 0), triangle.e2, t)
	AssertTupleEqual(Vector(0, 0, -1), triangle.normal, t)
}

func TestFindingTheNormalOnATriangle(t *testing.T) {
	triangle := MakeTriangle(Point(0, 1, 0), Point(-1, 0, 0), Point(1, 0, 0), Identity())
	n1 := triangle.NormalAt(Point(0, 0.5, 0))
	n2 := triangle.NormalAt(Point(-0.5, 0.75, 0))
	n3 := triangle.NormalAt(Point(0.5, 0.25, 0))

	AssertTupleEqual(n1, triangle.normal, t)
	AssertTupleEqual(n2, triangle.normal, t)
	AssertTupleEqual(n3, triangle.normal, t)

}
func TestIntersectingARayParallelToTheTriangle(t *testing.T) {
	triangle := MakeTriangle(Point(0, 1, 0), Point(-1, 0, 0), Point(1, 0, 0), Identity())
	r := Ray{Point(0, -1, -2), Vector(0, 1, 0)}

	Assert(len(triangle.Intersection(r)) == 0, "", t)
}
