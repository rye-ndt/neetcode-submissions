func exist2(board [][]byte, word string) bool {
	rows, cols := len(board), len(board[0])

	var backtrack func(r, c, index int) bool 

	backtrack = func(r, c, index int) bool {
		if index == len(word) {
			return true
		}

		if r < 0 || r >= rows || c < 0 || c >= cols || board[r][c] != word[index] {
			return false
		}

		temp := board[r][c]
		board[r][c] = '#'
		next := index+1

		found := backtrack(r+1, c, next) ||		
				 backtrack(r-1, c, next) || 
				 backtrack(r, c+1, next) ||
				 backtrack(r, c-1, next)

		board[r][c] = temp

		return found
	}

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if backtrack(r, c, 0) {
				return true
			}
		}
	}

	return false 
}

func exist(board [][]byte, word string) bool {
	rows := len(board)
	cols := len(board[0])

	var backtrack func(row, col, index int) bool

	backtrack = func(r, c, i int) bool {
		if i == len(word) {
			return true
		}

		rValid := r >= 0 && r < rows
		cValid := c >= 0 && c < cols

		if !rValid || !cValid || board[r][c] != word[i] {
			return false
		}

		nextIndex := i+1
		temp := board[r][c]
		board[r][c] = '#' // so it cannot be used anymore

		if  backtrack(r+1, c, nextIndex) || 
			backtrack(r, c+1, nextIndex) ||
			backtrack(r-1, c, nextIndex) || 
			backtrack(r, c-1, nextIndex) {
				return true
		}

		board[r][c] = temp
		
		return false
	}

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if backtrack(r, c, 0) {
				return true
			}// reset the index
		}
	}

	return false
}