func isValidBST(n *TreeNode) bool {
	var loop func(n *TreeNode, min, max int) bool 

	loop = func(n *TreeNode, min, max int) bool {
		switch {
			case n == nil: return true 
			case n.Val <= min || n.Val >= max: return false 
			default: return loop(n.Left, min, n.Val) && loop(n.Right, n.Val, max)
		}
	}

	return loop(n, -2147483647, 2147483647)
}
