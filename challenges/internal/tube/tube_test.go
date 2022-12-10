package tube

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckSignalStrength(t *testing.T) {
	tube := New([]string{})
	tube.checkSignalStrength()
	assert.Len(t, tube.SignalStrength, 0)
	tube.cycle = 20
	tube.checkSignalStrength()
	assert.Len(t, tube.SignalStrength, 1)
	tube.cycle = 21
	tube.checkSignalStrength()
	assert.Len(t, tube.SignalStrength, 1)
	tube.cycle = 40
	tube.checkSignalStrength()
	assert.Len(t, tube.SignalStrength, 1)
	tube.cycle = 60
	tube.checkSignalStrength()
	assert.Len(t, tube.SignalStrength, 2)
	tube.cycle = 61
	tube.checkSignalStrength()
	assert.Len(t, tube.SignalStrength, 2)
	tube.cycle = 100
	tube.checkSignalStrength()
	assert.Len(t, tube.SignalStrength, 3)

}
