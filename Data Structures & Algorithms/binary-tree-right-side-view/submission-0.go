/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// map all the values on each layer, in order  
// then return the last element of that layer

func rightSideView(root *TreeNode) []int {
	vals := [][]int{}

	var list func(n *TreeNode, level int) 

	list = func(n *TreeNode, level int) {
		if n == nil {
			return 
		}

		if level == len(vals) {
			vals = append(vals, []int{})
		}

		vals[level] = append(vals[level], n.Val)

		list(n.Left, level+1)
		list(n.Right, level+1)
	}

	list(root, 0)

	result := []int{}

	for _, v := range vals {
		result = append(result, v[len(v) - 1])
	}

	return result
}
