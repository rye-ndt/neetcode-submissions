func goodNodes(root *TreeNode) int {
	counter := 0

	var loop func(n *TreeNode, biggest int) 
	loop = func(n *TreeNode, biggest int) {
		if n == nil { return }

		if n.Val >= biggest { counter++ }

		loop(n.Left, max(biggest, n.Val))
		loop(n.Right, max(biggest, n.Val))
	}

	loop(root, root.Val-1)

	return counter 
}
