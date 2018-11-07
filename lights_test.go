package main

import "testing"

func TestAPointLightHasAPositionAndIntensity(t *testing.T) {
	intensity := Color{1, 1, 1}
	position := Point(0, 0, 0)

	light := PointLight{position, intensity}

	AssertTupleEqual(position, light.position, t)
	Assert(intensity == light.intensity, "", t)
}
