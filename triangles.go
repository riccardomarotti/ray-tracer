package main

import "math"

type Triangle struct {
	p1, p2, p3     Tuple
	e1, e2, normal Tuple
	material       Material
	transform      Matrix
}

func MakeTriangle(p1, p2, p3 Tuple, transofrm Matrix) Triangle {
	e1 := p2.Subtract(p1)
	e2 := p3.Subtract(p1)
	normal := e2.Cross(e1).Normalize()
	return Triangle{p1: p1, p2: p2, p3: p3, e1: e1, e2: e2, normal: normal, transform: transofrm}
}

func (t Triangle) Transform() Matrix {
	return t.transform
}
func (t Triangle) NormalAt(p Tuple) Tuple {
	return t.normal
}

func (t Triangle) Material() Material {
	return t.material
}
func (t Triangle) Intersection(r Ray) (intersection []Intersection) {
	transformedRay := r.Transform(t.Transform().Inverse())
	intersection = make([]Intersection, 0)

	dirCrossE2 := transformedRay.direction.Cross(t.e2)
	det := t.e2.Dot(dirCrossE2)

	if math.Abs(det) > Epsilon {
		intersection = append(intersection, Intersection{1, t})
	}
	return
}
