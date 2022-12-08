package challenges

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay08p1(t *testing.T) {
	d := Day08{}
	assert.NoError(t, d.Load("testfiles/day08.txt"))

	assert.True(t, d.IsVisible(0, 0))
	assert.True(t, d.IsVisible(1, 1))
	assert.False(t, d.IsVisible(2, 2))
	assert.False(t, d.IsVisible(3, 3))
	assert.False(t, d.IsVisible(1, 3))
	assert.True(t, d.IsVisible(2, 3))

	assert.Equal(t, "21", d.Part1())

}

func TestDay08p2(t *testing.T) {
	d := Day08{}
	assert.NoError(t, d.Load("testfiles/day08.txt"))

	assert.Equal(t, 4, d.score(1, 2))
	assert.Equal(t, 8, d.score(3, 2))

}
