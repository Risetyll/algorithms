package sort

import (
	"testing"
	"reflect"
	"math/rand"
	"sort"
)

func TestHoareSort1(t *testing.T) {
	slice := []int {5, 1, 6, 7, 8, 9, 1, 2, 3, 8}
	expected := []int {1, 1, 2, 3, 5, 6, 7, 8, 8, 9}
	HoareSort(slice, 0, len(slice) - 1)

	if !reflect.DeepEqual(slice, expected) {
		t.Fatal("Sorting error", slice)
	}
}
func TestHoareSort2(t *testing.T) {
	slice := []int {5}
	expected := []int {5}
	HoareSort(slice, 0, len(slice) - 1)

	if !reflect.DeepEqual(slice, expected) {
		t.Fatal("Sorting error", slice)
	}
}
func TestHoareSort3(t *testing.T) {
	slice := []int {5, 1}
	expected := []int {1, 5}
	HoareSort(slice, 0, len(slice) - 1)

	if !reflect.DeepEqual(slice, expected) {
		t.Fatal("Sorting error", slice)
	} 
}
func TestHoareSort4(t *testing.T) {
	slice := make([]int, 100000)
	for i := 0; i < len(slice); i++ {
		slice[i] = rand.Intn(len(slice))
	}

	HoareSort(slice, 0, len(slice) - 1)

	if !sort.SliceIsSorted(slice, func(i, j int) bool { return slice[i] < slice[j] }) {
		t.Fatal("Sorting error")
	}
}

func BenchmarkHoareSort(b *testing.B) {
	slice := make([]int, 100000)
	for i := 0; i < len(slice); i++ {
		slice[i] = rand.Intn(len(slice))
	}						
	b.ResetTimer()
	HoareSort(slice, 0, len(slice) - 1)
}