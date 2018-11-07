package main

import (
	"fmt"
	"math"
	"testing"
)

func AssertEqual(expected float64, actual float64, t *testing.T) {
	areEqual := areEqual(expected, actual)
	if !areEqual {
		t.Errorf("Expected value was %f, but received %f", expected, actual)
	}
}

func areEqual(expected float64, actual float64) bool {
	epsilon := 0.00001
	return math.Abs(expected-actual) < epsilon

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

func AssertMatrixEqual(expected Matrix, actual Matrix, t *testing.T) {
	message := fmt.Sprintf("Matrices differ.\nExpected was\n%v\n\nbut received\n%v\n\n", expected, actual)

	Assert(expected.rows == actual.rows, message, t)
	Assert(expected.cols == actual.cols, message, t)

	for i := 0; i < expected.rows; i++ {
		for j := 0; j < expected.cols; j++ {
			if !areEqual(expected.At(i, j), actual.At(i, j)) {
				t.Fatal(message)
			}
		}
	}
}

func Assert(trueCondition bool, message string, t *testing.T) {
	if !trueCondition {
		t.Errorf(message)
	}
}
