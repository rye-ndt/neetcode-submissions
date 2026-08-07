/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
	cloneMap := map[*Node]*Node{}

	dum := head 

	for dum != nil {
		cloneMap[dum] = &Node{
			Val: dum.Val,
		} 

		dum = dum.Next
	}

	dum = head 

	for dum != nil {
		cloneMap[dum].Next = cloneMap[dum.Next] 
		cloneMap[dum].Random = cloneMap[dum.Random]

		dum = dum.Next
	}

	return cloneMap[head]
}
