package main

type Material struct {
	color                                 Color
	ambient, diffuse, specular, shininess float64
}

func MakeMaterial() Material {
	return Material{Color{1, 1, 1}, 0.1, 0.9, 0.9, 200}
}
