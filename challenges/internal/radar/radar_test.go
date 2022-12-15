package radar

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseLine(t *testing.T) {
	s, err := parseLine("Sensor at x=2, y=18: closest beacon is at x=-2, y=15")
	assert.NoError(t, err)
	assert.Equal(t, coord{2, 18}, s.Position)
	assert.Equal(t, coord{-2, 15}, s.Beacon)
	assert.Equal(t, 7, s.Distance)

	s, err = parseLine("Sensor at x=9, y=16: closest beacon is at x=10, y=16")
	assert.NoError(t, err)
	assert.Equal(t, coord{9, 16}, s.Position)
	assert.Equal(t, coord{10, 16}, s.Beacon)
	assert.Equal(t, 1, s.Distance)

}
