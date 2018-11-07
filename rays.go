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
	transformedRay := r.Transform(s.Transform().Inverse())
	intersection = make([]Intersection, 0)

	sphereToRay := transformedRay.origin.Subtract(Point(0, 0, 0))
	b := 2 * transformedRay.direction.Dot(sphereToRay)
	a := transformedRay.direction.Dot(transformedRay.direction)
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

func (r Ray) Transform(T Matrix) Ray {
	return Ray{T.MultiplyByTuple(r.origin), T.MultiplyByTuple(r.direction)}
}
