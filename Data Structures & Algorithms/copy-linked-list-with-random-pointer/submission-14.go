/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(h *Node) *Node {
	mapper := map[*Node]*Node{}

	for i := h; i != nil; i = i.Next {
		mapper[i] = &Node{ Val: i.Val }
	}

	result := mapper[h]

	for i := result; h != nil; h = h.Next {
		i.Next = mapper[h.Next]
		i.Random = mapper[h.Random]
		i = i.Next
	}

	return result
}
