package main

type Plane struct {
	transofrm Matrix
	material  Material
}

func MakePlane(transform Matrix, material Material) Object {
	return Plane{transform, material}
}

func (p Plane) NormalAt(point Tuple) Tuple {
	return Vector(0, 1, 0)
}

func (p Plane) Transform() Matrix {
	return p.transofrm
}

func (p Plane) Material() Material {
	return p.material
}

func (p Plane) Intersection(r Ray) []Intersection {
	return make([]Intersection, 0)
}
