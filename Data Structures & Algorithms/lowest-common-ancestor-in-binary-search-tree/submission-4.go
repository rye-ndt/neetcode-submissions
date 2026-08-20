func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	switch {
		case p.Val < root.Val && q.Val < root.Val:
			return lowestCommonAncestor(root.Left, p, q)
		case p.Val > root.Val && q.Val > root.Val:
			return lowestCommonAncestor(root.Right, p, q)
		default:
			return root
	}
}