func lengthOfLongestSubstring(s string) int {
    res := 0

    for i := 0; i < len(s); i++ {
        charSet := make(map[byte]bool)
        for j := i; j < len(s); j++ {
            if charSet[s[j]] {
                break
            }
            charSet[s[j]] = true
        }
        if len(charSet) > res {
            res = len(charSet)
        }
    }
    return res
}