func maxProfit(prices []int) int {
	res := 0
	l := 0
	r := 1
	for r < len(prices) {
		if prices[r] > prices[l] {
			cur := prices[r] - prices[l]
			if cur > res {
				res = cur
			}
		} else {
			l = r
		}
		r++
	}
	return res
}
