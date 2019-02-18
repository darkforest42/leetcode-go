package ReverseLinkedList

type ListNode struct {
	Val int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	return tailRecursiveReverse(head, nil)
	//newHeader := &ListNode{}
	//for head != nil {
	//	next := head.Next
	//	head.Next = newHeader
	//	newHeader = head
	//	head = next
	//}
	//return newHeader
}

func tailRecursiveReverse(head, newHead *ListNode) *ListNode {
	if head == nil {
		return newHead
	}
	next := head.Next
	head.Next = newHead
	newHead = head
	return tailRecursiveReverse(next, newHead)
}
