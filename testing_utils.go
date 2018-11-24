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
	if reflect.TypeOf(expected) != reflect.TypeOf(actual) {
		t.Errorf("Triangles differ\nExpected: %v\nActual:   %v", expected, actual)
	}

	p1expected := parameterFromInterface("p1", &expected)
	p2expected := parameterFromInterface("p2", &expected)
	p3expected := parameterFromInterface("p3", &expected)

	p1actual := parameterFromInterface("p1", &actual)
	p2actual := parameterFromInterface("p2", &actual)
	p3actual := parameterFromInterface("p3", &actual)

	n1expected := parameterFromInterface("n1", &expected)
	n2expected := parameterFromInterface("n2", &expected)
	n3expected := parameterFromInterface("n3", &expected)

	n1actual := parameterFromInterface("n1", &actual)
	n2actual := parameterFromInterface("n2", &actual)
	n3actual := parameterFromInterface("n3", &actual)

	AssertTupleEqual(p1expected, p1actual, t)
	AssertTupleEqual(p2expected, p2actual, t)
	AssertTupleEqual(p3expected, p3actual, t)

	AssertTupleEqual(n1expected, n1actual, t)
	AssertTupleEqual(n2expected, n2actual, t)
	AssertTupleEqual(n3expected, n3actual, t)

}

func parameterFromInterface(parameterName string, i *interface{}) Tuple {
	x := reflect.ValueOf(i).Elem().Elem().FieldByName(parameterName).FieldByName("x").Float()
	y := reflect.ValueOf(i).Elem().Elem().FieldByName(parameterName).FieldByName("y").Float()
	z := reflect.ValueOf(i).Elem().Elem().FieldByName(parameterName).FieldByName("z").Float()

	return Point(x, y, z)
}
