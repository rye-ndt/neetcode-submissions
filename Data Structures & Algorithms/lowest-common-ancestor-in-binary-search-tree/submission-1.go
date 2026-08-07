/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func lowestCommonAncestor(root *TreeNode, p *TreeNode, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	fmt.Println("processing: ", root.Val)

	cur := root.Val
	qv := q.Val
	pv := p.Val 

    // find the position of both nodes first 
	if pv < cur && qv < cur {
		return lowestCommonAncestor(root.Left, p, q)
	}

	if pv > cur && qv > cur {
		return lowestCommonAncestor(root.Right, p, q)
	}

	return root
}


