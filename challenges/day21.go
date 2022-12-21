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
	m := yelling.New(d.rawInput)
	x := m["root"]
	x.Op = func(a, b int) int { return a - b }
	m["root"] = x

	res := 0
	//for i := -1000000; i < 1000000; i++ {
	for i := 3247317260000; i < 30000000000000; i++ {
		h := m["humn"]
		h.Result = i
		m["humn"] = h

		xres := yelling.Solve(m, "root")
		if xres == 0 {
			res = i
			break
		}
		//fmt.Printf("checking %d result: %d\n", i, xres)

	}

	return fmt.Sprintf("%d", res)
}
