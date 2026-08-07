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

	var traverse func(n *TreeNode, prevs []int) 

	traverse = func(n *TreeNode, prevs []int) {
		if n == nil {
			return
		}

		shouldInclude := true

		for _, v := range prevs {
			if v > n.Val {
				shouldInclude = false
			}
		}

		if shouldInclude {
			result += 1
		}

		prevs = append(prevs, n.Val)

		traverse(n.Left, prevs)
		traverse(n.Right, prevs)
	}

	traverse(root, []int{})

	return result
}
