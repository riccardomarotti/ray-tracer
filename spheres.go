package main

import "math"

type Sphere struct {
	baseObject BaseObject
	parent     *Group
}

func MakeSphereInGroup(transform Matrix, material Material, group *Group) Object {
	return Sphere{BaseObject{transform, material}, group}
}

func MakeSphere(transform Matrix, material Material) Object {
	return Sphere{BaseObject{transform, material}, nil}
}

func (s Sphere) Parent() *Group {
	return s.parent
}

func (s Sphere) Transform() Matrix {
	return s.baseObject.transform
}

func (s Sphere) Material() Material {
	return s.baseObject.material
}

func (s Sphere) NormalAt(p Tuple) Tuple {
	localNormalAt := func(p Tuple) Tuple {
		return p.Subtract(Point(0, 0, 0))
	}

	return s.baseObject.NormalAt(p, s, localNormalAt)
}

func (s Sphere) Intersection(r Ray) (intersection []Intersection) {
	localIntersect := func(r Ray) (intersection []Intersection) {
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

			i1 := Intersection{t: t1, object: s}
			i2 := Intersection{t: t2, object: s}

			intersection = []Intersection{i1, i2}
		}

		return
	}

	return s.baseObject.Intersection(r, localIntersect)
}

func MakeGlassSphere(transformation Matrix, refractiveIndex float64) Object {
	material := DefaultMaterial()
	material.transparency = 1.0
	material.refractiveIndex = refractiveIndex
	return MakeSphere(transformation, material)
}
