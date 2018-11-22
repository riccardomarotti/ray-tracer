package main

import (
	"fmt"
	"testing"
)

func TestIgnoringUnrecognizedLines(t *testing.T) {
	input := "not parsable things\nwith two lines"

	_, warn, _ := ParseObjFile(input)

	expectedWaring := "2 lines where discarded"
	Assert(warn == expectedWaring, fmt.Sprintf("\nExpected warning was: '%s'\nReceived: '%s'", expectedWaring, warn), t)
}
