package monkeymap

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	data, err := ioutil.ReadFile("testfiles/day22.txt")
	assert.NoError(t, err)
	m, instr := New(data)
	assert.Len(t, m, 96+4*16)
	assert.Equal(t, "10R5L5R10L4R5L5", instr)

}
