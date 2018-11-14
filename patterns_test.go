package main

import "testing"

func TestStripePatternIsConstantInY(t *testing.T) {
	pattern := MakeStripePattern(Color{1, 1, 1}, Color{0, 0, 0}, Identity())

	AssertColorEqual(Color{1, 1, 1}, pattern.ColorAt(Point(0, 0, 0), Identity()), t)
	AssertColorEqual(Color{1, 1, 1}, pattern.ColorAt(Point(0, 1, 0), Identity()), t)
	AssertColorEqual(Color{1, 1, 1}, pattern.ColorAt(Point(0, 2, 0), Identity()), t)
}

func TestStripePatternIsConstantInZ(t *testing.T) {
	pattern := MakeStripePattern(Color{0, 0, 0}, Color{1, 1, 1}, Identity())

	AssertColorEqual(Color{0, 0, 0}, pattern.ColorAt(Point(0, 0, 0), Identity()), t)
	AssertColorEqual(Color{0, 0, 0}, pattern.ColorAt(Point(0, 0, 1), Identity()), t)
	AssertColorEqual(Color{0, 0, 0}, pattern.ColorAt(Point(0, 0, 2), Identity()), t)
}

func TestStripePatternAlternatesInX(t *testing.T) {
	pattern := MakeStripePattern(Color{1, 1, 1}, Color{0, 0, 0}, Identity())

	AssertColorEqual(Color{1, 1, 1}, pattern.ColorAt(Point(0, 0, 0), Identity()), t)
	AssertColorEqual(Color{1, 1, 1}, pattern.ColorAt(Point(0.9, 0, 0), Identity()), t)
	AssertColorEqual(Color{0, 0, 0}, pattern.ColorAt(Point(1, 0, 0), Identity()), t)
	AssertColorEqual(Color{0, 0, 0}, pattern.ColorAt(Point(-0.1, 0, 0), Identity()), t)
	AssertColorEqual(Color{0, 0, 0}, pattern.ColorAt(Point(-1, 0, 0), Identity()), t)
	AssertColorEqual(Color{1, 1, 1}, pattern.ColorAt(Point(-1.1, 0, 0), Identity()), t)

}

func TestStripesWithAnObjectTransformation(t *testing.T) {
	pattern := MakeStripePattern(Color{1, 1, 1}, Color{0, 0, 0}, Identity())

	c := pattern.ColorAt(Point(1.5, 0, 0), Identity().Scale(2, 2, 2))

	AssertColorEqual(Color{1, 1, 1}, c, t)
}

func TestStripeWothAPatternTransformation(t *testing.T) {
	pattern := MakeStripePattern(Color{0, 0, 0}, Color{1, 1, 1}, Identity())

	c := pattern.ColorAt(Point(1.5, 0, 0), Identity().Scale(2, 2, 2))

	AssertColorEqual(Color{0, 0, 0}, c, t)
}

func TestStripesWithBothAnObjectAndAPatternTransformation(t *testing.T) {
	pattern := MakeStripePattern(Color{1, 1, 1}, Color{0, 0, 0}, Identity().Translate(0.5, 0, 0))

	c := pattern.ColorAt(Point(1.5, 0, 0), Identity().Scale(2, 2, 2))

	AssertColorEqual(Color{1, 1, 1}, c, t)
}

func TestGradientLieanrlyInterlpoatesBetweenColors(t *testing.T) {
	pattern := MakeGradientPattern(Color{0, 0, 0}, Color{1, 1, 1}, Identity())

	AssertColorEqual(Color{0, 0, 0}, pattern.ColorAt(Point(0, 0, 0), Identity()), t)
	AssertColorEqual(Color{0.25, 0.25, 0.25}, pattern.ColorAt(Point(0.25, 0, 0), Identity()), t)
	AssertColorEqual(Color{0.5, 0.5, 0.5}, pattern.ColorAt(Point(0.5, 0, 0), Identity()), t)
	AssertColorEqual(Color{0.75, 0.75, 0.75}, pattern.ColorAt(Point(0.75, 0, 0), Identity()), t)
}

