package challenges

import (
	"fmt"
	"strconv"

	"github.com/jcheyer/aoc2021/lib"
)

type Day08 struct {
	rawInput []string
	Forest   [][]int
	sizeX    int
	sizeY    int
}

func (d *Day08) Load(file string) error {

	if len(d.rawInput) > 0 {
		return nil
	}
	var err error
	d.rawInput, err = lib.ReadInputLines(file)
	if err != nil {
		return err
	}
	d.sizeY = len(d.rawInput)
	d.sizeX = len(d.rawInput[0])
	d.Forest = make([][]int, d.sizeY)
	for i := 0; i < d.sizeX; i++ {
		d.Forest[i] = make([]int, d.sizeX)
	}
	for x, line := range d.rawInput {
		for y, rune := range line {
			height, err := strconv.Atoi(string(rune))
			if err != nil {
				panic(err.Error())
			}
			d.Forest[x][y] = height
		}
	}

	return nil
}

func (d *Day08) IsVisible(x, y int) bool {
	height := d.Forest[x][y]

	if x == 0 || x == d.sizeX-1 || y == 0 || y == d.sizeY-1 {
		return true
	}

	v := true
	for i := 0; i < x; i++ {
		if d.Forest[i][y] >= height {
			v = false
			break
		}
	}
	if v {
		return true
	}

	v = true
	for i := x + 1; i < d.sizeX; i++ {
		if d.Forest[i][y] >= height {
			v = false
			break
		}
	}
	if v {
		return true
	}

	v = true
	for i := 0; i < y; i++ {
		if d.Forest[x][i] >= height {
			v = false
			break
		}
	}
	if v {
		return v
	}

	v = true
	for i := y + 1; i < d.sizeY; i++ {
		if d.Forest[x][i] >= height {
			v = false
			break
		}
	}

	return v
}

func (d *Day08) score(x, y int) int {
	height := d.Forest[x][y]

	if x == 0 || x == d.sizeX-1 || y == 0 || y == d.sizeY-1 {
		return 0
	}

	up := 0
	for i := x - 1; i >= 0; i-- {
		if d.Forest[i][y] >= height || i == 0 {
			up = x - i
			break
		}
	}

	left := 0
	for i := y - 1; i >= 0; i-- {
		if d.Forest[x][i] >= height || i == 0 {
			left = y - i
			break
		}
	}

	right := 0
	for i := y + 1; i < d.sizeY; i++ {
		if d.Forest[x][i] >= height || i == d.sizeX-1 {
			right = i - y
			break
		}
	}

	down := 0
	for i := x + 1; i < d.sizeX; i++ {
		if d.Forest[i][y] >= height || i == d.sizeX-1 {
			down = i - x
			break
		}
	}

	//fmt.Println(height, up, left, right, down)
	return up * left * right * down
}

func (d *Day08) Part1() string {
	res := int64(0)

	for x := 0; x < d.sizeX; x++ {
		for y := 0; y < d.sizeY; y++ {
			if d.IsVisible(x, y) {
				res++
			}
		}
	}

	return fmt.Sprintf("%d", res)
}

func (d *Day08) Part2() string {

	res := 0

	for x := 0; x < d.sizeX; x++ {
		for y := 0; y < d.sizeY; y++ {
			if d.score(x, y) > res {
				res = d.score(x, y)
			}
		}
	}

	return fmt.Sprintf("%d", res)
}
