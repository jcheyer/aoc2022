package challenges

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay11p1(t *testing.T) {
	d := Day11{}
	assert.NoError(t, d.Load("internal/monkeys/testfiles/day11.txt"))

	assert.Equal(t, "10605", d.Part1())

}

func TestDay11p2(t *testing.T) {
	d := Day11{}
	assert.NoError(t, d.Load("internal/monkeys/testfiles/day11.txt"))

	assert.Equal(t, "2713310158", d.Part2())

}
