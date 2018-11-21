package main

type Object interface {
	Transform() Matrix
	NormalAt(p Tuple) Tuple
	Material() Material
	Intersection(Ray) []Intersection
	Parent() *Group
}

type BaseObject struct {
	transform Matrix
	material  Material
}

func (o BaseObject) Intersection(ray Ray, localIntersect func(Ray) []Intersection) []Intersection {
	localRay := ray.Transform(o.transform.Inverse())
	return localIntersect(localRay)
}

func (o BaseObject) NormalAt(p Tuple, localNormalAt func(p Tuple) Tuple) Tuple {
	localPoint := o.transform.Inverse().MultiplyByTuple(p)
	localNormal := localNormalAt(localPoint)
	worldNormal := o.transform.Inverse().T().MultiplyByTuple(localNormal)
	worldNormal.w = 0

	return worldNormal.Normalize()
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
