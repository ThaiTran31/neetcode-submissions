func lengthOfLongestSubstring(s string) int {
    res := 0
    for i := 0; i < len(s); i++ {
        set := make(map[byte]bool)
        for j := i; j < len(s); j++ {
            if set[s[j]] {
                break
            } else {
                set[s[j]] = true
            }
        }
        if len(set) > res {
            res = len(set)
        }
    }
    return res
}
