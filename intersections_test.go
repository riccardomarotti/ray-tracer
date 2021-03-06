package main

import (
	"reflect"
	"testing"
)

func TestIntersecionEncapsulatesTAndSolid(t *testing.T) {
	s := MakeSphere(Identity(), DefaultMaterial())
	i := Intersection{t: 3.5, object: s}

	AssertEqual(i.t, 3.5, t)

	Assert(reflect.DeepEqual(i.object, s), "", t)
}

func TestAggregatingIntersections(t *testing.T) {
	s := MakeSphere(Identity(), DefaultMaterial())

	xs := []Intersection{{t: 1, object: s}, {t: 2, object: s}}

	Assert(len(xs) == 2, "", t)
	AssertEqual(1, xs[0].t, t)
	AssertEqual(2, xs[1].t, t)
}

func TestHitWhenAllIntersectionsHavePositiveT(t *testing.T) {
	s := MakeSphere(Identity(), DefaultMaterial())
	i1 := Intersection{t: 1, object: s}
	i2 := Intersection{t: 2, object: s}
	xs := []Intersection{i2, i1}

	Assert(reflect.DeepEqual(i1, Hit(xs)), "", t)
}

func TestHitWhenSomeIntersectionsHaveNegativeT(t *testing.T) {
	s := MakeSphere(Identity(), DefaultMaterial())
	i1 := Intersection{t: -1, object: s}
	i2 := Intersection{t: 1, object: s}
	xs := []Intersection{i2, i1}

	Assert(reflect.DeepEqual(i2, Hit(xs)), "", t)
}

func TestHitWhenAllIntersectionsHaveNegativeT(t *testing.T) {
	s := MakeSphere(Identity(), DefaultMaterial())
	i1 := Intersection{t: -1, object: s}
	i2 := Intersection{t: -2, object: s}
	xs := []Intersection{i2, i1}

	Assert(Intersection{} == Hit(xs), "", t)
}

func TestHitIsAlwaysTheLowestNonNegativeIntersection(t *testing.T) {
	s := MakeSphere(Identity(), DefaultMaterial())
	i1 := Intersection{t: 5, object: s}
	i2 := Intersection{t: 7, object: s}
	i3 := Intersection{t: -3, object: s}
	i4 := Intersection{t: 2, object: s}
	xs := []Intersection{i1, i2, i3, i4}

	Assert(reflect.DeepEqual(i4, Hit(xs)), "", t)
}
