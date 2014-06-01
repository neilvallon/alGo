package graph

import (
	"fmt"
)

type Graph struct {
	Nodes []*Node
}

type Node struct {
	val   int
	Edges []*Node
}

func (n *Node) String() string {
	return fmt.Sprintf("%d", n.val)
}

func (g *Graph) Connect(n1, n2 *Node) {
	n1.Edges = append(n1.Edges, n2)
}

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
