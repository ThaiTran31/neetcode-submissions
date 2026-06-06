func maxProfit(prices []int) int {
	minPrice := 9999999
	maxProfit := 0
	for i := 0; i < len(prices); i++ {
		minPrice = min(minPrice, prices[i])
		curProfit := prices[i] - minPrice
		maxProfit = max(maxProfit, curProfit)
	}
	return maxProfit
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
