package main

import (
	"math"
	"reflect"
	"testing"
)

func TestDefaultWorld(t *testing.T) {
	w := DefaultWorld()

	light := PointLight{Point(-10, 10, -10), Color{1, 1, 1}}
	s1 := MakeSphere(Identity(), Material{
		color:     Color{0.8, 1.0, 0.6},
		ambient:   0.1,
		diffuse:   0.7,
		specular:  0.2,
		shininess: 200,
	})
	s2 := MakeSphere(Identity().Scale(0.5, 0.5, 0.5), DefaultMaterial())

	Assert(light == w.light, "", t)
	Assert(reflect.DeepEqual(s1, w.objects[0]), "", t)
	Assert(reflect.DeepEqual(s2, w.objects[1]), "", t)
}

func TestIntersectWorldWithARay(t *testing.T) {
	w := DefaultWorld()
	ray := Ray{Point(0, 0, -5), Vector(0, 0, 1)}

	xs := w.Intersect(ray)

	Assert(len(xs) == 4, "", t)
	AssertEqual(4, xs[0].t, t)
	AssertEqual(4.5, xs[1].t, t)
	AssertEqual(5.5, xs[2].t, t)
	AssertEqual(6, xs[3].t, t)
}

func TestTheColorWhenARayMisses(t *testing.T) {
	world := DefaultWorld()
	ray := Ray{Point(0, 0, -5), Vector(0, 1, 0)}

	c := world.ColorAt(ray, 1)
	AssertColorEqual(Color{0, 0, 0}, c, t)
}

func TestTheColorWhenARayHits(t *testing.T) {
	world := DefaultWorld()
	ray := Ray{Point(0, 0, -5), Vector(0, 0, 1)}

	c := world.ColorAt(ray, 1)
	AssertColorEqual(Color{0.38066, 0.47583, 0.2855}, c, t)
}

func TestColotWithAnIntersectionBehindTheRay(t *testing.T) {
	world := WorldWithAmbientSetTo(1)

	ray := Ray{Point(0, 0, 0.75), Vector(0, 0, -1)}

	c := world.ColorAt(ray, 1)
	AssertColorEqual(world.objects[1].Material().color, c, t)
}

func TestNoShadowWhenNothingIsCollinearWithPointAndLight(t *testing.T) {
	world := DefaultWorld()
	p := Point(0, 10, 0)

	Assert(world.IsShadowed(p) == false, "The point shouln't be be shadowed", t)
}

func TestShadowWhenAnObjectIsBetweenThePointAndTheLight(t *testing.T) {
	world := DefaultWorld()
	p := Point(10, -10, 10)

	Assert(world.IsShadowed(p), "The point had to be shadowed", t)
}

func TestNoShadowWhenAnObjectIsBehindTheLight(t *testing.T) {
	world := DefaultWorld()
	p := Point(-20, 20, -20)

	Assert(world.IsShadowed(p) == false, "The point shouln't be be shadowed", t)
}

func TestNoShadowWhenAnObjectIsBehindThePoint(t *testing.T) {
	world := DefaultWorld()
	p := Point(-2, 2, -2)

	Assert(world.IsShadowed(p) == false, "The point shouln't be be shadowed", t)
}

func TestReflectedClorForNonReflectiveMaterial(t *testing.T) {
	world := WorldWithAmbientSetTo(.1)
	ray := Ray{Point(0, 0, 0), Vector(0, 0, 1)}
	shape := world.objects[1]
	hit := Intersection{t: 1, object: shape}

	hit = PrepareComputations(hit, ray, nil)

	AssertColorEqual(Color{0, 0, 0}, world.ReflectedColor(hit, 5), t)
}

func TestReflectedColorWithReflectiveMaterial(t *testing.T) {
	world := DefaultWorld()
	reflectiveMaterial := DefaultMaterial()
	reflectiveMaterial.reflective = 0.5
	plane := MakePlane(Identity().Translate(0, -1, 0), reflectiveMaterial)
	world.objects = append(world.objects, plane)
	ray := Ray{Point(0, 0, -3), Vector(0, -math.Sqrt(2)/2, math.Sqrt(2)/2)}
	hit := Intersection{t: math.Sqrt(2), object: plane}
	hit = PrepareComputations(hit, ray, nil)

	AssertColorEqual(Color{0.19032, -.2379, 0.14274}, world.ReflectedColor(hit, 5), t)
}

