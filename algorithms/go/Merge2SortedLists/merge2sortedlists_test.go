package Merge2SortedLists

import (
	"testing"
	"fmt"
)

func TestMergeTwoLists(t *testing.T){
	l1 := &ListNode{ 1, &ListNode{2, &ListNode{4, nil}}}
	l2 := &ListNode{ 1, &ListNode{3, &ListNode{4, nil}}}
	fmt.Print(mergeTwoLists(l1, l2))
}
