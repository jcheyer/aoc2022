package challenges

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay12p1(t *testing.T) {
	d := Day12{}
	assert.NoError(t, d.Load("testfiles/day12.txt"))

	assert.Equal(t, "31", d.Part1())

}
