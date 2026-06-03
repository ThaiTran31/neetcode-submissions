func trap(heights []int) int {
	res := 0
	for i, h := range heights {
		lHeight := findLeftHighest(heights, i)
		rHeight := findRightHighest(heights, i)
		cur := min(lHeight, rHeight) - h
		if cur > 0 {
			res += cur
		}
	}
	return res
}

func findLeftHighest(heights []int, idx int) int {
	res := 0
	for i := idx-1; i >= 0; i-- {
		if heights[i] > heights[idx] && heights[i] > res {
			res = heights[i]
		}
	}
	return res
}

func findRightHighest(heights []int, idx int) int {
	res := 0
	for i := idx+1; i < len(heights); i++ {
		if heights[i] > heights[idx] && heights[i] > res {
			res = heights[i]
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
