package PalindromeLinkedList

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
func isPalindrome(head *ListNode) bool {
	var intList []int
	for ;head != nil; {
		intList = append(intList, head.Val)
		head = head.Next
	}
	length := len(intList)
	for i := 0; i < length / 2; i++ {
		if intList[i] != intList[length - i -1] {
			return false
		}
	}
	return true
}
