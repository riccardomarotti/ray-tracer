package main

import "math"

type Cylinder struct {
	transform Matrix
	material  Material
}

func MakeCylinder(transform Matrix, material Material) Object {
	return Cylinder{transform, material}
}
func (c Cylinder) Transform() Matrix {
	return c.transform
}

func (c Cylinder) Material() Material {
	return c.material
}

func (c Cylinder) NormalAt(p Tuple) Tuple {
	// objectPoint := s.Transform().Inverse().MultiplyByTuple(p)
	// objectNormal := objectPoint.Subtract(Point(0, 0, 0))

	return Vector(0, 0, 0)

}

func (cylinder Cylinder) Intersection(r Ray) (intersection []Intersection) {
	transformedRay := r.Transform(cylinder.Transform().Inverse())
	intersection = make([]Intersection, 0)

	a := transformedRay.direction.x*transformedRay.direction.x + transformedRay.direction.z*transformedRay.direction.z
	if math.Abs(a) < Epsilon {
		return
	}

	b := 2*transformedRay.origin.x*transformedRay.direction.x + 2*transformedRay.origin.z*transformedRay.direction.z
	c := transformedRay.origin.x*transformedRay.origin.x + transformedRay.origin.z*transformedRay.origin.z - 1

	delta := b*b - 4*a*c

	if delta < 0 {
		return
	}

	t0 := (-b - math.Sqrt(delta)) / (2 * a)
	t1 := (-b + math.Sqrt(delta)) / (2 * a)

	intersection = []Intersection{
		{t: t0, object: cylinder},
		{t: t1, object: cylinder},
	}
	return
}
