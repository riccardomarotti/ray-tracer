package main

import (
	"math"
)

type Pattern interface {
	ColorAt(Tuple) Color
}

type StripePattern struct {
	a, b                       Color
	transform, objectTransform Matrix
}

func MakeStripePattern(colorA, colorB Color, transform, objectTransform Matrix) Pattern {
	return StripePattern{colorA, colorB, transform, objectTransform}
}

func (p StripePattern) colorAt(point Tuple) Color {
	mod := math.Mod(point.x, 2)
	if mod >= 1 && point.x >= 0 {
		return p.b
	}

	if point.x < 0 && mod >= -1 {
		return p.b
	}

	return p.a
}

func (p StripePattern) ColorAt(point Tuple) Color {
	objectPoint := p.objectTransform.Inverse().MultiplyByTuple(point)
	patternPoint := p.transform.Inverse().MultiplyByTuple(objectPoint)

	return p.colorAt(patternPoint)
}

type GradientPattern struct {
	a, b                       Color
	transform, objectTransform Matrix
}

func MakeGradientPattern(colorA, colorB Color, transform, objectTransform Matrix) Pattern {
	return GradientPattern{colorA, colorB, transform, objectTransform}
}

func (p GradientPattern) ColorAt(point Tuple) Color {
	objectPoint := p.objectTransform.Inverse().MultiplyByTuple(point)
	patternPoint := p.transform.Inverse().MultiplyByTuple(objectPoint)

	colorDiff := p.b.Subtract(p.a)
	gradient := patternPoint.x - math.Floor(patternPoint.x)
	return p.a.Add(colorDiff.By(gradient))
}

type RingPattern struct {
	a, b                       Color
	transform, objectTransform Matrix
}

func MakeRingPattern(colorA, colorB Color, transform, objectTransform Matrix) Pattern {
	return RingPattern{colorA, colorB, transform, objectTransform}
}

func (p RingPattern) ColorAt(point Tuple) Color {
	objectPoint := p.objectTransform.Inverse().MultiplyByTuple(point)
	patternPoint := p.transform.Inverse().MultiplyByTuple(objectPoint)

	x := patternPoint.x
	z := patternPoint.z
	if int(math.Floor(math.Sqrt((x*x)+(z*z))))%2 == 0 {
		return p.a
	}
	return p.b
}
