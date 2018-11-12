package main

import (
	"math"
	"testing"
)

func TestSphereTransformation(t *testing.T) {
	T := Identity().Translate(2, 3, 4)
	s := MakeSphere(T, DefaultMaterial())

	AssertMatrixEqual(T, s.Transform(), t)
}

func TestIntersectingAScaledSphereWithARay(t *testing.T) {
	r := Ray{Point(0, 0, -5), Vector(0, 0, 1)}
	s := MakeSphere(Identity().Scale(2, 2, 2), DefaultMaterial())

	xs := s.Intersection(r)

	Assert(len(xs) == 2, "", t)
	AssertEqual(3, xs[0].t, t)
	AssertEqual(7, xs[1].t, t)
}

func TestIntersectingATranslatedSphereWithARay(t *testing.T) {
	r := Ray{Point(0, 0, -5), Vector(0, 0, 1)}
	s := MakeSphere(Identity().Translate(5, 0, 0), DefaultMaterial())

	xs := s.Intersection(r)

	Assert(len(xs) == 0, "", t)
}

func TestNormalOnSphereAtAPointOnTheXAxis(t *testing.T) {
	s := MakeSphere(Identity(), DefaultMaterial())
	n := s.NormalAt(Point(1, 0, 0))

	AssertTupleEqual(Vector(1, 0, 0), n, t)
}

func TestNormalOnSphereAtAPointOnTheYAxis(t *testing.T) {
	s := MakeSphere(Identity(), DefaultMaterial())
	n := s.NormalAt(Point(0, 1, 0))

	AssertTupleEqual(Vector(0, 1, 0), n, t)
}

func TestNormalOnSphereAtAPointOnTheZAxis(t *testing.T) {
	s := MakeSphere(Identity(), DefaultMaterial())
	n := s.NormalAt(Point(0, 0, 1))

	AssertTupleEqual(Vector(0, 0, 1), n, t)
}

func TestNormalOnSphereAtANonAxialPoint(t *testing.T) {
	s := MakeSphere(Identity(), DefaultMaterial())
	xyz := math.Sqrt(3) / 3

	n := s.NormalAt(Point(xyz, xyz, xyz))

	AssertTupleEqual(Vector(xyz, xyz, xyz), n, t)
}

func TestNormalIsANormalizedVector(t *testing.T) {
	s := MakeSphere(Identity(), DefaultMaterial())
	xyz := math.Sqrt(3) / 3

	n := s.NormalAt(Point(xyz, xyz, xyz))

	AssertTupleEqual(n.Normalize(), n, t)
}

func TestComputingTheNormalOnATranslatedSphere(t *testing.T) {
	s := MakeSphere(Identity().Translate(0, 1, 0), DefaultMaterial())

	n := s.NormalAt(Point(0, 1.70711, -0.70711))

	AssertTupleEqual(Vector(0, 0.70711, -0.70711), n, t)
}

func TestComputingTheNormalOnAScaledSphere(t *testing.T) {
	s := MakeSphere(Identity().Scale(1, 0.5, 1), DefaultMaterial())

	n := s.NormalAt(Point(0, math.Sqrt(2)/2, -math.Sqrt(2)/2))

	AssertTupleEqual(Vector(0, 0.97014, -0.24254), n, t)
}

func TestSphereHasMaterial(t *testing.T) {
	m := DefaultMaterial()
	s := MakeSphere(Identity(), m)

	Assert(m == s.Material(), "", t)

}
