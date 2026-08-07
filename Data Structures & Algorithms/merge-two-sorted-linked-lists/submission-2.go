/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{}
	runner := result

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			runner.Next = l1 
			l1 = l1.Next
		} else {
			runner.Next = l2
			l2 = l2.Next
		}

		runner = runner.Next
	}

	if l1 != nil {
		runner.Next = l1
	} else {
		runner.Next = l2
	}

	return result.Next
}
