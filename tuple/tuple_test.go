package tuple

import (
	"math"
	"testing"
)

func assertEqual(expected float64, actual float64, t *testing.T) {
	epsilon := 0.00001

	areEqual := math.Abs(expected-actual) < epsilon
	if !areEqual {
		t.Errorf("Expected value was %f, but received %f", expected, actual)
	}
}

func assertTupleEqual(expected Tuple, actual Tuple, t *testing.T) {
	assertEqual(expected.x, actual.x, t)
	assertEqual(expected.y, actual.y, t)
	assertEqual(expected.z, actual.z, t)
	assertEqual(expected.w, actual.w, t)
}

func TestPointTuple(t *testing.T) {
	a := Tuple{4.3, -4.2, 3.1, 1.0}

	assertEqual(4.3, a.x, t)
	assertEqual(-4.2, a.y, t)
	assertEqual(3.1, a.z, t)

	if a.IsPoint() == false {
		t.Errorf("Tuple is expected to be a point.")
	}

	if a.IsVector() == true {
		t.Errorf("Tuple is not expected to be a vector.")
	}
}

func TestVectorTuple(t *testing.T) {
	a := Tuple{4.3, -4.2, 3.1, 0.0}

	assertEqual(4.3, a.x, t)
	assertEqual(-4.2, a.y, t)
	assertEqual(3.1, a.z, t)

	if a.IsVector() == false {
		t.Errorf("Tuple is expected to be a vector.")
	}

	if a.IsPoint() == true {
		t.Errorf("Tuple is not expected to be a point.")
	}
}

func TestPoint(t *testing.T) {
	p := Point(4, -4, 3)

	assertEqual(4.0, p.x, t)
	assertEqual(-4.0, p.y, t)
	assertEqual(3.0, p.z, t)
	assertEqual(1.0, p.w, t)
}

func TestVector(t *testing.T) {
	p := Vector(4, -4, 3)

	assertEqual(4.0, p.x, t)
	assertEqual(-4.0, p.y, t)
	assertEqual(3.0, p.z, t)
	assertEqual(0.0, p.w, t)
}

func TestAddTuples(t *testing.T) {
	p := Point(3, -2, 5)
	v := Vector(-2, 3, 1)

	pTranslated := p.Add(v)

	assertTupleEqual(Tuple{1, 1, 6, 1}, pTranslated, t)
}

func TestSubtractPoints(t *testing.T) {
	p1 := Point(3, 2, 1)
	p2 := Point(5, 6, 7)

	diff := p1.Subtract(p2)

	assertTupleEqual(Vector(-2, -4, -6), diff, t)
}

func TestSubtractAVectorFromAPoint(t *testing.T) {
	p := Point(3, 2, 1)
	v := Vector(5, 6, 7)

	diff := p.Subtract(v)

	assertTupleEqual(Point(-2, -4, -6), diff, t)
}

func TestSubtractVectors(t *testing.T) {
	v1 := Vector(3, 2, 1)
	v2 := Vector(5, 6, 7)

	diff := v1.Subtract(v2)

	assertTupleEqual(Vector(-2, -4, -6), diff, t)
}

func TestNegatingATuple(t *testing.T) {
	a := Tuple{1, -2, 3, -4}

	assertTupleEqual(Tuple{-1, 2, -3, 4}, a.Negate(), t)
}

func TestScalarMultiply(t *testing.T) {
	a := Tuple{1, -2, 3, -4}

	assertTupleEqual(Tuple{0.5, -1, 1.5, -2}, a.Multiply(0.5), t)
}

func TestScalarDivision(t *testing.T) {
	a := Tuple{1, -2, 3, -4}

	assertTupleEqual(Tuple{0.5, -1, 1.5, -2}, a.Divide(2), t)
}

func TestMagnitude(t *testing.T) {
	assertEqual(1, Vector(1, 0, 0).Magnitude(), t)
	assertEqual(1, Vector(0, 1, 0).Magnitude(), t)
	assertEqual(1, Vector(0, 0, 1).Magnitude(), t)

	assertEqual(math.Sqrt(14), Vector(1, 2, 3).Magnitude(), t)
	assertEqual(math.Sqrt(14), Vector(-1, -2, -3).Magnitude(), t)
}

func TestNormalization(t *testing.T) {
	assertTupleEqual(Vector(4, 0, 0).Normalize(), Vector(1, 0, 0), t)
	assertTupleEqual(Vector(1, 2, 3).Normalize(), Vector(1/math.Sqrt(14), 2/math.Sqrt(14), 3/math.Sqrt(14)), t)
	assertEqual(1, Vector(1, 2, 3).Normalize().Magnitude(), t)
}
