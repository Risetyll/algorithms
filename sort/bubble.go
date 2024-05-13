package sort 


// BubbleSort(slice []int) is a O(n^2) sort function
func BubbleSort(slice []int) []int {
	for i := 0; i < len(slice); i++ {
		for j := len(slice) - 1; j > i; j-- {
			if slice[j] < slice[j - 1] {
				slice[j], slice[j - 1] = slice[j - 1], slice[j]
			}
		}
	}

	return slice
}