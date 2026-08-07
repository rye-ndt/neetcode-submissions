/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isValidBST(n *TreeNode) bool {
	var loop func(n *TreeNode, lo, hi int) bool 

	loop = func(n *TreeNode, lo, hi int) bool {
		if n == nil {
			return true
		}

		v := n.Val 

		if v <= lo || v >= hi {
			return false
		}

		return loop(n.Left, lo, v) && loop(n.Right, v, hi)
	}

	return loop(n, -10000, 10000)
}
