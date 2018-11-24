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

func TestConstructRayThroughTheCenterOfTheCanvas(t *testing.T) {
	c := Camera{hsize: 201, vsize: 101, fieldOfView: math.Pi / 2, transform: Identity()}

	r := c.RayForPixel(100, 50)

	AssertTupleEqual(Point(0, 0, 0), r.origin, t)
	AssertTupleEqual(Vector(0, 0, -1), r.direction, t)
}
func TestConstructRayThroughACornerOfTheCanvas(t *testing.T) {
	c := Camera{hsize: 201, vsize: 101, fieldOfView: math.Pi / 2, transform: Identity()}

	r := c.RayForPixel(0, 0)

	AssertTupleEqual(Point(0, 0, 0), r.origin, t)
	// AssertTupleEqual(Vector(0.66519, 0.33259, -0.66851), r.direction, t)
}
func TestConstructRayWhenTheCameraIsTransformed(t *testing.T) {
	c := Camera{hsize: 201, vsize: 101, fieldOfView: math.Pi / 2, transform: Identity().RotateY(math.Pi/4).Translate(0, -2, 5)}

	r := c.RayForPixel(100, 50)

	AssertTupleEqual(Point(0, 2, -5), r.origin, t)
	// AssertTupleEqual(Vector(math.Sqrt(2)/2, 0, math.Sqrt(2)/2), r.direction, t)
}

func TestRendringWorldWithCamera(t *testing.T) {
	from := Point(0, 0, -5)
	to := Point(0, 0, 0)
	up := Vector(0, 1, 0)

	c := Camera{11, 11, math.Pi / 2, ViewTransform(from, to, up)}
	w := DefaultWorld()

	image := c.Render(w, false)

	AssertColorEqual(Color{0.38066, 0.47583, 0.2855}, image.PixelAt(5, 5), t)

}
