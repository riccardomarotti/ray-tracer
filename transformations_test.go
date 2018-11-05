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

func TestScalingPoint(t *testing.T) {
	S := Scaling(2, 3, 4)
	p := Point(-4, 6, 8)

	expectedScaledPoint := Point(-8, 18, 32)
	actualScaledPoint := S.MultiplyByTuple(p)

	AssertTupleEqual(expectedScaledPoint, actualScaledPoint, t)
}

func TestScalingVector(t *testing.T) {
	S := Scaling(2, 3, 4)
	v := Vector(-4, 6, 8)

	expectedScaledVector := Vector(-8, 18, 32)
	actualScaledVector := S.MultiplyByTuple(v)

	AssertTupleEqual(expectedScaledVector, actualScaledVector, t)
}

func TestScalingInverse(t *testing.T) {
	S := Scaling(2, 3, 4)
	Sinv := S.Inverse()
	p := Point(-4, 6, 8)

	expectedScaledPoint := Point(-2, 2, 2)
	actualScaledPoint := Sinv.MultiplyByTuple(p)

	AssertTupleEqual(expectedScaledPoint, actualScaledPoint, t)
}
