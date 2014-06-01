package graph

import (
	"math/rand"
	"time"
)

var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

func (n *Node) RemoveEdge(n2 *Node) {
	var j int
	for i := range n.Edges {
		if n.Edges[i] == n2 {
			n.Edges[j], n.Edges[i] = n.Edges[i], n.Edges[j]
			j++
		}
	}

	n.Edges = n.Edges[j:]
}

func (n *Node) ReplaceEdge(oedge, nedge *Node) {
	for i := range n.Edges {
		if n.Edges[i] == oedge {
			n.Edges[i] = nedge
		}
	}
}

func (g *Graph) RemoveNode(n *Node) {
	for i, _ := range g.Nodes {
		if g.Nodes[i] == n {
			g.Nodes[0], g.Nodes[i] = g.Nodes[i], g.Nodes[0]
		}
	}
	g.Nodes = g.Nodes[1:]
}

func (g *Graph) Contract() {
	n1 := g.Nodes[rnd.Intn(len(g.Nodes))]
	n2 := n1.Edges[rnd.Intn(len(n1.Edges))]

	n1.RemoveEdge(n2)
	n2.RemoveEdge(n1)

	for _, e := range n2.Edges {
		e.ReplaceEdge(n2, n1)
	}

	n1.Edges = append(n1.Edges, n2.Edges...)
	g.RemoveNode(n2)
}

func (g *Graph) MinCut() {
	for len(g.Nodes) > 2 {
		g.Contract()
	}
}
