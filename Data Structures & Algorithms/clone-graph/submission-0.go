/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

func cloneGraph(node *Node) *Node {
	seen := map[*Node]*Node{}

	var dfs func(cur *Node) *Node

	dfs = func(cur *Node) *Node{
		if cur == nil {
			return nil
		}

		if c, found := seen[cur]; found {
			return c
		}

		clone := &Node{
			Val: cur.Val,
		}

		seen[cur] = clone

		for _, n := range cur.Neighbors {
			clone.Neighbors = append(clone.Neighbors, dfs(n))
		}

		return clone
	}

	return dfs(node)
}
