// preorder = mid, left, right 
// inorder = left, mid, right 

func buildTree(preorder []int, inorder []int) *TreeNode {
	valToIndex := map[int]int{}

	for i, v := range inorder {
		valToIndex[v] = i
	}

	var loop func(pre, in []int, offset int) *TreeNode

	loop = func(pre, in []int, offset int) *TreeNode{
		if len(pre) == 0 { return nil }

		split := valToIndex[pre[0]] - offset

		return &TreeNode{ 
			Val: pre[0],
			Left: loop(pre[1:split+1], in[:split], offset),
			Right: loop(pre[split+1:], in[split+1:], offset + split + 1),
		}
	}

	return loop(preorder, inorder, 0)
}
