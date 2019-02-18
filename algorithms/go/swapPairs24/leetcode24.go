package swapPairs24

type ListNode struct {
	Val int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var l1, l2 *ListNode
	newHead := &ListNode{}
	newHead.Next = head
	prev := newHead
	for prev.Next != nil && prev.Next.Next != nil {
		l1 = prev.Next
		l2 = prev.Next.Next
		next := l2.Next
		l1.Next = next
		l2.Next = l1
		prev.Next = l2
		prev = l1
	}
	return newHead.Next
}
