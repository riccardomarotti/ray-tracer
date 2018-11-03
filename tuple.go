package main

import (
	"math"
)

// Tuple represents a tuple
type Tuple struct {
	x, y, z, w float64
}

// IsPoint checks if the tuple is a point
func (t Tuple) IsPoint() bool {
	return t.w == 1.0
}

// IsVector checks if the tuple is a vector
func (t Tuple) IsVector() bool {
	return t.w == 0.0
}

// Negate negates a tuple
func (t Tuple) Negate() Tuple {
	return Tuple{0, 0, 0, 0}.Subtract(t)
}

// Point creates a point tuple
func Point(x, y, z float64) Tuple {
	return Tuple{x, y, z, 1}
}

// Vector creates a vector tuple
func Vector(x, y, z float64) Tuple {
	return Tuple{x, y, z, 0}
}

// Add sums two tuples
func (t Tuple) Add(a Tuple) Tuple {
	return Tuple{t.x + a.x, t.y + a.y, t.z + a.z, t.w + a.w}
}

// Subtract subtracts two tuples
func (t Tuple) Subtract(a Tuple) Tuple {
	return Tuple{t.x - a.x, t.y - a.y, t.z - a.z, t.w - a.w}
}

// Multiply multiplies a tuple by a scalar
func (t Tuple) Multiply(s float64) Tuple {
	return Tuple{t.x * s, t.y * s, t.z * s, t.w * s}
}

// Divide divides a tuple by a scalar
func (t Tuple) Divide(s float64) Tuple {
	return Tuple{t.x / s, t.y / s, t.z / s, t.w / s}
}

// Magnitude calculates a tuple's magnitude
func (t Tuple) Magnitude() float64 {
	return math.Sqrt(t.x*t.x + t.y*t.y + t.z*t.z)
}

// Normalize normalizes a tuple
func (t Tuple) Normalize() Tuple {
	return Tuple{t.x / t.Magnitude(),
		t.y / t.Magnitude(),
		t.z / t.Magnitude(),
		t.w / t.Magnitude()}
}

// Dot calculates dot product
func (t Tuple) Dot(a Tuple) float64 {
	return t.x*a.x + t.y*a.y + t.z*a.z + t.w*a.w
}

// Cross calculates cross product
// It works only on three dimensional Tuples!
func (t Tuple) Cross(a Tuple) Tuple {
	return Vector(t.y*a.z-t.z*a.y,
		t.z*a.x-t.x*a.z,
		t.x*a.y-t.y*a.x)
}
