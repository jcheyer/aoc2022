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

const (
	d14Stone = 2
	d14Sand  = 1
	d14Air   = 0
)

type Day14 struct {
	rawInput      []string
	abyss         int
	abyssReached  bool
	infinityLevel int
	fullyFilled   bool
	cave          map[d14coord]int
	caveBackup    map[d14coord]int
	sandSource    d14coord
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
	maxY := 0
	for _, line := range d.rawInput {
		path := strings.Split(line, " -> ")
		for i := 0; i < len(path)-1; i++ {
			c1, c2 := d.sortCoords(d.string2coord(path[i]), d.string2coord(path[i+1]))
			// Set Stones and keep Track of Abyss and Infinity Reference
			for x := c1.x; x <= c2.x; x++ {
				for y := c1.y; y <= c2.y; y++ {
					d.cave[d14coord{x, y}] = d14Stone
					d.caveBackup[d14coord{x, y}] = d14Stone

					if y > maxY {
						maxY = y
					}
				}
			}
		}
	}
	d.abyss = maxY + 5
	d.infinityLevel = maxY + 2

	return nil
}

func (d Day14) sortCoords(c1, c2 d14coord) (d14coord, d14coord) {
	if c1.x == c2.x {
		return d14coord{x: c1.x, y: lib.LowInt(c1.y, c2.y)}, d14coord{x: c1.x, y: lib.HighInt(c1.y, c2.y)}
	}

	return d14coord{x: lib.LowInt(c1.x, c2.x), y: c1.y}, d14coord{x: lib.HighInt(c1.x, c2.x), y: c2.y}
}

func (d Day14) string2coord(s string) d14coord {
	cparts := strings.Split(s, ",")
	x, _ := strconv.Atoi(cparts[0])
	y, _ := strconv.Atoi(cparts[1])

	return d14coord{x: x, y: y}
}

func (d *Day14) getCaveVal(pos d14coord, withInfiniteWall bool) int {
	if withInfiniteWall && pos.y == d.infinityLevel {
		return d14Stone
	}

	v, ok := d.cave[pos]
	if ok {
		return v
	}

	return d14Air
}
func (d *Day14) dropSand(pos d14coord, withInfiniteWall bool) d14coord {
	if withInfiniteWall && d.getCaveVal(d.sandSource, true) == d14Sand {
		d.fullyFilled = true
		return pos
	}

	if pos.y >= d.abyss {
		d.abyssReached = true
		return pos
	}

	if d.getCaveVal(pos.nextDropCoord(1), withInfiniteWall) == d14Air {
		return d.dropSand(pos.nextDropCoord(1), withInfiniteWall)
	}
	if d.getCaveVal(pos.nextDropCoord(2), withInfiniteWall) == d14Air {
		return d.dropSand(pos.nextDropCoord(2), withInfiniteWall)
	}
	if d.getCaveVal(pos.nextDropCoord(3), withInfiniteWall) == d14Air {
		return d.dropSand(pos.nextDropCoord(3), withInfiniteWall)
	}

	d.cave[pos] = d14Sand

	return pos
}

func (p d14coord) nextDropCoord(i int) d14coord {
	switch i {
	case 1:
		return d14coord{x: p.x, y: p.y + 1} // down
	case 2:
		return d14coord{x: p.x - 1, y: p.y + 1} // down left
	}
	return d14coord{x: p.x + 1, y: p.y + 1} // down right
}

func (d *Day14) Part1() string {
	res := 0
	d.sandSource = d14coord{500, 0}

	for {
		_ = d.dropSand(d.sandSource, false)
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

		if d.fullyFilled {
			break
		}
		res++
	}

	return fmt.Sprintf("%d", res)
}
