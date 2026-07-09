func twoSum(nums []int, target int) []int {
    val2Idx := make(map[int]int)
	for i, num := range nums {
		val2Idx[num] = i
	}
	for i, num := range nums {
		j, ok := val2Idx[target-num]
		if ok && i != j {
			return []int{i, j}
		}
	}
	return []int{}
}
