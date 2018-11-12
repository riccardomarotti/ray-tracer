package main

import (
	"sort"
)

type World struct {
	light   PointLight
	objects []Object
}

func DefaultWorld() World {
	light := PointLight{Point(-10, 10, -10), Color{1, 1, 1}}
	s1 := MakeSphere(Identity(), Material{
		color:     Color{0.8, 1.0, 0.6},
		ambient:   0.1,
		diffuse:   0.7,
		specular:  0.2,
		shininess: 200,
	})
	s2 := MakeSphere(Identity().Scale(0.5, 0.5, 0.5), DefaultMaterial())

	return World{light, []Object{s1, s2}}
}

func WorldWithAmbientSetTo(a float64) World {
	light := PointLight{Point(-10, 10, -10), Color{1, 1, 1}}
	s1 := MakeSphere(Identity(), Material{
		color:     Color{0.8, 1.0, 0.6},
		ambient:   a,
		diffuse:   0.7,
		specular:  0.2,
		shininess: 200,
	})
	s2 := MakeSphere(Identity().Scale(0.5, 0.5, 0.5), Material{Color{1, 1, 1}, 1, 0.9, 0.9, 200})

	return World{light, []Object{s1, s2}}
}

func (w World) Intersect(ray Ray) []Intersection {
	var intersections []Intersection
	channel := make(chan []Intersection)
	for _, object := range w.objects {
		go func(o Object) {
			channel <- ray.Intersection(o)
		}(object)
	}

	for i := 0; i < len(w.objects); i++ {
		intersections = append(intersections, <-channel...)
	}

	sort.Slice(intersections, func(i, j int) bool {
		return intersections[i].t < intersections[j].t
	})

	return intersections
}

func (w World) ColorAt(r Ray) Color {
	color := Color{0, 0, 0}
	intersections := w.Intersect(r)
	intersection := Hit(intersections)

	hit := intersection != Intersection{}
	if hit {
		intersection = PrepareHit(intersection, r)
		color = intersection.Shade(w)
	}

	return color
}

func (w World) IsShadowed(p Tuple) bool {
	v := w.light.position.Subtract(p)
	r := Ray{p, v.Normalize()}
	intersection := Hit(w.Intersect(r))
	// panic(fmt.Sprintf("%v", intersection))

	return intersection != Intersection{} && intersection.t < v.Magnitude()
}
