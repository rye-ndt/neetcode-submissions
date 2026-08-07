/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	// just like they way we do plus by hand
	// align the row, plus from right to left -> the same order as the node list itself

	head := &ListNode{}
	runner := head
	carry := 0

	for l1 != nil || l2 != nil || carry != 0 { // as long as there is something to add 
		sum := 0 + carry // the remaining from the last step 

		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		runner.Next = &ListNode{
			Val: sum % 10,
		}

		carry = sum / 10
		runner = runner.Next
	}

	return head.Next
}


func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	runner := head 
	carry := 0

	for l1 != nil || l2 != nil || carry != 0 {
		total := carry 

		if l1 != nil {
			total += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			total += l2.Val 
			l2 = l2.Next
		}

		val := total % 10

		runner.Next = &ListNode{
			Val: val,
		}

		carry = total / 10

		runner = runner.Next
	}

	return head.Next
}
