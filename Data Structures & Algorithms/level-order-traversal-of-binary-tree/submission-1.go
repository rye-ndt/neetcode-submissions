func levelOrder(root *TreeNode) [][]int {
	result := [][]int{}

	var loop func(n *TreeNode, level int) 

	loop = func(n *TreeNode, level int) {
		if n == nil { return }

		if len(result) < level+1 {
			result = append(result, []int{})
		}

		result[level] = append(result[level], n.Val)

		loop(n.Left, level+1)
		loop(n.Right, level+1)
	}

	loop(root, 0)

	return result
}
