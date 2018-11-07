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
