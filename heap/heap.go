package heap

import "errors"

type Heap struct {
	store []int
	cmp   func(i, j int) bool
}

func New(less func(i, j int) bool) *Heap {
	return &Heap{cmp: less}
}

func (h *Heap) Push(elm int) {
	h.store = append(h.store, elm)
	h.bup(len(h.store))
}

func (h *Heap) bup(i int) {
	if i == 1 {
		return
	}
	n, p := i-1, (i/2)-1

	if h.cmp(h.store[n], h.store[p]) {
		h.store[p], h.store[n] = h.store[n], h.store[p]
		h.bup(p + 1)
	}
}

func (h *Heap) Pop() (elm int, err error) {
	if len(h.store) < 1 {
		return 0, errors.New("No elements remaining.")
	}

	l := len(h.store)

	elm = h.store[0]

	h.store[0] = h.store[l-1]
	h.store = h.store[:l-1]

	h.bdown(1)
	return
}

func (h *Heap) bdown(p int) {
	i := p
	l, r := i*2, (i*2)+1

	if l <= len(h.store) && h.cmp(h.store[l-1], h.store[i-1]) {
		i = l
	}
	if r <= len(h.store) && h.cmp(h.store[r-1], h.store[i-1]) {
		i = r
	}

	if p != i {
		h.store[p-1], h.store[i-1] = h.store[i-1], h.store[p-1]
		h.bdown(i)
	}
}
