package main

// Canvas represents a canvas
type Canvas struct {
	width, height int
	pixels        []Color
}

// MakeCanvas creates a new canvas of size iXj
func MakeCanvas(i, j int) Canvas {
	return Canvas{i, j, make([]Color, i*j)}
}

// WriteAt write pixel color at (i, j) position
func (c Canvas) WriteAt(i, j int, color Color) {
	c.pixels[c.width*i+j] = color
}

// PixelAt reads pixel color located in (i, j) position
func (c Canvas) PixelAt(i, j int) Color {
	return c.pixels[c.width*i+j]
}
