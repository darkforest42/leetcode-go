package Merge2SortedLists

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val int
	Next *ListNode
}
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil { return l2}
	if l2 == nil { return l1}
	NewList := &ListNode{}
	curser := NewList
	for l1 != nil || l2 != nil {
		if l1 == nil && l2 != nil {
			curser.Next = l2
			return NewList.Next
		}
		if l1 != nil && l2 == nil {
			curser.Next = l1
			return NewList.Next
		}
		if l1.Val <= l2.Val {
			curser.Next = l1
			curser = l1
			l1 = l1.Next
			continue
		}
		if l2.Val <= l1.Val{
			curser.Next = l2
			curser = l2
			l2 = l2.Next
			continue
		}
	}
	return NewList.Next
}
