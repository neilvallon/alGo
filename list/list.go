package list

import "fmt"

type List struct {
	first *Node
	last  *Node
	size  int
}

type Node struct {
	next *Node
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
		l.last.next = n
		l.last = n
	}

	l.size++
}
