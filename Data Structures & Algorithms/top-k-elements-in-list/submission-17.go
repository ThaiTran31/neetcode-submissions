func topKFrequent(nums []int, k int) []int {
	count := make(map[int]int)
	for _, num := range nums {
		count[num]++
	}

	freq := make([][]int, len(nums)+1)
	for num, f := range count {
		freq[f] = append(freq[f], num)
	}

	res := make([]int, 0)
	for i := len(freq)-1; i >= 0; i-- {
		for _, num := range freq[i] {
			if len(res) < k {
				res = append(res, num)
			}
			if len(res) == k {
				return res
			}
		}
	}

	return []int{}
}
