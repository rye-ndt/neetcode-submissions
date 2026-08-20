func rightSideView(root *TreeNode) []int {
	store := [][]int{}

	var loop func(n *TreeNode, level int) 
	loop = func(n *TreeNode, level int) {
		if n == nil { return }

		if len(store) < level + 1 {
			store = append(store, []int{})
		}

		store[level] = append(store[level], n.Val)

		loop(n.Left, level+1)
		loop(n.Right, level+1)
	}

	loop(root, 0)

	result := []int{}

	for _, level := range store {
		result = append(result, level[len(level)-1])
	}

	return result
}
