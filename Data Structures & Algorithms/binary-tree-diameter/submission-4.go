func diameterOfBinaryTree(n *TreeNode) int {
	result := 0

	var find func(n *TreeNode) int 

	find = func(n *TreeNode) int {
		if n == nil { return 0 }

		result = max(find(n.Left) + find(n.Right), result)

		if find(n.Left) > find(n.Right) { return 1 + find(n.Left )}

		return 1 + find(n.Right)
	}

	find(n)

	return result
}