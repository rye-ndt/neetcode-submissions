/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

    oldToNew := map[*Node]*Node{}

	dummy := head 

	for dummy != nil {
		oldToNew[dummy] = &Node{
			Val: dummy.Val,
		}

		dummy = dummy.Next
	}

	dummy = head

	for dummy != nil {
		oldToNew[dummy].Next = oldToNew[dummy.Next]
		oldToNew[dummy].Random = oldToNew[dummy.Random]

		dummy = dummy.Next
	}

	return oldToNew[head]
}
