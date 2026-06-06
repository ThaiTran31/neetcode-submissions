func maxProfit(prices []int) int {
	l, r, res := 0, 1, 0
	for r < len(prices) {
		cur := prices[r] - prices[l]
		if cur > res {
			res = cur
		}
		if prices[r] < prices[l] {
			l = r
		}
		r++
	}
	return res
}
