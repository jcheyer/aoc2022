package challenges

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestString2Coord(t *testing.T) {
	d := Day14{}
	assert.Equal(t, d14coord{x: 498, y: 4}, d.string2coord("498,4"))
}

func TestDay14Load(t *testing.T) {
	d := Day14{}
	assert.NoError(t, d.Load("testfiles/day14.txt"))
	assert.Equal(t, 14, d.abyss)
	assert.Equal(t, 11, d.infinityLevel)
}

func TestDay14Drop(t *testing.T) {
	d := Day14{}
	assert.NoError(t, d.Load("testfiles/day14.txt"))
	c := d.dropSand(d14coord{500, 0}, false)
	assert.Equal(t, d14coord{500, 8}, c)
	c = d.dropSand(d14coord{500, 0}, false)
	assert.Equal(t, d14coord{499, 8}, c)
	c = d.dropSand(d14coord{500, 0}, false)
	assert.Equal(t, d14coord{501, 8}, c)
	for i := 0; i <= 21; i++ {
		_ = d.dropSand(d14coord{500, 0}, false)
	}

	assert.Equal(t, true, d.abyssReached)

}

func TestDay14p1(t *testing.T) {
	d := Day14{}
	assert.NoError(t, d.Load("testfiles/day14.txt"))
	assert.Equal(t, "24", d.Part1())
}

func TestDay14DropP2(t *testing.T) {
	d := Day14{
		sandSource: d14coord{500, 0},
	}
	assert.NoError(t, d.Load("testfiles/day14.txt"))

	for i := 0; i <= 93; i++ {
		_ = d.dropSand(d14coord{500, 0}, true)
	}

	assert.Equal(t, true, d.fullyFilled)

}

func TestDay14p2(t *testing.T) {
	d := Day14{
		sandSource: d14coord{500, 0},
	}
	assert.NoError(t, d.Load("testfiles/day14.txt"))
	assert.Equal(t, "93", d.Part2())
}
