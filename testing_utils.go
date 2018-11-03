package main

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

func assertColorEqual(expected Color, actual Color, t *testing.T) {
	assertEqual(expected.r, actual.r, t)
	assertEqual(expected.g, actual.g, t)
	assertEqual(expected.b, actual.b, t)
}

func assertTupleEqual(expected Tuple, actual Tuple, t *testing.T) {
	assertEqual(expected.x, actual.x, t)
	assertEqual(expected.y, actual.y, t)
	assertEqual(expected.z, actual.z, t)
	assertEqual(expected.w, actual.w, t)
}
