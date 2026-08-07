func numIslands(grid [][]byte) int {
    result := 0
	rows := len(grid)
	cols := len(grid[0])

	var dfs func(r, c int)

	dfs = func(r, c int) {
		rValid := r >= 0 && r < rows
		cValid := c >= 0 && c < cols 

		if !rValid || !cValid || grid[r][c] == '0' {
			return
		}

		grid[r][c] = '0' // exclude it

		dfs(r+1, c)
		dfs(r, c+1)
		dfs(r-1, c)
		dfs(r, c-1)
	}

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == '1' {
				result+=1
				dfs(r, c)
			}
		}
	}

	return result 
}
