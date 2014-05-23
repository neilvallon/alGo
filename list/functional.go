package list

import "fmt"

func (l *List) Head() (i int, err error) {
	if l.IsEmpty() {
		err = fmt.Errorf("Cannot take head of empty list")
		return
	}
	i = l.first.val
	return
}

func (l *List) Tail() (rl *List, err error) {
	switch l.Size() {
	case 0:
		err = fmt.Errorf("Cannot take tail of empty list")
	case 1:
		// Return nil list. No error
	default:
		first, last, size := l.first.next.clone()
		rl = &List{first, last, size}
	}
	return
}

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
