func isValidSudoku(board [][]byte) bool {
	m := make(map[byte]bool)
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}
			if _, exist := m[board[i][j]]; exist {
				return false
			}
			m[board[i][j]] = true
		}
		m = make(map[byte]bool)
	}
	m = make(map[byte]bool)
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[j][i] == '.' {
				continue
			}
			if _, exist := m[board[j][i]]; exist {
				return false
			}
			m[board[j][i]] = true
		}
		m = make(map[byte]bool)
	}
	m = make(map[byte]bool)
	for i := 0; i <= 6; i+=3 {
		for j := 0; j <= 6; j+=3 {
			for i1 := i; i1 < i+3; i1++ {
				for j1 := j; j1 < j+3; j1++ {
					if board[i1][j1] == '.' {
						continue
					}
					if _, exist := m[board[i1][j1]]; exist {
						return false
					}
					m[board[i1][j1]] = true
				}
			}
			m = make(map[byte]bool)
		}
	}
	return true
}
