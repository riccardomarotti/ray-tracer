package main

import (
	"math"
)

type Pattern interface {
	ColorAt(Tuple, Matrix) Color
}

type StripePattern struct {
	a, b      Color
	transform Matrix
}

func MakeStripePattern(colorA, colorB Color, transform Matrix) Pattern {
	return StripePattern{colorA, colorB, transform}
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

func (p StripePattern) ColorAt(point Tuple, objectTransform Matrix) Color {
	objectPoint := objectTransform.Inverse().MultiplyByTuple(point)
	patternPoint := p.transform.Inverse().MultiplyByTuple(objectPoint)

	return p.colorAt(patternPoint)
}

type GradientPattern struct {
	a, b      Color
	transform Matrix
}

func MakeGradientPattern(colorA, colorB Color, transform Matrix) Pattern {
	return GradientPattern{colorA, colorB, transform}
}

func (p GradientPattern) ColorAt(point Tuple, objectTransform Matrix) Color {
	objectPoint := objectTransform.Inverse().MultiplyByTuple(point)
	patternPoint := p.transform.Inverse().MultiplyByTuple(objectPoint)

	colorDiff := p.b.Subtract(p.a)
	var gradient float64
	if patternPoint.x >= 0 {
		gradient = patternPoint.x - math.Floor(patternPoint.x)
	} else {
		gradient = -patternPoint.x + math.Floor(-patternPoint.x)
	}

	return p.a.Add(colorDiff.By(gradient))
}

type RingPattern struct {
	a, b      Color
	transform Matrix
}

func MakeRingPattern(colorA, colorB Color, transform Matrix) Pattern {
	return RingPattern{colorA, colorB, transform}
}

func (p RingPattern) ColorAt(point Tuple, objectTransform Matrix) Color {
	objectPoint := objectTransform.Inverse().MultiplyByTuple(point)
	patternPoint := p.transform.Inverse().MultiplyByTuple(objectPoint)

	x := patternPoint.x
	z := patternPoint.z
	if int(math.Floor(math.Sqrt((x*x)+(z*z))))%2 == 0 {
		return p.a
	}
	return p.b
}

type CheckersPattern struct {
	a, b      Color
	transform Matrix
}

func MakeCheckersPattern(colorA, colorB Color, transform Matrix) Pattern {
	return CheckersPattern{colorA, colorB, transform}
}

func (p CheckersPattern) ColorAt(point Tuple, objectTransform Matrix) Color {
	objectPoint := objectTransform.Inverse().MultiplyByTuple(point)
	patternPoint := p.transform.Inverse().MultiplyByTuple(objectPoint)

	x := int(math.Floor(patternPoint.x))
	y := int(math.Floor(patternPoint.y))
	z := int(math.Floor(patternPoint.z))

	if (x+y+z)%2 == 0 {
		return p.a
	}
	return p.b
}
