/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// just like they way we do plus by hand

	result := &ListNode{}
	dummy := result 
	carry := 0 // 9 + 9 = 18 -> carry = 1

	for l1 != nil || l2 != nil || carry != 0 {
		sum := carry 

		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		dummy.Next = &ListNode{
			Val: sum % 10,
		}

		carry = sum / 10
		dummy = dummy.Next
	}

	return result.Next
}
