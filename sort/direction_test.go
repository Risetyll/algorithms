package sort

import (
	"testing"
	"reflect"
	"math/rand"
)

func TestDirectionSort1(t *testing.T) {
	slice := []int {5, 1, 6, 7, 8, 9, 1, 2, 3, 8}
	expected := []int {1, 1, 2, 3, 5, 6, 7, 8, 8, 9}
	slice = DirectionSort(slice)

	if !reflect.DeepEqual(slice, expected) {
		t.Fatal("Sorting error")
	}
}
func TestDirectionSort2(t *testing.T) {
	slice := []int {5}
	expected := []int {5}
	slice = DirectionSort(slice)

	if !reflect.DeepEqual(slice, expected) {
		t.Fatal("Sorting error")
	}
}
func TestDirectionSort3(t *testing.T) {
	slice := []int {5, 1}
	expected := []int {1, 5}
	slice = DirectionSort(slice)

	if !reflect.DeepEqual(slice, expected) {
		t.Fatal("Sorting error")
	}
}

func BenchmarkDirectionSort(b *testing.B) {
	slice := make([]int, 100000)
	for i := 0; i < len(slice); i++ {
		slice[i] = rand.Intn(len(slice))
	}						
	b.ResetTimer()
	_ = DirectionSort(slice)
}