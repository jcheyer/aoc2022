package challenges

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay06p1(t *testing.T) {
	d := Day06{}
	assert.NoError(t, d.Load("testfiles/day06_1.txt"))
	assert.Equal(t, "5", d.Part1())
	d = Day06{}
	assert.NoError(t, d.Load("testfiles/day06_2.txt"))
	assert.Equal(t, "6", d.Part1())
	d = Day06{}
	assert.NoError(t, d.Load("testfiles/day06_3.txt"))
	assert.Equal(t, "10", d.Part1())
	d = Day06{}
	assert.NoError(t, d.Load("testfiles/day06_4.txt"))
	assert.Equal(t, "11", d.Part1())
}

func TestDay06p2(t *testing.T) {
	d := Day06{}
	assert.NoError(t, d.Load("testfiles/day06_5.txt"))
	assert.Equal(t, "19", d.Part2())
	d = Day06{}
	assert.NoError(t, d.Load("testfiles/day06_6.txt"))
	assert.Equal(t, "23", d.Part2())
	d = Day06{}
	assert.NoError(t, d.Load("testfiles/day06_7.txt"))
	assert.Equal(t, "23", d.Part2())
	d = Day06{}
	assert.NoError(t, d.Load("testfiles/day06_8.txt"))
	assert.Equal(t, "29", d.Part2())
	d = Day06{}
	assert.NoError(t, d.Load("testfiles/day06_9.txt"))
	assert.Equal(t, "26", d.Part2())
}
