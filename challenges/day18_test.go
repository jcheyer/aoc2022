package challenges

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay18Load(t *testing.T) {
	d := Day18{}
	assert.NoError(t, d.Load("testfiles/day18.txt"))
	cubes := d.generateCubes()
	assert.Len(t, cubes, 13)
	assert.Equal(t, d18Cube{2, 2, 4}, cubes[7])
}

func TestDay18p1(t *testing.T) {
	d := Day18{}
	assert.NoError(t, d.Load("testfiles/day18.txt"))
	assert.Equal(t, "64", d.Part1())
}
