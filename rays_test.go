package main

import (
	"reflect"
	"testing"
)

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

func TestRayIntersectsASphereAtTwoPoints(t *testing.T) {
	r := Ray{Point(0, 0, -5), Vector(0, 0, 1)}
	s := MakeSphere(Identity(), DefaultMaterial())

	xs := r.Intersection(s)

	Assert(2 == len(xs), "", t)
	AssertEqual(4, xs[0].t, t)
	AssertEqual(6, xs[1].t, t)
}

func TestRayIntersectsASphereAtATangent(t *testing.T) {
	r := Ray{Point(0, 1, -5), Vector(0, 0, 1)}
	s := MakeSphere(Identity(), DefaultMaterial())

	xs := r.Intersection(s)

	Assert(2 == len(xs), "", t)
	AssertEqual(5, xs[0].t, t)
	AssertEqual(5, xs[1].t, t)
}

func TestRayMissesASphere(t *testing.T) {
	r := Ray{Point(0, 2, -5), Vector(0, 0, 1)}
	s := MakeSphere(Identity(), DefaultMaterial())

	xs := r.Intersection(s)

	Assert(0 == len(xs), "Length of intersection had to be zero", t)
}

func TestRayOriginatesInsideASphere(t *testing.T) {
	r := Ray{Point(0, 0, 0), Vector(0, 0, 1)}
	s := MakeSphere(Identity(), DefaultMaterial())
	xs := r.Intersection(s)

	Assert(2 == len(xs), "", t)
	AssertEqual(-1, xs[0].t, t)
	AssertEqual(1, xs[1].t, t)
}

func TestSphereBehindARay(t *testing.T) {
	r := Ray{Point(0, 0, 5), Vector(0, 0, 1)}
	s := MakeSphere(Identity(), DefaultMaterial())
	xs := r.Intersection(s)

	Assert(2 == len(xs), "", t)
	AssertEqual(-6, xs[0].t, t)
	AssertEqual(-4, xs[1].t, t)
}

func TestIntersecSetsTheObjectOnTheIntersection(t *testing.T) {
	r := Ray{Point(0, 0, -5), Vector(0, 0, 1)}
	s := MakeSphere(Identity(), DefaultMaterial())

	xs := r.Intersection(s)

	Assert(2 == len(xs), "", t)
	Assert(reflect.DeepEqual(s, xs[0].object), "", t)
	Assert(reflect.DeepEqual(s, xs[1].object), "", t)
}

func TestTranslateRay(t *testing.T) {
	r := Ray{Point(1, 2, 3), Vector(0, 1, 0)}

	r2 := r.Transform(Identity().Translate(3, 4, 5))

	AssertTupleEqual(Point(4, 6, 8), r2.origin, t)
	AssertTupleEqual(Vector(0, 1, 0), r2.direction, t)

}

func TestScaleRay(t *testing.T) {
	r := Ray{Point(1, 2, 3), Vector(0, 1, 0)}

	r2 := r.Transform(Identity().Scale(2, 3, 4))

	AssertTupleEqual(Point(2, 6, 12), r2.origin, t)
	AssertTupleEqual(Vector(0, 3, 0), r2.direction, t)

}
