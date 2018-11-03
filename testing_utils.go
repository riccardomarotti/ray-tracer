package main

import (
	"math"
	"testing"
)

func AssertEqual(expected float64, actual float64, t *testing.T) {
	epsilon := 0.00001

	areEqual := math.Abs(expected-actual) < epsilon
	if !areEqual {
		t.Errorf("Expected value was %f, but received %f", expected, actual)
	}
}

func AssertColorEqual(expected Color, actual Color, t *testing.T) {
	AssertEqual(expected.r, actual.r, t)
	AssertEqual(expected.g, actual.g, t)
	AssertEqual(expected.b, actual.b, t)
}

func AssertTupleEqual(expected Tuple, actual Tuple, t *testing.T) {
	AssertEqual(expected.x, actual.x, t)
	AssertEqual(expected.y, actual.y, t)
	AssertEqual(expected.z, actual.z, t)
	AssertEqual(expected.w, actual.w, t)
}

func Assert(trueCondition bool, message string, t *testing.T) {
	if !trueCondition {
		t.Errorf(message)
	}
}
