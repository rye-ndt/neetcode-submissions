/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reorderList(head *ListNode) {
	// find the middle, using fast n slow 
	f, m := head, head 
	
	for f != nil && f.Next != nil {
		m = m.Next
		f = f.Next.Next
	}

	// reverse the second half
	secondHalf := m.Next
	m.Next = nil
	firstHalf := head 

	var prev *ListNode

	for secondHalf != nil {
		next := secondHalf.Next
		secondHalf.Next = prev
		prev = secondHalf
		secondHalf = next
	}

	secondHalf = prev

	dummy := &ListNode{} 
	result := dummy // do not use head 

	cur := 0

	for firstHalf != nil && secondHalf != nil {
		if cur == 0 {
			result.Next = firstHalf
			firstHalf = firstHalf.Next 
			cur = 1
		} else {
			result.Next = secondHalf
			secondHalf = secondHalf.Next
			cur = 0
		}

		result = result.Next
	}

	if firstHalf != nil {
		result.Next = firstHalf
	} else if secondHalf != nil {
		result.Next = secondHalf
	}

	head = result
}
