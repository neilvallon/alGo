package sort

func InsertionSort(ia []int) []int { // Return optional. Sorts in place
	for i := 1; i < len(ia); i++ {
		j, tmp := i, ia[i]
		for 0 < j && tmp < ia[j-1] {
			ia[j] = ia[j-1]
			j--
		}
		ia[j] = tmp
	}

	return ia
}
