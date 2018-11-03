package main

import "testing"

func TestCreateCanvas(t *testing.T) {
	c := MakeCanvas(10, 20)

	if 10 != c.width {
		t.Errorf("Expected width value was 10, but received %d", c.width)
	}

	if 20 != c.height {
		t.Errorf("Expected height value was 20, but received %d", c.height)
	}

	for _, pixel := range c.pixels {
		assertColorEqual(Color{0, 0, 0}, pixel, t)
	}
}

func TestWritePixel(t *testing.T) {
	c := MakeCanvas(10, 20)

	red := Color{1, 0, 0}

	c.WriteAt(2, 3, red)

	assertColorEqual(red, c.PixelAt(2, 3), t)
}
