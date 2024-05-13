package sort

import (
	"testing"
	"reflect"
	"math/rand"
)

func TestBubbleSort1(t *testing.T) {
	slice := []int {5, 1, 6, 7, 8, 9, 1, 2, 3, 8}
	expected := []int {1, 1, 2, 3, 5, 6, 7, 8, 8, 9}
	slice = BubbleSort(slice)

	if !reflect.DeepEqual(slice, expected) {
		t.Fatal("Sorting error", slice)
	}
}
func TestBubbleSort2(t *testing.T) {
	slice := []int {5}
	expected := []int {5}
	slice = BubbleSort(slice)

	if !reflect.DeepEqual(slice, expected) {
		t.Fatal("Sorting error", slice)
	}
}
func TestBubbleSort3(t *testing.T) {
	slice := []int {5, 1}
	expected := []int {1, 5}
	slice = BubbleSort(slice)

	if !reflect.DeepEqual(slice, expected) {
		t.Fatal("Sorting error", slice)
	}
}

func BenchmarkBubbleSort(b *testing.B) {
	slice := make([]int, 100000)
	for i := 0; i < len(slice); i++ {
		slice[i] = rand.Intn(len(slice))
	}						
	b.ResetTimer()
	_  = BubbleSort(slice)
}