/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// compare the longest left and right to see if they are equal

func isBalanced(root *TreeNode) bool {
	balanced := true

	var height func(r *TreeNode) int 

	height = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		left := height(node.Left)
		right := height(node.Right)

		if abs(left - right) > 1 {
			balanced = false
		}

		return max(left, right) + 1
	}

	height(root)

	return balanced
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func abs(a int) int {
	if a > 0 {
		return a 
	}

	return -a
}
