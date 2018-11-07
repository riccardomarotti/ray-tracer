package main

// Sphere represents a sphere
type Sphere struct {
	T Matrix
}

func MakeSphere(T Matrix) Object {
	return Sphere{T}
}

func (s Sphere) Transform() Matrix {
	return s.T
}

func (s Sphere) NormalAt(p Tuple) Tuple {
	objectPoint := s.Transform().Inverse().MultiplyByTuple(p)
	objectNormal := objectPoint.Subtract(Point(0, 0, 0))
	worldNormal := s.Transform().Inverse().T().MultiplyByTuple(objectNormal)
	worldNormal.w = 0
	return worldNormal.Normalize()
}
