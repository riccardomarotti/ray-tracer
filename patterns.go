package main

import (
	"math"
)

type Pattern interface {
	StripeAt(Tuple)
}

type StripePattern struct {
	a, b Color
}

func (p StripePattern) StripeAt(point Tuple) Color {
	mod := math.Mod(point.x, 2)
	if mod >= 1 && point.x >= 0 {
		return p.b
	}

	if point.x < 0 && mod >= -1 {
		return p.b
	}

	return p.a
}
