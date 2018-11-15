package main

import "math"

// Sphere represents a sphere
type Sphere struct {
	transform Matrix
	material  Material
}

func MakeSphere(transform Matrix, material Material) Object {
	return Sphere{transform, material}
}

func (s Sphere) Transform() Matrix {
	return s.transform
}

func (s Sphere) Material() Material {
	return s.material
}

func (s Sphere) NormalAt(p Tuple) Tuple {
	objectPoint := s.Transform().Inverse().MultiplyByTuple(p)
	objectNormal := objectPoint.Subtract(Point(0, 0, 0))
	worldNormal := s.Transform().Inverse().T().MultiplyByTuple(objectNormal)
	worldNormal.w = 0
	return worldNormal.Normalize()
}

func (s Sphere) Intersection(r Ray) (intersection []Intersection) {
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

		i1 := Intersection{t: t1, object: s}
		i2 := Intersection{t: t2, object: s}

		intersection = []Intersection{i1, i2}
	}

	return
}

func MakeGlassSphere(transofrmation Matrix, refractiveIndex float64) Object {
	material := DefaultMaterial()
	material.transparency = 1.0
	material.refractiveIndex = refractiveIndex
	return MakeSphere(transofrmation, material)
}
