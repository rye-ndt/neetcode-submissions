func maxDepth(n *TreeNode) int {
	if n == nil { return 0 }

	return 1 + max(maxDepth(n.Left), maxDepth(n.Right))
}
