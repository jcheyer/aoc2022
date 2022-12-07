package challenges

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay02p1(t *testing.T) {
	d := Day02{}
	assert.NoError(t, d.Load("testfiles/day02.txt"))
	assert.Equal(t, "15", d.Part1())
}

func TestDay02p2(t *testing.T) {
	d := Day02{}
	assert.NoError(t, d.Load("testfiles/day02.txt"))
	assert.Equal(t, "12", d.Part2())
}
