package sort

import (
	"testing"
	"reflect"
	"math/rand"
)

func TestInclusionSort1(t *testing.T) {
	slice := []int {5, 1, 6, 7, 8, 9, 1, 2, 3, 8}
	expected := []int {1, 1, 2, 3, 5, 6, 7, 8, 8, 9}
	slice = InclusionSort(slice)

	if !reflect.DeepEqual(slice, expected) {
		t.Fatal("Sorting error", slice)
	}
}
func TestInclusionSort2(t *testing.T) {
	slice := []int {5}
	expected := []int {5}
	slice = InclusionSort(slice)

	if !reflect.DeepEqual(slice, expected) {
		t.Fatal("Sorting error", slice)
	}
}
func TestInclusionSort3(t *testing.T) {
	slice := []int {5, 1}
	expected := []int {1, 5}
	slice = InclusionSort(slice)

	if !reflect.DeepEqual(slice, expected) {
		t.Fatal("Sorting error", slice)
	}
}

func BenchmarkInclusionSort(b *testing.B) {
	slice := make([]int, 100000)
	for i := 0; i < len(slice); i++ {
		slice[i] = rand.Intn(len(slice))
	}						
	b.ResetTimer()
	_  = ShellSort(slice, 3)
}