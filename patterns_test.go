package main

import "testing"

func TestStripePatternIsConstantInY(t *testing.T) {
	pattern := StripePattern{Color{1, 1, 1}, Color{0, 0, 0}}

	AssertColorEqual(Color{1, 1, 1}, pattern.ColorAt(Point(0, 0, 0)), t)
	AssertColorEqual(Color{1, 1, 1}, pattern.ColorAt(Point(0, 1, 0)), t)
	AssertColorEqual(Color{1, 1, 1}, pattern.ColorAt(Point(0, 2, 0)), t)
}

func TestStripePatternIsConstantInZ(t *testing.T) {
	pattern := StripePattern{Color{0, 0, 0}, Color{1, 1, 1}}

	AssertColorEqual(Color{0, 0, 0}, pattern.ColorAt(Point(0, 0, 0)), t)
	AssertColorEqual(Color{0, 0, 0}, pattern.ColorAt(Point(0, 0, 1)), t)
	AssertColorEqual(Color{0, 0, 0}, pattern.ColorAt(Point(0, 0, 2)), t)
}

func TestStripePatternAlternatesInX(t *testing.T) {
	pattern := StripePattern{Color{1, 1, 1}, Color{0, 0, 0}}

	AssertColorEqual(Color{1, 1, 1}, pattern.ColorAt(Point(0, 0, 0)), t)
	AssertColorEqual(Color{1, 1, 1}, pattern.ColorAt(Point(0.9, 0, 0)), t)
	AssertColorEqual(Color{0, 0, 0}, pattern.ColorAt(Point(1, 0, 0)), t)
	AssertColorEqual(Color{0, 0, 0}, pattern.ColorAt(Point(-0.1, 0, 0)), t)
	AssertColorEqual(Color{0, 0, 0}, pattern.ColorAt(Point(-1, 0, 0)), t)
	AssertColorEqual(Color{1, 1, 1}, pattern.ColorAt(Point(-1.1, 0, 0)), t)

}
