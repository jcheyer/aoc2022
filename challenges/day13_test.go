package challenges

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay13p1(t *testing.T) {
	d := Day13{}
	//assert.NoError(t, d.Load("testfiles/day13.txt"))

	assert.Equal(t, "0", d.Part1())

}
