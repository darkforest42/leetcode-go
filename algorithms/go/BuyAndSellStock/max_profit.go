package BuyAndSellStock

import "math"

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	maxPro := 0.0
	minPrice := float64(prices[0])
	for _, price := range prices {
		minPrice = math.Min(float64(price), minPrice)
		maxPro = math.Max(float64(price) - minPrice, maxPro)
	}
	return int(maxPro)
}
