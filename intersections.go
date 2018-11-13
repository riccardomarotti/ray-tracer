package main

type Intersection struct {
	t                                             float64
	object                                        Object
	point, eyeVector, normalVector, reflectVector Tuple
	inside                                        bool
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

func PrepareHit(i Intersection, r Ray) Intersection {
	point := r.Position(i.t)
	normalVector := i.object.NormalAt(point)
	point = point.Add(normalVector.Multiply(0.000009))
	eyeVector := r.direction.Multiply(-1)
	inside := false

	if normalVector.Dot(eyeVector) < 0 {
		inside = true
		normalVector = normalVector.Multiply(-1)
	}

	reflectVector := r.direction.Reflect(normalVector)

	return Intersection{
		t:             i.t,
		object:        i.object,
		point:         point,
		eyeVector:     eyeVector,
		normalVector:  normalVector,
		inside:        inside,
		reflectVector: reflectVector,
	}
}

func (i Intersection) Shade(world World) Color {
	return i.object.Material().Lighting(i.object.Transform(), world.light, i.point, i.eyeVector, i.normalVector, world.IsShadowed(i.point))
}
