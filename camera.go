package main

import (
	"math"
)

type Camera struct {
	hsize, vsize, fieldOfView float64
	transform                 Matrix
}

func (c Camera) halfWidthAndHeight() (float64, float64) {
	aspectRatio := c.hsize / c.vsize
	halfWidth := math.Tan(c.fieldOfView / 2)
	halfHeight := halfWidth

	if aspectRatio < 1 {
		halfWidth = halfWidth * aspectRatio
	} else {
		halfHeight = halfHeight / aspectRatio
	}

	return halfWidth, halfHeight
}

func (c Camera) PixelSize() float64 {
	halfWidth, _ := c.halfWidthAndHeight()
	return 2 * halfWidth / c.hsize
}

func (c Camera) RayForPixel(x, y float64) Ray {
	pixelSize := c.PixelSize()

	xOffset := (x + 0.5) * pixelSize
	yOffset := (y + 0.5) * pixelSize

	halfWidth, halfHeight := c.halfWidthAndHeight()
	worldX := halfWidth - xOffset
	worldY := halfHeight - yOffset

	inverse := c.transform.Inverse()
	pixel := inverse.MultiplyByTuple(Point(worldX, worldY, -1))
	origin := inverse.MultiplyByTuple(Point(0, 0, 0))
	direction := (pixel.Subtract(origin)).Normalize()

	return Ray{origin, direction}
}
