func numIslands(grid [][]byte) int {
	result := 0
	directions := [][]int{ {0, 1}, {1, 0}, {0, -1}, {-1, 0} }

	var flip func(x, y int) bool
	flip = func(x, y int) bool {
		if grid[y][x] == '0' { return false }
		grid[y][x] = '0'

		for _, dir := range directions {
			xValid := x + dir[0] >= 0 && x + dir[0] < len(grid[0])
			yValid := y + dir[1] >= 0 && y + dir[1] < len(grid)

			if xValid && yValid { flip(x + dir[0], y + dir[1]) }
		}

		return true
	}

	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if flip(x, y) { result++ }
		}
	}

	return result 
}