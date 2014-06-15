package graph

func WithPathCost(ajlst map[int]map[int]int) *Graph {
	g := &Graph{
		Nodes: make([]*Node, len(ajlst)),
	}

	// Create all nodes
	for n := range ajlst {
		g.Nodes[n-1] = &Node{
			val:       n,
			EdgeCosts: make(map[*Node]int),
		}
	}

	// Conect node edges
	for n, edges := range ajlst {
		for e, cost := range edges {
			g.Nodes[n-1].Edges = append(g.Nodes[n-1].Edges, g.Nodes[e-1])
			g.Nodes[n-1].EdgeCosts[g.Nodes[e-1]] = cost
		}
	}

	return g
}

func extractMin(ec map[*Node]int) (n *Node, minCost int) {
	minCost = 2147483647 // max int
	for e, c := range ec {
		if c < minCost {
			n, minCost = e, c
		}
	}

	delete(ec, n)
	return
}

func (g *Graph) PathsFrom(n *Node) map[*Node]int {
	front := map[*Node]int{n: 0}

	// Find all min path costs
	for len(front) != 0 {
		current, cost := extractMin(front)
		if current.explored {
			continue
		}
		current.explored = true
		current.minCost = cost

		for _, e := range current.Edges {
			pathCost := cost + current.EdgeCosts[e]
			c, ok := front[e]
			if !ok || pathCost < c {
				front[e] = pathCost
			}
		}
	}

	pathCosts := make(map[*Node]int)
	for _, n := range g.Nodes {
		pathCosts[n] = n.minCost
	}

	return pathCosts
}
