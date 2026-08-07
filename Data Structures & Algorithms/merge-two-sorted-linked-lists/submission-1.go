/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{}
	move := result // init 

	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			move.Next = l1
			l1 = l1.Next
		} else {
			move.Next = l2
			l2 = l2.Next
		}

		move = move.Next
	}


	if l1 != nil {
		move.Next = l1
	} else {
		move.Next = l2
	}

	return result.Next
}
