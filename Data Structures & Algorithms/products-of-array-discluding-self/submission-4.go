func productExceptSelf(nums []int) []int {
    prefixProducts := make([]int, len(nums))
	product := 1
	for i, _ := range nums {
		if i-1 < 0 {
			product = 1
		} else {
			product *= nums[i-1]
		}
		prefixProducts[i] = product
	}
	postfixProducts := make([]int, len(nums))
	for i := len(nums)-1; i >= 0; i-- {
		if i+1 == len(nums) {
			product = 1
		} else {
			product *= nums[i+1]
		}
		postfixProducts[i] = product
	}
	res := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		res[i] = prefixProducts[i] * postfixProducts[i]
	}
	return res
}