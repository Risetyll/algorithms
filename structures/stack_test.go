package structures

import (
	"reflect"
	"testing"
)

func TestStackInitialization(t *testing.T) {
	var s1 Stack

	if s1.top != nil {
		t.Errorf("stack initialization error: got: %v, want: %v", s1.top, nil)
	}

	var s2 = Stack{}
	if s2.top != nil {
		t.Errorf("stack initialization error: got %v, want  %v", s1.top, nil)
	}
}

func TestStackPush1(t *testing.T) {
	tests := []struct {
		name  string
		input int
		want  int
	}{
		{"check top node value #1", 5, 5},
		{"check top node value #2", 6, 6},
		{"check top node value #3", 1, 1},
		{"check top node value #4", 2, 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Stack{}
			s.Push(tt.input)

			if val := s.top.value; val != tt.want {
				t.Errorf("pushing error: want %v, got %v\n", tt.want, val)
			}
		})
	}
}

func TestStackPush2(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  []int
	}{
		{"check stack tail after push #1",
			[]int{5, 5, 1, 6, 7, 8},
			[]int{8, 7, 6, 1, 5, 5}},
		{"check stack tail after push #2",
			[]int{5},
			[]int{5},
		},
		{"check stack tail after push #3",
			[]int{1, 7},
			[]int{7, 1},
		},
		{"check stack tail after push #4",
			[]int{},
			[]int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Stack{}

			if s.top != nil {
				t.Errorf("stack top should be nil")
			}

			for _, v := range tt.input {
				s.Push(v)
			}

			tail := []int{}
			for s.top != nil {
				tail = append(tail, s.top.value)
				s.top = s.top.prev
			}

			if !reflect.DeepEqual(tail, tt.want) {
				t.Errorf("push error: got %v, want %v\n", tail, tt.want)
			}
		})
	}

}

func TestStackPeek(t *testing.T) {
	s := Stack{}
	s.Push(10)
	
	if s.Peek() != 10  && s.top.value != 10 {
		t.Errorf("peek error: got %v, want %v\n", s.Peek(), 10)
	}
}

func TestStackPop(t *testing.T) {
	s := Stack{}
	s.Push(10)

	if val := s.Pop(); val != 10 && s.top.value == 10 { 
		t.Errorf("pop error: got %v, want %v\n", val, 10)
	}
}
