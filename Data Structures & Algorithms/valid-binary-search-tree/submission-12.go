func isValidBST(n *TreeNode) bool {
	var loop func(n *TreeNode, min, max int) bool 

	loop = func(n *TreeNode, min, max int) bool {
		if n == nil { return true }

		if n.Val <= min || n.Val >= max { return false }

		return loop(n.Left, min, n.Val) && loop(n.Right, n.Val, max)
	}

	return loop(n, -2147483647, 2147483647)
}
