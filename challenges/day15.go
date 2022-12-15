package challenges

import (
	"fmt"

	"github.com/jcheyer/aoc2021/challenges/internal/radar"
	"github.com/jcheyer/aoc2021/lib"
)

type Day15 struct {
	rawInput             []string
	p1Line               int
	p2min, p2max, p2freq int
}

func (d *Day15) Load(file string) error {
	if len(d.rawInput) > 0 {
		return nil
	}
	var err error
	d.rawInput, err = lib.ReadInputLinesRaw(file)
	if err != nil {
		return err
	}

	d.p1Line = 2000000

	d.p2min = 0
	d.p2max = 4000000
	d.p2freq = 4000000

	return nil
}

func (d *Day15) Part1() string {
	//res := 0

	r, _ := radar.New(d.rawInput)

	return fmt.Sprintf("%d", r.Part1(d.p1Line))
}

func (d *Day15) Part2() string {
	//	res := 0

	r, _ := radar.New(d.rawInput)

	return fmt.Sprintf("%d", r.Part2(d.p2min, d.p2max, d.p2freq))
}
