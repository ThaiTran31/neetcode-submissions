func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	counter := make(map[rune]int)
	for _, r := range s {
		counter[r] += 1
	}
	for _, r := range t {
		counter[r] -= 1
	}
	for _, v := range counter {
		if v != 0 {
			return false
		}
	}
	return true
}
