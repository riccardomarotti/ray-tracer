package main

import "math"

type Cone struct {
	transform        Matrix
	material         Material
	minimum, maximum float64
	closed           bool
}

func MakeInfiniteCone(transform Matrix, material Material) Object {
	return Cone{transform, material, math.Inf(-1), math.Inf(1), false}
}

func MakeClosedCone(transform Matrix, material Material, minimum, maximum float64) Object {
	return Cone{transform, material, minimum, maximum, true}
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

	if cone.closed {
		intersections = append(intersections, cone.intersectConeCaps(r, intersections)...)
	}

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
		intersections = append(intersections, []Intersection{{t, cone}}...)

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

	y0 := r.origin.y + t0*r.direction.y
	y1 := r.origin.y + t1*r.direction.y

	if cone.closed {
		if y0 > cone.minimum && y0 < cone.maximum {
			intersections = append(intersections, []Intersection{Intersection{t0, cone}}...)
		}
		if y1 > cone.minimum && y1 < cone.maximum {
			intersections = append(intersections, []Intersection{Intersection{t1, cone}}...)
		}
	} else {
		intersections = []Intersection{Intersection{t0, cone}, Intersection{t1, cone}}
	}

	return
}

func checkConeCap(ray Ray, t, diameter float64) bool {
	x := ray.origin.x + t*ray.direction.x
	z := ray.origin.z + t*ray.direction.z
	return x*x+z*z <= diameter*diameter
}

func (cone Cone) intersectConeCaps(ray Ray, intersections []Intersection) []Intersection {
	var newIntersections []Intersection

	t := (cone.minimum - ray.origin.y) / ray.direction.y
	if checkConeCap(ray, t, cone.minimum) {
		newIntersections = append(newIntersections, Intersection{t, cone})
	}

	t = (cone.maximum - ray.origin.y) / ray.direction.y
	if checkConeCap(ray, t, cone.maximum) {
		newIntersections = append(newIntersections, Intersection{t, cone})
	}

	return append(intersections, newIntersections...)
}
