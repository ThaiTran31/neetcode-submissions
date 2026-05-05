func topKFrequent(nums []int, k int) []int {
	count := make(map[int]int)
	for _, num := range nums {
		count[num]++
	}

	buckets := make([][]int, len(nums) + 1)
	for num, freq := range count {
		buckets[freq] = append(buckets[freq], num)
	}
	
	res := make([]int, 0)
	i := len(buckets) - 1
	for len(res) < k {
		candidate := buckets[i]
		if candidate != nil {
			res = append(res, candidate...)
		}
		i--
	}

	return res
}
