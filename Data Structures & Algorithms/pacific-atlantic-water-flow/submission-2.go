func pacificAtlantic(h [][]int) [][]int {
	m, n := len(h), len(h[0])
	pac := make([][]bool, m)
	atl := make([][]bool, m)
	for i := range pac {
		pac[i] = make([]bool, n)
		atl[i] = make([]bool, n)
	}
	dirs := [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	var dfs func(r, c int, seen [][]bool)
	dfs = func(r, c int, seen [][]bool) {
		seen[r][c] = true
		for _, d := range dirs {
			nr, nc := r+d[0], c+d[1]
			if nr >= 0 && nr < m && nc >= 0 && nc < n &&
				!seen[nr][nc] && h[nr][nc] >= h[r][c] {
				dfs(nr, nc, seen)
			}
		}
	}

	for r := 0; r < m; r++ {
		dfs(r, 0, pac)
		dfs(r, n-1, atl)
	}
	for c := 0; c < n; c++ {
		dfs(0, c, pac)
		dfs(m-1, c, atl)
	}

	res := [][]int{}
	for r := 0; r < m; r++ {
		for c := 0; c < n; c++ {
			if pac[r][c] && atl[r][c] {
				res = append(res, []int{r, c})
			}
		}
	}
	return res
}