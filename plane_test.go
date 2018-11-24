package main

import (
	"reflect"
	"testing"
)

func TestTheNormalOfAPlaneIsConstantEverywhere(t *testing.T) {
	p := MakePlane(Identity(), DefaultMaterial())
	n1 := p.NormalAt(Point(0, 0, 0), Intersection{})
	n2 := p.NormalAt(Point(10, 0, -10), Intersection{})
	n3 := p.NormalAt(Point(-5, 0, 150), Intersection{})

	AssertTupleEqual(Vector(0, 1, 0), n1, t)
	AssertTupleEqual(Vector(0, 1, 0), n2, t)
	AssertTupleEqual(Vector(0, 1, 0), n3, t)
}

func TestIntersectWithARayParallelToThePlane(t *testing.T) {
	p := MakePlane(Identity(), DefaultMaterial())
	r := Ray{Point(0, 10, 0), Vector(0, 0, 1)}

	xs := p.Intersection(r)
	Assert(len(xs) == 0, "", t)
}

func TestIntersectWithACoplanarRay(t *testing.T) {
	p := MakePlane(Identity(), DefaultMaterial())
	r := Ray{Point(0, 0, 0), Vector(0, 0, 1)}

	xs := p.Intersection(r)
	Assert(len(xs) == 0, "", t)
}

func TestARayIntersectingAPlaneFromAbove(t *testing.T) {
	p := MakePlane(Identity(), DefaultMaterial())
	r := Ray{Point(0, 1, 0), Vector(0, -1, 0)}

	xs := p.Intersection(r)
	Assert(len(xs) == 1, "", t)
	AssertEqual(1, xs[0].t, t)
	Assert(reflect.DeepEqual(p, xs[0].object), "", t)
}

func TestARayIntersectingAPlaneFromBelow(t *testing.T) {
	p := MakePlane(Identity(), DefaultMaterial())
	r := Ray{Point(0, -1, 0), Vector(0, 1, 0)}

	xs := p.Intersection(r)
	Assert(len(xs) == 1, "", t)
	AssertEqual(1, xs[0].t, t)
	Assert(reflect.DeepEqual(p, xs[0].object), "", t)
}
