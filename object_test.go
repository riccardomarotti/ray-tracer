package main

import "testing"

func TestIntersecionEncapsulatesTAndSolid(t *testing.T) {
	s := Sphere{}
	i := Intersection{3.5, s}

	AssertEqual(i.t, 3.5, t)
	Assert(i.object == s, "", t)
}

func TestAggregatingIntersections(t *testing.T) {
	s := Sphere{}

	xs := []Intersection{Intersection{1, s}, Intersection{2, s}}

	Assert(len(xs) == 2, "", t)
	AssertEqual(1, xs[0].t, t)
	AssertEqual(2, xs[1].t, t)
}
