func twoSum(nums []int, target int) []int {
    numsWrapper := make([][2]int, 0)
	for i, num := range nums {
		numsWrapper = append(numsWrapper, [2]int{num, i})
	}

	sort.Slice(numsWrapper, func(i, j int) bool {
		return numsWrapper[i][0] < numsWrapper[j][0]
	})

	i, j := 0, len(nums)-1
	for i < j {
		curSum := numsWrapper[i][0] + numsWrapper[j][0]
		if curSum == target {
			if numsWrapper[i][1] < numsWrapper[j][1] {
				return []int{numsWrapper[i][1], numsWrapper[j][1]}
			} else {
				return []int{numsWrapper[j][1], numsWrapper[i][1]}
			}
		} else if curSum > target {
			j--
		} else {
			i++
		}
	}
	return []int{}
}
