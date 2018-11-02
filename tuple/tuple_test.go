package tuple

import (
	"testing"
)

func TestPointTuple(t *testing.T) {
	a := []float32{4.3, -4.2, 3.1, 1.0}

	if IsPoint(a) != true {
		t.Errorf("Tuple %v is expected to be a point.", a)
	}

	if IsVector(a) != false {
		t.Errorf("Tuple %v is not expected to be a vector.", a)
	}
}

func TestVectorTuple(t *testing.T) {
	a := []float32{4.3, -4.2, 3.1, 0.0}

	if IsVector(a) != true {
		t.Errorf("Tuple %v is expected to be a vector.", a)
	}

	if IsPoint(a) != false {
		t.Errorf("Tuple %v is not expected to be a point.", a)
	}
}
