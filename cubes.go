package main

import (
	"math"
)

type Cube struct {
	baseObject BaseObject
	parent     *Group
}

func MakeCube(transform Matrix, material Material) Object {
	return Cube{BaseObject{transform, material}, nil}
}

func MakeCubeInGroup(transform Matrix, material Material, group *Group) Object {
	c := Cube{BaseObject{transform, material}, group}
	group.AddChildren(c)
	return c
}

func (c Cube) Bounds() Bounds {
	localBounds := Bounds{Point(-1, -1, -1), Point(1, 1, 1)}
	return c.baseObject.Bounds(localBounds, c)
}
func (c Cube) Parent() *Group {
	return c.parent
}

func (c Cube) Transform() Matrix {
	return c.baseObject.transform
}

func (c Cube) Material() Material {
	return c.baseObject.material
}

func (c Cube) NormalAt(p Tuple) Tuple {
	localNormalAt := func(p Tuple) Tuple {
		maxC := math.Max(math.Abs(p.x), math.Max(math.Abs(p.y), math.Abs(p.z)))

		if maxC == math.Abs(p.x) {
			return Vector(p.x, 0, 0)
		} else if maxC == math.Abs(p.y) {
			return Vector(0, p.y, 0)
		}
		return Vector(0, 0, p.z)
	}

	return c.baseObject.NormalAt(p, c, localNormalAt)
}

func (c Cube) Intersection(ray Ray) []Intersection {
	localIntersection := func(r Ray) (intersection []Intersection) {
		intersection = []Intersection{}

		xTMin, xTMax := tMinMaxForAxis(r.origin.x, r.direction.x)
		yTMin, yTMax := tMinMaxForAxis(r.origin.y, r.direction.y)
		zTMin, zTMax := tMinMaxForAxis(r.origin.z, r.direction.z)

		tMin := math.Max(xTMin, math.Max(yTMin, zTMin))
		tMax := math.Min(xTMax, math.Min(yTMax, zTMax))

		hit := tMax > tMin
		if hit {
			intersection = []Intersection{Intersection{t: tMin, object: c}, Intersection{t: tMax, object: c}}
		}

		return
	}

	return c.baseObject.Intersection(ray, localIntersection)
}

func tMinMaxForAxis(origin, direction float64) (float64, float64) {
	tMinNumerator := -1 - origin
	tMaxNumrator := 1 - origin

	tMin := tMinNumerator / direction
	tMax := tMaxNumrator / direction

	if tMin > tMax {
		tMin, tMax = tMax, tMin
	}
	return tMin, tMax
}
