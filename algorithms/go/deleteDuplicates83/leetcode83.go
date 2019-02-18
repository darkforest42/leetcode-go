package deleteDuplicates83

type ListNode struct {
	Val int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	prev := head
	newHead := prev
	head  = head.Next
	for head != nil {
		if head.Val != prev.Val {
			prev.Next = head
			prev = head
		}
		head = head.Next
	}

	if prev.Next != nil {
		prev.Next = nil
	}
	return newHead
}
