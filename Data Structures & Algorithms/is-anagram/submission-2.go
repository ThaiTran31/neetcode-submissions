func isAnagram(s string, t string) bool {
	charCount := make(map[rune]int)
	sRunes := []rune(s)
	for _, r := range sRunes {
		charCount[r]++
	}
	tRunes := []rune(t)
	for _, r := range tRunes {
		charCount[r]--
	}
	for _, val := range charCount {
		if val != 0 {
			return false
		}
	}
	return true
}
