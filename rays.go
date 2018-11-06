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

func (r Ray) Intersection(s Solid) (intersection []float64) {
	intersection = make([]float64, 0)

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

		intersection = []float64{t1, t2}
	}

	return
}
