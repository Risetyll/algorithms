package structures

type Stack struct {
	top *node
}

func (s *Stack) IsEmpty() bool {
	return s.top == nil
}

func (s *Stack) Push(value int) {
	if s.IsEmpty() {
		s.top = &node{value: value}
		return
	}

	next := *s.top
	s.top = &node{value: value, next: &next}
}

func (s *Stack) Peek() int {
	if s.IsEmpty() {
		panic("empty stack")
	}

	return s.top.value
}

func (s *Stack) Pop() int {
	if s.IsEmpty() {
		panic("empty stack")
	}

	value := s.Peek()
	s.top = s.top.next

	return value
}
