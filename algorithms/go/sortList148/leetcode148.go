package sortList148


type ListNode struct {
	Val int
	Next *ListNode
}

//用快慢指针的方法找到链表中间节点，然后递归的对两个子链表排序，把两个排好序的子链表合并成一条有序的链表
//归并排序比较适合链表，它可以保证了最好和最坏时间复杂度都是 O(NlogN)，而且它在数组排序中广受诟病的空间复杂度在链表排序中也从O(n)降到了 O(1)
//因为链表快排中只能使用第一个节点作为枢纽，所以不能保证时间复杂度
//时间复杂度：最好/最坏 O(NlogN)
func merge(h1, h2 *ListNode) *ListNode{
	if h1 == nil {
		return h2
	}
	if h2 == nil {
		return h1
	}

}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	slowPtr, fastPtr := head, head
	for fastPtr.Next != nil && fastPtr.Next.Next != nil {
		fastPtr = fastPtr.Next.Next
		slowPtr = slowPtr.Next
	}
	mid := slowPtr.Next
	l1 := sortList(head)
	l2 := sortList(mid)
	return merge(l1, l2)
}
