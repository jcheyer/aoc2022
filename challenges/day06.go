package challenges

import (
	"fmt"

	"github.com/jcheyer/aoc2021/lib"
)

type Day06 struct {
	rawInput []string
}

func (d *Day06) Load(file string) error {

	if len(d.rawInput) > 0 {
		return nil
	}
	var err error
	d.rawInput, err = lib.ReadInputLinesRaw(file)
	if err != nil {
		return err
	}

	return nil
}

func unique(s string) bool {
	set := map[rune]bool{}
	for _, c := range s {
		set[c] = true
	}
	return len(set) == len(s)
}

func (d *Day06) Part1() string {

	//res := int64(0)
	l := d.rawInput[0]
	pos := 0
	for i := 3; i < len(l); i++ {
		if unique(l[i-3:i+1]) && pos == 0 {
			pos = i + 1
		}
	}

	return fmt.Sprintf("%d", pos)
}

func (d *Day06) Part2() string {

	l := d.rawInput[0]
	pos := 0
	for i := 13; i < len(l); i++ {
		if unique(l[i-13:i+1]) && pos == 0 {
			pos = i + 1
		}
	}
	return fmt.Sprintf("%d", pos)
}
