package main

import (
	"testing"
)

func TestAnObjectHasAParentAttribute(t *testing.T) {
	s := MakeSphere(Identity(), DefaultMaterial())

	Assert(s.Parent() == nil, "", t)
}
