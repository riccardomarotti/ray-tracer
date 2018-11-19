package main

import (
	"reflect"
	"testing"
)

func TestGroupWithAChild(t *testing.T) {
	s := MakeSphere(Identity(), DefaultMaterial())
	g := MakeGroup(Identity(), s)

	Assert(reflect.DeepEqual(s, g.children[0]), "", t)
}

func TestIntersectingARayWithAnEmptyGroup(t *testing.T) {
	g := MakeGroup(Identity())
	r := Ray{Point(0, 0, 0), Vector(0, 0, 1)}

	xs := g.Intersection(r)

	Assert(len(xs) == 0, "", t)
}

func TestIntersectingARayWithANonEmptyGroup(t *testing.T) {
	s1 := MakeSphere(Identity(), DefaultMaterial())
	s2 := MakeSphere(Identity().Translate(0, 0, -3), DefaultMaterial())
	s3 := MakeSphere(Identity().Translate(5, 0, 0), DefaultMaterial())

	g := MakeGroup(Identity(), s1, s2, s3)
	r := Ray{Point(0, 0, -5), Vector(0, 0, 1)}

	xs := g.Intersection(r)

	Assert(len(xs) == 4, "", t)
	Assert(reflect.DeepEqual(s2, xs[0].object), "", t)
	Assert(reflect.DeepEqual(s2, xs[1].object), "", t)
	Assert(reflect.DeepEqual(s1, xs[2].object), "", t)
	Assert(reflect.DeepEqual(s1, xs[3].object), "", t)
}

func TestIntersectingATransformedGroup(t *testing.T) {
	s := MakeSphere(Identity().Translate(5, 0, 0), DefaultMaterial())
	g := MakeGroup(Identity().Scale(2, 2, 2), s)
	r := Ray{Point(10, 0, -10), Vector(0, 0, 1)}

	xs := g.Intersection(r)

	Assert(len(xs) == 2, "", t)
}
