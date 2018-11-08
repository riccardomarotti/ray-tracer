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

	AssertColorEqual(Color{1.9, 1.9, 1.9}, m.Lighting(light, position, eyeVector, normalVector), t)
}
func TestLightingWithTheEyeBetweenTheLightAndTheSurfaceEyeOffset45Degrees(t *testing.T) {
	m := DefaultMaterial()
	position := Point(0, 0, 0)

	eyeVector := Vector(0, math.Sqrt(2)/2, -math.Sqrt(2)/2)
	normalVector := Vector(0, 0, -1)
	light := PointLight{Point(0, 0, -10), Color{1, 1, 1}}

	AssertColorEqual(Color{1.0, 1.0, 1.0}, m.Lighting(light, position, eyeVector, normalVector), t)
}

func TestLightingWithEyeOppositeSurfaceLightOffset45Degrees(t *testing.T) {
	m := DefaultMaterial()
	position := Point(0, 0, 0)

	eyeVector := Vector(0, 0, -1)
	normalVector := Vector(0, 0, -1)
	light := PointLight{Point(0, 10, -10), Color{1, 1, 1}}

	AssertColorEqual(Color{0.7364, 0.7364, 0.7364}, m.Lighting(light, position, eyeVector, normalVector), t)
}

func TestLightingWithEyeInThePathOfTheReflectionVector(t *testing.T) {
	m := DefaultMaterial()
	position := Point(0, 0, 0)

	eyeVector := Vector(0, -math.Sqrt(2)/2, -math.Sqrt(2)/2)
	normalVector := Vector(0, 0, -1)
	light := PointLight{Point(0, 10, -10), Color{1, 1, 1}}

	AssertColorEqual(Color{1.6364, 1.6364, 1.6364}, m.Lighting(light, position, eyeVector, normalVector), t)
}
func TestLightingWithTheLightBehindTheSurface(t *testing.T) {
	m := DefaultMaterial()
	position := Point(0, 0, 0)

	eyeVector := Vector(0, 0, -1)
	normalVector := Vector(0, 0, -1)
	light := PointLight{Point(0, 0, 10), Color{1, 1, 1}}

	AssertColorEqual(Color{0.1, 0.1, 0.1}, m.Lighting(light, position, eyeVector, normalVector), t)
}
