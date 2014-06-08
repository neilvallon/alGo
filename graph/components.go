package graph

var ftime = 0

func (g *Graph) dfs(n *Node) {
	n.explored = true
	for _, n2 := range n.Edges {
		if !n2.explored {
			g.dfs(n2)
		}
	}

	n.finishing = ftime
	g.order[ftime] = n

	ftime++
}

func (g *Graph) scc(n *Node) (SCC []*Node) {
	n.explored = true
	for _, n2 := range n.Edges {
		if !n2.explored {
			SCC = append(SCC, g.scc(n2)...)
		}
	}
	return append(SCC, n)
}

func (g *Graph) Components() (components [][]*Node) {
	for _, node := range g.Nodes {
		if !node.explored {
			g.dfs(node)
		}
	}

	// Transpose Graph
	g.order, g.Nodes = g.Nodes, g.order
	g.Transpose()
	//

	for i := len(g.Nodes) - 1; 0 <= i; i-- {
		node := g.Nodes[i]
		if !node.explored {
			components = append(components, g.scc(node))
		}
	}

	return
}

func (g *Graph) Transpose() {
	for _, node := range g.Nodes {
		node.explored = false // reset state
		node.Edges, node.inbound = node.inbound, node.Edges
	}
}
