package code

import (
	"fmt"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/jcheyer/aoc2021/lib"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	data, err := lib.ReadInputLines("./../../testfiles/day20.txt")

	assert.NoError(t, err)

	c := New(data)
	assert.NotNil(t, c)
	fmt.Println(c.String())
	spew.Dump(c)
	t.FailNow()
}
