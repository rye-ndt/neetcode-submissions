/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func maxDepth(root *TreeNode) int {
	// return the leaf
    if root == nil {
		return 0
	}

	left := maxDepth(root.Left)
	right := maxDepth(root.Right)	

	// combine
	return 1 + max(left, right)
}
