package radar

import (
	"fmt"
	"math"

	"github.com/jcheyer/aoc2021/lib"
)

type coord struct {
	x, y int
}

type Sensor struct {
	Position coord
	Beacon   coord
	Distance int
}

type Radar struct {
	Sensors []Sensor
	minX    int
	maxX    int
}

func New(data []string) (*Radar, error) {
	r := &Radar{
		minX: math.MaxInt,
		maxX: math.MinInt,
	}
	for _, line := range data {
		s, err := parseLine(line)
		if err != nil {
			return nil, err
		}
		r.Sensors = append(r.Sensors, s)
		r.minX = lib.LowInt(r.minX, s.Position.x-s.Distance)
		r.maxX = lib.HighInt(r.maxX, s.Position.x+s.Distance)
	}

	return r, nil
}

func (r Radar) Part1(lineNr int) int {
	return r.scanLine(lineNr)
}

func (r Radar) Part2(min, max, frequency int) int {
	pos := r.scanField(min, max)

	return pos.x*frequency + pos.y
}

func parseLine(line string) (Sensor, error) {
	s := Sensor{}
	_, err := fmt.Sscanf(line, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &s.Position.x, &s.Position.y, &s.Beacon.x, &s.Beacon.y)
	s.Distance = ManhattanDistance(s.Position, s.Beacon)
	return s, err

}

func ManhattanDistance(c1, c2 coord) int {
	x := lib.HighInt(c1.x, c2.x) - lib.LowInt(c1.x, c2.x)
	y := lib.HighInt(c1.y, c2.y) - lib.LowInt(c1.y, c2.y)
	return x + y
}

func (r Radar) scanLine(row int) int {

	total := 0
	for i := r.minX; i < r.maxX+1; i++ {
		testCoord := coord{i, row}

		inrange := false
		for _, s := range r.Sensors {
			if ManhattanDistance(s.Position, testCoord) <= s.Distance && testCoord != s.Beacon {
				inrange = true
				break
			}
		}

		if inrange {
			total++
		}
	}
	return total
}

func (r Radar) scanField(min, max int) coord {

	for y := min; y <= max; y++ {
		for x := min; x <= max; x++ {
			located := true
			for _, s := range r.Sensors {
				d := ManhattanDistance(s.Position, coord{x, y})
				if d <= s.Distance {
					skip := s.Distance - d
					x += skip
					located = false
					break
				}
			}

			if located {
				return coord{x, y}
			}
		}
	}

	return coord{0, 0}
}
