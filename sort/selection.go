package sort

func SelectionSort(ia []int) []int { // Return optional. Sorts in place
	l := len(ia)
	for i := 0; i < l-1; i++ {
		k := i
		for j := i + 1; j < l; j++ {
			if ia[j] < ia[k] {
				k = j
			}
		}
		ia[i], ia[k] = ia[k], ia[i]
	}

	return ia
}
