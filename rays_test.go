package main

import "testing"

func TestCreatingAndQueryingARay(t *testing.T) {
	origin := Point(1, 2, 3)
	direction := Vector(4, 5, 6)

	r := Ray{origin, direction}

	AssertTupleEqual(origin, r.origin, t)
	AssertTupleEqual(direction, r.direction, t)
}

func TestComputeAPointFromADistanceAlongTheRay(t *testing.T) {
	r := Ray{Point(2, 3, 4), Vector(1, 0, 0)}

	AssertTupleEqual(Point(2, 3, 4), r.Position(0), t)
	AssertTupleEqual(Point(3, 3, 4), r.Position(1), t)
	AssertTupleEqual(Point(1, 3, 4), r.Position(-1), t)
	AssertTupleEqual(Point(4.5, 3, 4), r.Position(2.5), t)
}
