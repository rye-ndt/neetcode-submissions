func isValidSudoku(board [][]byte) bool {
	// check row valid
	for _, row := range board {
		seen := map[byte]bool{}

		for _, num := range row {
			cur := num

			if cur == '.' {
				continue
			}

			if _, found := seen[num]; found {
				return false
			}

			seen[num] = true
		}
	}

	// check column valid 
	for i := 0; i < len(board[0]); i++ {
		seen := map[byte]bool{}

		for j := 0; j < len(board); j++ {
			cur := board[j][i]

			if cur == '.' {
				continue
			}

			if _, found := seen[cur]; found {
				return false
			}

			seen[cur] = true
		}
	}

	// check box valid 
	for boxRow := 0; boxRow < 9; boxRow+=3 {
		for boxCol := 0; boxCol < 9; boxCol+=3 {
			seen := map[byte]bool{}

			for i := boxRow; i < boxRow+3; i++ {
				for j := boxCol; j < boxCol+3; j++ {
					cur := board[i][j]

					if cur == '.' {
						continue
					}

					if _, found := seen[board[i][j]]; found {
						return false
					}

					seen[board[i][j]] = true
				}
			} 
		}
	}
	
	return true
}
