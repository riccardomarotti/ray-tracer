package main

import (
	"math"
)

type Plane struct {
	baseObject BaseObject
	parent     *Group
}

func MakePlane(transform Matrix, material Material) Object {
	return Plane{BaseObject{transform, material}, nil}
}

func (p Plane) Parent() *Group {
	return p.parent
}

func (p Plane) NormalAt(point Tuple) Tuple {
	localNormalAt := func(p Tuple) Tuple {
		return Vector(0, 1, 0)
	}

	return p.baseObject.NormalAt(point, p, localNormalAt)
}

func (p Plane) Transform() Matrix {
	return p.baseObject.transform
}

func (p Plane) Material() Material {
	return p.baseObject.material
}

func (p Plane) Intersection(r Ray) []Intersection {
	localIntersect := func(r Ray) (intersection []Intersection) {
		intersection = make([]Intersection, 0)
		if (math.Abs(r.direction.y)) >= Epsilon {
			t := -r.origin.y / r.direction.y
			i := Intersection{}
			i.t = t
			i.object = p
			intersection = []Intersection{i}
		}

		return
	}

	return p.baseObject.Intersection(r, localIntersect)
}
