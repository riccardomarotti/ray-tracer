package main

import (
	"log"
	"math"
	"os"
	"sync"
)

type Camera struct {
	hsize, vsize int
	fieldOfView  float64
	transform    Matrix
}

func (c Camera) halfWidthAndHeight() (float64, float64) {
	aspectRatio := float64(c.hsize) / float64(c.vsize)
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
	return 2 * halfWidth / float64(c.hsize)
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

func (c Camera) Render(w World, showProgress bool) Canvas {
	image := MakeCanvas(c.hsize, c.vsize)

	totalPixelCount := c.vsize * c.hsize
	currentPixelCount := 0
	var l *log.Logger
	if showProgress {
		l = log.New(os.Stderr, "", 0)
	} else {
		logfile, _ := os.Open(os.DevNull)
		defer logfile.Close()
		l = log.New(logfile, "", 0)
	}
	l.Printf("%d / %d", currentPixelCount, totalPixelCount)
	var wg sync.WaitGroup
	for y := 0; y < c.vsize; y++ {
		for x := 0; x < c.hsize; x++ {
			wg.Add(1)
			go func(x, y int) {
				defer func() {
					wg.Done()
					currentPixelCount++
					l.Printf("%d / %d", currentPixelCount, totalPixelCount)
				}()
				ray := c.RayForPixel(float64(x), float64(y))
				color := w.ColorAt(ray, 5)
				image.WriteAt(x, y, color)
			}(x, y)
		}
	}
	wg.Wait()
	return image
}
