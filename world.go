package main

type World struct {
	light   PointLight
	objects []Object
}

func DefaultWorld() World {
	light := PointLight{Point(-10, -10, -10), Color{1, 1, 1}}
	s1 := MakeSphere(Identity(), MakeMaterial(Color{0.8, 1.0, 0.6}, 0.9, 0.7, 0.2, 200))
	s2 := MakeSphere(Identity().Scale(0.5, 0.5, 0.5), DefaultMaterial())

	return World{light, []Object{s1, s2}}
}
