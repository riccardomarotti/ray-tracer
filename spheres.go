package main

// Sphere represents a sphere
type Sphere struct {
}

func MakeSphere() Object {
	return Sphere{}
}

func (s Sphere) Transform() Matrix {
	return Identity()
}
