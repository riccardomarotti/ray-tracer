package main

import (
	"fmt"
	"testing"
)

func TestIntersectingAConeWithARay(t *testing.T) {
	cone := MakeInfiniteCone(Identity(), DefaultMaterial())

	examples := [][2]Tuple{
		{Point(0, 0, -5), Vector(0, 0, 1)},
		{Point(0, 0, -5), Vector(1, 1, 1)},
		{Point(1, 1, -5), Vector(-0.5, -1, 1)},
	}

	expectedTs := [][2]float64{
		{5, 5},
		{8.66025, 8.66025},
		{4.55006, 49.44994},
	}

	for i := 0; i < len(examples); i++ {
		r := Ray{examples[i][0], examples[i][1].Normalize()}
		xs := cone.Intersection(r)

		Assert(len(xs) == 2, "There had to be two intersections", t)
		AssertEqual(expectedTs[i][0], xs[0].t, t)
		AssertEqual(expectedTs[i][1], xs[1].t, t)
	}
}

func TestIntersectingAConeWithARayParallelToOneOfItsHalves(t *testing.T) {
	cone := MakeInfiniteCone(Identity(), DefaultMaterial())

	r := Ray{Point(0, 0, -1), Vector(0, 1, 1).Normalize()}
	xs := cone.Intersection(r)

	Assert(len(xs) == 1, "", t)
	AssertEqual(0.35355, xs[0].t, t)
}

func TestIntersectingAConesEndCaps(t *testing.T) {
	cone := MakeClosedCone(Identity(), DefaultMaterial(), -.5, .5)

	examples := [][2]Tuple{
		{Point(0, 0, -5), Vector(0, 1, 0)},
		{Point(0, 0, -0.25), Vector(0, 1, 1)},
		{Point(0, 0, -0.25), Vector(0, 1, 0)},
	}

	expectedCounts := []int{0, 2, 4}

	for i := 0; i < len(examples); i++ {
		r := Ray{examples[i][0], examples[i][1].Normalize()}
		xs := cone.Intersection(r)

		actualCount := len(xs)
		Assert(actualCount == expectedCounts[i], fmt.Sprintf("Expected number of intersection: %d, but was %d", expectedCounts[i], actualCount), t)
	}
}
