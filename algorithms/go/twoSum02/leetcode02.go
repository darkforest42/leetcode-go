package twoSum02


//Definition for singly-linked list.
type ListNode struct {
	Val int
	Next *ListNode
}


func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	newHead := &ListNode{}
	current := newHead
	sum, carry := 0, 0

	for l1 != nil || l2 != nil || carry > 0 {
		sum = carry
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		carry = sum / 10
		current.Next = &ListNode{sum%10, nil}
		current = current.Next
	}
	return newHead.Next
}
