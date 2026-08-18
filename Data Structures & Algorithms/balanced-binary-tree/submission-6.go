func isBalanced(root *TreeNode) bool {
	result := true

	var height func(n *TreeNode) int

	height = func(n *TreeNode) int {
		if n == nil { return 0}

		l, r := height(n.Left), height(n.Right)

		if (l - r) * (l - r) > 1 { result = false }

		return 1 + max(l, r)
	}

	height(root)

	return result
}