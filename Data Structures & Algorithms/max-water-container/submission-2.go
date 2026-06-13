func maxArea(heights []int) int {
    res := 0
    for i := 0; i < len(heights); i++ {
        for j := i + 1; j < len(heights); j++ {
            area := min(heights[i], heights[j]) * (j - i)
            if area > res {
                res = area
            }
        }
    }
    return res
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}