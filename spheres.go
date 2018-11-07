package main

// Sphere represents a sphere
type Sphere struct {
	transform Matrix
	material  Material
}

func MakeSphere(transform Matrix, material Material) Object {
	return Sphere{transform, material}
}

func (s Sphere) Transform() Matrix {
	return s.transform
}

func (s Sphere) Material() Material {
	return s.material
}

func (s Sphere) NormalAt(p Tuple) Tuple {
	objectPoint := s.Transform().Inverse().MultiplyByTuple(p)
	objectNormal := objectPoint.Subtract(Point(0, 0, 0))
	worldNormal := s.Transform().Inverse().T().MultiplyByTuple(objectNormal)
	worldNormal.w = 0
	return worldNormal.Normalize()
}
