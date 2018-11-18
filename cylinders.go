package main

import "math"

type Cylinder struct {
	transform        Matrix
	material         Material
	minimum, maximum float64
}

func MakeInfiniteCylinder(transform Matrix, material Material) Object {
	return Cylinder{transform, material, math.Inf(-1), math.Inf(1)}
}

func MakeCylinder(transform Matrix, material Material, minimum, maximum float64) Object {
	return Cylinder{transform, material, minimum, maximum}
}
func (c Cylinder) Transform() Matrix {
	return c.transform
}

func (c Cylinder) Material() Material {
	return c.material
}

func (c Cylinder) NormalAt(p Tuple) Tuple {
	objectPoint := c.Transform().Inverse().MultiplyByTuple(p)

	return Vector(objectPoint.x, 0, objectPoint.z)

}

func (cylinder Cylinder) Intersection(r Ray) (intersections []Intersection) {
	transformedRay := r.Transform(cylinder.Transform().Inverse())
	intersections = make([]Intersection, 0)

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

	if t0 > t1 {
		t0, t1 = t1, t0
	}

	y0 := transformedRay.origin.y + t0*transformedRay.direction.y
	if (cylinder.minimum) < y0 && y0 < cylinder.maximum {
		intersections = append(intersections, Intersection{t0, cylinder})
	}

	y1 := transformedRay.origin.y + t1*transformedRay.direction.y
	if cylinder.minimum < y1 && y1 < cylinder.maximum {
		intersections = append(intersections, Intersection{t1, cylinder})
	}

	return
}
