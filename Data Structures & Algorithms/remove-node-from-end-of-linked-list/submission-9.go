/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// reverse -> remove -> reverse 

func reverse(head *ListNode) *ListNode {
	clone := head

	var left *ListNode 

	for clone != nil {
		right := clone.Next
		clone.Next = left
		left = clone
		clone = right
	}

	return left
}

func find(head *ListNode, pos int) *ListNode {
	clone := head

	for i := 1; i < pos && clone != nil; i++ {
		clone = clone.Next
	}

	return clone
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	clone := reverse(head)
	item := find(clone, n)

	anchor := clone

	for clone != nil {
		if clone.Next == item {
			clone.Next = clone.Next.Next 
			break
		}

		if clone == item {
			anchor = clone.Next
			break
		}	

		clone = clone.Next 
	}

	return reverse(anchor)
}
