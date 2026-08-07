/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func kthSmallest(root *TreeNode, k int) int {
	var sort func(n *TreeNode) []int

	sort = func(n *TreeNode) []int {
		if n == nil {
			return []int{}
		}

		merge := append(sort(n.Left), n.Val)

		return append(merge, sort(n.Right)...)
	}

	return sort(root)[k - 1]
}
