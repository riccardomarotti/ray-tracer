package main

type Plane struct {
	transofrm Matrix
	material  Material
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
