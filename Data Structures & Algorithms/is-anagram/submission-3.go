func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	charCount := make(map[rune]int)
	sRunes := []rune(s)
	tRunes := []rune(t)
	for i := 0; i < len(s); i++ {
		charCount[sRunes[i]]++
		charCount[tRunes[i]]--
	}
	for _, val := range charCount {
		if val != 0 {
			return false
		}
	}
	return true
}
