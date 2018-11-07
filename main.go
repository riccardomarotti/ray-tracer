package main

import (
	"fmt"
)

func main() {
	s := MakeSphere(Identity().Scale(0.5, 1, 0.5))
	canvasPixels := 100
	c := MakeCanvas(canvasPixels, canvasPixels)
	color := Color{0, .7, .3}

	rayOrigin := Point(0, 0, -10)
	wallZ := 10.0
	wallSize := 7.0
	half := wallSize / 2
	pixelSize := wallSize / 100

	for y := 0; y < canvasPixels; y++ {
		worldY := half - pixelSize*float64(y)
		for x := 0; x < canvasPixels; x++ {
			worldX := -half + pixelSize*float64(x)
			position := Point(worldX, worldY, wallZ)
			r := Ray{rayOrigin, position.Subtract(rayOrigin).Normalize()}
			xs := r.Intersection(s)

			hit := Hit(xs) != Intersection{}
			if hit {
				c.WriteAt(x, y, color)
			}
		}
	}

	fmt.Printf(c.PPM())
}
