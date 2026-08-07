/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func diameterOfBinaryTree(root *TreeNode) int {
	result := 0

	var height func(n *TreeNode) int

	height = func(node *TreeNode) int {
		// leaf
		if node == nil {
			return 0 
		}

		// left n right 
		left := height(node.Left)
		right := height(node.Right)

		// at each step, tryna find if there is a bigger result
		if left + right > result {
			result = left + right
		}

		// choose the bigger one to continue with the path
		return 1 + max(left, right)
	}

	height(root)

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
