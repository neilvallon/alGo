package sort

func merge(r []int, l []int) (ret []int) {
	rlen := len(r)
	llen := len(l)
	retlen := rlen + llen

	ret = make([]int, retlen)

	var j, k int
	for i := 0; i < retlen; i++ {
		if j == rlen {
			copy(ret[i:], l[k:])
			return
		} else if k == llen {
			copy(ret[i:], r[j:])
			return
		}

		if r[j] < l[k] {
			ret[i] = r[j]
			j++
		} else {
			ret[i] = l[k]
			k++
		}
	}

	return
}

func MergeSort(ia []int) []int {
	l := len(ia)
	if l <= 1 {
		return ia
	}

	ra := MergeSort(ia[:l/2])
	la := MergeSort(ia[l/2:])

	return merge(ra, la)
}
