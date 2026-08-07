func maxAreaOfIsland(grid [][]int) int {
   result := 0
   rows := len(grid)
   cols := len(grid[0])

   var dfs func(r, c int) int

   dfs = func(r, c int) int {
		rValid := r >= 0 && r < rows
		cValid := c >= 0 && c < cols

		if !rValid || !cValid || grid[r][c] == 0 {
			return 0
		}

		grid[r][c] = 0

		return 1 + dfs(r+1,c) + dfs(r-1,c) + dfs(r,c+1) + dfs(r,c-1)
   }

   for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == 1 {
				if temp := dfs(r, c); temp > result {
					result = temp
				}
			}
		}
   }

   return result
}
