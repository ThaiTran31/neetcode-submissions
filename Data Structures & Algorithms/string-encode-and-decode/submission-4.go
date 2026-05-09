type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	if (len(strs) == 0) {
		return ""
	}
	sizes := make([]string, 0)
	for _, str := range strs {
		sizes = append(sizes, fmt.Sprintf("%d", len(str)))
	}

	return strings.Join(sizes, ",") + "😉" + strings.Join(strs, "")
}

func (s *Solution) Decode(encoded string) []string {
	if (encoded == "") {
		return []string{}
	}
	parts := strings.Split(encoded, "😉")
	strSizes := strings.Split(parts[0], ",")
	i := 0
	res := []string{}
	for _, strSize := range strSizes {
		size, _ := strconv.Atoi(strSize)
		res = append(res, parts[1][i:i+size])
		i += size
	}
	return res
}
