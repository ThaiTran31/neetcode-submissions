func twoSum(numbers []int, target int) []int {
	for i, num := range numbers {
		compl := target - num
		// binary search for compl
		l, r := 0, len(numbers)-1
		for l <= r {
			mid := (l + r) / 2
			val := numbers[mid]
			if val == compl && mid != i {
				return []int{i+1, mid+1}
			} else if val > compl {
				r = mid - 1
			} else {
				l = mid + 1
			}
		}
	}
	return []int{}
}
