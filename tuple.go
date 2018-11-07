package main

import (
	"math"
)

type Tuple struct {
	x, y, z, w float64
}

func (t Tuple) IsPoint() bool {
	return t.w == 1.0
}

func (t Tuple) IsVector() bool {
	return t.w == 0.0
}

func (t Tuple) Negate() Tuple {
	return Tuple{0, 0, 0, 0}.Subtract(t)
}

func Point(x, y, z float64) Tuple {
	return Tuple{x, y, z, 1}
}

func Vector(x, y, z float64) Tuple {
	return Tuple{x, y, z, 0}
}

func (t Tuple) Add(a Tuple) Tuple {
	return Tuple{t.x + a.x, t.y + a.y, t.z + a.z, t.w + a.w}
}

func (t Tuple) Subtract(a Tuple) Tuple {
	return Tuple{t.x - a.x, t.y - a.y, t.z - a.z, t.w - a.w}
}

func (t Tuple) Multiply(s float64) Tuple {
	return Tuple{t.x * s, t.y * s, t.z * s, t.w * s}
}

func (t Tuple) Divide(s float64) Tuple {
	return Tuple{t.x / s, t.y / s, t.z / s, t.w / s}
}

func (t Tuple) Magnitude() float64 {
	return math.Sqrt(t.x*t.x + t.y*t.y + t.z*t.z)
}

func (t Tuple) Normalize() Tuple {
	return Tuple{t.x / t.Magnitude(),
		t.y / t.Magnitude(),
		t.z / t.Magnitude(),
		t.w / t.Magnitude()}
}

func (t Tuple) Dot(a Tuple) float64 {
	return t.x*a.x + t.y*a.y + t.z*a.z + t.w*a.w
}

func (t Tuple) Cross(a Tuple) Tuple {
	return Vector(t.y*a.z-t.z*a.y,
		t.z*a.x-t.x*a.z,
		t.x*a.y-t.y*a.x)
}

func (t Tuple) Reflect(normal Tuple) Tuple {
	return t.Subtract(normal.Multiply(2).Multiply(t.Dot(normal)))
}
