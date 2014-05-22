package list

import "fmt"

type List struct {
	first *Node
	last  *Node
	size  int
}

type Node struct {
	next *Node
	prev *Node
	val  int
}

func (l *List) String() string {
	if l.IsEmpty() {
		return "List()"
	} else {
		return fmt.Sprintf("List(%s)", l.first.Stringl())
	}
}

func (n *Node) String() string {
	return fmt.Sprintf("%d", n.val)
}

func (n *Node) Stringl() string {
	if n.next == nil {
		return fmt.Sprintf("%q", n)
	} else {
		return fmt.Sprintf("%q, %s", n, n.next.Stringl())
	}
}

func (l *List) IsEmpty() bool {
	return l.size == 0
}

func (l *List) Push(i int) {
	n := &Node{val: i}

	if l.IsEmpty() {
		l.first = n
		l.last = n
	} else {
		n.prev = l.last
		l.last.next = n
		l.last = n
	}

	l.size++
}

func (l *List) Pop() (i int, err error) {
	if l.IsEmpty() {
		err = fmt.Errorf("Pop from empty list")
		return
	}

	i = l.last.val
	l.last = l.last.prev
	if l.last != nil {
		l.last.next = nil
	}

	l.size--
	return
}

func (l *List) Shift() (i int, err error) {
	if l.IsEmpty() {
		err = fmt.Errorf("Shift from empty list")
		return
	}

	i = l.first.val
	l.first = l.first.next
	if l.first != nil {
		l.first.prev = nil
	}

	l.size--
	return
}

func (l *List) Unshift(i int) {
	n := &Node{val: i}

	if l.IsEmpty() {
		l.first = n
		l.last = n
	} else {
		n.next = l.first
		l.first.prev = n
		l.first = n
	}

	l.size++
}
