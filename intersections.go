package main

import (
	"math"
	"reflect"
)

const Epsilon = 0.00001

type Intersection struct {
	t, n1, n2                                                 float64
	object                                                    Object
	point, eyeVector, normalVector, reflectVector, underPoint Tuple
	inside                                                    bool
}

func Hit(i []Intersection) (hit Intersection) {
	hit = Intersection{}

	for k := 0; k < len(i); k++ {
		if i[k].t >= 0 && (i[k].t <= hit.t || hit.t == 0) {
			hit = i[k]
		}
	}

	return
}

func PrepareComputations(i Intersection, r Ray, allIntersections []Intersection) Intersection {
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
				i.n1 = 1.0
			} else {
				lastIndex := len(containers) - 1
				i.n1 = containers[lastIndex].Material().refractiveIndex
			}
		}

		if x := contains(containers, intersection.object); x != -1 {
			containers = append(containers[:x], containers[x+1:]...)
		} else {
			containers = append(containers, intersection.object)
		}

		if areIntersectionsEqual(intersection, i) {
			if len(containers) == 0 {
				i.n2 = 1.0
			} else {
				lastIndex := len(containers) - 1
				i.n2 = containers[lastIndex].Material().refractiveIndex
			}

			break
		}
	}

	return Intersection{
		t:             i.t,
		object:        i.object,
		point:         point,
		eyeVector:     eyeVector,
		normalVector:  normalVector,
		inside:        inside,
		reflectVector: reflectVector,
		underPoint:    underPoint,
		n1:            i.n1,
		n2:            i.n2,
	}
}

func (i Intersection) Shade(world World, remaining int) Color {
	surface := i.object.Material().Lighting(i.object.Transform(), world.light, i.point, i.eyeVector, i.normalVector, world.IsShadowed(i.point))
	reflected := world.ReflectedColor(i)
	refracted := world.RefractedColor(i, remaining)
	return surface.Add(reflected).Add(refracted)
}

func contains(array []Object, o Object) int {
	for i := 0; i < len(array); i++ {
		if reflect.DeepEqual(o, array[i]) {
			return i
		}
	}

	return -1
}

func areIntersectionsEqual(i1, i2 Intersection) bool {
	return (math.Abs(i1.t-i2.t) < Epsilon) && reflect.DeepEqual(i1.object, i2.object)
}
