package main

import (
	"reflect"
	"testing"
)

func TestDefaultWorld(t *testing.T) {
	w := DefaultWorld()

	light := PointLight{Point(-10, -10, -10), Color{1, 1, 1}}
	s1 := MakeSphere(Identity(), MakeMaterial(Color{0.8, 1.0, 0.6}, 0.9, 0.7, 0.2, 200))
	s2 := MakeSphere(Identity().Scale(0.5, 0.5, 0.5), DefaultMaterial())

	Assert(light == w.light, "", t)
	Assert(reflect.DeepEqual(s1, w.objects[0]), "", t)
	Assert(reflect.DeepEqual(s2, w.objects[1]), "", t)
}
