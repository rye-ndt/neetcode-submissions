/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */



func copyRandomList(head *Node) *Node {
	// clone all nodes 
	clone := map[*Node]*Node{}

	preserve := head 

	for head != nil {
		clone[head] = &Node{
			Val: head.Val,
		}

		head = head.Next
	}

	head = preserve 

	for head != nil {
		clone[head].Next = clone[head.Next]
		clone[head].Random = clone[head.Random]
		
		head = head.Next
	}

	return clone[preserve]
}
