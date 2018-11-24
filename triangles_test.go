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
	n1 := triangle.NormalAt(Point(0, 0.5, 0), Intersection{})
	n2 := triangle.NormalAt(Point(-0.5, 0.75, 0), Intersection{})
	n3 := triangle.NormalAt(Point(0.5, 0.25, 0), Intersection{})

	AssertTupleEqual(triangle.normal, n1, t)
	AssertTupleEqual(triangle.normal, n2, t)
	AssertTupleEqual(triangle.normal, n3, t)

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

func TestTriangleUsesUandVToInterpolateNormals(t *testing.T) {
	p1 := Point(0, 1, 0)
	p2 := Point(-1, 0, 0)
	p3 := Point(1, 0, 0)
	n1 := Vector(0, 1, 0)
	n2 := Vector(-1, 0, 0)
	n3 := Vector(1, 0, 0)

	tri := MakeSmoothTriangle(p1, p2, p3, n1, n2, n3, Identity(), DefaultMaterial(), nil)
	r := Ray{Point(-0.2, 0.3, -2), Vector(0, 0, 1)}

	i := tri.Intersection(r)[0]

	n := tri.NormalAt(Point(0, 0, 0), i)

	AssertTupleEqual(Vector(-0.5547, 0.83205, 0), n, t)
}
