package main

import (
	"strings"
	"testing"
)

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

func TestPPMConversionHeader(t *testing.T) {
	c := MakeCanvas(5, 3)

	ppmHeader := strings.Join(strings.Split(c.PPM(), "\n")[:3], "\n")

	expected := "P3\n5 3\n255"

	if expected != ppmHeader {
		t.Errorf("Expected ppm '%s', but was '%s'", expected, ppmHeader)
	}
}

func TestPPMConversionData(t *testing.T) {
	c := MakeCanvas(5, 3)
	c1 := Color{1.5, 0, 0}
	c2 := Color{0, 0.5, 0}
	c3 := Color{-0.5, 0, 1}

	c.WriteAt(0, 0, c1)
	c.WriteAt(2, 1, c2)
	c.WriteAt(4, 2, c3)

	ppm := c.PPM()

	expected := "P3\n5 3\n255\n255 0 0 0 0 0 0 0 0 0 0 0 0 0 0\n0 0 0 0 0 0 0 127 0 0 0 0 0 0 0\n0 0 0 0 0 0 0 0 0 0 0 0 0 0 255\n"

	if expected != ppm {
		t.Errorf("Expected ppm '%s', but was '%s'", expected, ppm)
	}
}
