package lib

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChunks(t *testing.T) {
	c1, err := ReadInputLinesWithEmptyLines("testfiles/chunks1.txt")
	assert.NoError(t, err)
	assert.Len(t, c1, 1032)
	c2 := Chunks(c1)
	assert.Len(t, c2, 37)
	assert.Len(t, c2[36], 27)
	assert.Equal(t, "--- scanner 36 ---", c2[36][0])
	assert.Equal(t, "721,-680,490", c2[36][1])
	assert.Equal(t, "446,-275,-410", c2[36][11])
	assert.Equal(t, "-809,-714,758", c2[36][26])

	c3, err := ReadInputLinesWithEmptyLines("testfiles/chunks2.txt")
	assert.NoError(t, err)
	assert.Len(t, c3, 102)

	c4 := Chunks(c3)
	assert.Len(t, c4, 2)
	assert.Len(t, c4[0], 1)
	assert.Len(t, c4[0][0], 512)
	assert.Len(t, c4[1], 100)

}
