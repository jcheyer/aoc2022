package challenges

import (
	"fmt"
	"sort"

	"github.com/jcheyer/aoc2021/lib"
)

type Day01 struct {
	rawInput []string
	calories map[int]int64
}

func (d *Day01) Load(file string) error {
	if len(d.rawInput) > 0 {
		return nil
	}
	var err error
	d.rawInput, err = lib.ReadInputLinesWithEmptyLines(file)
	if err != nil {
		return err
	}

	d.calories, err = lib.ParseAndAggregateLines2Int(d.rawInput)
	if err != nil {
		return err
	}
	return nil
}

func (d *Day01) Part1() string {
	maxCal := int64(0)
	for _, cal := range d.calories {
		if cal > maxCal {
			maxCal = cal
		}
	}
	return fmt.Sprintf("%d", maxCal)
}

func (d *Day01) Part2() string {

	topCal := make([]int, 0, len(d.calories))
	for k := range d.calories {
		topCal = append(topCal, int(d.calories[k]))
	}
	sort.Ints(topCal)

	for _, k := range topCal {
		fmt.Println(k)
	}
	top3 := topCal[len(topCal)-1] + topCal[len(topCal)-2] + topCal[len(topCal)-3]

	return fmt.Sprintf("%d", top3)
}
