func lengthOfLongestSubstring(s string) int {
	charSet := make(map[byte]bool)
	res := 0
	for i := 0; i < len(s); i++ {
		cur := 0
		clear(charSet)
		for j := i; j < len(s); j++ {
			if charSet[s[j]] {
				if cur > res {
					res = cur
				}
				break
			}
			charSet[s[j]] = true
			cur++
		}
		if cur > res {
			res = cur
		}
	}
	return res
}
