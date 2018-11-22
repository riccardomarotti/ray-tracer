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

				vertices := []Tuple{}
				for i := 1; i < len(values); i++ {
					tIndex, _ := strconv.ParseInt(values[i], 10, 64)
					vertices = append(vertices, output.vertices[tIndex])
				}

				fanTriangulation(vertices, &output.defaultGroup)
			default:
				discardedLinesCount++
			}
		} else {
			discardedLinesCount++
		}
	}

	return
}

func fanTriangulation(vertices []Tuple, g *Group) (triangles []Triangle) {
	triangles = []Triangle{}

	for i := 0; i < len(vertices)-1; i++ {
		triangles = append(triangles, MakeTriangleInGroup(vertices[1], vertices[i], vertices[i+1], Identity(), DefaultMaterial(), g))
	}

	return
}
