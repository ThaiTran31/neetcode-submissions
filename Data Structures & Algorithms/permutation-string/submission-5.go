func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}
	sortedS1 := sortStr(s1)
	len1 := len(s1)
	for i := 0; i <= len(s2) - len1; i++ {
		if isPermutation(sortedS1, s2[i:i+len1]) {
			return true
		}
	}
	return false
}

func isPermutation(sortedS1, s2 string) bool {
	sortedS2 := sortStr(s2)
	return sortedS1 == sortedS2
}

func sortStr(s string) string {
	r := []rune(s)
	sort.Slice(r, func(i, j int) bool {
		return r[i] < r[j]
	})
	return string(r)
}
