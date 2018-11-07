package main

import (
	"math"
)

type Ray struct {
	origin, direction Tuple
}

func (r Ray) Position(d float64) Tuple {
	return r.direction.Multiply(d).Add(r.origin)
}

func (r Ray) Intersection(s Object) (intersection []Intersection) {
	intersection = make([]Intersection, 0)

	sphereToRay := r.origin.Subtract(Point(0, 0, 0))
	b := 2 * r.direction.Dot(sphereToRay)
	a := r.direction.Dot(r.direction)
	c := sphereToRay.Dot(sphereToRay) - 1

	delta := b*b - 4*a*c

	if delta >= 0 {
		t1 := (-b - math.Sqrt(delta)) / (2 * a)
		t2 := (-b + math.Sqrt(delta)) / (2 * a)
		if t1 > t2 {
			t1, t2 = t2, t1
		}

		i1 := Intersection{t1, s}
		i2 := Intersection{t2, s}

		intersection = []Intersection{i1, i2}
	}

	return
}

func (r Ray) Translate(x, y, z float64) Ray {
	T := Identity().Translate(x, y, z)

	return Ray{T.MultiplyByTuple(r.origin), T.MultiplyByTuple(r.direction)}
}

func (r Ray) Scale(x, y, z float64) Ray {
	T := Identity().Scale(x, y, z)

	return Ray{T.MultiplyByTuple(r.origin), T.MultiplyByTuple(r.direction)}
}
