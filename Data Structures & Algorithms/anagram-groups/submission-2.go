func groupAnagrams(strs []string) [][]string {
	m := make(map[[26]int][]string)
	for _, s := range strs {
		arr := [26]int{}
		for _, r := range s {
			arr[r-'a']++
		}
		m[arr] = append(m[arr], s)
	}
	res := make([][]string, 0)
	for _, v := range m {
		res = append(res, v)
	}
	return res
}
