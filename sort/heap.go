package sort

//HeapSort(slice []int) is a O(nlogn) sort function
func HeapSort(slice []int) []int {
	buildMaxHeap(slice)

	for i := len(slice) - 1; i >= 0; i-- {
		slice[0], slice[i] = slice[i], slice[0]
		heapify(slice[0:i], 0)
	}

	return slice
}

func buildMaxHeap(slice []int) {
	for i := len(slice) / 2; i >= 0; i-- {
		heapify(slice, i)
	}
}

func heapify(slice []int, index int) {
	left, right := 2*index+1, 2*index+2
	var max int

	if left < len(slice) && slice[left] > slice[index] {
		max = left
	} else {
		max = index
	}

	if right < len(slice) && slice[right] > slice[max] {
		max = right
	}

	if max != index {
		slice[index], slice[max] = slice[max], slice[index]
		heapify(slice, max)
	}
}
