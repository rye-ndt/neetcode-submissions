/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func invertTree(root *TreeNode) *TreeNode {
    // on leaf
	if root == nil {
		return nil
	} 

	// left and right 
	left := invertTree(root.Left)
	right := invertTree(root.Right)

	// combine
	root.Left = right 
	root.Right = left

	return root
}
