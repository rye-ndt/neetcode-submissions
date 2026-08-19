func isSubtree(n, sub *TreeNode) bool {
	var loop func(n, sub *TreeNode, start bool) bool

	loop = func(n, sub *TreeNode, start bool) bool {
		if (n == nil && sub != nil) || (n != nil && sub == nil) { 
			return false 
		}

		if n == nil && sub == nil { return true }

		if !start {
			if n.Val == sub.Val && loop(n.Left, sub.Left, true) && loop(n.Right, sub.Right, true) {
				// start new here 
				return true
			} else {
				return loop(n.Left, sub, false) || loop(n.Right, sub, false)
			}
		} else {
			return n.Val == sub.Val && loop(n.Left, sub.Left, true) && loop(n.Right, sub.Right, true)
		}
	}

	return loop(n, sub, false)
}