package main

type Object interface {
	Transform() Matrix
	NormalAt(p Tuple) Tuple
	Material() Material
	Intersection(Ray) []Intersection
}
