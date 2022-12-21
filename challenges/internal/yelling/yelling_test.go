package yelling

import (
	"testing"

	"github.com/jcheyer/aoc2021/lib"
	"github.com/stretchr/testify/assert"
)

func TestParseLine(t *testing.T) {
	n, m := parseLine("root: pppw + sjmn")
	assert.Equal(t, "root", n)
	assert.Equal(t, "root", m.name)
	assert.Equal(t, "pppw", m.v1)
	assert.Equal(t, "sjmn", m.v2)
	assert.Equal(t, 4, m.Op(1, 3))

	n, m = parseLine("dbpl: 5")
	assert.Equal(t, "dbpl", n)
	assert.Equal(t, 5, m.Result)

	_, m = parseLine("ptdq: humn - dvpt")
	assert.Equal(t, 6, m.Op(8, 2))

	_, m = parseLine("sjmn: drzm * dbpl")
	assert.Equal(t, 8, m.Op(2, 4))

	_, m = parseLine("pppw: cczh / lfqf")
	assert.Equal(t, 4, m.Op(12, 3))

}

func TestSolve(t *testing.T) {
	data, err := lib.ReadInputLines("testfiles/day21.txt")
	assert.NoError(t, err)
	m := New(data)
	res := Solve(m, "root")
	assert.Equal(t, 152, res)
}
