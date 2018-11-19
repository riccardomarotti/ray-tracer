package main

import "sort"

type Group struct {
	transform Matrix
	children  []Object
}

func MakeGroup(transform Matrix, children ...Object) Group {
	return Group{transform, children}
}

func (g Group) Transform() Matrix {
	return g.transform
}

func (g Group) NormalAt(p Tuple) Tuple {
	return Vector(0, 0, 0)
}

func (g Group) Material() Material {
	return Material{}
}

func (g Group) Intersection(r Ray) (intersections []Intersection) {
	intersections = make([]Intersection, 0)
	transformedRay := r.Transform(g.Transform().Inverse())

	for _, child := range g.children {
		intersections = append(intersections, child.Intersection(transformedRay)...)
	}

	sort.Slice(intersections, func(i, j int) bool {
		return intersections[i].t < intersections[j].t
	})

	return intersections
}
