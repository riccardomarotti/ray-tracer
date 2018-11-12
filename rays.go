package main

type Ray struct {
	origin, direction Tuple
}

func (r Ray) Position(d float64) Tuple {
	return r.direction.Multiply(d).Add(r.origin)
}

func (r Ray) Transform(T Matrix) Ray {
	return Ray{T.MultiplyByTuple(r.origin), T.MultiplyByTuple(r.direction)}
}
