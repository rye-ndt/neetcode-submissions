/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isSameTree(p *TreeNode, q *TreeNode) bool {
    if p == nil && q == nil {
		fmt.Println("both nodes nil")
		return true
	}

	if (p == nil && q != nil) || (p != nil && q == nil) {
		fmt.Println("false for compare ", p, q)
		return false
	}

	if p.Val != q.Val {
		fmt.Println("comparing cur: ", p.Val, q.Val)
		return false
	}

	fmt.Println("comparing ", p.Left, q.Left, p.Right, q.Right)

	left := isSameTree(p.Left, q.Left)
	right := isSameTree(p.Right, q.Right)


	fmt.Println("left and right", left, right)

	return left && right
}
