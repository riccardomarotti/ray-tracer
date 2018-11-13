package main

import "testing"

func TestStripePatternIsConstantInY(t *testing.T) {
	pattern := MakeStripePattern(Color{1, 1, 1}, Color{0, 0, 0}, Identity(), Identity())

	AssertColorEqual(Color{1, 1, 1}, pattern.ColorAt(Point(0, 0, 0)), t)
	AssertColorEqual(Color{1, 1, 1}, pattern.ColorAt(Point(0, 1, 0)), t)
	AssertColorEqual(Color{1, 1, 1}, pattern.ColorAt(Point(0, 2, 0)), t)
}

func TestStripePatternIsConstantInZ(t *testing.T) {
	pattern := MakeStripePattern(Color{0, 0, 0}, Color{1, 1, 1}, Identity(), Identity())

	AssertColorEqual(Color{0, 0, 0}, pattern.ColorAt(Point(0, 0, 0)), t)
	AssertColorEqual(Color{0, 0, 0}, pattern.ColorAt(Point(0, 0, 1)), t)
	AssertColorEqual(Color{0, 0, 0}, pattern.ColorAt(Point(0, 0, 2)), t)
}

func TestStripePatternAlternatesInX(t *testing.T) {
	pattern := MakeStripePattern(Color{1, 1, 1}, Color{0, 0, 0}, Identity(), Identity())

	AssertColorEqual(Color{1, 1, 1}, pattern.ColorAt(Point(0, 0, 0)), t)
	AssertColorEqual(Color{1, 1, 1}, pattern.ColorAt(Point(0.9, 0, 0)), t)
	AssertColorEqual(Color{0, 0, 0}, pattern.ColorAt(Point(1, 0, 0)), t)
	AssertColorEqual(Color{0, 0, 0}, pattern.ColorAt(Point(-0.1, 0, 0)), t)
	AssertColorEqual(Color{0, 0, 0}, pattern.ColorAt(Point(-1, 0, 0)), t)
	AssertColorEqual(Color{1, 1, 1}, pattern.ColorAt(Point(-1.1, 0, 0)), t)

}

func TestStripesWithAnObjectTransformation(t *testing.T) {
	pattern := MakeStripePattern(Color{1, 1, 1}, Color{0, 0, 0}, Identity(), Identity().Scale(2, 2, 2))

	c := pattern.ColorAt(Point(1.5, 0, 0))

	AssertColorEqual(Color{1, 1, 1}, c, t)
}

func TestStripeWothAPatternTransformation(t *testing.T) {
	pattern := MakeStripePattern(Color{0, 0, 0}, Color{1, 1, 1}, Identity(), Identity().Scale(2, 2, 2))

	c := pattern.ColorAt(Point(1.5, 0, 0))

	AssertColorEqual(Color{0, 0, 0}, c, t)
}

func TestStripesWithBothAnObjectAndAPatternTransformation(t *testing.T) {
	pattern := MakeStripePattern(Color{1, 1, 1}, Color{0, 0, 0}, Identity().Translate(0.5, 0, 0), Identity().Scale(2, 2, 2))

	c := pattern.ColorAt(Point(1.5, 0, 0))

	AssertColorEqual(Color{1, 1, 1}, c, t)
}

func TestGradientLieanrlyInterlpoatesBetweenColors(t *testing.T) {
	pattern := MakeGradientPattern(Color{0, 0, 0}, Color{1, 1, 1}, Identity(), Identity())

	AssertColorEqual(Color{0, 0, 0}, pattern.ColorAt(Point(0, 0, 0)), t)
	AssertColorEqual(Color{0.25, 0.25, 0.25}, pattern.ColorAt(Point(0.25, 0, 0)), t)
	AssertColorEqual(Color{0.5, 0.5, 0.5}, pattern.ColorAt(Point(0.5, 0, 0)), t)
	AssertColorEqual(Color{0.75, 0.75, 0.75}, pattern.ColorAt(Point(0.75, 0, 0)), t)
}
