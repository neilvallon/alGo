package list

import "fmt"

type Stack interface {
	Push(int)
	Pop() (int, error)
	IsEmpty() bool
	Size() int
}

type Queue interface {
	Push(int)
	Shift() (int, error)
	IsEmpty() bool
	Size() int
}

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
		return fmt.Sprintf("List(%s)", l.first.Stringrec())
	}
}

func (n *Node) String() string {
	return fmt.Sprintf("%d", n.val)
}

// Stringrec returns the values of all nodes connected to the current.
// "n", "n1", "n2", "n3", ...
func (n *Node) Stringrec() string {
	if n.next == nil {
		return fmt.Sprintf("%q", n)
	} else {
		return fmt.Sprintf("%q, %s", n, n.next.Stringrec())
	}
}

func (l *List) IsEmpty() bool {
	return l.size == 0
}

func (l *List) Size() int {
	return l.size
}

// Push adds an element to the end of the list.
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

// Pop removes and returns the last element of the list.
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

// Shift removes and returns the first element of the list.
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

// Unshift adds an element to the front of the list.
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
