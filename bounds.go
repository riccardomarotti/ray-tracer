package main

import "math"

type Bounds struct {
	min, max Tuple
}

func (b Bounds) Intersection(ray Ray, o Object) []Intersection {
	intersection := []Intersection{}

	xTMin, xTMax := b.tMinMaxForAxis(ray.origin.x, ray.direction.x, b.min.x, b.max.x)
	yTMin, yTMax := b.tMinMaxForAxis(ray.origin.y, ray.direction.y, b.min.y, b.max.y)
	zTMin, zTMax := b.tMinMaxForAxis(ray.origin.z, ray.direction.z, b.min.z, b.max.z)

	tMin := math.Max(xTMin, math.Max(yTMin, zTMin))
	tMax := math.Min(xTMax, math.Min(yTMax, zTMax))

	hit := tMax > tMin
	if hit {
		intersection = []Intersection{Intersection{t: tMin, object: o}, Intersection{t: tMax, object: o}}
	}

	return intersection
}

func (b Bounds) tMinMaxForAxis(origin, direction, min, max float64) (float64, float64) {
	tMinNumerator := min - origin
	tMaxNumrator := max - origin

	tMin := tMinNumerator / direction
	tMax := tMaxNumrator / direction

	if tMin > tMax {
		tMin, tMax = tMax, tMin
	}
	return tMin, tMax
}
