func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	s1Count := make([]int, 26, 26)
	s2Count := make([]int, 26, 26)
	for i := 0; i < len(s1); i++ {
		s1Count[s1[i]-'a']++
		s2Count[s2[i]-'a']++
	}

	matches := 0
	for i := 0; i < 26; i++ {
		if s1Count[i] == s2Count[i] {
			matches++
		}
	}

	l := 0
	for r := len(s1); r < len(s2); r++ {
		if matches == 26 {
			return true
		}

		removedPos := s2[l] - 'a'
		s2Count[removedPos]--
		if s2Count[removedPos] == s1Count[removedPos] {
			matches++
		} else if s1Count[removedPos]-1 == s2Count[removedPos] {
			matches--
		}

		addedPos := s2[r] - 'a'
		s2Count[addedPos]++
		if s2Count[addedPos] == s1Count[addedPos] {
			matches++
		} else if s1Count[addedPos]+1 == s2Count[addedPos] {
			matches--
		}

		l++
	}

	return matches == 26
}
