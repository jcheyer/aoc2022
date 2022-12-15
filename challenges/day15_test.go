package challenges

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay15p1(t *testing.T) {
	d := Day15{}
	assert.NoError(t, d.Load("internal/radar/testfiles/day15.txt"))
	d.p1Line = 10
	assert.Equal(t, "26", d.Part1())

}

func TestDay15p2(t *testing.T) {
	d := Day15{}
	assert.NoError(t, d.Load("internal/radar/testfiles/day15.txt"))
	d.p2min = 0
	d.p2max = 20
	assert.Equal(t, "56000011", d.Part2())

}
