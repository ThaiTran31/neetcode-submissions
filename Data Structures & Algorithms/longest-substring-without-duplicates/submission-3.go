func lengthOfLongestSubstring(s string) int {
	l, r := 0, 0
	set := make(map[byte]bool)
	res := 0
	for r < len(s) {
		if !set[s[r]] {
			set[s[r]] = true
		} else {
			if len(set) > res {
				res = len(set)
			}
			for l < r {
				if s[l] != s[r] {
					delete(set, s[l])
					l++
				} else {
					l++
					break
				}
			}
		}
		r++
	}
	if len(set) > res {
		return len(set)
	}
	return res
}
