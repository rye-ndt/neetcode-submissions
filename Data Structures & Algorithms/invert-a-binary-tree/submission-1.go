/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func invertTree(n *TreeNode) *TreeNode {
	if n == nil { return nil }

	left := n.Left
	n.Left = invertTree(n.Right)
	n.Right = invertTree(left)

	return n
}
