package main

import (
	"testing"
)

func TestRayMissesTheCylinder(t *testing.T) {
	cylinder := MakeCylinder(Identity(), DefaultMaterial())

	examples := map[int][2]Tuple{
		0: {Point(1, 0, 0), Vector(0, 1, 0)},
		1: {Point(0, 0, 0), Vector(0, 1, 0)},
		2: {Point(0, 0, -5), Vector(1, 1, 1)},
	}

	for i := 0; i < len(examples); i++ {
		direction := examples[i][1].Normalize()
		ray := Ray{examples[i][0], direction}

		xs := cylinder.Intersection(ray)

		Assert(len(xs) == 0, "Instersection had to be empty", t)
	}
}

func TestRayStrikesTheCylinder(t *testing.T) {
	cylinder := MakeCylinder(Identity(), DefaultMaterial())

	exampleRays := [][2]Tuple{
		{Point(1, 0, -5), Vector(0, 0, 1)},
		{Point(0, 0, -5), Vector(0, 0, 1)},
		{Point(0.5, 0, -5), Vector(0.1, 1, 1)},
	}

	expectedTs := [][2]float64{
		{5, 5},
		{4, 6},
		{6.80798, 7.08872},
	}

	for i := 0; i < len(exampleRays); i++ {
		r := Ray{exampleRays[i][0], exampleRays[i][1].Normalize()}
		xs := cylinder.Intersection(r)

		Assert(len(xs) == 2, "", t)
		AssertEqual(expectedTs[i][0], xs[0].t, t)
		AssertEqual(expectedTs[i][1], xs[1].t, t)
	}
}
