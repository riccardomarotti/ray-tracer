package main

import (
	"math"
)

type Computations struct {
	object                                                    Object
	t, n1, n2                                                 float64
	point, eyeVector, normalVector, reflectVector, underPoint Tuple
	inside                                                    bool
}

func PrepareComputations(i Intersection, r Ray, allIntersections []Intersection) (comps Computations) {
	comps = Computations{}
	comps.t = i.t
	comps.object = i.object

	rawPoint := r.Position(i.t)
	normalVector := i.object.NormalAt(rawPoint)
	point := rawPoint.Add(normalVector.Multiply(Epsilon))
	underPoint := rawPoint.Subtract(normalVector.Multiply(Epsilon))

	eyeVector := r.direction.Multiply(-1)
	inside := false

	if normalVector.Dot(eyeVector) < 0 {
		inside = true
		normalVector = normalVector.Multiply(-1)
	}

	reflectVector := r.direction.Reflect(normalVector)

	containers := []Object{}

	for _, intersection := range allIntersections {
		if areIntersectionsEqual(intersection, i) {
			if len(containers) == 0 {
				comps.n1 = 1.0
			} else {
				lastIndex := len(containers) - 1
				comps.n1 = containers[lastIndex].Material().refractiveIndex
			}
		}

		if x := contains(containers, intersection.object); x != -1 {
			containers = append(containers[:x], containers[x+1:]...)
		} else {
			containers = append(containers, intersection.object)
		}

		if areIntersectionsEqual(intersection, i) {
			if len(containers) == 0 {
				comps.n2 = 1.0
			} else {
				lastIndex := len(containers) - 1
				comps.n2 = containers[lastIndex].Material().refractiveIndex
			}

			break
		}
	}

	comps.point = point
	comps.eyeVector = eyeVector

	comps.normalVector = normalVector
	comps.inside = inside
	comps.reflectVector = reflectVector
	comps.underPoint = underPoint

	return
}

func (c Computations) Shade(world World, remaining int) Color {
	surface := c.object.Material().Lighting(c.object.Transform(), world.light, c.point, c.eyeVector, c.normalVector, world.IsShadowed(c.point))
	reflected := world.ReflectedColor(c, remaining-1)
	refracted := world.RefractedColor(c, remaining-1)

	material := c.object.Material()
	if material.reflective > 0 && material.transparency > 0 {
		reflectance := c.Schlick()
		return surface.Add(reflected.Multiply(reflectance)).Add(refracted.Multiply(1 - reflectance))
	}
	return surface.Add(reflected).Add(refracted)
}

func (i Computations) Schlick() float64 {
	cos := i.eyeVector.Dot(i.normalVector)

	if i.n1 > i.n2 {
		n := i.n1 / i.n2
		sin2Theta := n * n * (1 - cos*cos)
		if sin2Theta > 1 {
			return 1
		}
		cos = math.Sqrt(1 - sin2Theta)
	}

	r0 := math.Pow(((i.n1 - i.n2) / (i.n1 + i.n2)), 2)
	return r0 + (1-r0)*math.Pow(1-cos, 5)
}
