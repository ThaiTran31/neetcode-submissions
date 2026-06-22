func lengthOfLongestSubstring(s string) int {
    char2Idx := make(map[byte]int)
    res := 0
    l := 0
    for r := 0; r < len(s); r++ {
        if i, exist := char2Idx[s[r]]; exist {
            if i+1 > l {
                l = i+1
            }
        }
        char2Idx[s[r]] = r
        if r - l + 1 > res {
            res = r - l + 1
        }
    }
    return res
}
