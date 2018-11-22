package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Obj struct {
	vertices []Tuple
	groups   []Group
}

func ObjFileToGroups(filename string, transofrm Matrix, material Material) (groups []Group, discardeLineCount int, errors string) {
	dat, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	output, linesDiscarded, e := ParseObjString(string(dat), transofrm, material)

	groups = output.groups
	discardeLineCount = linesDiscarded
	errors = e

	return
}

func ParseObjString(input string, transform Matrix, material Material) (output Obj, discardedLinesCount int, errors string) {
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
	defaultGroup := MakeGroup(transform)
	output.groups = []Group{*defaultGroup}

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
				groupsCount := len(output.groups)
				for i := 1; i < len(vertices)-1; i++ {
					MakeTriangleInGroup(vertices[0], vertices[i], vertices[i+1], Identity(), material, &output.groups[groupsCount-1])
				}
			case 'g':
				output.groups = append(output.groups, *MakeGroup(transform))

			default:
				discardedLinesCount++
			}
		} else {
			discardedLinesCount++
		}
	}

	return
}
