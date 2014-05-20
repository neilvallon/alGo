package sort

func partition(lst []int, pivot int) (left []int, right []int) {
	for _, n := range lst {
		if n <= pivot {
			left = append(left, n)
		} else {
			right = append(right, n)
		}
	}
	return
}

func qsmerge(l []int, r []int, pivot int) []int {
	s := len(l) + len(r) + 1
	ret := make([]int, s)

	copy(ret, l)
	ret[len(l)] = pivot
	copy(ret[len(l)+1:], r)

	return ret
}

func QuickSort(lst []int) []int {
	if len(lst) <= 1 {
		return lst
	}

	pivot := lst[0]
	l, r := partition(lst[1:], pivot)

	return qsmerge(QuickSort(l), QuickSort(r), pivot)
}
