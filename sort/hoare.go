package sort

func HoareSort(slice []int, left, right int) {
	i, j := left, right
	for i <= j {
		for slice[i] < slice[left] {
			i++
		}
		for slice[j] > slice[left] {
			j--
		}

		if i <= j {
			slice[i], slice[j] = slice[j], slice[i]
			i++
			j--
		}
	}

	if left < j {
		HoareSort(slice, left, j)
	}
	if i < right {
		HoareSort(slice, i, right)
	}
}
