/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
	// 2 parts: store the values - wire them together 

	mapper := map[*Node]*Node{}

	dummy := head // copy the whole address. do not use head directly, to preserve the head 

	// clone
	for dummy != nil {
		mapper[dummy] = &Node{
			Val: dummy.Val,
		}

		dummy = dummy.Next
	}

	// wire 
	dummy = head // reset

	for dummy != nil {
		mapper[dummy].Next = mapper[dummy.Next] // since the next node is stored in the last step
		mapper[dummy].Random = mapper[dummy.Random]

		dummy = dummy.Next
	}

	return mapper[head] // return the clone, wired. 
}
