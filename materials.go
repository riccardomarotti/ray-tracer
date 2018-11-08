package main

import (
	"math"
)

type Material struct {
	color                                 Color
	ambient, diffuse, specular, shininess float64
}

func DefaultMaterial() Material {
	return Material{Color{1, 1, 1}, 0.1, 0.9, 0.9, 200}
}

func MakeMaterial(color Color, ambient, diffuse, specular, shininess float64) Material {
	return Material{color, ambient, diffuse, specular, shininess}
}

func (m Material) Lighting(light PointLight, position Tuple, eyeVector Tuple, normalVector Tuple) Color {
	diffuse := Color{0, 0, 0}
	specular := Color{0, 0, 0}

	actualColor := m.color.Schur(light.intensity)
	lightVector := light.position.Subtract(position).Normalize()
	ambient := actualColor.By(m.ambient)
	lightDotNormal := lightVector.Dot(normalVector)

	if lightDotNormal > 0 {
		diffuse = actualColor.By(m.diffuse).By(lightDotNormal)
		reflectVector := lightVector.Multiply(-1).Reflect(normalVector)
		reflectDotEye := math.Pow(reflectVector.Dot(eyeVector), m.shininess)

		if reflectDotEye > 0 {
			specular = light.intensity.By(m.specular).By(reflectDotEye)
		}
	}

	return ambient.Add(diffuse).Add(specular)
}
