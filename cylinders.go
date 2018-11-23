package main

import "math"

type Cylinder struct {
	baseObject       BaseObject
	minimum, maximum float64
	closed           bool
	parent           *Group
}

func MakeInfiniteCylinder(transform Matrix, material Material) Object {
	return Cylinder{BaseObject{transform, material}, math.Inf(-1), math.Inf(1), false, nil}
}

func MakeClosedCylinder(transform Matrix, material Material, minimum, maximum float64) Object {
	return Cylinder{BaseObject{transform, material}, minimum, maximum, true, nil}
}

func MakeCylinder(transform Matrix, material Material, minimum, maximum float64, closed bool) Object {
	return Cylinder{BaseObject{transform, material}, minimum, maximum, closed, nil}
}

func MakeCylinderInGroup(transform Matrix, material Material, minimum, maximum float64, closed bool, g *Group) Object {
	c := Cylinder{BaseObject{transform, material}, minimum, maximum, closed, g}
	g.AddChildren(c)
	return c
}

func (cylinder Cylinder) Bounds() Bounds {
	localBounds := Bounds{Point(-1, cylinder.minimum, -1), Point(1, cylinder.maximum, 1)}
	return cylinder.baseObject.Bounds(localBounds, cylinder)
}

func (cylinder Cylinder) Parent() *Group {
	return cylinder.parent
}

func (cylinder Cylinder) Transform() Matrix {
	return cylinder.baseObject.transform
}

func (cylinder Cylinder) Material() Material {
	return cylinder.baseObject.material
}

func (cylinder Cylinder) NormalAt(p Tuple) Tuple {
	localNormalAt := func(p Tuple) Tuple {
		distance := p.x*p.x + p.z*p.z

		if distance < 1 && p.y > cylinder.maximum-Epsilon {
			return Vector(0, 1, 0)
		} else if distance < 1 && p.y <= cylinder.minimum+Epsilon {
			return Vector(0, -1, 0)
		}

		return Vector(p.x, 0, p.z)
	}

	return cylinder.baseObject.NormalAt(p, cylinder, localNormalAt)
}

func (cylinder Cylinder) Intersection(r Ray) []Intersection {
	localIntersect := func(r Ray) (intersections []Intersection) {
		intersections = make([]Intersection, 0)

		if cylinder.closed {
			intersections = cylinder.intersectCaps(r, intersections)
		}

		a := r.direction.x*r.direction.x + r.direction.z*r.direction.z
		if math.Abs(a) < Epsilon {
			return
		}

		b := 2*r.origin.x*r.direction.x + 2*r.origin.z*r.direction.z
		c := r.origin.x*r.origin.x + r.origin.z*r.origin.z - 1

		delta := b*b - 4*a*c

		if delta < 0 {
			return
		}

		t0 := (-b - math.Sqrt(delta)) / (2 * a)
		t1 := (-b + math.Sqrt(delta)) / (2 * a)

		if t0 > t1 {
			t0, t1 = t1, t0
		}

		y0 := r.origin.y + t0*r.direction.y
		if (cylinder.minimum) < y0 && y0 < cylinder.maximum {
			intersections = append(intersections, Intersection{t0, cylinder})
		}

		y1 := r.origin.y + t1*r.direction.y
		if cylinder.minimum < y1 && y1 < cylinder.maximum {
			intersections = append(intersections, Intersection{t1, cylinder})
		}

		return
	}

	return cylinder.baseObject.Intersection(r, localIntersect)
}

func checkCap(ray Ray, t float64) bool {
	x := ray.origin.x + t*ray.direction.x
	z := ray.origin.z + t*ray.direction.z

	return (x*x + z*z) <= 1
}

func (cylinder Cylinder) intersectCaps(ray Ray, intersections []Intersection) []Intersection {
	var newIntersections []Intersection

	t := (cylinder.minimum - ray.origin.y) / ray.direction.y
	if checkCap(ray, t) {
		newIntersections = append(newIntersections, Intersection{t, cylinder})
	}

	t = (cylinder.maximum - ray.origin.y) / ray.direction.y
	if checkCap(ray, t) {
		newIntersections = append(newIntersections, Intersection{t, cylinder})
	}

	return append(intersections, newIntersections...)

}
