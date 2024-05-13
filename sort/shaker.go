package sort

//ShakerSort(slice []int) is a O(n^2) modification of BubbleSort 
func ShakerSort(slice []int) []int {
	left, right := 0, len(slice)

	for left < right {
		for i := right - 1; i > left; i-- {
			if slice[i] <= slice[i-1] {
				slice[i], slice[i-1] = slice[i-1], slice[i]
			}
		}
		left++

		for i := left; i < right - 1; i++ {
			if slice[i] >= slice[i+1] {
				slice[i], slice[i+1] = slice[i+1], slice[i]
			}
		}
		right--
	}
	return slice
}