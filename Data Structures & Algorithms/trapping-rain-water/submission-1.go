func trap(heights []int) int {
	maxPrefixHeights := make([]int, len(heights))
	for i := 1; i < len(heights); i++ {
		candidate := 0
		if heights[i-1] > maxPrefixHeights[i-1] {
			candidate = heights[i-1]
		} else {
			candidate = maxPrefixHeights[i-1]
		}
		if candidate > heights[i] {
			maxPrefixHeights[i] = candidate
		} else {
			maxPrefixHeights[i] = 0
		}
	}
	maxPostfixHeights := make([]int, len(heights))
	for i := len(heights) - 2; i >= 0; i-- {
		candidate := 0
		if heights[i+1] > maxPostfixHeights[i+1] {
			candidate = heights[i+1]
		} else {
			candidate = maxPostfixHeights[i+1]
		}
		if candidate > heights[i] {
			maxPostfixHeights[i] = candidate
		} else {
			maxPostfixHeights[i] = 0
		}
	}
	res := 0
	for i := 0; i < len(heights); i++ {
		cur := min(maxPrefixHeights[i], maxPostfixHeights[i]) - heights[i]
		if cur > 0 {
			res += cur
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
