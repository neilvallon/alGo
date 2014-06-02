package sort

func QuickSelect(lst []int, ith int) int {
	if ith <= 0 || len(lst) < ith {
		panic("Selection element out of range.")
	}

	rp := rnd.Intn(len(lst))
	lst[0], lst[rp] = lst[rp], lst[0]

	p := lst[0]
	j := 1
	for i := 1; i < len(lst); i++ {
		if lst[i] < p {
			lst[j], lst[i] = lst[i], lst[j]
			j++
		}
	}

	switch {
	case ith < j:
		return QuickSelect(lst[1:j], ith)
	case ith > j:
		return QuickSelect(lst[j:], ith-j)
	case ith == j:
		return p
	}

	panic("Something went horribly wrong")
}
