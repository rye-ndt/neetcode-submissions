/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// figured this out myself
// at each node of the original tree, check if the current value is equivalent to the subroot 
// if no, go to left and right of the original tree
// if yes, then start comparing 2 trees. then return if the compare returns true

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
   if root == nil {
		return false
   }
   
   if root.Val == subRoot.Val && sameTree(root, subRoot) {
		return true
   }

   return isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}

func sameTree(a, b *TreeNode) bool {
	if a == nil && b == nil {
		return true
	}

	if b == nil || a == nil {
		return false
	}

	return a.Val == b.Val && sameTree(a.Left, b.Left) && sameTree(a.Right, b.Right)
}
