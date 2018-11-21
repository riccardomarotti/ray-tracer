package main

import (
	"math"
)

type Cube struct {
	transform Matrix
	material  Material
	parent    *Group
}

func MakeCube(transform Matrix, material Material) Object {
	return Cube{transform, material, nil}
}

func (c Cube) Parent() *Group {
	return c.parent
}

func (c Cube) Transform() Matrix {
	return c.transform
}

func (c Cube) Material() Material {
	return c.material
}

func (c Cube) NormalAt(p Tuple) Tuple {
	objectPoint := c.Transform().Inverse().MultiplyByTuple(p)

	maxC := math.Max(math.Abs(objectPoint.x), math.Max(math.Abs(objectPoint.y), math.Abs(objectPoint.z)))

	if maxC == math.Abs(objectPoint.x) {
		return Vector(objectPoint.x, 0, 0)
	} else if maxC == math.Abs(objectPoint.y) {
		return Vector(0, objectPoint.y, 0)
	}
	return Vector(0, 0, objectPoint.z)
}

func (c Cube) Intersection(ray Ray) (intersection []Intersection) {
	transformedRay := ray.Transform(c.Transform().Inverse())
	intersection = []Intersection{}

	xTMin, xTMax := tMinMaxForAxis(transformedRay.origin.x, transformedRay.direction.x)
	yTMin, yTMax := tMinMaxForAxis(transformedRay.origin.y, transformedRay.direction.y)
	zTMin, zTMax := tMinMaxForAxis(transformedRay.origin.z, transformedRay.direction.z)

	tMin := math.Max(xTMin, math.Max(yTMin, zTMin))
	tMax := math.Min(xTMax, math.Min(yTMax, zTMax))

	hit := tMax > tMin
	if hit {
		intersection = []Intersection{Intersection{tMin, c}, Intersection{tMax, c}}
	}

	return
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
