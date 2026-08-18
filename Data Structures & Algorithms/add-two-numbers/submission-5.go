func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{}
	anchor := result
	carry := 0

	for l1 != nil || l2 != nil || carry != 0 {
		val := carry

		if l1 != nil { 
			val += l1.Val 
			l1 = l1.Next 
		}

		if l2 != nil { 
			val += l2.Val 
			l2 = l2.Next 
		}

		carry = val / 10

		result.Next = &ListNode{
			Val: val % 10 ,
		}

		result = result.Next
	}

	return anchor.Next
}
