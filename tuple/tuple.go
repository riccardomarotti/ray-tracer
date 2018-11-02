package tuple

// Tuple represents a tuple
type Tuple struct {
	x, y, z, w float64
}

// IsPoint checks if the tuple is a point
func (t *Tuple) IsPoint() bool {
	return t.w == 1.0
}

// IsVector checks if the tuple is a vector
func (t *Tuple) IsVector() bool {
	return t.w == 0.0
}
