package challenges

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jcheyer/aoc2021/lib"
)

type d14coord struct {
	x, y int
}

const d14Stone = 2
const d14Sand = 1
const d14Air = 0

type Day14 struct {
	rawInput      []string
	abyss         int
	abyssReached  bool
	infinityLevel int
	fullyFilled   bool
	cave          map[d14coord]int
	caveBackup    map[d14coord]int
	sandSource    d14coord
	minX          int
	maxX          int
}

func (d *Day14) Load(file string) error {
	if len(d.rawInput) > 0 {
		return nil
	}
	var err error
	d.rawInput, err = lib.ReadInputLinesRaw(file)
	if err != nil {
		return err
	}
	d.cave = make(map[d14coord]int)
	d.caveBackup = make(map[d14coord]int)
	d.abyss = 0
	d.minX = 10000
	d.maxX = 0
	maxY := 0
	for _, line := range d.rawInput {
		path := strings.Split(line, " -> ")
		for i := 0; i < len(path)-1; i++ {
			c1 := d.string2coord(path[i])
			c2 := d.string2coord(path[i+1])
			c1, c2 = d.optSwap(c1, c2)
			for x := c1.x; x <= c2.x; x++ {
				for y := c1.y; y <= c2.y; y++ {
					d.cave[d14coord{x, y}] = d14Stone
					d.caveBackup[d14coord{x, y}] = d14Stone

					if y > maxY {
						maxY = y
					}
				}
				if d.minX > x {
					d.minX = x
				}
				if d.maxX < x {
					d.maxX = x
				}
			}
		}
	}
	d.abyss = maxY + 5
	d.infinityLevel = maxY + 2
	fmt.Printf("%s\n", d)

	return nil
}

func (d Day14) optSwap(c1, c2 d14coord) (d14coord, d14coord) {
	if c1.x > c2.x {
		tmpx := c2.x
		c2.x = c1.x
		c1.x = tmpx
	}

	if c1.y > c2.y {
		tmpy := c2.y
		c2.y = c1.y
		c1.y = tmpy
	}
	return c1, c2
}

func (d Day14) string2coord(s string) d14coord {
	cparts := strings.Split(s, ",")
	x, err := strconv.Atoi(cparts[0])
	if err != nil {
		panic(err.Error())
	}
	y, err := strconv.Atoi(cparts[1])
	if err != nil {
		panic(err.Error())
	}
	return d14coord{x: x, y: y}
}

func (d *Day14) getCaveVal(pos d14coord, withInfiniteWall bool) int {
	if withInfiniteWall {
		if pos.y == d.infinityLevel {
			return d14Stone
		}
	}
	v, ok := d.cave[pos]
	if ok {
		return v
	}
	return d14Air
}
func (d *Day14) dropSand(pos d14coord, withInfiniteWall bool) d14coord {
	if withInfiniteWall {
		if d.getCaveVal(d.sandSource, true) == d14Sand {
			d.fullyFilled = true
			return pos
		}

	}

	if pos.y >= d.abyss {
		d.abyssReached = true
		return pos
	}

	if d.getCaveVal(pos.dropCoord(1), withInfiniteWall) == d14Air {
		return d.dropSand(pos.dropCoord(1), withInfiniteWall)
	}
	if d.getCaveVal(pos.dropCoord(2), withInfiniteWall) == d14Air {
		return d.dropSand(pos.dropCoord(2), withInfiniteWall)
	}
	if d.getCaveVal(pos.dropCoord(3), withInfiniteWall) == d14Air {
		return d.dropSand(pos.dropCoord(3), withInfiniteWall)
	}

	d.cave[pos] = d14Sand
	if pos.x < d.minX {
		d.minX = pos.x
	}
	if pos.x > d.maxX {
		d.maxX = pos.x
	}

	return pos
}

func (d *Day14) String() string {
	var s strings.Builder
	s.WriteString(fmt.Sprintf("minX: %d maxX:%d\n", d.minX, d.maxX))
	for i := 0; i <= d.abyss; i++ {
		for j := d.minX; j <= d.maxX; j++ {
			_, ok := d.cave[d14coord{x: j, y: i}]
			if !ok {
				d.cave[d14coord{x: j, y: i}] = d14Air
			}
			switch d.cave[d14coord{x: j, y: i}] {
			case d14Air:
				s.WriteString(".")
			case d14Stone:
				s.WriteString("#")
			case d14Sand:
				s.WriteString("o")
			}
		}
		s.WriteString("\n")
	}
	s.WriteString("\n")
	return s.String()
}

func (p d14coord) dropCoord(i int) d14coord {
	switch i {
	case 1:
		return d14coord{x: p.x, y: p.y + 1}
	case 2:
		return d14coord{x: p.x - 1, y: p.y + 1}
	}
	return d14coord{x: p.x + 1, y: p.y + 1}
}

func (d *Day14) Part1() string {
	res := 0
	d.sandSource = d14coord{500, 0}

	for {
		_ = d.dropSand(d.sandSource, false)
		//fmt.Printf("%s\n", d)
		if d.abyssReached {
			break
		}
		res++
	}

	return fmt.Sprintf("%d", res)
}

func (d *Day14) Part2() string {
	res := 0
	d.cave = make(map[d14coord]int)
	for k, v := range d.caveBackup {
		d.cave[k] = v
	}

	for {
		_ = d.dropSand(d.sandSource, true)
		//fmt.Printf("%s\n", d)

		if d.fullyFilled {
			break
		}
		res++
	}

	return fmt.Sprintf("%d", res)
}
