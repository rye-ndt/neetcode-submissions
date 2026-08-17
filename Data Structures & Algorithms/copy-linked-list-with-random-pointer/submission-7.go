/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
	clone := head 

	cloneMap := map[*Node]*Node{}

	for clone != nil {
		cloneMap[clone] = &Node{
			Val: clone.Val,
		}

		clone = clone.Next
	}

	mover := cloneMap[head]
	var result *Node

	for head != nil {
		mover.Next = cloneMap[head.Next]
		mover.Random = cloneMap[head.Random]

		if result == nil { result = mover }

		head = head.Next
		mover = mover.Next
	}


	return result
}
