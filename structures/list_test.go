package structures

import (
	"testing"
	//"reflect"
)

func TestListPush(t *testing.T) {
	list := List{}
	slice := []int{5, 1, 6, 7, 8, 1, 7, 9, 3, 6, 4}

	for i := 0; i < len(slice); i++ {
		list.Push(slice[i])
	}

	for i := 0; i < len(slice); i++ {
		if slice[i] != list.head.value {
			t.Errorf("push error: got %v, want %v", list.head.value, slice[i])
		}
		list.head = list.head.next
	}
}

func TestListPopFirst(t *testing.T) {
	list := List{}
	slice := []int{5, 1, 6, 7, 8, 1, 7, 9, 3, 6, 4}

	for i := 0; i < len(slice); i++ {
		list.Push(slice[i])
	}

	for i := 0; i < len(slice); i++ {
		if value := list.PopFirst(); slice[i] != value {
			t.Errorf("pop first error: got %v, want %v", value, slice[i])
		}
	}
}

func TestListPopLast(t *testing.T) {
	list := List{}
	slice := []int{5, 1, 6, 7, 8, 1, 7, 9, 3, 6, 4}

	for i := 0; i < len(slice); i++ {
		list.Push(slice[i])
	}

	for i := len(slice) - 1; i >= 0; i-- {
		if value := list.PopLast(); slice[i] != value {
			t.Errorf("pop last error: got %v, want %v", value, slice[i])
		}
	}
}

func TestListIsEmpty(t *testing.T) {
	list := List{}

	if list.IsEmpty() == false {
		t.Errorf("is empty error: got %v, want %v", list.IsEmpty(), true)
	}

	list.Push(5)
	if list.IsEmpty() == true {
		t.Errorf("is empty error: got %v, want %v", list.IsEmpty(), false)
	}
}

func TestListGetValue(t *testing.T) {
	list := List{}
	slice := []int{5, 1, 6, 7, 8, 1, 7, 9, 3, 6, 4}

	for i := 0; i < len(slice); i++ {
		list.Push(slice[i])
	}

	if list.GetValue(0) != slice[0] {
		t.Errorf("get value error: got %v, want %v", list.GetValue(0), slice[0])
	}

	if list.GetValue(3) != slice[3] {
		t.Errorf("get value error: got %v, want %v", list.GetValue(3), slice[3])
	}

	if list.GetValue(10) != slice[10] {
		t.Errorf("get value error: got %v, want %v", list.GetValue(10), slice[10])
	}
}

func TestListCount(t *testing.T) {
	list := List{}

	if list.Count() != 0 {
		t.Errorf("count error: got %v, want %v", list.Count(), 0)
	}

	list.Push(6)
	list.Push(4)
	list.Push(3)

	if list.Count() != 3 {
		t.Errorf("count error: got %v, want %v", list.Count(), 3)
	}
}
