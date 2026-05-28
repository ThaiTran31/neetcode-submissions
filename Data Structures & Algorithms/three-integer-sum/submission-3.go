func threeSum(nums []int) [][]int {
	numCount := make(map[int]int)
	sort.Ints(nums)
	for _, num := range nums {
		numCount[num] += 1
	}
	res := [][]int{}
	for i := 0; i < len(nums); i++ {
		numCount[nums[i]]--
		if i > 0 && nums[i-1] == nums[i] {
			continue
		}
		for j := i+1; j < len(nums); j++ {
			numCount[nums[j]]--
			if j > i+1 && nums[j-1] == nums[j] {
				continue
			}
			compl := -(nums[i] + nums[j])
			if numCount[compl] > 0 {
				res = append(res, []int{nums[i], nums[j], compl})
			}
		}
		for j := i+1; j < len(nums); j++ {
			numCount[nums[j]]++
		}
	}
	return res
}
