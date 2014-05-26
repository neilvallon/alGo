package sort

func InplaceQS(lst []int) []int {
	if len(lst) <= 1 {
		return lst
	}

	p := lst[0]
	j := 1
	for i := 1; i < len(lst); i++ {
		if lst[i] < p {
			lst[j], lst[i] = lst[i], lst[j]
			j++
		}
	}
	lst[0], lst[j-1] = lst[j-1], lst[0]

	InplaceQS(lst[:j-1])
	InplaceQS(lst[j:])
	return lst
}
