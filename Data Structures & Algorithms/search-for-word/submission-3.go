func exist(board [][]byte, word string) bool {
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
