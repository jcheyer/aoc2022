package challenges

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay05p1(t *testing.T) {
	d := Day05{}
	assert.NoError(t, d.Load("testfiles/day05.txt"))
	assert.Equal(t, "CMZ", d.Part1())
}

func TestDay05p2(t *testing.T) {
	d := Day05{}
	assert.NoError(t, d.Load("testfiles/day05.txt"))
	assert.Equal(t, "MCD", d.Part2())
}
