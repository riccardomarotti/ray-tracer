package main

import "sort"

type Group struct {
	transform Matrix
	children  []Object
	parent    *Group
}

func MakeGroup(transform Matrix) Group {
	return Group{transform, []Object{}, nil}
}

func MakeGroupInGroup(transform Matrix, parent *Group) Group {
	return Group{transform, []Object{}, parent}
}

func (g Group) Parent() *Group {
	return g.parent
}

func (g *Group) AddChildren(objects ...Object) {
	g.children = append(g.children, objects...)
}

func (g Group) Transform() Matrix {
	return g.transform
}

func (g Group) NormalAt(p Tuple) Tuple {
	panic("NormalAt should n ot be called on a Group")
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
