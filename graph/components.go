package graph

func (g *Graph) dfs(n *Node, fin *[]*Node) {
	n.explored = true
	for _, n2 := range n.Edges {
		if !n2.explored {
			g.dfs(n2, fin)
		}
	}

	*fin = append(*fin, n)
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
	finOrd := make([]*Node, 0, len(g.Nodes))
	for _, node := range g.Nodes {
		if !node.explored {
			g.dfs(node, &finOrd)
		}
	}

	g.Transpose()
	for i := len(finOrd) - 1; 0 <= i; i-- {
		node := finOrd[i]
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
