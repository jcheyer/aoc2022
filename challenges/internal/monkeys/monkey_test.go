package monkeys

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseMonkeyNumber(t *testing.T) {
	assert.Equal(t, 1, parseMonkeyNum("Monkey 1:"))
	assert.Equal(t, 3, parseMonkeyNum("Monkey 3:"))
}

func TestParseMonkeyItems(t *testing.T) {
	assert.Equal(t, []int64{79, 98}, parseMonkeyItems("  Starting items: 79, 98"))
	assert.Equal(t, []int64{54, 65, 75, 74}, parseMonkeyItems("Starting items: 54, 65, 75, 74"))
}

func TestParseMonkeyOperation(t *testing.T) {
	f := parseMonkeyOperation("  Operation: new = old * 19")
	assert.Equal(t, int64(38), f(2))

	f = parseMonkeyOperation("  Operation: new = old * old")
	assert.Equal(t, int64(4), f(2))

	f = parseMonkeyOperation("  Operation: new = old + 123")
	assert.Equal(t, int64(125), f(2))
}

func TestParseMonkeyTest(t *testing.T) {
	assert.Equal(t, int64(19), parseMonkeyTest("  Test: divisible by 19"))
}

func TestParseMonkeyAction(t *testing.T) {
	assert.Equal(t, 1, parseMonkeyAction("    If true: throw to monkey 1", "true"))
	assert.Equal(t, 3, parseMonkeyAction("    If false: throw to monkey 3", "false"))
}

func TestMonkeyParser(t *testing.T) {

}
