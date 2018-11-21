package main

import "math"

type Triangle struct {
	p1, p2, p3     Tuple
	e1, e2, normal Tuple
	material       Material
	transform      Matrix
}

func MakeTriangle(p1, p2, p3 Tuple, transofrm Matrix, material Material) Triangle {
	e1 := p2.Subtract(p1)
	e2 := p3.Subtract(p1)
	normal := e2.Cross(e1).Normalize()
	return Triangle{p1: p1, p2: p2, p3: p3, e1: e1, e2: e2, normal: normal, transform: transofrm, material: material}
}

func (t Triangle) Transform() Matrix {
	return t.transform
}

func (t Triangle) NormalAt(p Tuple) Tuple {
	worldNormal := t.Transform().Inverse().T().MultiplyByTuple(t.normal)
	worldNormal.w = 0
	return worldNormal
}

func (t Triangle) Material() Material {

	return t.material
}

func (t Triangle) Intersection(r Ray) (intersection []Intersection) {
	transformedRay := r.Transform(t.Transform().Inverse())
	intersection = make([]Intersection, 0)

	dirCrossE2 := transformedRay.direction.Cross(t.e2)
	det := t.e1.Dot(dirCrossE2)

	if math.Abs(det) > Epsilon {
		f := 1 / det
		p1ToOrigin := transformedRay.origin.Subtract(t.p1)
		u := f * p1ToOrigin.Dot(dirCrossE2)
		if u >= 0 && u <= 1 {
			originCrossE1 := p1ToOrigin.Cross(t.e1)
			v := f * transformedRay.direction.Dot(originCrossE1)
			if v >= 0 && (u+v) <= 1 {

				intersection = append(intersection, Intersection{t: f * t.e2.Dot(originCrossE1), object: t})
			}
		}
	}

	return
}
