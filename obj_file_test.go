package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestEmptyInout(t *testing.T) {
	input := ""

	output, discardedLinesCount, err := ParseObjString(input, Identity(), DefaultMaterial())

	Assert(discardedLinesCount == 0, fmt.Sprintf("Discarded lines had to be 0, but where %d", discardedLinesCount), t)
	Assert(err == "", err, t)
	Assert(reflect.DeepEqual(Obj{}, output), fmt.Sprintf("output had to be empty, insted it was %v", output), t)
}

func TestIgnoringUnrecognizedLines(t *testing.T) {
	input := "not parsable things\nwith two lines"

	_, discardedLinesCount, err := ParseObjString(input, Identity(), DefaultMaterial())

	Assert(discardedLinesCount == 2, fmt.Sprintf("Discarded lines had to be 2, but where %d", discardedLinesCount), t)
	Assert(err == "", "Errors occurred during parsing", t)
}

func TestVertexRecords(t *testing.T) {
	input := `v -1 1 0
v -1.0000 0.5000 0.0000
v 1 0 0
v 1 1 0`

	output, discardedLinesCount, err := ParseObjString(input, Identity(), DefaultMaterial())

	Assert(discardedLinesCount == 0, fmt.Sprintf("Discarded lines had to be 0, but where %d", discardedLinesCount), t)
	Assert(err == "", err, t)

	AssertTupleEqual(Point(-1, 1, 0), output.vertices[1], t)
}

func TestParsingTriangleFaces(t *testing.T) {
	input := `v -1 1 0
v -1.0000 0.5000 0.0000
v 1 0 0
v 1 1 0

f 1 2 3
f 1 3 4
`
	output, discardedLinesCount, err := ParseObjString(input, Identity(), DefaultMaterial())

	Assert(discardedLinesCount == 2, fmt.Sprintf("Discarded lines had to be 2, but where %d", discardedLinesCount), t)
	Assert(err == "", err, t)

	g := output.groups[0]
	t1 := g.children[0]
	t2 := g.children[1]

	p1 := Point(-1, 1, 0)
	p2 := Point(-1, .5, 0)
	p3 := Point(1, 0, 0)
	p4 := Point(1, 1, 0)
	expectedT1 := MakeTriangleInGroup(p1, p2, p3, Identity(), DefaultMaterial(), g)
	expectedT2 := MakeTriangleInGroup(p1, p3, p4, Identity(), DefaultMaterial(), g)

	AssertTrianglesEqual(expectedT1, t1, t)
	AssertTrianglesEqual(expectedT2, t2, t)
	AssertMatrixEqual(Identity(), g.Transform(), t)
}

func TestParsingPolygonData(t *testing.T) {
	input := `v -1 1 0
v -1.0000 0.5000 0.0000
v 1 0 0
v 1 1 0
v 0 2 0

f 1 2 3 4 5
`
	output, _, _ := ParseObjString(input, Identity(), DefaultMaterial())

	g := output.groups[0]
	t1 := g.children[0]
	t2 := g.children[1]
	t3 := g.children[2]

	p1 := Point(-1, 1, 0)
	p2 := Point(-1, .5, 0)
	p3 := Point(1, 0, 0)
	p4 := Point(1, 1, 0)
	p5 := Point(0, 2, 0)

	expectedT1 := MakeTriangleInGroup(p1, p2, p3, Identity(), DefaultMaterial(), g)
	expectedT2 := MakeTriangleInGroup(p1, p3, p4, Identity(), DefaultMaterial(), g)
	expectedT3 := MakeTriangleInGroup(p1, p4, p5, Identity(), DefaultMaterial(), g)

	AssertTrianglesEqual(expectedT1, t1, t)
	AssertTrianglesEqual(expectedT2, t2, t)
	AssertTrianglesEqual(expectedT3, t3, t)
}

func TestTrianglesInGroup(t *testing.T) {
	output, _, _ := ObjFileToGroups("triangles.obj", Identity(), DefaultMaterial())

	g1 := output[1]
	g2 := output[2]
	t1 := g1.children[0]
	t2 := g2.children[0]

	p1 := Point(-1, 1, 0)
	p2 := Point(-1, 0, 0)
	p3 := Point(1, 0, 0)
	p4 := Point(1, 1, 0)

	expectedT1 := MakeTriangleInGroup(p1, p2, p3, Identity(), DefaultMaterial(), g1)
	expectedT2 := MakeTriangleInGroup(p1, p3, p4, Identity(), DefaultMaterial(), g2)

	AssertTrianglesEqual(expectedT1, t1, t)
	AssertTrianglesEqual(expectedT2, t2, t)
	AssertMatrixEqual(Identity(), g1.Transform(), t)
	AssertMatrixEqual(Identity(), g2.Transform(), t)
}
