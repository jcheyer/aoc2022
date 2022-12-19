package challenges

import (
	"fmt"

	"github.com/jcheyer/aoc2021/lib"
)

type Day18 struct {
	rawInput []string
}

type d18Cube struct {
	x int
	y int
	z int
}

func (d *Day18) Load(file string) error {
	if len(d.rawInput) > 0 {
		return nil
	}
	var err error
	d.rawInput, err = lib.ReadInputLinesRaw(file)
	if err != nil {
		return err
	}

	return nil
}

func (d Day18) generateCubes() []d18Cube {
	ret := []d18Cube{}
	for _, line := range d.rawInput {
		cube := d18Cube{}
		fmt.Sscanf(line, "%d,%d,%d", &cube.x, &cube.y, &cube.z)
		ret = append(ret, cube)
	}
	return ret
}

func (c d18Cube) Neighbors() []d18Cube {
	n := []d18Cube{}
	for _, i := range []int{-1, 1} {
		n = append(n, d18Cube{c.x + i, c.y, c.z})
		n = append(n, d18Cube{c.x, c.y + i, c.z})
		n = append(n, d18Cube{c.x, c.y, c.z + i})

	}

	return n
}

func (d *Day18) Part1() string {
	res := 0

	cubes := d.generateCubes()
	MegaCube := make(map[d18Cube]bool)
	for _, cube := range cubes {
		MegaCube[cube] = true
	}

	for _, cube := range cubes {
		for _, neighbor := range cube.Neighbors() {
			if !MegaCube[neighbor] {
				res++
			}
		}
	}

	return fmt.Sprintf("%d", res)
}

func (d *Day18) Part2() string {
	res := 0

	return fmt.Sprintf("%d", res)
}
