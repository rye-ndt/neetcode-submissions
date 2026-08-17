/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
	runner := head 
	cloneMap := map[*Node]*Node{}

	for runner != nil {
		cloneMap[runner] = &Node{ Val: runner.Val }
		runner = runner.Next
	}

	result := cloneMap[head]

	for head != nil {
		runner = cloneMap[head]
		runner.Next = cloneMap[head.Next]
		runner.Random = cloneMap[head.Random]

		head = head.Next
		runner = runner.Next
	}


	return result
}
