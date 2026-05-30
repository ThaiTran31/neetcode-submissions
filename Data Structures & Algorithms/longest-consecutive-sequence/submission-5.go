func longestConsecutive(nums []int) int {
    count := make(map[int]bool)
    for _, num := range nums {
        count[num] = true
    }
    res := 0
    for _, num := range nums {
        if _, ok := count[num-1]; ok {
            continue
        }
        curMax := 1
        i := 1
        for {
            _, ok := count[num+i]
            if !ok {
                if curMax > res {
                    res = curMax
                }
                break
            }
            curMax++
            i++
        }
    }
    return res
}
