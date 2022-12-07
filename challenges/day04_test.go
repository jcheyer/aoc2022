package challenges

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay04p1(t *testing.T) {
	d := Day04{}
	assert.NoError(t, d.Load("testfiles/day04.txt"))
	assert.Equal(t, "2", d.Part1())
}

func TestDay04p2(t *testing.T) {
	d := Day04{}
	assert.NoError(t, d.Load("testfiles/day04.txt"))
	assert.Equal(t, "4", d.Part2())
}
