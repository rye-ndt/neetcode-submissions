func numIslands(grid [][]byte) int {
    result := 0
	rows := len(grid)
	cols := len(grid[0])

	var dfs func(r, c int) 

	dfs = func(r, c int) {
		// to mark everything as visited
		rValid := r >= 0 && r < rows
		cValid := c >= 0 && c < cols 

		if !rValid || !cValid || grid[r][c] == '0' {
			return
		}

		grid[r][c] = '0'

		dfs(r+1, c)
		dfs(r, c+1)
		dfs(r-1, c)
		dfs(r, c-1)
	}

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			cur := grid[r][c]

			if cur == '1' {
				dfs(r, c)
				result++
			}
		}
	}

	return result
}
