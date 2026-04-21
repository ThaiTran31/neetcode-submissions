func twoSum(nums []int, target int) []int {
	mNums := make(map[int]int)
	for i, num := range nums {
		complement := target - num
		if j, exist := mNums[complement]; exist {
			return []int{j, i}
		} else {
			mNums[num] = i
		}
	}
	return []int{}
}