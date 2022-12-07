package stack

type Stack[C any] []C

func (s *Stack[C]) Push(v C) {
	*s = append(*s, v)
}

func (s *Stack[C]) PushN(v []C) {
	*s = append(*s, v...)
}

func (s *Stack[C]) Pop() C {
	if len(*s) == 0 {
		panic("stack is empty")
	}
	v := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return v
}

func (s *Stack[C]) PopN(n int) []C {
	if len(*s) < n {
		panic("stack is empty")
	}
	v := (*s)[len(*s)-n:]
	*s = (*s)[:len(*s)-n]
	return v
}

func (s *Stack[C]) Peek() C {
	return (*s)[len(*s)-1]
}

func (s *Stack[C]) SneakPush(v C) {
	*s = append([]C{v}, *s...)
}
