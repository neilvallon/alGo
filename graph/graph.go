package graph

import (
	"fmt"
)

type Graph struct {
	Nodes []*Node
	order []*Node
}

type Node struct {
	val   int
	Edges []*Node

	inbound  []*Node
	explored bool
}

func (n *Node) String() string {
	return fmt.Sprintf("%d", n.val)
}

func (g *Graph) Connect(n1, n2 *Node) {
	n1.Edges = append(n1.Edges, n2)
}

// Creates a new graph structure from an adjacency list.
func New(ajlst map[int][]int) *Graph {
	g := &Graph{
		Nodes: make([]*Node, len(ajlst)),
	}

	// Create all nodes
	for n := range ajlst {
		g.Nodes[n-1] = &Node{val: n}
	}

	// Conect node edges
	for n1, v := range ajlst {
		for _, n2 := range v {
			g.Connect(g.Nodes[n1-1], g.Nodes[n2-1])
		}
	}

	return g
}

func NewDirected(ajlst map[int][]int) *Graph {
	nodes := make([]*Node, len(ajlst))

	for n1v, e := range ajlst {
		if n1v > len(nodes) {
			newnodes := make([]*Node, n1v)
			copy(newnodes, nodes)
			nodes = newnodes
		}
		n1 := nodes[n1v-1]
		if n1 == nil {
			n1 = &Node{val: n1v}
			nodes[n1v-1] = n1
		}
		for _, n2v := range e {
			if n2v > len(nodes) {
				newnodes := make([]*Node, n2v)
				copy(newnodes, nodes)
				nodes = newnodes
			}
			n2 := nodes[n2v-1]
			if n2 == nil {
				n2 = &Node{val: n2v}
				nodes[n2v-1] = n2
			}
			n1.Edges = append(n1.Edges, n2)
			n2.inbound = append(n2.inbound, n1)
		}
	}

	return &Graph{
		Nodes: nodes,
		order: make([]*Node, len(nodes)),
	}
}
