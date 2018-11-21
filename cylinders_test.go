package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestRayMissesTheCylinder(t *testing.T) {
	cylinder := MakeInfiniteCylinder(Identity(), DefaultMaterial())

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
	cylinder := MakeInfiniteCylinder(Identity(), DefaultMaterial())

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

func TestNormalVectorOnACylinder(t *testing.T) {
	cylinder := MakeInfiniteCylinder(Identity(), DefaultMaterial())

	examples := [][2]Tuple{
		{Point(1, 0, 0), Vector(1, 0, 0)},
		{Point(0, 5, -1), Vector(0, 0, -1)},
		{Point(0, -2, 1), Vector(0, 0, 1)},
		{Point(-1, 1, 0), Vector(-1, 0, 0)},
	}

	for i := 0; i < len(examples); i++ {
		n := cylinder.NormalAt(examples[i][0])

		AssertTupleEqual(examples[i][1], n, t)
	}
}

func TestIntersectingAConstrainedCylinder(t *testing.T) {
	cylidner := MakeCylinder(Identity(), DefaultMaterial(), 1, 2, false)

	examples := [][2]Tuple{
		{Point(0, 1.5, 0), Vector(0.1, 1, 0)},
		{Point(0, 3, -5), Vector(0, 0, 1)},
		{Point(0, 0, -5), Vector(0, 0, 1)},
		{Point(0, 2, -5), Vector(0, 0, 1)},
		{Point(0, 1, -5), Vector(0, 0, 1)},
		{Point(0, 1.5, -2), Vector(0, 0, 1)},
	}

	expectedCounts := []int{0, 0, 0, 0, 0, 2}

	for i := 0; i < len(examples); i++ {
		r := Ray{examples[i][0], examples[i][1].Normalize()}
		xs := cylidner.Intersection(r)

		expectedCount := expectedCounts[i]
		actualCount := len(xs)

		Assert(expectedCount == actualCount, fmt.Sprintf("Expected count: %d, but was: %d", expectedCount, actualCount), t)
	}
}

func TestIntersectingTheCapsOfAClosedCylinder(t *testing.T) {
	cylidner := MakeClosedCylinder(Identity(), DefaultMaterial(), 1, 2)

	examples := [][2]Tuple{
		{Point(0, 3, 0), Vector(0, -1, 0)},
		{Point(0, 3, -2), Vector(0, -1, 2)},
		{Point(0, 4, -2), Vector(0, -1, 1)},
		{Point(0, 0, -2), Vector(0, 1, 2)},
		{Point(0, -1, -2), Vector(0, 1, 1)},
	}

	expectedCounts := []int{2, 2, 2, 2, 2}

	for i := 0; i < len(examples); i++ {
		r := Ray{examples[i][0], examples[i][1].Normalize()}
		xs := cylidner.Intersection(r)

		expectedCount := expectedCounts[i]
		actualCount := len(xs)

		Assert(expectedCount == actualCount, fmt.Sprintf("Expected count: %d, but was: %d", expectedCount, actualCount), t)
	}
}

func TestNormalVectorOnACylindersEndCap(t *testing.T) {
	cylidner := MakeClosedCylinder(Identity(), DefaultMaterial(), 1, 2)

	examples := [][2]Tuple{
		{Point(0, 1, 0), Vector(0, -1, 0)},
		{Point(0.5, 1, 0), Vector(0, -1, 0)},
		{Point(0, 1, 0.5), Vector(0, -1, 0)},
		{Point(0, 2, 0), Vector(0, 1, 0)},
		{Point(0.5, 2, 0), Vector(0, 1, 0)},
		{Point(0, 2, 0.5), Vector(0, 1, 0)},
	}

	for i := 0; i < len(examples); i++ {
		n := cylidner.NormalAt(examples[i][0])

		AssertTupleEqual(examples[i][1], n, t)
	}
}

func TestCylinderInAGroup(t *testing.T) {
	g := MakeGroup(Identity())
	c := MakeCylinderInGroup(Identity(), DefaultMaterial(), 0, 1, false, g)

	Assert(reflect.DeepEqual(c, g.children[0]), "", t)
	Assert(c.Parent() == g, "", t)
}

func TestCylinderMaterial(t *testing.T) {
	m := DefaultMaterial()
	c := MakeInfiniteCylinder(Identity(), m)

	Assert(m == c.Material(), "", t)
}
