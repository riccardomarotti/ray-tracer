package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Obj struct {
	vertices     []Tuple
	defaultGroup Group
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
		if len(line) > 0 {
			switch line[0] {
			case 'v':
				values := strings.Split(line, " ")
				x, _ := strconv.ParseFloat(values[1], 64)
				y, _ := strconv.ParseFloat(values[2], 64)
				z, _ := strconv.ParseFloat(values[3], 64)
				output.vertices = append(output.vertices, Point(x, y, z))
			case 'f':
				values := strings.Split(line, " ")
				t1index, _ := strconv.ParseInt(values[1], 10, 64)
				t2index, _ := strconv.ParseInt(values[2], 10, 64)
				t3index, _ := strconv.ParseInt(values[3], 10, 64)

				MakeTriangleInGroup(output.vertices[t1index], output.vertices[t2index], output.vertices[t3index], Identity(), DefaultMaterial(), &output.defaultGroup)
			default:
				discardedLinesCount++
			}
		} else {
			discardedLinesCount++
		}
	}

	return
}
