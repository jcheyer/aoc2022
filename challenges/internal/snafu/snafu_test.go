package snafu

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestS2D(t *testing.T) {
	assert.Equal(t, 1747, Snafu2Dec("1=-0-2"))
	assert.Equal(t, 906, Snafu2Dec("12111"))
}

func TestD2S(t *testing.T) {
	assert.Equal(t, "1=-0-2", Dec2Snafu(1747))
}
