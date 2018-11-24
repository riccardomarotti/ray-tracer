package main

import "testing"

func TestConstructingTheTriangle(t *testing.T) {
	p1 := Point(0, 1, 0)
	p2 := Point(-1, 0, 0)
	p3 := Point(1, 0, 0)

	triangle := MakeTriangle(p1, p2, p3, Identity(), DefaultMaterial())

	AssertTupleEqual(p1, triangle.p1, t)
	AssertTupleEqual(p2, triangle.p2, t)
	AssertTupleEqual(p3, triangle.p3, t)
	AssertTupleEqual(Vector(-1, -1, 0), triangle.e1, t)
	AssertTupleEqual(Vector(1, -1, 0), triangle.e2, t)
	AssertTupleEqual(Vector(0, 0, -1), triangle.normal, t)
}

func TestFindingTheNormalOnATriangle(t *testing.T) {
	triangle := MakeTriangle(Point(0, 1, 0), Point(-1, 0, 0), Point(1, 0, 0), Identity(), DefaultMaterial())
	n1 := triangle.NormalAt(Point(0, 0.5, 0))
	n2 := triangle.NormalAt(Point(-0.5, 0.75, 0))
	n3 := triangle.NormalAt(Point(0.5, 0.25, 0))

	AssertTupleEqual(n1, triangle.normal, t)
	AssertTupleEqual(n2, triangle.normal, t)
	AssertTupleEqual(n3, triangle.normal, t)

}
func TestIntersectingARayParallelToTheTriangle(t *testing.T) {
	triangle := MakeTriangle(Point(0, 1, 0), Point(-1, 0, 0), Point(1, 0, 0), Identity(), DefaultMaterial())
	r := Ray{Point(0, -1, -2), Vector(0, 1, 0)}

	Assert(len(triangle.Intersection(r)) == 0, "Intersection ha d to be empty", t)
}

func TestARayMissesThep1p3Edge(t *testing.T) {
	triangle := MakeTriangle(Point(0, 1, 0), Point(-1, 0, 0), Point(1, 0, 0), Identity(), DefaultMaterial())
	r := Ray{Point(1, 1, -2), Vector(0, 0, 1)}

	Assert(len(triangle.Intersection(r)) == 0, "Intersection ha d to be empty", t)
}

func TestARayMissesThep1p2Edge(t *testing.T) {
	triangle := MakeTriangle(Point(0, 1, 0), Point(-1, 0, 0), Point(1, 0, 0), Identity(), DefaultMaterial())
	r := Ray{Point(-1, 1, -2), Vector(0, 0, 1)}

	Assert(len(triangle.Intersection(r)) == 0, "Intersection ha d to be empty", t)
}

func TestARayMissesThep2p3Edge(t *testing.T) {
	triangle := MakeTriangle(Point(0, 1, 0), Point(-1, 0, 0), Point(1, 0, 0), Identity(), DefaultMaterial())
	r := Ray{Point(0, -1, -2), Vector(0, 0, 1)}

	Assert(len(triangle.Intersection(r)) == 0, "Intersection ha d to be empty", t)
}

func TestARayStrikesTheTriangle(t *testing.T) {
	triangle := MakeTriangle(Point(0, 1, 0), Point(-1, 0, 0), Point(1, 0, 0), Identity(), DefaultMaterial())
	r := Ray{Point(0, 0.5, -2), Vector(0, 0, 1)}

	AssertEqual(2, triangle.Intersection(r)[0].t, t)
}

func TestTriangleMaterial(t *testing.T) {
	m := DefaultMaterial()
	triangle := MakeTriangle(Point(0, 0, 0), Point(1, 2, 3), Point(4, 5, 6), Identity(), m)

	Assert(m == triangle.Material(), "", t)
}

func TestIntersectionStoresUandV(t *testing.T) {
	p1 := Point(0, 1, 0)
	p2 := Point(-1, 0, 0)
	p3 := Point(1, 0, 0)

	tri := MakeTriangle(p1, p2, p3, Identity(), DefaultMaterial())
	r := Ray{Point(-0.2, 0.3, -2), Vector(0, 0, 1)}

	xs := tri.Intersection(r)

	AssertEqual(0.45, xs[0].u, t)
	AssertEqual(0.25, xs[0].v, t)
}
