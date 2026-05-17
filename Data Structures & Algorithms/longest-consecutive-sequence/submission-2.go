func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	res := 1
	curMax := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] - nums[i-1] == 0 {
			continue
		} else if nums[i] - nums[i-1] == 1 {
			curMax++
		} else {
			if curMax > res {
				res = curMax
			}
			curMax = 1
		}
	}
	if curMax > res {
		res = curMax
	}

	return res
}
