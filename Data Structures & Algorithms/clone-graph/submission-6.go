/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

func cloneGraph(node *Node) *Node {
	clones := map[*Node]*Node{}

	var clone func(n *Node) *Node 
	clone = func(n *Node) *Node {
		if n == nil || clones[n] != nil { return clones[n] }

		newNode := &Node{ 
			Val: n.Val,
			Neighbors: []*Node{},
		}

		clones[n] = newNode

		for _, c := range n.Neighbors {
			newNode.Neighbors = append(newNode.Neighbors, clone(c))
		}

		return newNode
	}
	
	return clone(node)
}