func maxProfit(prices []int) int {
	l, r := 0, 1
	res := 0
	for r < len(prices) {
		if prices[r] < prices[l] {
			l = r
		} else {
			cur := prices[r] - prices[l]
			if cur > res {
				res = cur
			}
		}
		r++
	}
	return res
}
