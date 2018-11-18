package main

import (
	"fmt"
	"testing"
)

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

func TestRayMissesTheCube(t *testing.T) {
	c := MakeCube(Identity(), DefaultMaterial())

	examples := map[int][2]Tuple{
		0: {Point(-2, 0, 0), Vector(0.2673, 0.5345, 0.8018)},
		1: {Point(0, -2, 0), Vector(0.8018, 0.2673, 0.5345)},
		2: {Point(0, 0, -2), Vector(0.5345, 0.8018, 0.2673)},
		3: {Point(2, 0, 2), Vector(0, 0, -1)},
		4: {Point(0, 2, 2), Vector(0, -1, 0)},
		5: {Point(2, 2, 0), Vector(-1, 0, 0)},
	}

	for i := 0; i < len(examples); i++ {
		r := Ray{examples[i][0], examples[i][1]}
		xs := c.Intersection(r)

		Assert(len(xs) == 0, fmt.Sprintf("Cube intersections had to be zero for ray %v", r), t)
	}
}

func TestNormalOnACube(t *testing.T) {
	c := MakeCube(Identity(), DefaultMaterial())

	examples := map[int][2]Tuple{
		0: {Point(1, 0.5, -0.8), Vector(1, 0, 0)},
		1: {Point(-1, -0.2, 0.9), Vector(-1, 0, 0)},
		2: {Point(-0.4, 1, -0.1), Vector(0, 1, 0)},
		3: {Point(0.3, -1, -0.7), Vector(0, -1, 0)},
		4: {Point(-0.6, 0.3, 1), Vector(0, 0, 1)},
		5: {Point(0.4, 0.4, -1), Vector(0, 0, -1)},
		6: {Point(1, 1, 1), Vector(1, 0, 0)},
		7: {Point(-1, -1, -1), Vector(-1, 0, 0)},
	}

	for i := 0; i < len(examples); i++ {
		p := examples[i][0]

		AssertTupleEqual(examples[i][1], c.NormalAt(p), t)
	}
}
