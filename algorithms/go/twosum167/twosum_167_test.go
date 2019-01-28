package twosum

import (
	"fmt"
	"testing"
)

func TestTwoSum(t *testing.T) {
	example := []int{2, 7, 11, 15}
	result := TwoSum(example, 9)
	fmt.Print(result)
}
