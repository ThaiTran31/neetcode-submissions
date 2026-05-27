func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	res := make([][]int, 0)
	resCount := make(map[[3]int]bool)
	for idx, num := range nums {
		// two sum with sorted arry and target = -num
		i := 0
		j := len(nums) - 1
		for i < j {
			sum := nums[i] + nums[j]
			if sum > -num {
				j--
			} else if sum < -num {
				i++
			} else {
				if i != idx &&
					j != idx &&
					nums[i] + nums[j] == -num {
					triplet := [3]int{num, nums[i], nums[j]}
					sort.Ints(triplet[:])
					if _, ok := resCount[triplet]; !ok {
						res = append(res,triplet[:])
						resCount[triplet] = true
					}
				}
				i++
				j--
			}
		}
	}
	return res
}
