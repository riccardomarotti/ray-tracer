package main

type Intersection struct {
	t      float64
	object Object
}

type HitData struct {
	point, eyeVector, normalVector Tuple
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

func PrepareHit(i Intersection, r Ray) HitData {
	point := r.Position(i.t)
	return HitData{point: point, eyeVector: r.direction.Multiply(-1), normalVector: i.object.NormalAt(point)}
}
