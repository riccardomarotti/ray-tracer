package main

import (
	"math"
	"sort"
)

type Group struct {
	baseObject        BaseObject
	children          []Object
	parent            *Group
	bounds            Bounds
	ignoreBoundingBox bool
}

func MakeGroup(transform Matrix) *Group {
	return &Group{BaseObject{transform, DefaultMaterial()}, []Object{}, nil, Bounds{}, false}
}

func MakeGroupInGroup(transform Matrix, parent *Group) *Group {
	g := &Group{BaseObject{transform, DefaultMaterial()}, []Object{}, parent, Bounds{}, false}
	parent.AddChildren(g)
	return g
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
	panic("Groups don't have a Material")
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

	if !g.ignoreBoundingBox {
		boundingBoxIntersections := g.bounds.Intersection(r, g)
		if len(boundingBoxIntersections) == 0 {
			return []Intersection{}
		}
	}

	return g.baseObject.Intersection(r, localIntersect)
}

func (g Group) Bounds() Bounds {
	return g.bounds
}

func (g *Group) CalculateBounds() {
	allBounds := []Bounds{}
	for _, child := range g.children {
		childBounds := child.Bounds()
		allBounds = append(allBounds, Bounds{childBounds.min, childBounds.max})
	}

	minx := math.Inf(1)
	miny := math.Inf(1)
	minz := math.Inf(1)
	maxx := math.Inf(-1)
	maxy := math.Inf(-1)
	maxz := math.Inf(-1)

	for _, bounds := range allBounds {
		if bounds.min.x < minx {
			minx = bounds.min.x
		}
		if bounds.min.y < miny {
			miny = bounds.min.y
		}
		if bounds.min.z < minz {
			minz = bounds.min.z
		}

		if bounds.max.x > maxx {
			maxx = bounds.max.x
		}

		if bounds.max.y > maxy {
			maxy = bounds.max.y
		}

		if bounds.max.z > maxz {
			maxz = bounds.max.z
		}
	}

	g.bounds = Bounds{Point(minx, miny, minz), Point(maxx, maxy, maxz)}
}
