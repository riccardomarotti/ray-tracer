package main

type Triangle struct {
	p1, p2, p3     Tuple
	e1, e2, normal Tuple
	material       Material
}

func MakeTriangle(p1, p2, p3 Tuple) Triangle {
	e1 := p2.Subtract(p1)
	e2 := p3.Subtract(p1)
	normal := e2.Cross(e1).Normalize()
	return Triangle{p1: p1, p2: p2, p3: p3, e1: e1, e2: e2, normal: normal}
}

func (t Triangle) Transform() Matrix {
	return Identity()
}
func (t Triangle) NormalAt(p Tuple) Tuple {
	return t.normal
}

func (t Triangle) Material() Material {
	return t.material
}
func (t Triangle) Intersection(Ray) []Intersection {
	return []Intersection{}
}
