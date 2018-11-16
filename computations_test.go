package main

import (
	"math"
	"testing"
)

func TestShadingAnIntersection(t *testing.T) {
	world := DefaultWorld()
	ray := Ray{Point(0, 0, -5), Vector(0, 0, 1)}
	shape := world.objects[0]

	hit := Intersection{t: 4, object: shape}
	comps := PrepareComputations(hit, ray, nil)

	c := comps.Shade(world, 5)

	AssertColorEqual(Color{0.38066, 0.47583, 0.2855}, c, t)
}

func TestShadingAnIntersectionFromTheInside(t *testing.T) {
	world := DefaultWorld()
	world.light = PointLight{Point(0, 0.25, 0), Color{1, 1, 1}}
	ray := Ray{Point(0, 0, 0), Vector(0, 0, 1)}
	shape := world.objects[1]

	intersections := shape.Intersection(ray)
	hit := Hit(intersections)
	comps := PrepareComputations(hit, ray, intersections)

	c := comps.Shade(world, 5)

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
	comps := PrepareComputations(hit, r, intersections)

	c := comps.Shade(world, 5)

	AssertColorEqual(Color{0.1, 0.1, 0.1}, c, t)
}

func TestThePointIdOffset(t *testing.T) {
	ray := Ray{Point(0, 0, -5), Vector(0, 0, 1)}
	shape := MakeSphere(Identity(), DefaultMaterial())

	intersections := shape.Intersection(ray)
	hit := Hit(intersections)
	comps := PrepareComputations(hit, ray, intersections)

	Assert(comps.point.z > -1.1 && comps.point.z < -1, "", t)

}
func TestPrecomputingTheReflectionVector(t *testing.T) {
	shape := MakePlane(Identity(), DefaultMaterial())
	ray := Ray{Point(0, 1, -1), Vector(0, -math.Sqrt(2)/2, math.Sqrt(2)/2)}
	hit := Intersection{t: math.Sqrt(2), object: shape}

	comps := PrepareComputations(hit, ray, nil)

	AssertTupleEqual(Vector(0, math.Sqrt(2)/2, math.Sqrt(2)/2), comps.reflectVector, t)
}

func TestShadeWithReflectiveMaterial(t *testing.T) {
	world := DefaultWorld()
	reflectiveMaterial := DefaultMaterial()
	reflectiveMaterial.reflective = .5
	plane := MakePlane(Identity().Translate(0, -1, 0), reflectiveMaterial)
	world.objects = append(world.objects, plane)

	ray := Ray{Point(0, 0, -3), Vector(0, -math.Sqrt(2)/2, math.Sqrt(2)/2)}
	hit := Intersection{t: math.Sqrt(2), object: plane}
	comps := PrepareComputations(hit, ray, nil)

	AssertColorEqual(Color{0.87677, 0.92436, 0.82918}, comps.Shade(world, 5), t)
}

func TestN1AndN2AtVariousIntersections(t *testing.T) {
	A := MakeGlassSphere(Identity().Scale(2, 2, 2), 1.5)
	B := MakeGlassSphere(Identity().Translate(0, 0, -0.25), 2.0)
	C := MakeGlassSphere(Identity().Translate(0, 0, 0.25), 2.5)

	ray := Ray{Point(0, 0, -4), Vector(0, 0, 1)}
	xs := []Intersection{
		{t: 2, object: A},
		{t: 2.75, object: B},
		{t: 3.25, object: C},
		{t: 4.75, object: B},
		{t: 5.25, object: C},
		{t: 6, object: A},
	}

	examples := map[int][]float64{
		0: {1.0, 1.5},
		1: {1.5, 2.0},
		2: {2.0, 2.5},
		3: {2.5, 2.5},
		4: {2.5, 1.5},
		5: {1.5, 1.0},
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

func TestShadeWithATransparentMaterial(t *testing.T) {
	floorMaterial := DefaultMaterial()
	floorMaterial.transparency = 0.5
	floorMaterial.refractiveIndex = 1.5
	floor := MakePlane(Identity().Translate(0, -1, 0), floorMaterial)

	ballMaterial := DefaultMaterial()
	ballMaterial.color = Color{1, 0, 0}
	ballMaterial.ambient = 0.5
	ballMaterial.transparency = 0.1
	ball := MakeSphere(Identity().Translate(0, -3.5, -0.5), ballMaterial)

	w := DefaultWorld()
	w.objects = append(w.objects, []Object{floor, ball}...)

	r := Ray{Point(0, 0, -3), Vector(0, -math.Sqrt(2)/2, math.Sqrt(2)/2)}
	xs := []Intersection{{t: math.Sqrt(2), object: floor}}

	intersection := PrepareComputations(xs[0], r, xs)

	AssertColorEqual(Color{0.93642, 0.68642, 0.68642}, intersection.Shade(w, 5), t)
}

func TestSchlickApproximationUnderTotalInternalReflection(t *testing.T) {
	shape := MakeGlassSphere(Identity(), 1.5)
	r := Ray{Point(0, 0, math.Sqrt(2)/2), Vector(0, 1, 0)}
	xs := []Intersection{{t: -math.Sqrt(2) / 2, object: shape}, {t: math.Sqrt(2) / 2, object: shape}}

	hitData := PrepareComputations(xs[1], r, xs)

	AssertEqual(1, hitData.Schlick(), t)
}

func TestSchlickApproximationWithAPerpendicularViewingAngle(t *testing.T) {
	shape := MakeGlassSphere(Identity(), 1.5)
	r := Ray{Point(0, 0, 0), Vector(0, 1, 0)}
	xs := []Intersection{{t: -1, object: shape}, {t: 1, object: shape}}

	hitData := PrepareComputations(xs[1], r, xs)

	AssertEqual(.04, hitData.Schlick(), t)
}

func TestSchlickApproximationWithSmallAncgleAndN2GreaterThanN1(t *testing.T) {
	shape := MakeGlassSphere(Identity(), 1.5)
	r := Ray{Point(0, 0.99, -2), Vector(0, 0, 1)}
	xs := []Intersection{{t: 1.8589, object: shape}}

	hitData := PrepareComputations(xs[0], r, xs)

	AssertEqual(.48873, hitData.Schlick(), t)
}

func TestShadeWithReflectiveAndTransparentMaterial(t *testing.T) {
	w := DefaultWorld()
	r := Ray{Point(0, 0, -3), Vector(0, -math.Sqrt(2)/2, math.Sqrt(2)/2)}

	floorMaterial := DefaultMaterial()
	floorMaterial.reflective = 0.5
	floorMaterial.transparency = 0.5
	floorMaterial.ambient = 0
	floorMaterial.refractiveIndex = 1.5
	floor := MakePlane(Identity().Translate(0, -1, 0), floorMaterial)

	ballMaterial := DefaultMaterial()
	ballMaterial.color = Color{1, 0, 0}
	ballMaterial.ambient = 0.3146
	ball := MakeSphere(Identity().Translate(0, -3.5, -0.5), ballMaterial)

	w.objects = append(w.objects, []Object{floor, ball}...)

	xs := []Intersection{{t: math.Sqrt(2) / 2, object: floor}}
	hitData := PrepareComputations(xs[0], r, xs)

	AssertColorEqual(Color{0.93391, 0.69643, 0.60243}, hitData.Shade(w, 5), t)
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
