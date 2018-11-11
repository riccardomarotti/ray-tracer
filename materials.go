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

func (m Material) Lighting(light PointLight, position Tuple, eyeVector Tuple, normalVector Tuple, inShadow bool) Color {
	diffuse := Color{0, 0, 0}
	specular := Color{0, 0, 0}

	actualColor := m.color.Schur(light.intensity)
	lightVector := light.position.Subtract(position).Normalize()
	ambient := actualColor.By(m.ambient)
	lightDotNormal := lightVector.Dot(normalVector)

	if lightDotNormal > 0 && !inShadow {
		diffuse = actualColor.By(m.diffuse).By(lightDotNormal)
		reflectVector := lightVector.Multiply(-1).Reflect(normalVector)
		reflectDotEye := math.Pow(reflectVector.Dot(eyeVector), m.shininess)
		specular = light.intensity.By(m.specular).By(reflectDotEye)
	}

	return ambient.Add(diffuse).Add(specular)
}
