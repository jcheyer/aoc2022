package challenges

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay09bla(t *testing.T) {
	d := Day09{}
	assert.NoError(t, d.Load("testfiles/day09.txt"))
	//assert.Equal(t, 5, d.grid.size.x)
	//assert.Equal(t, 4, d.grid.size.y)
	d.grid.move("R 4")
	//assert.Equal(t, 1, 2)
}

func TestDay09p1(t *testing.T) {
	d := Day09{}
	assert.NoError(t, d.Load("testfiles/day09.txt"))
	assert.Equal(t, "13", d.Part1())
}

func TestDay09p2(t *testing.T) {
	d := Day09{}
	assert.NoError(t, d.Load("testfiles/day09_2.txt"))
	//assert.Equal(t, "13", d.Part1())
	assert.NotEqual(t, "fdsafdsa", d.Part1())
	assert.Equal(t, "36", d.Part2())

}
