package main

import (
	"math"
)

type Pattern interface {
	ColorAt(Tuple) Color
	ColorAtObject(Object, Tuple) Color
}

type StripePattern struct {
	a, b      Color
	transform Matrix
}

func MakeStripePattern(colorA, colorB Color, transform Matrix) Pattern {
	return StripePattern{colorA, colorB, transform}
}

func (p StripePattern) ColorAt(point Tuple) Color {
	mod := math.Mod(point.x, 2)
	if mod >= 1 && point.x >= 0 {
		return p.b
	}

	if point.x < 0 && mod >= -1 {
		return p.b
	}

	return p.a
}

func (p StripePattern) ColorAtObject(object Object, point Tuple) Color {
	objectPoint := object.Transform().Inverse().MultiplyByTuple(point)
	patternPoint := p.transform.Inverse().MultiplyByTuple(objectPoint)

	return p.ColorAt(patternPoint)
}
