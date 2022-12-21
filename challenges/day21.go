package challenges

import (
	"fmt"

	"github.com/jcheyer/aoc2021/challenges/internal/yelling"
	"github.com/jcheyer/aoc2021/lib"
)

type Day21 struct {
	rawInput []string
}

func (d *Day21) Load(file string) error {
	if len(d.rawInput) > 0 {
		return nil
	}
	var err error
	d.rawInput, err = lib.ReadInputLines(file)
	if err != nil {
		return err
	}

	return nil
}

func (d *Day21) Part1() string {
	m := yelling.New(d.rawInput)

	return fmt.Sprintf("%d", yelling.Solve(m, "root"))
}

func (d *Day21) Part2() string {
	res := 0

	return fmt.Sprintf("%d", res)
}
