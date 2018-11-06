package main

import "testing"

func TestCreatingAndQueryingARay(t *testing.T) {
	origin := Point(1, 2, 3)
	direction := Vector(4, 5, 6)

	r := Ray{origin, direction}

	AssertTupleEqual(origin, r.origin, t)
	AssertTupleEqual(direction, r.direction, t)
}
