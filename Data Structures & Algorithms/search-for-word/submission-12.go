func exist(board [][]byte, word string) bool {
	var backtrack func(path [][2]int) bool
	options := [][]int{ {0, 1}, {1, 0}, {0, -1}, {-1, 0} }

	backtrack = func(path [][2]int) bool {
		str := stringify(path, board)

		if str == word { return true }

		if str != word[:len(str)] || len(path) == len(word) { 
			return false 
		}

		x := path[len(path)-1][0]
		y := path[len(path)-1][1]

		for _, o := range options {
			newX, newY := x + o[0], y + o[1]
			validX := valid(newX, len(board[0]))
			validY := valid(newY, len(board))
			usedBefore := used(path, newX, newY)

			if validX && validY && !usedBefore && backtrack(append(path, [2]int{newX, newY})) {
				return true
			}
		}

		return false
	}

	for y := 0; y < len(board); y++ {
		for x := 0; x < len(board[y]); x++ {
			if board[y][x] != word[0] { continue }
			if backtrack([][2]int{{x, y}}) { return true }
		}
	}

	return false
}

func stringify(path [][2]int, board[][]byte) string {
	result := ""
	for _, p := range path {
		result += string(board[p[1]][p[0]])
	}
	return result
}

func used(useList [][2]int, x, y int) bool {
	for _, item := range useList {
		if item[0] == x && item[1] == y { return true }
	}

	return false
}

func valid(v int, upper int) bool {
	return v >= 0 && v < upper 
}