/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{}
	head := result

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			result.Next = &ListNode{ Val: l1.Val, Next: &ListNode{} }
			l1 = l1.Next
		} else {
			result.Next = &ListNode{ Val: l2.Val, Next: &ListNode{} }
			l2 = l2.Next
		}

		result = result.Next
	}

	fmt.Println("result", result)

	switch {
		case l1 != nil: result.Next = l1
		case l2 != nil: result.Next = l2
	}

	return head.Next
}
