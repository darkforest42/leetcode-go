package PalindromeLinkedList

import (
	"testing"
	"fmt"
)

func TestIsPalindrome(t *testing.T) {
	input := &ListNode{-1, &ListNode{2, &ListNode{-1, nil}}}
	fmt.Println(isPalindrome(input))
}