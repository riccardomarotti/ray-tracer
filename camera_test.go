package main

import (
	"math"
	"testing"
)

func TestPixelSizeOfHorizontalCanvas(t *testing.T) {
	c := Camera{hsize: 200, vsize: 125, fieldOfView: math.Pi / 2, transform: Identity()}

	AssertEqual(0.01, c.PixelSize(), t)
}

func TestPixelSizeOfVerticalCanvas(t *testing.T) {
	c := Camera{hsize: 125, vsize: 200, fieldOfView: math.Pi / 2, transform: Identity()}

	AssertEqual(0.01, c.PixelSize(), t)
}
