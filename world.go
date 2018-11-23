package main

import (
	"math"
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
	s2 := MakeSphere(Identity().Scale(0.5, 0.5, 0.5), Material{color: Color{1, 1, 1}, ambient: 1, diffuse: 0.9, specular: 0.9, shininess: 200})

	return World{light, []Object{s1, s2}}
}

func (w World) Intersect(ray Ray) []Intersection {
	var intersections []Intersection
	channel := make(chan []Intersection)
	for _, object := range w.objects {
		go func(o Object) {
			channel <- o.Intersection(ray)
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

func (w World) ColorAt(r Ray, remaining int) Color {
	color := Color{0, 0, 0}
	intersections := w.Intersect(r)
	intersection := Hit(intersections)

	hit := intersection != Intersection{}
	if hit {
		comps := PrepareComputations(intersection, r, intersections)
		color = comps.Shade(w, remaining)
	}

	return color
}

func (w World) IsShadowed(p Tuple) bool {
	v := w.light.position.Subtract(p)
	r := Ray{p, v.Normalize()}
	intersection := Hit(w.Intersect(r))

	return intersection != Intersection{} && intersection.t < v.Magnitude()
}

func (w World) ReflectedColor(c Computations, remaining int) Color {
	if c.i.object.Material().reflective == 0 || remaining <= 0 {
		return Color{0, 0, 0}
	}

	reflectedRay := Ray{c.point, c.reflectVector}
	color := w.ColorAt(reflectedRay, remaining-1)

	return color.Multiply(c.i.object.Material().reflective)
}

func (w World) RefractedColor(c Computations, remaining int) Color {
	nRatio := c.n1 / c.n2
	cosThetaI := c.eyeVector.Dot(c.normalVector)
	sinThetaT := nRatio * nRatio * (1 - (cosThetaI * cosThetaI))

	if c.i.object.Material().transparency == 0 || remaining <= 0 || sinThetaT > 1 {
		return Color{0, 0, 0}
	}

	cosThetaT := math.Sqrt(1 - sinThetaT)
	direction := c.normalVector.Multiply(nRatio*cosThetaI - cosThetaT).Subtract(c.eyeVector.Multiply(nRatio))
	refractRay := Ray{c.underPoint, direction}

	return w.ColorAt(refractRay, remaining-1).Multiply(c.i.object.Material().transparency)
}
