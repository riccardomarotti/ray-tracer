package main

import "math"

type Bounds struct {
	min, max Tuple
}

func (b Bounds) Intersection(ray Ray, o Object) []Intersection {
	intersection := []Intersection{}

	localRay := ray.Transform(o.Transform().Inverse())

	xTMin, xTMax := b.tMinMaxForAxis(localRay.origin.x, localRay.direction.x, b.min.x, b.max.x)
	yTMin, yTMax := b.tMinMaxForAxis(localRay.origin.y, localRay.direction.y, b.min.y, b.max.y)
	zTMin, zTMax := b.tMinMaxForAxis(localRay.origin.z, localRay.direction.z, b.min.z, b.max.z)

	tMin := math.Max(xTMin, math.Max(yTMin, zTMin))
	tMax := math.Min(xTMax, math.Min(yTMax, zTMax))

	hit := tMax > tMin
	if hit {
		intersection = []Intersection{Intersection{tMin, o}, Intersection{tMax, o}}
	}

	return intersection
}

func (b Bounds) tMinMaxForAxis(origin, direction, min, max float64) (float64, float64) {
	tMinNumerator := min - origin
	tMaxNumrator := max + origin

	tMin := tMinNumerator / direction
	tMax := tMaxNumrator / direction

	if tMin > tMax {
		tMin, tMax = tMax, tMin
	}
	return tMin, tMax
}
