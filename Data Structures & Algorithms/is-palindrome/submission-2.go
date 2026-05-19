func isPalindrome(s string) bool {
	i := 0;
	j := len(s) - 1;
	lowerS := strings.ToLower(s)
	for i <= j {
		for i < len(s) && !isAlphanumeric(lowerS[i]) {
			i++
		}
		for j >= 0 && !isAlphanumeric(lowerS[j]) {
			j--
		}
		if i >= len(s) || j < 0 {
			break
		}
		if lowerS[i] != lowerS[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func isAlphanumeric(r byte) bool {
    return unicode.IsLetter(rune(r)) || unicode.IsNumber(rune(r))
}
