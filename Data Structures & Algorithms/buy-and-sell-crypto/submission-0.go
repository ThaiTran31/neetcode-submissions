func maxProfit(prices []int) int {
	res := 0
	for i := 0; i < len(prices); i++ {
		for j := i+1; j < len(prices); j++ {
			if prices[i] < prices[j] {
				cur := prices[j] - prices[i]
				if cur > res {
					res = cur
				}
			}
		}
	}
	return res
}
