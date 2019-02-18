package detectCycle142

type ListNode struct {
	Val int
	Next *ListNode
}
func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	slowPtr, fastPtr := head, head
	hasCycle := false
	for fastPtr.Next != nil && fastPtr.Next.Next != nil {
		slowPtr = slowPtr.Next
		fastPtr = fastPtr.Next.Next
		if slowPtr == fastPtr {
			hasCycle = true
			break
		}
	}
	if hasCycle {
		for head != slowPtr {
			head = head.Next
			slowPtr = slowPtr.Next
		}
		return head
	}else{
		return nil
	}
}

//func detectCycle(head *ListNode) *ListNode {
//	if head == nil || head.Next == nil {
//		return nil
//	}
//	nodesMap := make(map[*ListNode]bool)
//	for head != nil {
//		if _, ok := nodesMap[head]; ok {
//			return head
//		}
//		nodesMap[head] = true
//		head = head.Next
//	}
//	return nil
//}
