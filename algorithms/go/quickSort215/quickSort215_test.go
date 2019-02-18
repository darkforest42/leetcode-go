package quickSort215

import (
	"fmt"
	"testing"
)

func TestFindKthLargest(t *testing.T) {
	example := []int{3,2,1,5,6,4}
	result := FindKthLargest(example, 2)
	fmt.Print(result)
}