func cloneGraph(node *Node) *Node {
	clones := map[*Node]*Node{}

	var clone func(n *Node) *Node 
	clone = func(n *Node) *Node {
		if n == nil || clones[n] != nil { return clones[n] }

		newNode := &Node{ Val: n.Val }
		clones[n] = newNode

		for _, c := range n.Neighbors {
			newNode.Neighbors = append(newNode.Neighbors, clone(c))
		}

		return newNode
	}
	
	return clone(node)
}