/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func levelOrder(root *TreeNode) [][]int {
    result := make([][]int, 0)

	var traverse func(n *TreeNode, level int) 

	traverse = func(n *TreeNode, level int) {
		if n == nil {
			return
		}

		if level == len(result) {
			result = append(result, []int{})
		}
		
		result[level] = append(result[level], n.Val)

		nextLevel := level+1

		traverse(n.Left, nextLevel)
		traverse(n.Right, nextLevel)

		return
	}

	traverse(root, 0)

	return result
}
