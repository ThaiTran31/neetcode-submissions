func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int)
	for _, num := range nums {
		m[num]++
	}
	orderedNums := make([]int, 0)
	for num, _ := range m {
		orderedNums = append(orderedNums, num)
	}
	sort.Slice(orderedNums, func(i, j int) bool {
		return m[orderedNums[i]] > m[orderedNums[j]]
	})
	return orderedNums[:k]
}
