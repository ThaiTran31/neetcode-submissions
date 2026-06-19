func lengthOfLongestSubstring(s string) int {
	charSet := make(map[byte]bool)
	res, cur := 0, 0
	for i := 0; i < len(s); i++ {
		for j := i; j < len(s); j++ {
			if charSet[s[j]] {
				cur = 0
				clear(charSet)
				break
			}
			charSet[s[j]] = true
			cur++
			if cur > res {
				res = cur
			}
		}
	}
	return res
}
