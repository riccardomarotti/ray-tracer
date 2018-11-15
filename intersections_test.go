package main

import (
	"math"
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

	intersections := shape.Intersection(ray)
	hit := Hit(intersections)

	hitData := PrepareComputations(hit, ray, intersections)

	AssertTupleEqual(Point(0, 0, -1), hitData.point, t)
	AssertTupleEqual(Vector(0, 0, -1), hitData.eyeVector, t)
	AssertTupleEqual(Vector(0, 0, -1), hitData.normalVector, t)
}

func TestIntersectionOutside(t *testing.T) {
	ray := Ray{Point(0, 0, -5), Vector(0, 0, 1)}
	shape := MakeSphere(Identity(), DefaultMaterial())

	intersections := shape.Intersection(ray)
	hit := Hit(intersections)

	hitData := PrepareComputations(hit, ray, intersections)

	Assert(hitData.inside == false, "", t)
}

func TestIntersectionInside(t *testing.T) {
	ray := Ray{Point(0, 0, 0), Vector(0, 0, 1)}
	shape := MakeSphere(Identity(), DefaultMaterial())

	intersections := shape.Intersection(ray)
	hit := Hit(intersections)

	hitData := PrepareComputations(hit, ray, intersections)

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
	hit = PrepareComputations(hit, ray, nil)

	c := hit.Shade(world)

	AssertColorEqual(Color{0.38066, 0.47583, 0.2855}, c, t)
}

func TestShadingAnIntersectionFromTheInside(t *testing.T) {
	world := DefaultWorld()
	world.light = PointLight{Point(0, 0.25, 0), Color{1, 1, 1}}
	ray := Ray{Point(0, 0, 0), Vector(0, 0, 1)}
	shape := world.objects[1]

	intersections := shape.Intersection(ray)
	hit := Hit(intersections)
	hit = PrepareComputations(hit, ray, intersections)

	c := hit.Shade(world)

	AssertColorEqual(Color{.1, .1, .1}, c, t)
}

func TestShadeIsGivenAnIntersectionInShadow(t *testing.T) {
	light := PointLight{Point(0, 0, -10), Color{1, 1, 1}}
	s1 := MakeSphere(Identity(), DefaultMaterial())
	s2 := MakeSphere(Identity().Translate(0, 0, 10), DefaultMaterial())
	world := World{light, []Object{s1, s2}}

	r := Ray{Point(0, 0, 5), Vector(0, 0, 1)}
	intersections := s2.Intersection(r)
	hit := Hit(intersections)
	hit = PrepareComputations(hit, r, intersections)

	c := hit.Shade(world)

	AssertColorEqual(Color{0.1, 0.1, 0.1}, c, t)
}

func TestThePointIdOffset(t *testing.T) {
	ray := Ray{Point(0, 0, -5), Vector(0, 0, 1)}
	shape := MakeSphere(Identity(), DefaultMaterial())

	intersections := shape.Intersection(ray)
	hit := Hit(intersections)
	hit = PrepareComputations(hit, ray, intersections)

	Assert(hit.point.z > -1.1 && hit.point.z < -1, "", t)

}
func TestPrecomputingTheReflectionVector(t *testing.T) {
	shape := MakePlane(Identity(), DefaultMaterial())
	ray := Ray{Point(0, 1, -1), Vector(0, -math.Sqrt(2)/2, math.Sqrt(2)/2)}
	hit := Intersection{t: math.Sqrt(2), object: shape}

	hit = PrepareComputations(hit, ray, nil)

	AssertTupleEqual(Vector(0, math.Sqrt(2)/2, math.Sqrt(2)/2), hit.reflectVector, t)
}

func TestShadeWithReflectiveMaterial(t *testing.T) {
	world := WorldWithAmbientSetTo(.47013)
	reflectiveMaterial := DefaultMaterial()
	reflectiveMaterial.reflective = 1
	plane := MakePlane(Identity().Translate(0, -1, 0), reflectiveMaterial)
	world.objects = append(world.objects, plane)

	ray := Ray{Point(0, 0, -3), Vector(0, -math.Sqrt(2)/2, math.Sqrt(2)/2)}
	hit := Intersection{t: math.Sqrt(2), object: plane}
	hit = PrepareComputations(hit, ray, nil)

	AssertColorEqual(Color{0.87677, 0.92436, 0.82918}, hit.Shade(world), t)
}

func TestN1AndN2AtVariousIntersections(t *testing.T) {
	A := MakeGlassSphere(Identity().Scale(2, 2, 2), 1.5)
	B := MakeGlassSphere(Identity().Translate(0, 0, -0.25), 2.0)
	C := MakeGlassSphere(Identity().Translate(0, 0, 0.25), 2.5)

	ray := Ray{Point(0, 0, -4), Vector(0, 0, 1)}
	xs := []Intersection{
		Intersection{t: 2, object: A},
		Intersection{t: 2.75, object: B},
		Intersection{t: 3.25, object: C},
		Intersection{t: 4.75, object: B},
		Intersection{t: 5.25, object: C},
		Intersection{t: 6, object: A},
	}

	examples := map[int][]float64{
		0: []float64{1.0, 1.5},
		1: []float64{1.5, 2.0},
		2: []float64{2.0, 2.5},
		3: []float64{2.5, 2.5},
		4: []float64{2.5, 1.5},
		5: []float64{1.5, 1.0},
	}

	for index := 0; index < len(xs); index++ {
		hit := PrepareComputations(xs[index], ray, xs)

		AssertEqual(examples[index][0], hit.n1, t)
		AssertEqual(examples[index][1], hit.n2, t)
	}
}

func TestTheUnderPointIsOffsetBelowTheSurface(t *testing.T) {
	r := Ray{Point(0, 0, -5), Vector(0, 0, 1)}
	shape := MakeGlassSphere(Identity().Translate(0, 0, 1), 1.5)

	i := Intersection{t: 5, object: shape}
	xs := []Intersection{i}

	hitData := PrepareComputations(i, r, xs)

	Assert(hitData.underPoint.z > Epsilon/2, "", t)
}
