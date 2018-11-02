package color

// Color represents a color
type Color struct {
	r, g, b float64
}

// Add sums two colors
func (c Color) Add(a Color) Color {
	return Color{c.r + a.r, c.g + a.g, c.b + a.b}
}

// Subtract subtracts two colors
func (c Color) Subtract(a Color) Color {
	return Color{c.r - a.r, c.g - a.g, c.b - a.b}
}

// By multiplies a color by a scalar
func (c Color) By(s float64) Color {
	return Color{c.r * s, c.g * s, c.b * s}
}

// Schur calculates the Schur product of colors
func (c Color) Schur(a Color) Color {
	return Color{c.r * a.r, c.g * a.g, c.b * a.b}
}
