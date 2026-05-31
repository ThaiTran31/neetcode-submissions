func maxArea(heights []int) int {
	l := 0
	r := len(heights) - 1
	res := 0
	for l < r {
		minHeight, isLeftLess := getMinHeightAndPos(heights[l], heights[r])
		cur := (r-l) * minHeight
		if cur > res {
			res = cur
		}
		if isLeftLess {
			l++
		} else {
			r--
		}
	}
	return res
}

func getMinHeightAndPos(h1, h2 int) (int, bool) {
	if h1 <= h2 {
		return h1, true
	}
	return h2, false
}
