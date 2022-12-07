package challenges

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay01P1(t *testing.T) {
	d := Day01{}
	assert.NoError(t, d.Load("testfiles/day01.txt"))
	assert.Len(t, d.calories, 5)
}
