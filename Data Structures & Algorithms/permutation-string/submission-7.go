func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}
	len1 := len(s1)
	for i := 0; i <= len(s2) - len1; i++ {
		if isPermutation(s1, s2[i:i+len1]) {
			return true
		}
	}
	return false
}

func isPermutation(s1, s2 string) bool {
	var charCount [26]int
	for _, r := range s1 {
		charCount[r-'a']++
	}
	for _, r := range s2 {
		charCount[r-'a']--
	}
	for _, v := range charCount {
		if v != 0 {
			return false
		}
	}
	return true
}
