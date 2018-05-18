package BuyAndSellStock

import (
	"testing"
	"fmt"
)

func TestMaxProfit(t *testing.T) {
	prices := []int{7,1,5,3,6,4}
	fmt.Print(maxProfit(prices))
}
