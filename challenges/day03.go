package challenges

import (
	"fmt"
	"strings"

	"github.com/jcheyer/aoc2021/lib"
)

type Day03 struct {
	rawInput []string
}

var items = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func (d *Day03) Load(file string) error {

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

func (d *Day03) Part1() string {

	res := int64(0)

	for _, line := range d.rawInput {
		substr := line[len(line)/2:]
		for i := 0; i < len(line)/2; i++ {

			if strings.Contains(substr, string(line[i])) {
				sub := string(line[i])
				res += int64(strings.Index(items, sub) + 1)

				break
			}
		}
	}

	return fmt.Sprintf("%d", res)
}

func (d *Day03) Part2() string {

	res := int64(0)

	for offset := 0; offset < len(d.rawInput); offset = offset + 3 {

		for _, item := range items {
			sub := string(item)
			if strings.Contains(d.rawInput[offset], sub) &&
				strings.Contains(d.rawInput[offset+1], sub) &&
				strings.Contains(d.rawInput[offset+2], sub) {
				res += int64(strings.Index(items, sub) + 1)

				break
			}
		}
	}

	return fmt.Sprintf("%d", res)
}
