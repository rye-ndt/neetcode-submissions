/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func goodNodes(root *TreeNode) int {
	result := 0

	var traverse func(n *TreeNode, max int) 

	traverse = func(n *TreeNode, max int) {
		if n == nil {
			return
		}

		if n.Val >= max {
			result += 1
			max = n.Val
		} 

		traverse(n.Left, max)
		traverse(n.Right, max)
	}

	traverse(root, root.Val - 1)

	return result
}
