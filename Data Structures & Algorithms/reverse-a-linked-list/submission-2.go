/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode // nil 

	for head != nil {
		next := head.Next 
		head.Next = prev 

		// loop
		prev = head 
		head = next
	}

	return prev
}
