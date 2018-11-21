package main

import (
	"math"
	"reflect"
	"testing"
)

func TestGroupWithAChild(t *testing.T) {
	g := MakeGroup(Identity())
	s := MakeSphereInGroup(Identity(), DefaultMaterial(), &g)
	g.AddChildren(s)

	Assert(reflect.DeepEqual(s, g.children[0]), "", t)
	Assert(s.Parent() == &g, "", t)
}

func TestIntersectingARayWithAnEmptyGroup(t *testing.T) {
	g := MakeGroup(Identity())
	r := Ray{Point(0, 0, 0), Vector(0, 0, 1)}

	xs := g.Intersection(r)

	Assert(len(xs) == 0, "", t)
}

func TestIntersectingARayWithANonEmptyGroup(t *testing.T) {
	g := MakeGroup(Identity())
	s1 := MakeSphereInGroup(Identity(), DefaultMaterial(), &g)
	s2 := MakeSphereInGroup(Identity().Translate(0, 0, -3), DefaultMaterial(), &g)
	s3 := MakeSphereInGroup(Identity().Translate(5, 0, 0), DefaultMaterial(), &g)

	g.AddChildren(s1, s2, s3)

	r := Ray{Point(0, 0, -5), Vector(0, 0, 1)}

	xs := g.Intersection(r)

	Assert(len(xs) == 4, "", t)
	Assert(reflect.DeepEqual(s2, xs[0].object), "", t)
	Assert(reflect.DeepEqual(s2, xs[1].object), "", t)
	Assert(reflect.DeepEqual(s1, xs[2].object), "", t)
	Assert(reflect.DeepEqual(s1, xs[3].object), "", t)
}

func TestIntersectingATransformedGroup(t *testing.T) {
	g := MakeGroup(Identity().Scale(2, 2, 2))
	s := MakeSphereInGroup(Identity().Translate(5, 0, 0), DefaultMaterial(), &g)
	g.AddChildren(s)

	r := Ray{Point(10, 0, -10), Vector(0, 0, 1)}

	xs := g.Intersection(r)

	Assert(len(xs) == 2, "", t)
}

func TestConvertingAPointFromWorldToObjectSpace(t *testing.T) {
	g1 := MakeGroup(Identity().RotateY(math.Pi / 2))
	g2 := MakeGroupInGroup(Identity().Scale(2, 2, 2), &g1)

	g1.AddChildren(g2)

	s := MakeSphereInGroup(Identity().Translate(5, 0, 0), DefaultMaterial(), &g2)

	p := WorldToObject(s, Point(-2, 0, -10))

	AssertTupleEqual(Point(0, 0, -1), p, t)
}

func TestConvertingANormalFroObjectToWorldSpace(t *testing.T) {
	g1 := MakeGroup(Identity().RotateY(math.Pi / 2))
	g2 := MakeGroupInGroup(Identity().Scale(1, 2, 3), &g1)
	g1.AddChildren(g2)

	s := MakeSphereInGroup(Identity().Translate(5, 0, 0), DefaultMaterial(), &g2)

	n := NormalToWorld(s, Vector(math.Sqrt(3)/3, math.Sqrt(3)/3, math.Sqrt(3)/3))

	AssertTupleEqual(Vector(0.28571, 0.42857, -0.85714), n, t)
}

// func TestFindTheNormalOnAnObjectGroup(t *testing.T) {
// 	g1 := MakeGroup(Identity().RotateY(math.Pi / 2))
// 	g2 := MakeGroupInGroup(Identity().Scale(1, 2, 3), &g1)
// 	s := MakeSphereInGroup(Identity().Translate(5, 0, 0), DefaultMaterial(), &g2)

// 	n := s.NormalAt(Point(1.7321, 1.1547, -5.5774))
// 	AssertTupleEqual(Vector(0.2857, 0.4286, -0.8571), n, t)
// }
