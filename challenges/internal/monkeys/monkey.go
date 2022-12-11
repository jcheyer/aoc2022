package monkeys

import (
	"fmt"
	"strconv"
	"strings"
)

type Monkey struct {
	Items       []int64
	Operation   func(x int64) int64
	Test        int64
	TrueTarget  int
	FalseTarget int
	Counter     int
}

func New(data []string) map[int]*Monkey {
	mm := make(map[int]*Monkey)
	num := 0
	items := []int64{}
	op := func(x int64) int64 { return x }
	test := int64(0)
	trueAction := 0
	falseAction := 0

	for i := 0; i < len(data); i++ {
		switch i % 7 {
		case 0:
			num = parseMonkeyNum(data[i])
		case 1:
			items = parseMonkeyItems(data[i])
		case 2:
			op = parseMonkeyOperation(data[i])
		case 3:
			test = parseMonkeyTest(data[i])
		case 4:
			trueAction = parseMonkeyAction(data[i], "true")
		case 5:
			falseAction = parseMonkeyAction(data[i], "false")
			monkey := &Monkey{
				Items: items,

				Operation:   op,
				Test:        test,
				TrueTarget:  trueAction,
				FalseTarget: falseAction,
				Counter:     0,
			}
			mm[num] = monkey
		}
	}
	return mm
}

func parseMonkeyNum(line string) int {
	num := 0
	_, err := fmt.Sscanf(line, "Monkey %d:", &num)
	if err != nil {
		panic(err.Error())
	}
	return num
}

func parseMonkeyItems(line string) []int64 {
	res := make([]int64, 0)
	p := strings.Split(line, ":")
	items := strings.Split(p[1], ",")
	for _, item := range items {
		intItem, err := strconv.ParseInt(strings.Trim(item, " "), 10, 64)

		if err != nil {
			panic(err.Error())
		}
		res = append(res, intItem)
	}
	return res
}

func parseMonkeyOperation(line string) func(x int64) int64 {
	op := strings.Split(line, ":")

	op2 := strings.ReplaceAll(strings.Trim(op[1], " "), "* old", "^ 2")

	realOp := ""
	val := int64(0)

	_, err := fmt.Sscanf(op2, "new = old %s %d", &realOp, &val)
	if err != nil {
		panic(err.Error())
	}

	switch realOp {
	case "+":
		return func(x int64) int64 { return x + val }
	case "*":
		return func(x int64) int64 { return x * val }
	case "^":
		return func(x int64) int64 { return x * x }
	}

	return nil
}

func parseMonkeyTest(line string) int64 {
	test := int64(0)
	_, err := fmt.Sscanf(line, "  Test: divisible by %d", &test)
	if err != nil {
		panic(line + " " + err.Error())
	}
	return test
}

func parseMonkeyAction(line string, a string) int {
	dest := 0
	_, err := fmt.Sscanf(line, "    If "+a+": throw to monkey %d", &dest)
	if err != nil {
		panic("*" + line + "* " + err.Error())
	}
	return dest
}
