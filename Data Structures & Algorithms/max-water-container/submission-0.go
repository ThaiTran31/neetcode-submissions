func maxArea(heights []int) int {
	res := 0
	for i, h1 := range heights {
		for j := i+1; j < len(heights); j++ {
			curMax := getMinHeight(h1, heights[j]) * (j-i)
			if curMax > res {
				res = curMax
			}
		}
	}
	return res
}

func getMinHeight(h1, h2 int) int {
	if h1 <= h2 {
		return h1
	}
	return h2
}
