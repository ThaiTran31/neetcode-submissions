func productExceptSelf(nums []int) []int {
	zeroCount := 0
	for _, num := range nums {
		if num == 0 {
			zeroCount++
		}
		if zeroCount == 2 {
			break
		}
	}
	res := make([]int, len(nums))
	if zeroCount == 2 {
		return res
	} else if zeroCount == 1 {
		product := 1
		zeroIdx := 0
		for i, num := range nums {
			if num != 0 {
				product *= num
			} else {
				zeroIdx = i
			}
		}
		res[zeroIdx] = product
		return res
	} else {
		product := 1
		for _, num := range nums {
			product *= num
		}
		for i, num := range nums {
			res[i] = product / num
		}
		return res
	}
}
