package main

import (
	"math"
)

type Camera struct {
	hsize, vsize, fieldOfView float64
	transform                 Matrix
}

func (c Camera) PixelSize() float64 {
	halfView := math.Tan(c.fieldOfView / 2)
	aspectRatio := c.hsize / c.vsize

	canvasWidth := 2 * halfView / c.hsize

	if aspectRatio < 1 {
		canvasWidth = canvasWidth * aspectRatio
	}

	return canvasWidth
}
