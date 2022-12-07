package challenges

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay03p1(t *testing.T) {
	d := Day03{}
	assert.NoError(t, d.Load("testfiles/day03.txt"))
	assert.Equal(t, "157", d.Part1())
}

func TestDay03p2(t *testing.T) {
	d := Day03{}
	assert.NoError(t, d.Load("testfiles/day03.txt"))
	assert.Equal(t, "70", d.Part2())
}
