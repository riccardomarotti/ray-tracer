package main

import "testing"

func TestStripePatternIsConstantInY(t *testing.T) {
	pattern := MakeStripePattern(Color{1, 1, 1}, Color{0, 0, 0}, Identity())

	AssertColorEqual(Color{1, 1, 1}, pattern.ColorAt(Point(0, 0, 0)), t)
	AssertColorEqual(Color{1, 1, 1}, pattern.ColorAt(Point(0, 1, 0)), t)
	AssertColorEqual(Color{1, 1, 1}, pattern.ColorAt(Point(0, 2, 0)), t)
}

func TestStripePatternIsConstantInZ(t *testing.T) {
	pattern := MakeStripePattern(Color{0, 0, 0}, Color{1, 1, 1}, Identity())

	AssertColorEqual(Color{0, 0, 0}, pattern.ColorAt(Point(0, 0, 0)), t)
	AssertColorEqual(Color{0, 0, 0}, pattern.ColorAt(Point(0, 0, 1)), t)
	AssertColorEqual(Color{0, 0, 0}, pattern.ColorAt(Point(0, 0, 2)), t)
}

func TestStripePatternAlternatesInX(t *testing.T) {
	pattern := MakeStripePattern(Color{1, 1, 1}, Color{0, 0, 0}, Identity())

	AssertColorEqual(Color{1, 1, 1}, pattern.ColorAt(Point(0, 0, 0)), t)
	AssertColorEqual(Color{1, 1, 1}, pattern.ColorAt(Point(0.9, 0, 0)), t)
	AssertColorEqual(Color{0, 0, 0}, pattern.ColorAt(Point(1, 0, 0)), t)
	AssertColorEqual(Color{0, 0, 0}, pattern.ColorAt(Point(-0.1, 0, 0)), t)
	AssertColorEqual(Color{0, 0, 0}, pattern.ColorAt(Point(-1, 0, 0)), t)
	AssertColorEqual(Color{1, 1, 1}, pattern.ColorAt(Point(-1.1, 0, 0)), t)

}

func TestStripesWithAnObjectTransformation(t *testing.T) {
	object := MakeSphere(Identity().Scale(2, 2, 2), DefaultMaterial())
	pattern := MakeStripePattern(Color{1, 1, 1}, Color{0, 0, 0}, Identity())

	c := pattern.ColorAtObject(object, Point(1.5, 0, 0))

	AssertColorEqual(Color{1, 1, 1}, c, t)
}

func TestStripeWothAPatternTransformation(t *testing.T) {
	object := MakeSphere(Identity(), DefaultMaterial())
	pattern := MakeStripePattern(Color{0, 0, 0}, Color{1, 1, 1}, Identity().Scale(2, 2, 2))

	c := pattern.ColorAtObject(object, Point(1.5, 0, 0))

	AssertColorEqual(Color{0, 0, 0}, c, t)
}

func TestStripesWithBothAnObjectAndAPatternTransformation(t *testing.T) {
	object := MakeSphere(Identity().Scale(2, 2, 2), DefaultMaterial())
	pattern := MakeStripePattern(Color{1, 1, 1}, Color{0, 0, 0}, Identity().Translate(0.5, 0, 0))

	c := pattern.ColorAtObject(object, Point(1.5, 0, 0))

	AssertColorEqual(Color{1, 1, 1}, c, t)
}
