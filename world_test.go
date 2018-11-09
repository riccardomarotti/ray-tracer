package main

import (
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

	c := world.ColorAt(ray)
	AssertColorEqual(Color{0, 0, 0}, c, t)
}

func TestTheColorWhenARayHits(t *testing.T) {
	world := DefaultWorld()
	ray := Ray{Point(0, 0, -5), Vector(0, 0, 1)}

	c := world.ColorAt(ray)
	AssertColorEqual(Color{0.38066, 0.47583, 0.2855}, c, t)
}

func TestColotWithAnIntersectionBehindTheRay(t *testing.T) {
	world := WorldWithAmbientSetTo(1)

	ray := Ray{Point(0, 0, 0.75), Vector(0, 0, -1)}

	c := world.ColorAt(ray)
	AssertColorEqual(world.objects[1].Material().color, c, t)
}
