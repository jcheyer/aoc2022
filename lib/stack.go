package lib

import "errors"

type stack struct {
	top  int
	data []rune
}

func NewStack() *stack {
	s := stack{-1, []rune{}}
	return &s
}

func (s *stack) IsEmpty() bool {
	return s.top == -1
}

func (s *stack) Push(v rune) {
	if s.top+1 == len(s.data) {
		s.data = append(s.data, v)
	} else {
		s.data[s.top+1] = v
	}
	s.top++
}

func (s *stack) Pop() (rune, error) {
	if s.top == -1 {
		return rune(0), errors.New("stack is empty")
	}
	s.top--
	return s.data[s.top+1], nil
}
