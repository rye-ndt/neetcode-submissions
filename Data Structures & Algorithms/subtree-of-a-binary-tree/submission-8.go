func null(a *TreeNode) bool { return a == nil }

func isSubtree(n, sub *TreeNode) bool {
	var loop func(n, sub *TreeNode, flag bool) bool

	loop = func(n, sub *TreeNode, flag bool) bool {
		if (null(n) && !null(sub)) || (!null(n) && null(sub)) { return false }

		if null(n) && null(sub) { return true }

		match := n.Val == sub.Val && loop(n.Left, sub.Left, true) && loop(n.Right, sub.Right, true)

		if !flag && !match { 
			return loop(n.Left, sub, flag) || loop(n.Right, sub, flag)
		}
	
		return match
	}

	return loop(n, sub, false)
}