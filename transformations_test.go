package main

import (
	"math"
	"testing"
)

func TestTranslation(t *testing.T) {
	translation := Translate(5, -3, 2)
	p := Point(-3, 4, 5)

	expected := Point(2, 1, 7)
	actual := translation.MultiplyByTuple(p)

	AssertTupleEqual(expected, actual, t)

	invTranslation := translation.Inverse()
	expectedInverseTrnslatedPoint := Point(-8, 7, 3)

	AssertTupleEqual(expectedInverseTrnslatedPoint, invTranslation.MultiplyByTuple(p), t)
}

func TestTranslationDoesNotAffectVectors(t *testing.T) {
	translation := Translate(5, -3, 2)
	v := Vector(-3, 4, 5)

	AssertTupleEqual(v, translation.MultiplyByTuple(v), t)
}

func TestScalingPoint(t *testing.T) {
	S := Scale(2, 3, 4)
	p := Point(-4, 6, 8)

	expectedScaledPoint := Point(-8, 18, 32)
	actualScaledPoint := S.MultiplyByTuple(p)

	AssertTupleEqual(expectedScaledPoint, actualScaledPoint, t)
}

func TestScalingVector(t *testing.T) {
	S := Scale(2, 3, 4)
	v := Vector(-4, 6, 8)

	expectedScaledVector := Vector(-8, 18, 32)
	actualScaledVector := S.MultiplyByTuple(v)

	AssertTupleEqual(expectedScaledVector, actualScaledVector, t)
}

func TestScalingInverse(t *testing.T) {
	S := Scale(2, 3, 4)
	Sinv := S.Inverse()
	p := Point(-4, 6, 8)

	expectedScaledPoint := Point(-2, 2, 2)
	actualScaledPoint := Sinv.MultiplyByTuple(p)

	AssertTupleEqual(expectedScaledPoint, actualScaledPoint, t)
}

func TestReflectionIsScalingBtNegativeNumber(t *testing.T) {
	R := Scale(-1, 1, 1)
	p := Point(2, 3, 4)

	AssertTupleEqual(Point(-2, 3, 4), R.MultiplyByTuple(p), t)
}

func TestRotatePointAroundXAxis(t *testing.T) {
	p := Point(0, 1, 0)

	halfQuarterRotation := RotateX(math.Pi / 4)
	quarterRotation := RotateX(math.Pi / 2)

	AssertTupleEqual(Point(0, math.Sqrt(2)/2, math.Sqrt(2)/2), halfQuarterRotation.MultiplyByTuple(p), t)
	AssertTupleEqual(Point(0, 0, 1), quarterRotation.MultiplyByTuple(p), t)
}

func TestInverseOfRotationRatatesInTheOppositeDirection(t *testing.T) {
	p := Point(0, 1, 0)

	halfQuarterRotation := RotateX(math.Pi / 4)
	inverseHalfQuarterRotation := halfQuarterRotation.Inverse()

	AssertTupleEqual(Point(0, math.Sqrt(2)/2, -math.Sqrt(2)/2), inverseHalfQuarterRotation.MultiplyByTuple(p), t)
}

func TestRotatePointAroundYAxis(t *testing.T) {
	p := Point(0, 0, 1)

	halfQuarterRotation := RotateY(math.Pi / 4)
	quarterRotation := RotateY(math.Pi / 2)

	AssertTupleEqual(Point(math.Sqrt(2)/2, 0, math.Sqrt(2)/2), halfQuarterRotation.MultiplyByTuple(p), t)
	AssertTupleEqual(Point(1, 0, 0), quarterRotation.MultiplyByTuple(p), t)
}

func TestRotatePointAroundZAxis(t *testing.T) {
	p := Point(0, 1, 0)

	halfQuarterRotation := RotateZ(math.Pi / 4)
	quarterRotation := RotateZ(math.Pi / 2)

	AssertTupleEqual(Point(-math.Sqrt(2)/2, math.Sqrt(2)/2, 0), halfQuarterRotation.MultiplyByTuple(p), t)
	AssertTupleEqual(Point(-1, 0, 0), quarterRotation.MultiplyByTuple(p), t)
}

func TestShearingXtoY(t *testing.T) {
	shearing := Shear(1, 0, 0, 0, 0, 0)
	p := Point(2, 3, 4)

	AssertTupleEqual(Point(5, 3, 4), shearing.MultiplyByTuple(p), t)
}

func TestShearingXtoZ(t *testing.T) {
	shearing := Shear(0, 1, 0, 0, 0, 0)
	p := Point(2, 3, 4)

	AssertTupleEqual(Point(6, 3, 4), shearing.MultiplyByTuple(p), t)
}

func TestShearingYtoX(t *testing.T) {
	shearing := Shear(0, 0, 1, 0, 0, 0)
	p := Point(2, 3, 4)

	AssertTupleEqual(Point(2, 5, 4), shearing.MultiplyByTuple(p), t)
}

func TestShearingYtoZ(t *testing.T) {
	shearing := Shear(0, 0, 0, 1, 0, 0)
	p := Point(2, 3, 4)

	AssertTupleEqual(Point(2, 7, 4), shearing.MultiplyByTuple(p), t)
}

func TestShearingZtoX(t *testing.T) {
	shearing := Shear(0, 0, 0, 0, 1, 0)
	p := Point(2, 3, 4)

	AssertTupleEqual(Point(2, 3, 6), shearing.MultiplyByTuple(p), t)
}

func TestShearingZtoY(t *testing.T) {
	shearing := Shear(0, 0, 0, 0, 0, 1)
	p := Point(2, 3, 4)

	AssertTupleEqual(Point(2, 3, 7), shearing.MultiplyByTuple(p), t)
}

func TestChainingTransformation(t *testing.T) {
	p := Point(1, 0, 1)
	rotationX := RotateX(math.Pi / 2)
	scaling := Scale(5, 5, 5)
	translation := Translate(10, 5, 7)

	pRotated := rotationX.MultiplyByTuple(p)
	AssertTupleEqual(Point(1, -1, 0), pRotated, t)

	pScaled := scaling.MultiplyByTuple(pRotated)
	AssertTupleEqual(Point(5, -5, 0), pScaled, t)

	pTranslated := translation.MultiplyByTuple(pScaled)
	AssertTupleEqual(Point(15, 0, 7), pTranslated, t)

	chainedTransformation := translation.Multiply(scaling).Multiply(rotationX)
	AssertTupleEqual(Point(15, 0, 7), chainedTransformation.MultiplyByTuple(p), t)
}
