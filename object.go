package main

type Object interface {
	Transform() Matrix
}

type Intersection struct {
	t      float64
	object Object
}

func Hit(i []Intersection) (hit Intersection) {
	hit = Intersection{}

	for k := 0; k < len(i); k++ {
		if i[k].t >= 0 && (i[k].t <= hit.t || hit.t == 0) {
			hit = i[k]
		}
	}

	return
}
