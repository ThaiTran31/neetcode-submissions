func lengthOfLongestSubstring(s string) int {
	count := make(map[uint8]bool)
	res, cur := 0, 0
	for i := 0; i < len(s); i++ {
		cur = 1
		count[s[i]] = true
		for j := i+1; j < len(s); j++ {
			if !count[s[j]] {
				cur++
				count[s[j]] = true
			} else {
				if cur > res {
					res = cur
				}
				clear(count)
				break
			}
		}
		if cur > res {
			res = cur
		}
	}
	return res
}
