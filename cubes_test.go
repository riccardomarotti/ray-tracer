package main

import "testing"

func TestRayIntersectsCubeOnXPositiveAxis(t *testing.T) {
	c := MakeCube(Identity(), DefaultMaterial())
	r := Ray{Point(5, 0.5, 0), Vector(-1, 0, 0)}
	xs := c.Intersection(r)

	Assert(len(xs) == 2, "Ther had to be two intersections", t)
	AssertEqual(4, xs[0].t, t)
	AssertEqual(6, xs[1].t, t)
}

func TestRayIntersectsCubeOnXNegativeAxis(t *testing.T) {
	c := MakeCube(Identity(), DefaultMaterial())
	r := Ray{Point(-5, 0.5, 0), Vector(1, 0, 0)}
	xs := c.Intersection(r)

	Assert(len(xs) == 2, "Ther had to be two intersections", t)
	AssertEqual(4, xs[0].t, t)
	AssertEqual(6, xs[1].t, t)
}

func TestRayIntersectsCubeOnYPositiveAxis(t *testing.T) {
	c := MakeCube(Identity(), DefaultMaterial())
	r := Ray{Point(0.5, 5, 0), Vector(0, -1, 0)}
	xs := c.Intersection(r)

	Assert(len(xs) == 2, "Ther had to be two intersections", t)
	AssertEqual(4, xs[0].t, t)
	AssertEqual(6, xs[1].t, t)
}

func TestRayIntersectsCubeOnYNegativeAxis(t *testing.T) {
	c := MakeCube(Identity(), DefaultMaterial())
	r := Ray{Point(0.5, -5, 0), Vector(0, 1, 0)}
	xs := c.Intersection(r)

	Assert(len(xs) == 2, "Ther had to be two intersections", t)
	AssertEqual(4, xs[0].t, t)
	AssertEqual(6, xs[1].t, t)
}

func TestRayIntersectsCubeOnZPositiveAxis(t *testing.T) {
	c := MakeCube(Identity(), DefaultMaterial())
	r := Ray{Point(0.5, 0, 5), Vector(0, 0, -1)}
	xs := c.Intersection(r)

	Assert(len(xs) == 2, "Ther had to be two intersections", t)
	AssertEqual(4, xs[0].t, t)
	AssertEqual(6, xs[1].t, t)
}

func TestRayIntersectsCubeOnZNegativeAxis(t *testing.T) {
	c := MakeCube(Identity(), DefaultMaterial())
	r := Ray{Point(0.5, 0, -5), Vector(0, 0, 1)}
	xs := c.Intersection(r)

	Assert(len(xs) == 2, "Ther had to be two intersections", t)
	AssertEqual(4, xs[0].t, t)
	AssertEqual(6, xs[1].t, t)
}

func TestRayIntersectsCubeOnInside(t *testing.T) {
	c := MakeCube(Identity(), DefaultMaterial())
	r := Ray{Point(0, 0.5, 0), Vector(0, 0, 1)}
	xs := c.Intersection(r)

	Assert(len(xs) == 2, "Ther had to be two intersections", t)
	AssertEqual(-1, xs[0].t, t)
	AssertEqual(1, xs[1].t, t)
}
