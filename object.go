package main

type Object interface {
	Transform() Matrix
	NormalAt(p Tuple) Tuple
	Material() Material
	Intersection(Ray) []Intersection
	Parent() *Group
	Bounds() Bounds
}

type BaseObject struct {
	transform Matrix
	material  Material
}

func (o BaseObject) Bounds(localBounds Bounds, object Object) Bounds {
	return BoundsToWorld(object, Bounds{localBounds.min, localBounds.max})
}

func (o BaseObject) Intersection(ray Ray, localIntersect func(Ray) []Intersection) []Intersection {
	localRay := ray.Transform(o.transform.Inverse())
	return localIntersect(localRay)
}

func (o BaseObject) NormalAt(p Tuple, object Object, localNormalAt func(p Tuple) Tuple) Tuple {
	localPoint := WorldToObject(object, p)
	localNormal := localNormalAt(localPoint)

	return NormalToWorld(object, localNormal)
}

func WorldToObject(o Object, p Tuple) Tuple {
	hasParent := o.Parent() != nil
	if hasParent {
		parent := o.Parent()
		p = WorldToObject(*parent, p)
	}

	return o.Transform().Inverse().MultiplyByTuple(p)
}

func NormalToWorld(o Object, v Tuple) Tuple {
	normal := o.Transform().Inverse().T().MultiplyByTuple(v)
	normal.w = 0
	normal = normal.Normalize()

	if o.Parent() != nil {
		normal = NormalToWorld(o.Parent(), normal)
	}

	return normal
}

func BoundsToWorld(o Object, b Bounds) Bounds {
	min := o.Transform().MultiplyByTuple(b.min)
	max := o.Transform().MultiplyByTuple(b.max)

	bounds := Bounds{min, max}

	if o.Parent() != nil {
		bounds = BoundsToWorld(o.Parent(), bounds)
	}

	return bounds
}