func TestMutuallyReflectiveSurfaces(t *testing.T) {
	light := PointLight{Point(-10, 10, -10), Color{1, 1, 1}}

	material := DefaultMaterial()
	material.reflective = 1

	lower := MakePlane(Identity().Translate(0, -1, 0), material)
	uppler := MakePlane(Identity().Translate(0, 1, 0), material)

	world := World{light, []Object{lower, uppler}}

	world.ColorAt(Ray{Point(0, 0, 0), Vector(0, 1, 0)}, 5)
}
func TestTheRefractedColorWithAnOpaqueSurface(t *testing.T) {
	w := DefaultWorld()
	shape := w.objects[0]
	r := Ray{Point(0, 0, -5), Vector(0, 0, 1)}
	xs := []Intersection{{t: 4, object: shape}, {t: 6, object: shape}}

	hitData := PrepareComputations(xs[0], r, xs)

	AssertColorEqual(Color{0, 0, 0}, w.RefractedColor(hitData, 5), t)
}

func TestTheRefractedColotAtTheMaximumRecursiveDepth(t *testing.T) {
	light := PointLight{Point(-10, 10, -10), Color{1, 1, 1}}
	shape := MakeSphere(Identity(), Material{
		color:           Color{0.8, 1.0, 0.6},
		ambient:         0.1,
		diffuse:         0.7,
		specular:        0.2,
		shininess:       200,
		transparency:    1.0,
		refractiveIndex: 1.5,
	})
	w := World{light, []Object{shape}}

	r := Ray{Point(0, 0, -5), Vector(0, 0, 1)}
	xs := []Intersection{{t: 4, object: shape}, {t: 5, object: shape}}

	hitData := PrepareComputations(xs[0], r, xs)

	AssertColorEqual(Color{0, 0, 0}, w.RefractedColor(hitData, 0), t)
}

func TestRefractedColorUnderTotalInternalRefrlection(t *testing.T) {
	light := PointLight{Point(-10, 10, -10), Color{1, 1, 1}}
	shape := MakeSphere(Identity(), Material{
		color:           Color{0.8, 1.0, 0.6},
		ambient:         0.1,
		diffuse:         0.7,
		specular:        0.2,
		shininess:       200,
		transparency:    1.0,
		refractiveIndex: 1.5,
	})

	otherShape := MakeSphere(Identity().Scale(0.5, 0.5, 0.5), DefaultMaterial())

	w := World{light, []Object{shape, otherShape}}
	r := Ray{Point(0, 0, math.Sqrt(2)/2), Vector(0, 1, 0)}
	xs := []Intersection{{t: -math.Sqrt(2) / 2, object: shape}, {t: math.Sqrt(2) / 2, object: shape}}

	hitData := PrepareComputations(xs[1], r, xs)

	AssertColorEqual(Color{0, 0, 0}, w.RefractedColor(hitData, 5), t)
}

func TestRefractedColorWithRefractedRay(t *testing.T) {
	light := PointLight{Point(-10, 10, -10), Color{1, 1, 1}}
	materialA := Material{
		color:     Color{0.8, 1.0, 0.6},
		ambient:   1,
		diffuse:   0.7,
		specular:  0.2,
		shininess: 200,
		pattern:   MakeTestPattern(Identity()),
	}

	materialB := DefaultMaterial()
	materialB.transparency = 1
	materialB.refractiveIndex = 1.5

	A := MakeSphere(Identity(), materialA)
	B := MakeSphere(Identity(), materialB)

	w := World{light, []Object{A, B}}

	r := Ray{Point(0, 0, 0.1), Vector(0, 1, 0)}
	xs := []Intersection{
		{t: -0.9899, object: A},
		{t: -0.4899, object: B},
		{t: 0.4899, object: B},
		{t: 0.9899, object: A},
	}

	hitData := PrepareComputations(xs[2], r, xs)

	AssertColorEqual(Color{0, 0.99878, 0.04724}, w.RefractedColor(hitData, 5), t)
}
