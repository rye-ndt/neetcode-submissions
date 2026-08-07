/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func diameterOfBinaryTree(root *TreeNode) int {
	diameter := 0 

	var height func(n *TreeNode) int

	height = func(node *TreeNode) int {
		// leaf 
		if node == nil {
			return 0
		}

		// left and right 
		left := height(node.Left)
		right := height(node.Right)

		// combine 
		if left + right > diameter {
			diameter = left + right
		}

		if left > right {
			return left + 1
		}

		return right + 1
	}

	height(root)

	return diameter
}
