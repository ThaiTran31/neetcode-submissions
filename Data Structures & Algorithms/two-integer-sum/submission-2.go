func twoSum(nums []int, target int) []int {
	res := make([]int, 2)
    m := toMap(nums)
	for i, num := range nums {
		neededNum := target - num
		if j, exist := m[neededNum]; exist && i != j {
			res[0] = i
			res[1] = j
			return res
		}
	}
	return res
}

func toMap(fromNums []int) map[int]int {
	res := make(map[int]int)
	for i, num := range fromNums {
		res[num] = i
	}
	return res
}
