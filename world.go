package main

import "sort"

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
