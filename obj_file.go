package main

import (
	"fmt"
	"strings"
)

func ParseObjFile(input string) (output, warn, err string) {
	lines := len(strings.Split(input, "\n"))

	warn = fmt.Sprintf("%d lines where discarded", lines)

	return
}
