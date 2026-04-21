func hasDuplicate(nums []int) bool {
    numSet := make(map[int]bool)
	for _, num := range nums {
		if _, ok := numSet[num]; ok {
			return true
		}
		numSet[num] = true
	}
	return false
}
