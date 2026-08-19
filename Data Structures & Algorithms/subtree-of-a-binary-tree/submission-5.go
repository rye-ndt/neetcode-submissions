func isSubtree(n, sub *TreeNode) bool {
	var loop func(n, sub *TreeNode, flag bool) bool

	loop = func(n, sub *TreeNode, flag bool) bool {
		if (n == nil && sub != nil) || (n != nil && sub == nil) { return false }
		if n == nil && sub == nil { return true }

		totallyEqual := n.Val == sub.Val && loop(n.Left, sub.Left, true) && loop(n.Right, sub.Right, true)

		if !flag && !totallyEqual { 
			return loop(n.Left, sub, false) || loop(n.Right, sub, false)
		}
	
		return totallyEqual
	}

	return loop(n, sub, false)
}