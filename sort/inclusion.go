package sort 

//InclusionSort(slice []int) is a O(n^2) sort function
func InclusionSort(slice []int) []int {
	for i := 1; i < len(slice); i++ {
		for j := i; j > 0; j-- { 
			if slice[j - 1] > slice[j] {
				slice[j - 1], slice[j] = slice[j], slice[j - 1]
			}
		}
	}

	return slice
}