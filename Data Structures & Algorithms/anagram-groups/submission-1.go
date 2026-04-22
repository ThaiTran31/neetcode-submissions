func groupAnagrams(strs []string) [][]string {
	anaMap := make(map[[26]int][]string)
	for _, str := range strs {
		arr := alphaStr2Arr(str)
		if anaStrs, exist := anaMap[arr]; exist {
			anaStrs = append(anaStrs, str)
			anaMap[arr] = anaStrs
		} else {
			anaMap[arr] = []string{str}
		}
	}
	res := make([][]string, 0)
	for _, v := range anaMap {
		res = append(res, v)
	}
	return res
}

func alphaStr2Arr(s string) [26]int {
	var res [26]int
	for _, r := range s {
		res[int(r) % 26] += 1
	}
	return res
}
