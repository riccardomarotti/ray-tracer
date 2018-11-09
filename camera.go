package main

import (
	"math"
)

type Camera struct {
	hsize, vsize, fieldOfView float64
	transform                 Matrix
}

func (c Camera) halfWidth() float64 {
	aspectRatio := c.hsize / c.vsize
	halfWidth := math.Tan(c.fieldOfView / 2)

	if aspectRatio < 1 {
		halfWidth = halfWidth * aspectRatio
	}

	return halfWidth
}

func (c Camera) halfHeight() float64 {
	aspectRatio := c.hsize / c.vsize
	halfHeight := math.Tan(c.fieldOfView / 2)

	if aspectRatio >= 1 {
		halfHeight = halfHeight / aspectRatio
	}

	return halfHeight
}

func (c Camera) PixelSize() float64 {
	return 2 * c.halfWidth() / c.hsize
}

func (c Camera) RayForPixel(x, y float64) Ray {
	pixelSize := c.PixelSize()

	xOffset := (x + 0.5) * pixelSize
	yOffset := (y + 0.5) * pixelSize

	worldX := c.halfWidth() - xOffset
	worldY := c.halfHeight() - yOffset

	inverse := c.transform.Inverse()
	pixel := inverse.MultiplyByTuple(Point(worldX, worldY, -1))
	origin := inverse.MultiplyByTuple(Point(0, 0, 0))
	direction := (pixel.Subtract(origin)).Normalize()

	return Ray{origin, direction}
}
