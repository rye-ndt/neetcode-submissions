/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func hasCycle(head *ListNode) bool {
	s, f := head, head 

	for f != nil && f.Next != nil && f.Next.Next != nil {
		f = f.Next.Next 
		s = s.Next 

		if f == s { return true }
	}

	return false
}
