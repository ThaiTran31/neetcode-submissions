func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	sMap := buildMapFromString(s)
	tMap := buildMapFromString(t)
	for key, val := range sMap {
		if (tMap[key] != val) {
			return false
		}
	}
	return true
}

func buildMapFromString(s string) map[rune]int {
	res := make(map[rune]int)
	for _, val := range s {
		res[val] += 1
	}
	return res
}
