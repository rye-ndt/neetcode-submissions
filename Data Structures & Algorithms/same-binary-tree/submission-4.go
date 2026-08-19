func isSameTree(a, b *TreeNode) bool {
	if a == nil && b == nil { return true }

	switch {
		case a == nil && b != nil: return false
		case a != nil && b == nil: return false 
		case a.Val != b.Val: return false
		default: 
			return isSameTree(a.Left, b.Left) && isSameTree(a.Right, b.Right)
	}
}
