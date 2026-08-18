func invertTree(n *TreeNode) *TreeNode {
	if n == nil { return n }

	left := n.Left
	n.Left = invertTree(n.Right)
	n.Right = invertTree(left)

	return n
}
