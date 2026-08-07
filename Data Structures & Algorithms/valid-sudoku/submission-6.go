const blockSize = 3
const boardSize = 9

func isValidSudoku(board [][]byte) bool {
	for r := 0; r < boardSize; r++ {
		row := board[r]

		seen := map[byte]bool{}

		for c := 0; c < boardSize; c++ {
			if seen[row[c]] {
				return false
			}

			if row[c] != '.' {
				seen[row[c]] = true
			}
		}
	}

	for c := 0; c < boardSize; c++ {
		seen := map[byte]bool{}

		for r := 0; r < boardSize; r++ {
			item := board[r][c]

			if seen[item] {
				return false
			}

			if item != '.' {
				seen[item] = true
			}
		}
	}

	for rBlock := 0; rBlock < boardSize - blockSize; rBlock += blockSize {
		for cBlock := 0; cBlock < boardSize - blockSize; cBlock += blockSize {
			seen := map[byte]bool{}

			for r := rBlock; r < rBlock + blockSize; r++ {
				for c := cBlock; c < cBlock + blockSize; c++ {
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
