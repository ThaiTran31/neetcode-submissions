func characterReplacement(s string, k int) int {
	res := 0
	for i := 0; i < len(s); i++ {
		charCount := make(map[byte]int)
		maxFreq := 0
		for j := i; j < len(s); j++ {
			charCount[s[j]]++
			if charCount[s[j]] > maxFreq {
				maxFreq = charCount[s[j]]
			}
			if j - i + 1 - maxFreq > k {
				break
			}
			if j - i + 1 > res {
				res = j - i + 1
			}
		}
	}
	return res
}
