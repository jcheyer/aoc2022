package challenges

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay10p1(t *testing.T) {
	d := Day10{}
	assert.NoError(t, d.Load("internal/tube/testfiles/day10.txt"))

	assert.Equal(t, "13140", d.Part1())

	//fmt.Printf("%s", d.tube.Screen)
}
