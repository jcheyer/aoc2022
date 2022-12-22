package monkeymap

import (
	"regexp"
	"strings"
)

type coord struct{ x, y int }

type mmap map[coord]int

func New(data []byte) (mmap, string) {
	m := make(mmap)
	parts := strings.Split(string(data), "\n\n")
	lines := strings.Split(parts[0], "\n")

	for i, line := range lines {
		for j, chr := range line {
			val := 0
			switch chr {
			case ' ':
				val = 0
			case '.':
				val = 1
			case '#':
				val = 2
			}
			m[coord{j, i}] = val
		}
	}
	return m, parts[1]
}

func parseInstructions(s string) {
	_ = regexp.MustCompile(`(?m)[0-9]+|L|R`)
	return
}
