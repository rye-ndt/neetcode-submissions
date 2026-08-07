/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	nilNode := &ListNode{
		Next: head,
	}

	s, f := nilNode, nilNode 

	for i := 0; i < n; i++ {
		// move the f first
		f = f.Next
	}

	// now move both til the end 
	for f.Next != nil {
		s = s.Next 
		f = f.Next 
	}

	// now truncate
	s.Next = s.Next.Next 

	return nilNode.Next
}
