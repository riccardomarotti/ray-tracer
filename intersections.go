package main

import (
	"math"
	"reflect"
)

const Epsilon = 0.00001

type Intersection struct {
	t, u, v float64
	object  Object
}

func Hit(i []Intersection) (hit Intersection) {
	hit = Intersection{}

	for k := 0; k < len(i); k++ {
		if i[k].t > 4*Epsilon && (i[k].t <= hit.t || hit.t == 0) {
			hit = i[k]
		}
	}

	return
}

func contains(array []Object, o Object) int {
	for i := 0; i < len(array); i++ {
		if reflect.DeepEqual(o, array[i]) {
			return i
		}
	}

	return -1
}

func areIntersectionsEqual(i1, i2 Intersection) bool {
	return (math.Abs(i1.t-i2.t) < Epsilon) && reflect.DeepEqual(i1.object, i2.object)
}
