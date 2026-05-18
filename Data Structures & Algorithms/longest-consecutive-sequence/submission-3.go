func longestConsecutive(nums []int) int {
	res := 0
	m := make(map[int]bool)
	for _, num := range nums {
		m[num] = true
	}
	for _, num := range nums {
		if _, exist := m[num-1]; exist {
			continue
		}
		curMax := 1
		next := num + 1
		for {
			_, exist := m[next]
			if !exist {
				break
			}
			curMax++
			next++
		}
		if curMax > res {
			res = curMax
		}
	}
	return res
}
