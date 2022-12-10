package challenges

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"
	"github.com/jcheyer/aoc2021/challenges/internal/tube"
	"github.com/jcheyer/aoc2021/lib"
)

type Day10 struct {
	rawInput []string
	tube     *tube.Tube
}

func (d *Day10) Load(file string) error {

	if len(d.rawInput) > 0 {
		return nil
	}
	var err error
	d.rawInput, err = lib.ReadInputLines(file)
	if err != nil {
		return err
	}
	d.tube = tube.New(d.rawInput)
	spew.Dump(d.tube.SignalStrength)
	return nil
}

func (d *Day10) Part1() string {
	res := int64(0)
	for _, signalStrength := range d.tube.SignalStrength {

		res += int64(signalStrength)
		spew.Dump(res)

	}

	return fmt.Sprintf("%d", res)
}

func (d *Day10) Part2() string {

	res := 0

	fmt.Printf("%s\n", d.tube.Screen)

	return fmt.Sprintf("%d", res)
}
