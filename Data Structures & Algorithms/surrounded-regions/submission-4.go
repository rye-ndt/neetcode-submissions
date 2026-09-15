func solve(board [][]byte) {
	dirs := [][2]int{ {1, 0}, {0, 1}, {-1, 0}, {0,-1} }
	reachOutside := map[[2]int]bool{}
	os := map[[2]int]bool{}

	var dfs func(x, y int)
	dfs = func(x, y int) {
		for _, d := range dirs {
			newX, newY := x + d[0], y + d[1]
			xValid := newX >= 0 && newX < len(board[0])-1
			yValid := newY >= 0 && newY < len(board)-1

			if !xValid || !yValid || reachOutside[[2]int{newX, newY}]{ continue }
			if board[newY][newX] == 'X' { continue }
			reachOutside[[2]int{newX, newY}] = true
			dfs(newX, newY)
		}
	}

	for x := 0; x < len(board[0]); x++ {
		for y := 0; y < len(board); y++ {
			if board[y][x] == 'X' { continue }
			if x == 0 || x == len(board[0]) - 1 || y == 0 || y == len(board)-1 {
				reachOutside[[2]int{x, y}] = true
				dfs(x, y)
			} else { 
				os[[2]int{x, y}] = true
 			}
		}
	}

	for k, _ := range os {
		if reachOutside[[2]int{k[0], k[1]}] { continue }
		board[k[1]][k[0]] = 'X'
	}
}