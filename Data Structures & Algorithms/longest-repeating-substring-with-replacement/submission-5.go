func characterReplacement(s string, k int) int {
	res := 0
	charSet := make(map[byte]bool)
	for i := 0; i < len(s); i++ {
		charSet[s[i]] = true
	}
	for c := range charSet {
		l := 0
		count := 0
		for r := 0; r < len(s); r++ {
			if s[r] == c {
				count++
			}
			if r-l+1 - count <= k {
				if r-l+1 > res {
					res = r-l+1
				}
			} else {
				for r-l+1 - count > k {
					if s[l] == c {
						count--
					}
					l++
				}
			}
		}
	}
	return res
}
