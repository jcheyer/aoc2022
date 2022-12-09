package challenges

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/jcheyer/aoc2021/lib"
)

type Day09 struct {
	rawInput []string
	grid     d09grid
}

type pos struct {
	x, y int
}

func (p pos) String() string {
	return fmt.Sprintf("%d:%d", p.x, p.y)
}

type d09grid struct {
	field   map[pos]int
	h       pos
	t1      pos
	t2      pos
	t3      pos
	t4      pos
	t5      pos
	t6      pos
	t7      pos
	t8      pos
	t9      pos
	t9visit map[pos]int
}

func (d *Day09) Load(file string) error {

	if len(d.rawInput) > 0 {
		return nil
	}
	var err error
	d.rawInput, err = lib.ReadInputLines(file)
	if err != nil {
		return err
	}
	d.grid = d09grid{
		h:  pos{x: 0, y: 0},
		t1: pos{x: 0, y: 0},
		t2: pos{x: 0, y: 0},
		t3: pos{x: 0, y: 0},
		t4: pos{x: 0, y: 0},
		t5: pos{x: 0, y: 0},
		t6: pos{x: 0, y: 0},
		t7: pos{x: 0, y: 0},
		t8: pos{x: 0, y: 0},
		t9: pos{x: 0, y: 0},
	}

	d.grid.field = make(map[pos]int)
	d.grid.t9visit = make(map[pos]int)

	d.grid.field[pos{0, 0}] = 1

	return nil
}

func (g *d09grid) move(s string) {
	parts := strings.Split(s, " ")
	dx := 0
	dy := 0
	switch parts[0] {
	case "U":
		dy = 1
	case "D":
		dy = -1
	case "L":
		dx = -1
	case "R":
		dx = +1
	}
	count, _ := strconv.Atoi(parts[1])
	for i := 0; i < count; i++ {
		g.h.x += dx
		g.h.y += dy
		g.t1 = g.tfollowh(g.h, g.t1)
		g.field[g.t1] += 1
		// too lazy to make a loop (the aoc-for-allergy kicks in......)
		g.t2 = g.tfollowh(g.t1, g.t2)
		g.t3 = g.tfollowh(g.t2, g.t3)
		g.t4 = g.tfollowh(g.t3, g.t4)
		g.t5 = g.tfollowh(g.t4, g.t5)
		g.t6 = g.tfollowh(g.t5, g.t6)
		g.t7 = g.tfollowh(g.t6, g.t7)
		g.t8 = g.tfollowh(g.t7, g.t8)
		g.t9 = g.tfollowh(g.t8, g.t9)
		g.t9visit[g.t9] += 1

	}
}

func (g *d09grid) tfollowh(pos1, pos2 pos) pos {

	xDist := math.Abs(float64(pos1.x - pos2.x))
	yDist := math.Abs(float64(pos1.y - pos2.y))

	if xDist < 2 && yDist < 2 {
		return pos2
	}

	xMoveDir := float64(pos1.x-pos2.x) / 2
	yMoveDir := float64(pos1.y-pos2.y) / 2

	if xMoveDir < 0 {
		xMoveDir = math.Floor(xMoveDir)
	} else {
		xMoveDir = math.Ceil(xMoveDir)
	}

	if yMoveDir < 0 {
		yMoveDir = math.Floor(yMoveDir)
	} else {
		yMoveDir = math.Ceil(yMoveDir)
	}

	pos2.x += int(xMoveDir)
	pos2.y += int(yMoveDir)
	return pos2
}

func (d *Day09) Part1() string {
	res := int64(0)

	for _, line := range d.rawInput {
		d.grid.move(line)
		//fmt.Printf("LineNr: %d H: %s T: %s T9: %s Movement: %s\n", i, d.grid.h, d.grid.t1, d.grid.t9, line)
	}

	for _, visits := range d.grid.field {
		if visits > 0 {
			res++
		}

	}

	return fmt.Sprintf("%d", res)
}

func (d *Day09) Part2() string {

	res := 0

	for _, visits := range d.grid.t9visit {
		if visits > 0 {
			res++
		}

	}

	return fmt.Sprintf("%d", res)
}
