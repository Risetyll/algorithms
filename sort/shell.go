package sort

//ShellSort(slice []int, shift int) is a O(n^1.2) sort modification of InclusionSort algorithm 
func ShellSort(slice []int, shift int) []int {
	for k := shift; k > 0; k-- {
		for i := k; i < len(slice); i += k {
			for j := i; j > 0; j -= k {
				if slice[j-k] > slice[j] {
					slice[j-k], slice[j] = slice[j], slice[j-k]
				}
			}
		}
	}

	return slice
}
