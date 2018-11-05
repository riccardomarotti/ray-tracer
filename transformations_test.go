package main

import "testing"

func TestTranslation(t *testing.T) {
	translation := Translation(5, -3, 2)
	p := Point(-3, 4, 5)

	expected := Point(2, 1, 7)
	actual := translation.MultiplyByTuple(p)

	AssertTupleEqual(expected, actual, t)

	invTranslation := translation.Inverse()
	expectedInverseTrnslatedPoint := Point(-8, 7, 3)

	AssertTupleEqual(expectedInverseTrnslatedPoint, invTranslation.MultiplyByTuple(p), t)
}

func TestTranslationDoesNotAffectVectors(t *testing.T) {
	translation := Translation(5, -3, 2)
	v := Vector(-3, 4, 5)

	AssertTupleEqual(v, translation.MultiplyByTuple(v), t)
}
