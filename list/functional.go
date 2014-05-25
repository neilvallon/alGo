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
