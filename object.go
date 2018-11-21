package main

type Object interface {
	Transform() Matrix
	NormalAt(p Tuple) Tuple
	Material() Material
	Intersection(Ray) []Intersection
	Parent() *Group
}

func WorldToObject(o Object, p Tuple) Tuple {
	hasParent := o.Parent() != nil
	if hasParent {
		parent := o.Parent()
		p = WorldToObject(*parent, p)
	}

	return o.Transform().Inverse().MultiplyByTuple(p)
}
