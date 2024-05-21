package structures

type Stack struct {
	top *node
}

type node struct {
	value int
	prev  *node
}

func (s *Stack) IsEmpty() bool {
	return s.top == nil
}

func (s *Stack) Push(value int) {
	if s.IsEmpty() {
		s.top = &node{value: value}
		return
	}

	prev := *s.top
	s.top = &node{value: value, prev: &prev}
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
	s.top = s.top.prev

	return value
}

/* func (s *Stack) Pop() int {
}
*/

/* func (s *Stack) Push(value int) {
	prev := *s
	prev.value = value
}

func (s *Stack) Peek() int {
	return s.value
}

func (s *Stack) Pop() int {
	top := s.Peek()
	*s = *s.prev

	return top
}
*/
