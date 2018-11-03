package main

import (
	"fmt"
	"strings"
)

// Canvas represents a canvas
type Canvas struct {
	width, height int
	pixels        []Color
}

// MakeCanvas creates a new canvas of size iXj
func MakeCanvas(i, j int) Canvas {
	return Canvas{i, j, make([]Color, i*j)}
}

func (c Canvas) flatten(i, j int) int {
	return c.width*j + i
}

// WriteAt write pixel color at (i, j) position
func (c Canvas) WriteAt(i, j int, color Color) {
	c.pixels[c.flatten(i, j)] = color
}

// PixelAt reads pixel color located in (i, j) position
func (c Canvas) PixelAt(i, j int) Color {
	return c.pixels[c.flatten(i, j)]
}

func toIntColorValue(value float64) int {
	if value < 0 {
		return 0
	}
	intValue := int(value * 255)
	if intValue >= 255 {
		return 255
	}
	return intValue
}

// PPM resturns the PPM text representing the canvas
func (c Canvas) PPM() string {
	maxColorValue := 255
	var values string
	for j := 0; j < c.height; j++ {
		for i := 0; i < c.width; i++ {
			p := c.PixelAt(i, j)
			red := toIntColorValue(p.r)
			green := toIntColorValue(p.g)
			blue := toIntColorValue(p.b)

			values += fmt.Sprintf("%d %d %d ", red, green, blue)
		}
		values = strings.Trim(values, " ") + "\n"

	}
	return fmt.Sprintf("P3\n%d %d\n%d\n%s", c.width, c.height, maxColorValue, values)
}
