/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func index(s []int, x int) int {
	for i, v := range s {
		if v == x {
			return i
		}
	}

	return -1
}

func buildTree(pre []int, in []int) *TreeNode {
	if len(pre) == 0 {
		return nil
	}

	nodeVal := pre[0]

	mid := index(in, nodeVal)

	if mid == 0 && len(pre) == 1 {
		return &TreeNode{
			Val: nodeVal,
		}
	}

	return &TreeNode{
		Val: nodeVal,
		Left: buildTree(pre[1:mid+1], in[:mid]),
		Right: buildTree(pre[mid+1:], in[mid+1:]),
	}
}
