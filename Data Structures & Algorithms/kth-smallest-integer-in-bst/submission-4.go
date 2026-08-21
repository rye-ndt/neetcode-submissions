func kthSmallest(root *TreeNode, k int) int {
	var loop func(n *TreeNode) []int 
	loop = func (n *TreeNode) []int {
		if n == nil { return nil }

		return append(append(append([]int{}, loop(n.Left)...), n.Val), loop(n.Right)...)
	}

	return loop(root)[k-1]
}