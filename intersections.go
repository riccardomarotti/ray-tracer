package main

type Intersection struct {
	t      float64
	object Object
}

type HitData struct {
	point,
	eyeVector, normalVector Tuple
	inside bool
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
	normalVector := i.object.NormalAt(point)
	eyeVector := r.direction.Multiply(-1)
	inside := false

	if normalVector.Dot(eyeVector) < 0 {
		inside = true
		normalVector = normalVector.Multiply(-1)
	}

	return HitData{
		point:        point,
		eyeVector:    eyeVector,
		normalVector: normalVector,
		inside:       inside,
	}
}
