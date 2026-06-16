func characterReplacement(s string, k int) int {
	res := 0
	for i := 0; i < len(s); i++ {
		freqs := make(map[byte]int)
		maxFreq := 0
		for j := i; j < len(s); j++ {
			freqs[s[j]]++
			if freqs[s[j]] > maxFreq {
				maxFreq = freqs[s[j]]
			}
			if j - i + 1 - maxFreq <= k && j - i + 1 > res {
				res = j - i + 1
			}
		}
	}
	return res
}
