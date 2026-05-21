func isPalindrome(s string) bool {
	normalizedS := ""
	for _, r := range s {
		if (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') {
			normalizedS += string(r)
		}
		if r >= 'A' && r <= 'Z' {
			normalizedS += string(r + 'a' - 'A')
		}
	}
	reversedS := reverseString(normalizedS)
	return reversedS == normalizedS
}

func reverseString(s string) string {
	runes := []rune(s)
	length := len(s)
	for i := 0; i < length / 2; i++ {
		runes[i], runes[length-1-i] = runes[length-1-i], runes[i]
	}
	return string(runes)
}