func TestRingShouldExtendInBothXandZ(t *testing.T) {
	pattern := MakeRingPattern(Color{0, 0, 0}, Color{1, 1, 1}, Identity())

	AssertColorEqual(Color{0, 0, 0}, pattern.ColorAt(Point(0, 0, 0), Identity()), t)
	AssertColorEqual(Color{1, 1, 1}, pattern.ColorAt(Point(1, 0, 0), Identity()), t)
	AssertColorEqual(Color{1, 1, 1}, pattern.ColorAt(Point(0, 0, 1), Identity()), t)
	AssertColorEqual(Color{1, 1, 1}, pattern.ColorAt(Point(0.708, 0, 0.708), Identity()), t)
}

func TestCheckersShouldRepeatInX(t *testing.T) {
	pattern := MakeCheckersPattern(Color{0, 0, 0}, Color{1, 1, 1}, Identity())

	AssertColorEqual(Color{0, 0, 0}, pattern.ColorAt(Point(0, 0, 0), Identity()), t)
	AssertColorEqual(Color{0, 0, 0}, pattern.ColorAt(Point(0.99, 0, 0), Identity()), t)
	AssertColorEqual(Color{1, 1, 1}, pattern.ColorAt(Point(1.01, 0, 0), Identity()), t)
}

func TestCheckersShouldRepeatInY(t *testing.T) {
	pattern := MakeCheckersPattern(Color{0, 0, 0}, Color{1, 1, 1}, Identity())

	AssertColorEqual(Color{0, 0, 0}, pattern.ColorAt(Point(0, 0, 0), Identity()), t)
	AssertColorEqual(Color{0, 0, 0}, pattern.ColorAt(Point(0, 0.99, 0), Identity()), t)
	AssertColorEqual(Color{1, 1, 1}, pattern.ColorAt(Point(0, 1.01, 0), Identity()), t)
}

func TestCheckersShouldRepeatInZ(t *testing.T) {
	pattern := MakeCheckersPattern(Color{0, 0, 0}, Color{1, 1, 1}, Identity())

	AssertColorEqual(Color{0, 0, 0}, pattern.ColorAt(Point(0, 0, 0), Identity()), t)
	AssertColorEqual(Color{0, 0, 0}, pattern.ColorAt(Point(0, 0, 0.99), Identity()), t)
	AssertColorEqual(Color{1, 1, 1}, pattern.ColorAt(Point(0, 0, 1.01), Identity()), t)
}

func TestBlendPatternBLendsTwoPatterns(t *testing.T) {
	pattern1 := MakeGradientPattern(Color{0, 0, 0}, Color{1, 1, 1}, Identity())
	pattern2 := MakeCheckersPattern(Color{.5, .6, .7}, Color{1, 1, 1}, Identity())

	AssertColorEqual(Color{0, .5, 0}, pattern1.ColorAt(Point(.5, .5, .5), Identity()), t)
	AssertColorEqual(Color{0.5, 0.6, 0.7}, pattern2.ColorAt(Point(.5, .5, .5), Identity()), t)

	blended := MakeBlendPattern(pattern1, pattern2)

	AssertColorEqual(Color{.25, .3, .35}, blended.ColorAt(Point(.5, .5, .5), Identity()), t)
}

func TestPerturbPatterSlightlyModifiesPerturbedPattern(t *testing.T) {
	pattern := MakeGradientPattern(Color{0, 0, 0}, Color{1, 1, 1}, Identity())

	AssertColorEqual(Color{.3, .3, .3}, pattern.ColorAt(Point(.3, 0, 0), Identity()), t)
	AssertColorEqual(Color{.2, .2, .2}, pattern.ColorAt(Point(.2, 0, 0), Identity()), t)
	AssertColorEqual(Color{.1, .1, .1}, pattern.ColorAt(Point(.1, 0, 0), Identity()), t)

	perturb := MakePerturbPattern(pattern)

	AssertColorEqual(Color{.21581, .21581, .21581}, perturb.ColorAt(Point(.3, 0, 0), Identity()), t)
	AssertColorEqual(Color{.13003, .13003, .13003}, perturb.ColorAt(Point(.2, 0, 0), Identity()), t)
	AssertColorEqual(Color{.03965, .03965, .03965}, perturb.ColorAt(Point(.1, 0, 0), Identity()), t)
}
