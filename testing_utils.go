package main

import (
	"fmt"
	"math"
	"reflect"
	"testing"
)

func AssertEqual(expected float64, actual float64, t *testing.T) {
	areEqual := areEqual(expected, actual)
	if !areEqual {
		t.Errorf("Expected value was %f, but received %f", expected, actual)
	}
}

func areEqual(expected float64, actual float64) bool {
	return math.Abs(expected-actual) < 1.05*Epsilon

}

func AssertColorEqual(expected Color, actual Color, t *testing.T) {
	errorTemplate := "Color differ.\nExpected was:\n%v\n\nBut was:\n%v\n\n"

	ok := areEqual(expected.r, actual.r) ||
		areEqual(expected.r, actual.r) ||
		areEqual(expected.g, actual.g) ||
		areEqual(expected.b, actual.b)

	if !ok {
		t.Fatal(fmt.Sprintf(errorTemplate, expected, actual))
	}
}

func AssertTupleEqual(expected Tuple, actual Tuple, t *testing.T) {
	errorTemplate := "Tuple differ.\nExpected was:\n%v\n\nBut was:\n%v\n"
	equalX := areEqual(expected.x, actual.x)
	equalY := areEqual(expected.y, actual.y)
	equalZ := areEqual(expected.z, actual.z)
	equalW := areEqual(expected.w, actual.w)
	if !equalX || !equalY || !equalZ || !equalW {
		t.Fatal(fmt.Sprintf(errorTemplate, expected, actual))
	}

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

func AssertTrianglesEqual(expected, actual interface{}, t *testing.T) {
	areEqual := reflect.TypeOf(expected) == reflect.TypeOf(actual)

	p1expected := reflect.ValueOf(&expected).Elem().Elem().FieldByName("p1").String()
	p2expected := reflect.ValueOf(&expected).Elem().Elem().FieldByName("p2").String()
	p3expected := reflect.ValueOf(&expected).Elem().Elem().FieldByName("p3").String()
	p1actual := reflect.ValueOf(&actual).Elem().Elem().FieldByName("p1").String()
	p2actual := reflect.ValueOf(&actual).Elem().Elem().FieldByName("p2").String()
	p3actual := reflect.ValueOf(&actual).Elem().Elem().FieldByName("p3").String()

	areEqual = areEqual && p1expected == p1actual && p2expected == p2actual && p3expected == p3actual

	if !areEqual {
		t.Errorf("Triangles differ\nExpected: %v\nActual:   %v", expected, actual)
	}
}
