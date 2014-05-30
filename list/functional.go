package list

import "fmt"

// Head returns the first value at the front of a list.
func (l *List) Head() (i int, err error) {
	if l.IsEmpty() {
		err = fmt.Errorf("Cannot take head of empty list")
		return
	}
	i = l.first.val
	return
}


// Tail returns a new list containing all elements other than the Head.
// If no elements remain an empty list is returned.
func (l *List) Tail() (rl *List, err error) {
	switch l.Size() {
	case 0:
		err = fmt.Errorf("Cannot take tail of empty list")
	case 1:
		rl = new(List)
	default:
		first, last, size := l.first.next.clone()
		rl = &List{first, last, size}
	}
	return
}

// Clone coppies the current node value and all child nodes connected by 'next'
// into a new linked list.
func (n *Node) clone() (first *Node, last *Node, size int) {
	first = &Node{val: n.val}
	if n.next == nil {
		return first, first, 1
	}

	next, last, size := n.next.clone()
	first.next, next.prev = next, first
	size++
	return
}

// Map returns a new list with function f applied to each element.
func (l *List) Map(f func(int) int) (rl *List) {
	rl = new(List)
	if l.first == nil {
		return
	}

	prev := &Node{val: f(l.first.val)}
	rl.first = prev
	for next := l.first.next; next != nil; next = next.next {
		prev.next = &Node{val: f(next.val), prev: prev}
		prev = prev.next
	}

	rl.last, rl.size = prev, l.size
	return
}

// FlatMap returns a new list consisting of the concatenation of all lists
// returned by function f.
func (l *List) FlatMap(f func(int) []int) (rl *List) {
	rl = new(List)
	if l.first == nil {
		return
	}

	fake := new(Node)
	prev := fake
	for next := l.first; next != nil; next = next.next {
		for _, v := range f(next.val) {
			prev.next = &Node{val: v, prev: prev}
			prev = prev.next
			rl.size++
		}
	}
	fake.next.prev = nil

	rl.first, rl.last = fake.next, prev
	return
}

// Filter returns a new list containing only elements where function f is true.
func (l *List) Filter(f func(int) bool) (rl *List) {
	rl = new(List)
	if l.first == nil {
		return
	}

	fake := new(Node)
	prev := fake
	for next := l.first; next != nil; next = next.next {
		if n := next.val; f(n) {
			prev.next = &Node{val: n, prev: prev}
			prev = prev.next
			rl.size++
		}
	}
	fake.next.prev = nil

	rl.first, rl.last = fake.next, prev
	return
}

// Exists checks if function f holds for at least one element.
func (l *List) Exists(f func(int) bool) bool {
	if l.IsEmpty() {
		return false
	}

	for next := l.first; next != nil; next = next.next {
		if f(next.val) {
			return true
		}
	}

	return false
}

// ForAll checks if function f holds for all elements.
func (l *List) ForAll(f func(int) bool) bool {
	if l.IsEmpty() {
		return false
	}

	for next := l.first; next != nil; next = next.next {
		if !f(next.val) {
			return false
		}
	}

	return true
}
