package main

import "math"

type Cone struct {
	transform Matrix
	material  Material
}

func MakeInfiniteCone(transform Matrix, material Material) Object {
	return Cone{transform, material}
}

func (cone Cone) Transform() Matrix {
	return cone.transform
}

func (cone Cone) Material() Material {
	return cone.material
}

func (cone Cone) NormalAt(p Tuple) Tuple {
	// objectPoint := cone.Transform().Inverse().MultiplyByTuple(p)
	// distance := objectPoint.x*objectPoint.x + objectPoint.z*objectPoint.z

	return Vector(0, 0, 0)
}

func (cone Cone) Intersection(r Ray) (intersections []Intersection) {
	transformedRay := r.Transform(cone.Transform().Inverse())
	intersections = make([]Intersection, 0)

	dx := transformedRay.direction.x
	dy := transformedRay.direction.y
	dz := transformedRay.direction.z
	ox := transformedRay.origin.x
	oy := transformedRay.origin.y
	oz := transformedRay.origin.z

	a := dx*dx - dy*dy + dz*dz
	b := 2*ox*dx - 2*oy*dy + 2*oz*dz

	if math.Abs(a) < Epsilon && math.Abs(b) < Epsilon {
		return
	}

	c := ox*ox - oy*oy + oz*oz

	if math.Abs(a) < Epsilon && math.Abs(b) > Epsilon {
		t := -c / (2 * b)
		intersections = []Intersection{{t, cone}}
		return
	}

	delta := b*b - 4*a*c

	if delta < 0 {
		return
	}

	t0 := (-b - math.Sqrt(delta)) / (2 * a)
	t1 := (-b + math.Sqrt(delta)) / (2 * a)

	if t0 > t1 {
		t0, t1 = t1, t0
	}

	intersections = []Intersection{Intersection{t0, cone}, Intersection{t1, cone}}

	return
}
