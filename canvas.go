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

func truncateLineAt(line string, i int) string {
	if len(line) > i {
		lastSpaceIndex := strings.LastIndex(line[:i+1], " ")
		if lastSpaceIndex != -1 {
			return line[:lastSpaceIndex] + "\n" + line[lastSpaceIndex+1:]
		}
	}
	return line
}

// PPM resturns the PPM text representing the canvas
func (c Canvas) PPM() string {
	maxColorValue := 255
	var values string
	for j := 0; j < c.height; j++ {
		var currentLine string
		for i := 0; i < c.width; i++ {
			p := c.PixelAt(i, j)
			red := toIntColorValue(p.r)
			green := toIntColorValue(p.g)
			blue := toIntColorValue(p.b)

			currentLine += fmt.Sprintf("%d %d %d ", red, green, blue)
		}
		currentLine = strings.Trim(currentLine, " ")
		currentLine = truncateLineAt(currentLine, 70)
		values += currentLine + "\n"
	}
	return fmt.Sprintf("P3\n%d %d\n%d\n%s", c.width, c.height, maxColorValue, values)
}
