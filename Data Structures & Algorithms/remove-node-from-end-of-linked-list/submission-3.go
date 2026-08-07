/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{
		Next: head,
	}

	// start by -1 
	s, f := dummy, dummy

	for i := 0; i < n; i++ {
		f = f.Next
	}

	// move both until the end - 1
	// why f.Next instead of f? -> to stop at end - 1
	for f.Next != nil {
		f = f.Next
		s = s.Next
	}

	// now s is at n-1 node 
	s.Next = s.Next.Next

	return dummy.Next
}
