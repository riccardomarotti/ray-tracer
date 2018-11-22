package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestEmptyInout(t *testing.T) {
	input := ""

	output, discardedLinesCount, err := ParseObjFile(input)

	Assert(discardedLinesCount == 0, fmt.Sprintf("Discarded lines had to be 0, but where %d", discardedLinesCount), t)
	Assert(err == "", err, t)
	Assert(reflect.DeepEqual(Obj{}, output), fmt.Sprintf("output had to be empty, insted it was %v", output), t)
}

func TestIgnoringUnrecognizedLines(t *testing.T) {
	input := "not parsable things\nwith two lines"

	_, discardedLinesCount, err := ParseObjFile(input)

	Assert(discardedLinesCount == 2, fmt.Sprintf("Discarded lines had to be 2, but where %d", discardedLinesCount), t)
	Assert(err == "", "Errors occurred during parsing", t)
}

func TestVertexRecords(t *testing.T) {
	input := `v -1 1 0
v -1.0000 0.5000 0.0000
v 1 0 0
v 1 1 0`

	output, discardedLinesCount, err := ParseObjFile(input)

	Assert(discardedLinesCount == 0, fmt.Sprintf("Discarded lines had to be 0, but where %d", discardedLinesCount), t)
	Assert(err == "", "Errors occurred during parsing", t)

	AssertTupleEqual(Point(-1, 1, 0), output.vertices[1], t)
}
