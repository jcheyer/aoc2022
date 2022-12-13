package challenges

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay13p1(t *testing.T) {
	d := Day13{}
	assert.NoError(t, d.Load("testfiles/day13.txt"))

	x1, err := d.s2a("[1,1,3,1,1]")
	assert.NoError(t, err)
	x2, err := d.s2a("[1,1,5,1,1]")
	assert.NoError(t, err)
	assert.Equal(t, -2, d.solve(x1, x2))

	x1, err = d.s2a("[[1],[2,3,4]]")
	assert.NoError(t, err)
	x2, err = d.s2a("[[1],4]")
	assert.NoError(t, err)
	assert.Equal(t, -2, d.solve(x1, x2))

	x1, err = d.s2a("[9]")
	assert.NoError(t, err)
	x2, err = d.s2a("[[8,7,6]]")
	assert.NoError(t, err)
	assert.Equal(t, 1, d.solve(x1, x2))

	x1, err = d.s2a("[[4,4],4,4]")
	assert.NoError(t, err)
	x2, err = d.s2a("[[4,4],4,4,4]")
	assert.NoError(t, err)
	assert.Equal(t, -1, d.solve(x1, x2))

	x1, err = d.s2a("[7,7,7,7]")
	assert.NoError(t, err)
	x2, err = d.s2a("[7,7,7]")
	assert.NoError(t, err)
	assert.Equal(t, 1, d.solve(x1, x2))

	x1, err = d.s2a("[]")
	assert.NoError(t, err)
	x2, err = d.s2a("[3]")
	assert.NoError(t, err)
	assert.Equal(t, -1, d.solve(x1, x2))

	x1, err = d.s2a("[[[]]]")
	assert.NoError(t, err)
	x2, err = d.s2a("[[]]")
	assert.NoError(t, err)
	assert.Equal(t, 1, d.solve(x1, x2))

	x1, err = d.s2a("[1,[2,[3,[4,[5,6,7]]]],8,9]")
	assert.NoError(t, err)
	x2, err = d.s2a("[1,[2,[3,[4,[5,6,0]]]],8,9]")
	assert.NoError(t, err)
	assert.Equal(t, 7, d.solve(x1, x2))

	assert.Equal(t, "13", d.Part1())
}

func TestDay13p2(t *testing.T) {
	d := Day13{}
	assert.NoError(t, d.Load("testfiles/day13.txt"))

	assert.Equal(t, "140", d.Part2())
}
