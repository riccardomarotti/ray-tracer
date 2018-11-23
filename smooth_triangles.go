package main

type SmoothTriangle struct {
	p1, p2, p3, n1, n2, n3 Tuple
}

func MakeSmoothTriangle(p1, p2, p3, n1, n2, n3 Tuple) SmoothTriangle {
	return SmoothTriangle{p1, p2, p3, n1, n2, n3}
}
