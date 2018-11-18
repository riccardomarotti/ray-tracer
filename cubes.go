package main

import (
	"math"
)

type Cube struct {
	transform Matrix
	material  Material
}

func MakeCube(transform Matrix, material Material) Object {
	return Cube{transform, material}
}

func (c Cube) Transform() Matrix {
	return c.transform
}

func (c Cube) Material() Material {
	return c.material
}

func (c Cube) NormalAt(p Tuple) Tuple {
	return Vector(0, 0, 0)
}

func (c Cube) Intersection(ray Ray) []Intersection {
	xTMin, xTMax := tMinMaxForAxis(ray.origin.x, ray.direction.x)
	yTMin, yTMax := tMinMaxForAxis(ray.origin.y, ray.direction.y)
	zTMin, zTMax := tMinMaxForAxis(ray.origin.z, ray.direction.z)

	tMin := math.Max(xTMin, math.Max(yTMin, zTMin))
	tMax := math.Min(xTMax, math.Min(yTMax, zTMax))
	return []Intersection{Intersection{tMin, c}, Intersection{tMax, c}}
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
