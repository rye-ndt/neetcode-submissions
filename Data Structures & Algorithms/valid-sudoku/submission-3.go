func isValidSudoku(board [][]byte) bool {
	// row -> vertical
	for r := 0; r < 9; r++ {
		row := board[r]

		seen := map[byte]bool{}

		for c := 0; c < 9; c++ {
			if seen[row[c]] {
				return false
			}

			if row[c] != '.' {
				seen[row[c]] = true
			}
		}
	}

	// col -> horizontal	
	for c := 0; c < 9; c++ {
		seen := map[byte]bool{}

		for r := 0; r < 9; r++ {
			item := board[r][c]

			if seen[item] {
				return false
			}

			if item != '.' {
				seen[item] = true
			}
		}
	}

	for verticalStart := 0; verticalStart < 6; verticalStart += 3 {
		for horizontalStart := 0; horizontalStart < 6; horizontalStart += 3 {
			seen := map[byte]bool{}

			for r := verticalStart; r < verticalStart+3; r++ {
				for c := horizontalStart; c < horizontalStart+3; c++ {
					if seen[board[r][c]] {
						return false
					}

					if board[r][c] != '.' {
						seen[board[r][c]] = true
					}
				}
			}
		}
	}

	return true
}
