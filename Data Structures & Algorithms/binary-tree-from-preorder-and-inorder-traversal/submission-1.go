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

func buildTree(preorder []int, inorder []int) *TreeNode {

	var find func(pre, in []int) *TreeNode

	find = func(pre, in []int) *TreeNode {
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

		lIn := in[:mid]
		rIn := in[mid+1:]

		lPre := pre[1:mid+1]
		rPre := pre[mid+1:]

		return &TreeNode{
			Val: nodeVal,
			Left: find(lPre, lIn),
			Right: find(rPre, rIn),
		}
	}
	
	return find(preorder, inorder)
}
