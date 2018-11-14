package main

import (
	"math"
	"testing"
)

func TestDefaultMaterial(t *testing.T) {
	m := DefaultMaterial()

	Assert(Color{1, 1, 1} == m.color, "", t)
	AssertEqual(0.1, m.ambient, t)
	AssertEqual(0.9, m.diffuse, t)
	AssertEqual(0.9, m.specular, t)
	AssertEqual(200, m.shininess, t)
}

func TestLightingWithTheEyeBetweenTheLightAndTheSurface(t *testing.T) {
	m := DefaultMaterial()
	position := Point(0, 0, 0)

	eyeVector := Vector(0, 0, -1)
	normalVector := Vector(0, 0, -1)
	light := PointLight{Point(0, 0, -10), Color{1, 1, 1}}

	AssertColorEqual(Color{1.9, 1.9, 1.9}, m.Lighting(Identity(), light, position, eyeVector, normalVector, false), t)
}
func TestLightingWithTheEyeBetweenTheLightAndTheSurfaceEyeOffset45Degrees(t *testing.T) {
	m := DefaultMaterial()
	position := Point(0, 0, 0)

	eyeVector := Vector(0, math.Sqrt(2)/2, -math.Sqrt(2)/2)
	normalVector := Vector(0, 0, -1)
	light := PointLight{Point(0, 0, -10), Color{1, 1, 1}}

	AssertColorEqual(Color{1.0, 1.0, 1.0}, m.Lighting(Identity(), light, position, eyeVector, normalVector, false), t)
}

func TestLightingWithEyeOppositeSurfaceLightOffset45Degrees(t *testing.T) {
	m := DefaultMaterial()
	position := Point(0, 0, 0)

	eyeVector := Vector(0, 0, -1)
	normalVector := Vector(0, 0, -1)
	light := PointLight{Point(0, 10, -10), Color{1, 1, 1}}

	AssertColorEqual(Color{0.7364, 0.7364, 0.7364}, m.Lighting(Identity(), light, position, eyeVector, normalVector, false), t)
}

func TestLightingWithEyeInThePathOfTheReflectionVector(t *testing.T) {
	m := DefaultMaterial()
	position := Point(0, 0, 0)

	eyeVector := Vector(0, -math.Sqrt(2)/2, -math.Sqrt(2)/2)
	normalVector := Vector(0, 0, -1)
	light := PointLight{Point(0, 10, -10), Color{1, 1, 1}}

	AssertColorEqual(Color{1.6364, 1.6364, 1.6364}, m.Lighting(Identity(), light, position, eyeVector, normalVector, false), t)
}
func TestLightingWithTheLightBehindTheSurface(t *testing.T) {
	m := DefaultMaterial()
	position := Point(0, 0, 0)

	eyeVector := Vector(0, 0, -1)
	normalVector := Vector(0, 0, -1)
	light := PointLight{Point(0, 0, 10), Color{1, 1, 1}}

	AssertColorEqual(Color{0.1, 0.1, 0.1}, m.Lighting(Identity(), light, position, eyeVector, normalVector, false), t)
}

func TestLightingWithTheSurfaceInShadow(t *testing.T) {
	m := DefaultMaterial()
	position := Point(0, 0, 0)

	eyeVector := Vector(0, 0, -1)
	normalVector := Vector(0, 0, -1)
	light := PointLight{Point(0, 0, -10), Color{1, 1, 1}}
	inShadow := true

	result := m.Lighting(Identity(), light, position, eyeVector, normalVector, inShadow)
	AssertColorEqual(Color{0.1, 0.1, 0.1}, result, t)
}

func TestLightingWithPatternApplied(t *testing.T) {
	m := DefaultMaterial()
	m.pattern = StripePattern{Color{1, 1, 1}, Color{0, 0, 0}, Identity()}
	m.ambient = 1
	m.diffuse = 0
	m.specular = 0
	eyeVector := Vector(0, 0, -1)
	normalVector := Vector(0, 0, -1)
	light := PointLight{Point(0, 0, -10), Color{1, 1, 1}}

	c1 := m.Lighting(Identity(), light, Point(0.9, 0, 0), eyeVector, normalVector, false)
	c2 := m.Lighting(Identity(), light, Point(1.1, 0, 0), eyeVector, normalVector, false)

	AssertColorEqual(Color{1, 1, 1}, c1, t)
	AssertColorEqual(Color{0, 0, 0}, c2, t)
}

func TestReflectivityForDefaultMaterial(t *testing.T) {
	m := DefaultMaterial()

	AssertEqual(0, m.reflective, t)
}

func TestTransparencyForDeaultMaterial(t *testing.T) {
	m := DefaultMaterial()

	AssertEqual(0, m.transparency, t)
}

func TestRefractiveIndexForDefaultMaterial(t *testing.T) {
	m := DefaultMaterial()

	AssertEqual(1, m.refractiveIndex, t)
}
