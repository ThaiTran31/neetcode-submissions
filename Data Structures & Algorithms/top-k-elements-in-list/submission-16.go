func topKFrequent(nums []int, k int) []int {
    count := make(map[int]int)
    for _, num := range nums {
        count[num]++
    }
    freqNums := make([]int, 0)
    for num, _ := range count {
        freqNums = append(freqNums, num)
    }
    sort.Slice(freqNums, func(i, j int) bool {
        return count[freqNums[i]] > count[freqNums[j]]
    })
    return freqNums[:k]
}
