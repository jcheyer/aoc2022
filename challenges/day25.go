package challenges

import (
	"fmt"

	"github.com/jcheyer/aoc2021/challenges/internal/snafu"
	"github.com/jcheyer/aoc2021/lib"
)

type Day25 struct {
	rawInput []string
}

func (d *Day25) Load(file string) error {
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

func (d *Day25) Part1() string {
	res := 0
	for _, line := range d.rawInput {
		res += snafu.Snafu2Dec(line)
	}

	return snafu.Dec2Snafu(res)
}

func (d *Day25) Part2() string {

	return fmt.Sprintf("%d", 0)
}
