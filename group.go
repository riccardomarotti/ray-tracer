package main

import "sort"

type Group struct {
	baseObject BaseObject
	children   []Object
	parent     *Group
}

func MakeGroup(transform Matrix) Group {
	return Group{BaseObject{transform, DefaultMaterial()}, []Object{}, nil}
}

func MakeGroupInGroup(transform Matrix, parent *Group) Group {
	return Group{BaseObject{transform, DefaultMaterial()}, []Object{}, parent}
}

func (g Group) Parent() *Group {
	return g.parent
}

func (g *Group) AddChildren(objects ...Object) {
	g.children = append(g.children, objects...)
}

func (g Group) Transform() Matrix {
	return g.baseObject.transform
}

func (g Group) NormalAt(p Tuple) Tuple {
	panic("NormalAt should n ot be called on a Group")
}

func (g Group) Material() Material {
	return g.baseObject.material
}

func (g Group) Intersection(r Ray) []Intersection {
	localIntersect := func(r Ray) (intersections []Intersection) {
		for _, child := range g.children {
			intersections = append(intersections, child.Intersection(r)...)
		}

		sort.Slice(intersections, func(i, j int) bool {
			return intersections[i].t < intersections[j].t
		})

		return intersections
	}

	return g.baseObject.Intersection(r, localIntersect)
}
