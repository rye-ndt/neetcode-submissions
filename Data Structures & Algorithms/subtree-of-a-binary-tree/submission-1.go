/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
    var find func(a, b *TreeNode)

	var compare func(a, b *TreeNode) bool 

	compare = func(a, b *TreeNode) bool {
		if a == nil && b == nil {
			fmt.Println("case 1")
			return true
		}

		if (a != nil && b == nil) || (b != nil && a == nil) {
			fmt.Println("case 2")
			return false
		}

		if a.Val != b.Val {
			fmt.Println("case 3")
			return false
		}

		fmt.Println("case final")

		return compare(a.Left, b.Left) && compare(a.Right, b.Right)
	}

	found := false

	find = func(a, b *TreeNode) {
		if a == nil && b == nil {
			return
		}

		if a != nil {
			if a.Val == subRoot.Val {
				fmt.Println("found & will try at a ", a.Val)
				if compare(a, subRoot) {
					found = true
					return
				}
			} 

			fmt.Println("cannot find at a ", a.Val, " going left and right")
			find(a.Left, a.Right)
		}

		if b != nil {
			if b.Val == subRoot.Val {
				fmt.Println("found & will try at b ", b.Val)
				if compare(b, subRoot) {
					found = true
					return
				}
			}
			
			fmt.Println("cannot find at b", b.Val, " going left and right")
			find(b.Left, b.Right)
		}		
	}

	find(root, nil)

	return found
}
