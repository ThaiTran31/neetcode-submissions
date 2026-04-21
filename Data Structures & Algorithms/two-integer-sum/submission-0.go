func twoSum(nums []int, target int) []int {
    res := make([]int, 2)
	for i, num := range nums {
		neededNum := target - num
		for j := i + 1; j < len(nums); j++ {
			if nums[j] == neededNum {
				res[0] = i
				res[1] = j
				return res
			}
		}
	}
	return res
}
