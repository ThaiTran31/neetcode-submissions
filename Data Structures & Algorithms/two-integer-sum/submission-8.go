func twoSum(nums []int, target int) []int {
	val2Idx := make(map[int]int)
	for i, num := range nums {
		compl := target - num
		if j, ok := val2Idx[compl]; ok {
			return []int{j, i}
		} else {
			val2Idx[num] = i
		}
	}
	return []int{}
}
