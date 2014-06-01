package graph

import (
	"math/rand"
	"time"
)

var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

// RemoveEdge deleates all edges connected to the given node
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

// ReplaceEdge Replaces all edges conected to the old node to the new node that
// it is combining with.
func (n *Node) ReplaceEdge(oedge, nedge *Node) {
	for i := range n.Edges {
		if n.Edges[i] == oedge {
			n.Edges[i] = nedge
		}
	}
}

// Removes node from the graph.
// Note: Does not remove edges conected to the given node.
func (g *Graph) RemoveNode(n *Node) {
	for i, _ := range g.Nodes {
		if g.Nodes[i] == n {
			g.Nodes[0], g.Nodes[i] = g.Nodes[i], g.Nodes[0]
		}
	}
	g.Nodes = g.Nodes[1:]
}

// Contract selects a random edge of the graph and combines the conected nodes.
// The first node selected will swallow the edges of one of its neighbors.
func (g *Graph) Contract() {
	// Pick a random node and edge to that node
	n1 := g.Nodes[rnd.Intn(len(g.Nodes))]
	n2 := n1.Edges[rnd.Intn(len(n1.Edges))]

	// Remove edges that will connect the node to itself
	n1.RemoveEdge(n2)
	n2.RemoveEdge(n1)

	// Re-link edges connected to the second node to the first
	for _, e := range n2.Edges {
		e.ReplaceEdge(n2, n1)
	}

	n1.Edges = append(n1.Edges, n2.Edges...)
	g.RemoveNode(n2)
}

// MinCut contracts random edges until the graph size is 2 nodes.
func (g *Graph) MinCut() {
	for len(g.Nodes) > 2 {
		g.Contract()
	}
}
