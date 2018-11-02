package color

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
func TestAddColors(t *testing.T) {
	c1 := Color{0.9, 0.6, 0.75}
	c2 := Color{0.7, 0.1, 0.25}

	assertColorEqual(Color{1.6, 0.7, 1.0}, c1.Add(c2), t)
}

func TestSubtractColors(t *testing.T) {
	c1 := Color{0.9, 0.6, 0.75}
	c2 := Color{0.7, 0.1, 0.25}

	assertColorEqual(Color{0.2, 0.5, 0.5}, c1.Subtract(c2), t)
}

func TestMultiplyColorByScalar(t *testing.T) {
	c := Color{0.2, 0.3, 0.4}

	assertColorEqual(Color{0.4, 0.6, 0.8}, c.By(2), t)
}

func TestColorMultiplication(t *testing.T) {
	c1 := Color{1, 0.2, 0.4}
	c2 := Color{0.9, 1, 0.1}

	assertColorEqual(Color{0.9, 0.2, 0.04}, c1.Schur(c2), t)
}
