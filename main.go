package main

import (
	"fmt"
	"sync"
)

func main() {
	material := MakeMaterial()
	material.color = Color{0.2, 1, 1}
	s := MakeSphere(Identity().Scale(2, 1, 2), material)

	lightPosition := Point(-10, 10, -10)
	lightColor := Color{1, 1, 1}
	light := PointLight{lightPosition, lightColor}

	canvasPixels := 1000
	c := MakeCanvas(canvasPixels, canvasPixels)

	rayOrigin := Point(0, 0, -10)
	wallZ := 20.0
	wallSize := 14.0
	half := wallSize / 2
	pixelSize := wallSize / float64(canvasPixels)

	var wg sync.WaitGroup

	for y := 0; y < canvasPixels; y++ {
		worldY := half - pixelSize*float64(y)
		for x := 0; x < canvasPixels; x++ {
			worldX := -half + pixelSize*float64(x)
			wg.Add(1)
			go writePixel(worldX, worldY, wallZ, rayOrigin, s, c, x, y, light, &wg)
		}
	}

	wg.Wait()

	fmt.Printf(c.PPM())
}

func writePixel(worldX, worldY, wallZ float64, rayOrigin Tuple, s Object, c Canvas, x, y int, light PointLight, wg *sync.WaitGroup) {
	defer wg.Done()

	position := Point(worldX, worldY, wallZ)
	r := Ray{rayOrigin, position.Subtract(rayOrigin).Normalize()}
	xs := r.Intersection(s)

	intersection := Hit(xs)
	hit := intersection != Intersection{}
	if hit {
		point := r.Position(intersection.t)
		normal := intersection.object.NormalAt(point)
		eye := r.direction.Multiply(-1)
		color := intersection.object.Material().Lighting(light, point, eye, normal)
		c.WriteAt(x, y, color)
	}
}
