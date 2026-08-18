func diameterOfBinaryTree(n *TreeNode) int {
	result := 0

	var find func(n *TreeNode) int 

	find = func(n *TreeNode) int {
		if n == nil { return 0 }

		l, r := find(n.Left), find(n.Right)

		result = max(l + r, result)

		switch {
			case l > r: return 1 + l 
			default: return 1 + r
		}
	}

	find(n)

	return result
}