package main

import (
	"math"
)

type Plane struct {
	transofrm Matrix
	material  Material
}

func MakePlane(transform Matrix, material Material) Object {
	return Plane{transform, material}
}

func (p Plane) NormalAt(point Tuple) Tuple {
	objectNormal := Vector(0, 1, 0)
	worldNormal := p.Transform().Inverse().T().MultiplyByTuple(objectNormal)
	worldNormal.w = 0

	return worldNormal.Normalize()
}

func (p Plane) Transform() Matrix {
	return p.transofrm
}

func (p Plane) Material() Material {
	return p.material
}

func (p Plane) Intersection(r Ray) (intersection []Intersection) {
	localRay := r.Transform(p.Transform().Inverse())
	intersection = make([]Intersection, 0)
	intersection = make([]Intersection, 0)

	if (math.Abs(localRay.direction.y)) >= Epsilon {
		t := -localRay.origin.y / localRay.direction.y
		i := Intersection{}
		i.t = t
		i.object = p
		return []Intersection{i}
	}

	return
}
