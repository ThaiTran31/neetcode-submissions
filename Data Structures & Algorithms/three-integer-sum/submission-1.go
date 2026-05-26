func threeSum(nums []int) [][]int {
	m := make(map[int][]int)
	for i, num := range nums {
		m[num] = append(m[num], i)
	}
	res := make([][]int, 0)
	tripletCount := make(map[[3]int]bool)
	for i, num := range nums {
		for j, num1 := range nums {
			if i == j {
				continue
			}
			compl := -num - num1
			idxes, ok := m[compl]
			if ok {
				for _, idx := range idxes {
					if idx != i && idx != j {
						triplet := [3]int{nums[i], nums[j], nums[idx]}
						sortedTripl := triplet[:]
						sort.Ints(sortedTripl)
						if _, ok := tripletCount[triplet]; !ok {
							tripletCount[triplet] = true
							res = append(res, sortedTripl)
						}
					}
				}
			}
		}
	}
	return res
}
