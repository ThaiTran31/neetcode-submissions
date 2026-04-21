func groupAnagrams(strs []string) [][]string {
	anaMap := make(map[string][]string)
	for _, str := range strs {
		stdAna := sortString(str)
		if anas, exist := anaMap[stdAna]; exist {
			anas = append(anas, str)
			anaMap[stdAna] = anas
		} else {
			anaMap[stdAna] = []string{str}
		}
	}
	res := make([][]string, 0)
	for _, v := range anaMap {
		res = append(res, v)
	}
	return res
}

func sortString (s string) string {
	r := []rune(s)
	sort.Slice(r, func(i, j int) bool { return r[i] < r[j] })
	return string(r)
}
