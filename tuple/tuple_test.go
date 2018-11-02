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
