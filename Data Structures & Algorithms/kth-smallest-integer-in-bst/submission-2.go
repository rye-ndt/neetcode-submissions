func kthSmallest(root *TreeNode, k int) int {
	var loop func(n *TreeNode) []int 

	loop = func (n *TreeNode) []int {
		if n == nil {
			return nil
		}

		l, r := loop(n.Left), loop(n.Right)

		return append(append(append([]int{}, l...), n.Val), r...)
	}

	return loop(root)[k-1]
}


/*

6
|   \ 
4    7 
| \ 
3  5
|
2

*/