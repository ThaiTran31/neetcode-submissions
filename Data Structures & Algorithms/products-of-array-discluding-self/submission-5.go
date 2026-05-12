func productExceptSelf(nums []int) []int {
	res := make([]int, len(nums))
	for i := range nums {
		product := 1
		for j := 0; j < len(nums); j++ {
			if i == j {
				continue
			}
			product *= nums[j]
		}
		res[i] = product
	}
	fmt.Println(res)
	return res
}
