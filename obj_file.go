package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Obj struct {
	vertices []Tuple
}

func ParseObjFile(input string) (output Obj, discardedLinesCount int, errors string) {
	lines := strings.Split(input, "\n")
	discardedLinesCount = 0

	defer func() {
		if err := recover(); err != nil {
			errors = fmt.Sprintf("Fatal error during parsing: %s (%T)", err, err)
		}
	}()

	if len(lines) <= 1 {
		return
	}

	output.vertices = []Tuple{Tuple{}}

	for _, line := range lines {
		switch line[0] {
		case 'v':
			values := strings.Split(line, " ")
			x, _ := strconv.ParseFloat(values[1], 64)
			y, _ := strconv.ParseFloat(values[2], 64)
			z, _ := strconv.ParseFloat(values[3], 64)
			output.vertices = append(output.vertices, Point(x, y, z))

		default:
			discardedLinesCount++
		}
	}

	return
}
