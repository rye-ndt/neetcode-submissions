func isBalanced(n *TreeNode) bool {
	balance := true

	var loop func(n *TreeNode) int

	loop = func(n *TreeNode) int {
		if n == nil { return 0}

		l, r := loop(n.Left), loop(n.Right)

		if (l - r) * (l - r) > 1 { balance = false }

		return 1 + max(l, r)
	}

	loop(n)

	return balance
}