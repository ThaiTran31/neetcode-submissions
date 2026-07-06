func hasDuplicate(nums []int) bool {
    countSet := make(map[int]struct{})
	for _, num := range nums {
		countSet[num] = struct{}{}
	}
	return len(countSet) != len(nums)
}
