func topKFrequent(nums []int, k int) []int {
	count := make(map[int]int)
	for _, num := range nums {
		count[num]++
	}
	orderedNums := make([]int, 0)
	for num, _ := range count {
		orderedNums = append(orderedNums, num)
	}
	sort.Slice(orderedNums, func(i, j int) bool {
		return count[orderedNums[i]] > count[orderedNums[j]]
	})
	return orderedNums[:k]
}
