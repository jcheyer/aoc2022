package challenges

import (
	"fmt"
	"strings"

	"github.com/jcheyer/aoc2021/challenges/internal/stack"
	"github.com/jcheyer/aoc2021/lib"
)

type Day05 struct {
	rawInput []string
	stack    []stack.Stack[string]
}

func (d *Day05) Load(file string) error {

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

func prepLine(l string) string {
	l = strings.ReplaceAll(l, "   ", " ")
	l = strings.ReplaceAll(l, "  ", " ")
	l = strings.ReplaceAll(l, "[", "")
	l = strings.ReplaceAll(l, "]", "")
	return l
}

func (d *Day05) buildStack() (read int) {
	l := d.rawInput[0]
	l = prepLine(l)
	d.stack = make([]stack.Stack[string], len(l))
	lcount := 0
	for _, line := range d.rawInput {
		if strings.Contains(line, "1") {
			break
		}

		for pointer := 1; pointer < len(line); pointer += 4 {
			if line[pointer] != ' ' {
				letter := string(rune(line[pointer]))

				number := pointer / 4
				d.stack[number].SneakPush(letter)
			}
		}

		lcount++

	}
	//fmt.Printf("Stacks: %+v\n", d.stack)
	lcount++
	lcount++

	return lcount
}

func (d *Day05) movep1(from int) {
	commands := d.rawInput[from:]
	for _, command := range commands {
		count := 0
		from := 0
		to := 0

		fmt.Sscanf(command, "move %d from %d to %d\n", &count, &from, &to)

		for x := 0; x < count; x++ {
			d.stack[to-1].Push(d.stack[from-1].Pop())
		}

	}

}

func (d *Day05) movep2(from int) {
	commands := d.rawInput[from:]
	for _, command := range commands {
		count := 0
		from := 0
		to := 0

		fmt.Sscanf(command, "move %d from %d to %d\n", &count, &from, &to)

		d.stack[to-1].PushN(d.stack[from-1].PopN(count))

	}

}

func (d *Day05) Part1() string {

	read := d.buildStack()

	d.movep1(read)

	res := ""

	for st := range d.stack {
		res += string(d.stack[st].Peek())
	}

	return res
}

func (d *Day05) Part2() string {

	read := d.buildStack()

	d.movep2(read)

	res := ""

	for st := range d.stack {
		res += string(d.stack[st].Peek())
	}

	return res
}
