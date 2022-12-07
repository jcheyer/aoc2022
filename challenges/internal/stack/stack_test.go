package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {
	s := Stack[int]{}
	s.Push(9)
	assert.Len(t, s, 1)
	assert.Equal(t, 9, s.Pop())
	assert.Len(t, s, 0)
	s.Push(8)
	s.Push(7)
	assert.Equal(t, 7, s.Peek())
	assert.Len(t, s, 2)
	assert.Equal(t, 7, s.Pop())
	assert.Len(t, s, 1)
	assert.Equal(t, 8, s.Pop())
	assert.Len(t, s, 0)
	s.Push(6)
	s.Push(5)
	s.Push(4)
	drei := s.PopN(3)
	assert.Len(t, drei, 3)
	assert.Len(t, s, 0)
	assert.Equal(t, []int{6, 5, 4}, drei)
	s.PushN(drei)
	assert.Len(t, s, 3)
	assert.Equal(t, 4, s.Pop())
	assert.Equal(t, 5, s.Pop())
	assert.Equal(t, 6, s.Pop())

	assert.Len(t, s, 0)
	s.Push(4)
	s.SneakPush(3)
	assert.Len(t, s, 2)
	assert.Equal(t, 4, s.Peek())

	assert.Equal(t, 4, s.Pop())
	assert.Equal(t, 3, s.Pop())
	assert.Len(t, s, 0)

	s.SneakPush(2)
	assert.Len(t, s, 1)
	assert.Equal(t, 2, s.Pop())
}
