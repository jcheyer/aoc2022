package challenges

import (
	"fmt"

	"github.com/jcheyer/aoc2021/challenges/internal/mountainmap"
	"github.com/jcheyer/aoc2021/lib"
)

type Day12 struct {
	rawInput    []string
	mountainmap *mountainmap.Mountainmap
}

func (d *Day12) Load(file string) error {

	if len(d.rawInput) > 0 {
		return nil
	}
	var err error
	d.rawInput, err = lib.ReadInputLinesRaw(file)
	if err != nil {
		return err
	}
	d.mountainmap, err = mountainmap.New(d.rawInput)
	if err != nil {
		return err
	}

	return nil
}

func (d *Day12) Part1() string {

	//spew.Dump(d.mountainmap.D)

	return fmt.Sprintf("%d", d.mountainmap.Part1())
}

func (d *Day12) Part2() string {
	//res := 0

	return fmt.Sprintf("%d", d.mountainmap.Part2())
}
