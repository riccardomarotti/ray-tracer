package tuple

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
	return Subtract(Tuple{0, 0, 0, 0}, t)
}

// Point creates a point tuple
func Point(x, y, z float64) Tuple {
	return Tuple{x, y, z, 1}
}

// Vector creates a vector tuple
func Vector(x, y, z float64) Tuple {
	return Tuple{x, y, z, 0}
}

// Sum sums two tuples
func Sum(a1 Tuple, a2 Tuple) Tuple {
	return Tuple{a1.x + a2.x, a1.y + a2.y, a1.z + a2.z, a1.w + a2.w}
}

// Subtract subtracts two tuples
func Subtract(a1 Tuple, a2 Tuple) Tuple {
	return Tuple{a1.x - a2.x, a1.y - a2.y, a1.z - a2.z, a1.w - a2.w}
}

// Multiply multiplies a tuple by a scalar
func Multiply(a Tuple, s float64) Tuple {
	return Tuple{a.x * s, a.y * s, a.z * s, a.w * s}
}

// Divide divides a tuple by a scalar
func Divide(a Tuple, s float64) Tuple {
	return Tuple{a.x / s, a.y / s, a.z / s, a.w / s}
}
