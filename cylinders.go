package main

import "math"

type Cylinder struct {
	transform        Matrix
	material         Material
	minimum, maximum float64
	closed           bool
	parent           *Group
}

func MakeInfiniteCylinder(transform Matrix, material Material) Object {
	return Cylinder{transform, material, math.Inf(-1), math.Inf(1), false, nil}
}

func MakeClosedCylinder(transform Matrix, material Material, minimum, maximum float64) Object {
	return Cylinder{transform, material, minimum, maximum, true, nil}
}

func MakeCylinder(transform Matrix, material Material, minimum, maximum float64, closed bool) Object {
	return Cylinder{transform, material, minimum, maximum, closed, nil}
}

func (cylinder Cylinder) Parent() *Group {
	return cylinder.parent
}

func (cylinder Cylinder) Transform() Matrix {
	return cylinder.transform
}

func (cylinder Cylinder) Material() Material {
	return cylinder.material
}

func (cylinder Cylinder) NormalAt(p Tuple) Tuple {
	objectPoint := cylinder.Transform().Inverse().MultiplyByTuple(p)
	distance := objectPoint.x*objectPoint.x + objectPoint.z*objectPoint.z

	if distance < 1 && objectPoint.y > cylinder.maximum-Epsilon {
		return Vector(0, 1, 0)
	} else if distance < 1 && objectPoint.y <= cylinder.minimum+Epsilon {
		return Vector(0, -1, 0)
	}

	return Vector(objectPoint.x, 0, objectPoint.z)
}

func (cylinder Cylinder) Intersection(r Ray) (intersections []Intersection) {
	transformedRay := r.Transform(cylinder.Transform().Inverse())
	intersections = make([]Intersection, 0)

	if cylinder.closed {
		intersections = cylinder.intersectCaps(transformedRay, intersections)
	}

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

func checkCap(ray Ray, t float64) bool {
	x := ray.origin.x + t*ray.direction.x
	z := ray.origin.z + t*ray.direction.z

	return (x*x + z*z) <= 1
}

func (cylinder Cylinder) intersectCaps(ray Ray, intersections []Intersection) []Intersection {
	var newIntersections []Intersection

	t := (cylinder.minimum - ray.origin.y) / ray.direction.y
	if checkCap(ray, t) {
		newIntersections = append(newIntersections, Intersection{t, cylinder})
	}

	t = (cylinder.maximum - ray.origin.y) / ray.direction.y
	if checkCap(ray, t) {
		newIntersections = append(newIntersections, Intersection{t, cylinder})
	}

	return append(intersections, newIntersections...)

}
