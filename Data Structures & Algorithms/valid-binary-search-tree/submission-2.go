/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isValidBST(n *TreeNode) bool {
	var loop func(n *TreeNode, min, max int) bool

	loop = func(n *TreeNode, min, max int) bool {
		if n == nil {
			return true
		}

		v := n.Val

		if v <= min || v >= max {
			return false
		}

		return loop(n.Left, min, v) && loop(n.Right, v, max)
	}

	return loop(n, -10000, 10000)
}
