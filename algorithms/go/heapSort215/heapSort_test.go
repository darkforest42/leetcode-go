package heapSort215

import (
	"fmt"
	"testing"
)

func TestFindKthLargest(t *testing.T) {
	example := []int{3,2,3,1,2,4,5,5,6}
	result := FindKthLargest(example, 4)
	fmt.Println(example)
	fmt.Print(result)
}
