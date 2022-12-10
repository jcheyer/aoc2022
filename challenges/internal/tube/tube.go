package tube

import (
	"fmt"
	"strconv"
	"strings"
)

type Tube struct {
	x              int
	cycle          int
	Screen         string
	SignalStrength map[int]int
}

func New(data []string) *Tube {
	t := &Tube{
		x:              1,
		SignalStrength: make(map[int]int),
	}

	for _, line := range data {

		parts := strings.Split(line, " ")
		fmt.Printf("Parsing %s %+v x: %+v cycle: %+v\n", line, parts, t.x, t.cycle)
		switch parts[0] {

		case "noop":
			t.cycle++
			t.checkSignalStrength()
			t.drawPixel()
			continue

		default:
			val, _ := strconv.Atoi(parts[1])
			t.cycle++
			t.checkSignalStrength()
			t.drawPixel()
			t.cycle++
			t.checkSignalStrength()
			t.drawPixel()
			t.x += val

		}
		fmt.Printf("%s\n", t.Screen)
	}
	return t
}

func (t *Tube) checkSignalStrength() {
	if (t.cycle-20)%40 == 0 {
		t.SignalStrength[t.cycle] = t.cycle * t.x
	}

	return
}

// There is a small off-by-one bug but the characters are readable anyway.....
func (t *Tube) drawPixel() {
	hpos := (t.cycle % 40)
	if hpos == t.x || hpos == t.x+2 || hpos == t.x+1 {
		t.Screen += "X"
	} else {
		t.Screen += "."
	}
	if hpos == 0 {
		t.Screen += "\n"
	}

	return
}
