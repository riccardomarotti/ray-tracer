package main

import "testing"

func TestSphereTransformation(t *testing.T) {
	T := Identity().Translate(2, 3, 4)
	s := MakeSphere(T)

	AssertMatrixEqual(T, s.Transform(), t)
}

func TestIntersectingAScaledSphereWithARay(t *testing.T) {
	r := Ray{Point(0, 0, -5), Vector(0, 0, 1)}
	s := MakeSphere(Identity().Scale(2, 2, 2))

	xs := r.Intersection(s)

	Assert(len(xs) == 2, "", t)
	AssertEqual(3, xs[0].t, t)
	AssertEqual(7, xs[1].t, t)
}

func TestIntersectingATranslatedSphereWithARay(t *testing.T) {
	r := Ray{Point(0, 0, -5), Vector(0, 0, 1)}
	s := MakeSphere(Identity().Translate(5, 0, 0))

	xs := r.Intersection(s)

	Assert(len(xs) == 0, "", t)
}
