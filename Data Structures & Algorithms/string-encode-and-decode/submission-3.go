type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	sizes := make([]string, 0)
	for _, str := range strs {
		sizes = append(sizes, fmt.Sprintf("%d", len(str)))
	}
	res := strings.Join(sizes, ",")
	res += "#"
	for _, str := range strs {
		res += str
	}
	return res
}

func (s *Solution) Decode(encoded string) []string {
	idx := strings.Index(encoded, "#")
	if idx == 0 {
		return []string{}
	}
	res := make([]string, 0)
	sizesStr := encoded[:idx]
	decodedStr := encoded[idx+1:]
	sizes := strings.Split(sizesStr, ",")
	accumSize := 0
	for _, sizeStr := range sizes {
		size, _ := strconv.Atoi(sizeStr)
		res = append(res, decodedStr[accumSize:accumSize+size])
		accumSize += size
	}
	return res
}
