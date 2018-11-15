package main

import (
	"math"
)

type Material struct {
	color                                                                            Color
	ambient, diffuse, specular, shininess, reflective, transparency, refractiveIndex float64
	pattern                                                                          Pattern
}

func DefaultMaterial() Material {
	return Material{
		color:           Color{1, 1, 1},
		ambient:         0.1,
		diffuse:         0.9,
		specular:        0.9,
		shininess:       200,
		refractiveIndex: 1}
}

func (m Material) Lighting(objectTransform Matrix, light PointLight, position Tuple, eyeVector Tuple, normalVector Tuple, inShadow bool) Color {
	diffuse := Color{0, 0, 0}
	specular := Color{0, 0, 0}

	hasPattern := m.pattern != nil
	var actualColor Color
	if hasPattern {
		actualColor = m.pattern.ColorAt(position, objectTransform)
	} else {
		actualColor = m.color
	}

	actualColor = actualColor.Schur(light.intensity)
	lightVector := light.position.Subtract(position).Normalize()
	ambient := actualColor.Multiply(m.ambient)
	lightDotNormal := lightVector.Dot(normalVector)

	if lightDotNormal > 0 && !inShadow {
		diffuse = actualColor.Multiply(m.diffuse).Multiply(lightDotNormal)
		reflectVector := lightVector.Multiply(-1).Reflect(normalVector)
		reflectDotEye := math.Pow(reflectVector.Dot(eyeVector), m.shininess)
		specular = light.intensity.Multiply(m.specular).Multiply(reflectDotEye)
	}

	return ambient.Add(diffuse).Add(specular)
}
