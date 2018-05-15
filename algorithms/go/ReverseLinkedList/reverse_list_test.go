package ReverseLinkedList

import (
	"testing"
	"fmt"
)
func TestReverseList(t *testing.T){
	linkedList := &ListNode{ 1, &ListNode{2, &ListNode{3, nil}}}
	fmt.Print(reverseList(linkedList))
}