package tuple

// IsPoint checks if the passed tuple is a point
func IsPoint(tuple []float32) bool {
	return tuple[len(tuple)-1] == 1.0
}

// IsVector checks if the passed tuple is a vector
func IsVector(tuple []float32) bool {
	return tuple[len(tuple)-1] == 0.0
}
