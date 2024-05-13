package sort

//DirectionSort(slice []int) is a O(n^2) sort function
func DirectionSort(slice []int) []int {
	for i := 0; i < len(slice); i++ {
		min_index := i
		for j := i + 1; j < len(slice); j++ {
			if slice[j] <= slice[min_index] {
				min_index = j	
			}		
		}
		slice[i], slice[min_index] = slice[min_index], slice[i]
	}

	return slice
}