package lib

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {
	s := NewStack()
	assert.True(t, s.IsEmpty())
	_, err := s.Pop()
	assert.Error(t, err)
	s.Push('a')
	s.Push('b')

	r, err := s.Pop()
	assert.NoError(t, err)
	assert.Equal(t, 'b', r)

	r, err = s.Pop()
	assert.NoError(t, err)
	assert.Equal(t, 'a', r)

	s.Push('c')

	r, err = s.Pop()
	assert.NoError(t, err)
	assert.Equal(t, 'c', r)

}
