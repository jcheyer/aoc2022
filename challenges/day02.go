package challenges

import (
	"fmt"

	"github.com/jcheyer/aoc2021/lib"
)

type Day02 struct {
	rawInput []string
	calc     map[string]int64
	strategy map[string]int64
}

const (
	win  = 6
	draw = 3
	loss = 0
)

func (d *Day02) Load(file string) error {

	if len(d.rawInput) > 0 {
		return nil
	}
	var err error
	d.rawInput, err = lib.ReadInputLines(file)
	if err != nil {
		return err
	}

	d.calc = make(map[string]int64)

	d.calc["A X"] = 1 + draw
	d.calc["B X"] = 1 + loss
	d.calc["C X"] = 1 + win

	d.calc["A Y"] = 2 + win
	d.calc["B Y"] = 2 + draw
	d.calc["C Y"] = 2 + loss

	d.calc["A Z"] = 3 + loss
	d.calc["B Z"] = 3 + win
	d.calc["C Z"] = 3 + draw

	d.strategy = make(map[string]int64)

	d.strategy["A X"] = 3 + loss
	d.strategy["B X"] = 1 + loss
	d.strategy["C X"] = 2 + loss

	d.strategy["A Y"] = 1 + draw
	d.strategy["B Y"] = 2 + draw
	d.strategy["C Y"] = 3 + draw

	d.strategy["A Z"] = 2 + win
	d.strategy["B Z"] = 3 + win
	d.strategy["C Z"] = 1 + win

	return nil
}

func (d *Day02) Part1() string {

	res := int64(0)
	for _, line := range d.rawInput {
		res += d.calc[line]

	}

	return fmt.Sprintf("%d", res)
}

func (d *Day02) Part2() string {

	res := int64(0)
	for _, line := range d.rawInput {
		res += d.strategy[line]

	}

	return fmt.Sprintf("%d", res)
}
