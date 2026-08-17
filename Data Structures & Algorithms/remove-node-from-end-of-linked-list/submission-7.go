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
	fmt.Println("clone", clone)

	item := find(clone, n)
	fmt.Println("item", item)

	anchor := clone

	for clone != nil {
		switch {
			case clone.Next == item:
				clone.Next = clone.Next.Next
				break
			case clone == item:
				if clone.Next == nil { return nil }
				clone = clone.Next
				anchor = clone
			default: clone = clone.Next
		}
	}

	return reverse(anchor)
}
