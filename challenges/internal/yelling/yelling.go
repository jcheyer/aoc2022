package yelling

import (
	"fmt"
	"strings"
)

type MonkeyOp func(a, b int) int

type Monkey struct {
	name   string
	v1, v2 string
	op     MonkeyOp
	result int
	done   bool
}

type Monkeys map[string]Monkey

func New(data []string) Monkeys {
	monkeys := make(Monkeys)
	for _, line := range data {
		name, monkey := parseLine(line)
		monkeys[name] = monkey
	}
	return monkeys
}

func parseLine(line string) (string, Monkey) {
	m := Monkey{}

	line = strings.ReplaceAll(line, ":", " : ")

	_, err := fmt.Sscanf(line, "%s : %d", &m.name, &m.result)
	if err == nil {
		m.done = true
		return m.name, m
	}

	op := ""
	_, err = fmt.Sscanf(line, "%s : %s %s %s", &m.name, &m.v1, &op, &m.v2)
	if err != nil {
		panic(err.Error())
	}
	switch op {
	case "*":
		m.op = func(a, b int) int { return a * b }
	case "+":
		m.op = func(a, b int) int { return a + b }
	case "-":
		m.op = func(a, b int) int { return a - b }
	case "/":
		m.op = func(a, b int) int { return a / b }
	}
	return m.name, m
}

func Solve(m Monkeys, monkey string) int {
	return m[monkey].Solve(m)
}

func (monkey Monkey) Solve(m Monkeys) int {
	if monkey.done {
		return monkey.result
	}

	return monkey.op(Solve(m, monkey.v1), Solve(m, monkey.v2))
}
