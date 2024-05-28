package structures

type List struct {
	head *node
}

func (l *List) IsEmpty() bool {
	return l.head == nil
}

func (l *List) Push(val int) {
	new_node := &node{value: val}

	if l.head == nil {
		l.head = new_node
		return
	}

	buf := l.head

	for buf.next != nil {
		buf = buf.next
	}

	buf.next = new_node
}

func (l *List) PopFirst() int {
	if l.head == nil {
		panic("empty list")
	}

	firstVal := l.head.value
	l.head = l.head.next

	return firstVal
}

func (l *List) PopLast() int {
	if l.head == nil {
		panic("empty list")
	}

	if l.head.next == nil {
		value := l.head.value
		l.head = nil

		return value
	}

	buf := l.head

	for buf.next.next != nil {
		buf = buf.next
	}

	value := buf.next.value

	buf.next = nil

	return value
}

func (l *List) GetValue(index int) int {
	if index > l.Count() {
		panic("out of range")
	}

	buf := l.head

	for i := 0; i < index; i++ {
		buf = buf.next
	}

	return buf.value
}

func (l *List) Count() int {
	count := 0

	buf := l.head

	if buf == nil {
		return count
	}

	for buf != nil {
		buf = buf.next
		count++
	}

	return count
}
