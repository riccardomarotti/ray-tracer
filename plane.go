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
	return Vector(0, 1, 0)
}

func (p Plane) Transform() Matrix {
	return p.transofrm
}

func (p Plane) Material() Material {
	return p.material
}

func (p Plane) Intersection(r Ray) (intersection []Intersection) {
	intersection = make([]Intersection, 0)
	if (math.Abs(r.direction.y)) >= 0.0001 {
		t := -r.origin.y / r.direction.y
		i := Intersection{}
		i.t = t
		i.object = p
		return []Intersection{i}
	}

	return
}
