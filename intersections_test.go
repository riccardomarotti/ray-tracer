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

	xs := []Intersection{Intersection{t: 1, object: s}, Intersection{t: 2, object: s}}

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

func TestPrecomputingTheStateOfAnIntersection(t *testing.T) {
	ray := Ray{Point(0, 0, -5), Vector(0, 0, 1)}
	shape := MakeSphere(Identity(), DefaultMaterial())

	hit := Hit(ray.Intersection(shape))

	hitData := PrepareHit(hit, ray)

	AssertTupleEqual(Point(0, 0, -1), hitData.point, t)
	AssertTupleEqual(Vector(0, 0, -1), hitData.eyeVector, t)
	AssertTupleEqual(Vector(0, 0, -1), hitData.normalVector, t)
}

func TestIntersectionOutside(t *testing.T) {
	ray := Ray{Point(0, 0, -5), Vector(0, 0, 1)}
	shape := MakeSphere(Identity(), DefaultMaterial())

	hit := Hit(ray.Intersection(shape))

	hitData := PrepareHit(hit, ray)

	Assert(hitData.inside == false, "", t)
}

func TestIntersectionInside(t *testing.T) {
	ray := Ray{Point(0, 0, 0), Vector(0, 0, 1)}
	shape := MakeSphere(Identity(), DefaultMaterial())

	hit := Hit(ray.Intersection(shape))

	hitData := PrepareHit(hit, ray)

	AssertTupleEqual(Point(0, 0, 1), hitData.point, t)
	AssertTupleEqual(Vector(0, 0, -1), hitData.eyeVector, t)

	Assert(hitData.inside == true, "Hit ha dto be inside", t)
	AssertTupleEqual(Vector(0, 0, -1), hitData.normalVector, t)
}

func TestShadingAnIntersection(t *testing.T) {
	world := DefaultWorld()
	ray := Ray{Point(0, 0, -5), Vector(0, 0, 1)}
	shape := world.objects[0]

	hit := Intersection{t: 4, object: shape}
	hit = PrepareHit(hit, ray)

	c := hit.Shade(world)

	AssertColorEqual(Color{0.38066, 0.47583, 0.2855}, c, t)
}

func TestShadingAnIntersectionFromTheInside(t *testing.T) {
	world := DefaultWorld()
	world.light = PointLight{Point(0, 0.25, 0), Color{1, 1, 1}}
	ray := Ray{Point(0, 0, 0), Vector(0, 0, 1)}
	shape := world.objects[1]

	hit := Hit(ray.Intersection(shape))
	hit = PrepareHit(hit, ray)

	c := hit.Shade(world)

	AssertColorEqual(Color{.1, .1, .1}, c, t)
}

func TestShadeIsGivenAnIntersectionInShadow(t *testing.T) {
	light := PointLight{Point(0, 0, -10), Color{1, 1, 1}}
	s1 := MakeSphere(Identity(), DefaultMaterial())
	s2 := MakeSphere(Identity().Translate(0, 0, 10), DefaultMaterial())
	world := World{light, []Object{s1, s2}}

	r := Ray{Point(0, 0, 5), Vector(0, 0, 1)}
	hit := Hit(r.Intersection(s2))
	hit = PrepareHit(hit, r)

	c := hit.Shade(world)

	AssertColorEqual(Color{0.1, 0.1, 0.1}, c, t)
}

func TestThePointIdOffset(t *testing.T) {
	ray := Ray{Point(0, 0, -5), Vector(0, 0, 1)}
	shape := MakeSphere(Identity(), DefaultMaterial())

	hit := Hit(ray.Intersection(shape))
	hit = PrepareHit(hit, ray)

	Assert(hit.point.z > -1.1 && hit.point.z < -1, "", t)

}
