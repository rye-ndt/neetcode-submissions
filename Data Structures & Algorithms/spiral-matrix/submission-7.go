func spiralOrder(matrix [][]int) []int {
	visited := map[[2]int]int{}

	valid := func(r, c int) bool {
		if r < 0 || c < 0 || r >= len(matrix) || c >= len(matrix[0]) { return false } 
		if _, ok := visited[[2]int{r, c}]; ok { return false }
		return true 
	}

	r, c, d := 0, 0, 0
	dirs := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

	for len(visited) < len(matrix) * len(matrix[0]) {
		visited[[2]int{r, c}] = len(visited)

		if !valid(r + dirs[d][0], c + dirs[d][1]) {
			d = (d + 1) % 4
		}

		r, c = r + dirs[d][0], c + dirs[d][1]
	}

	result := make([]int, len(visited)) 
	for k, v := range visited {
		result[v] = matrix[k[0]][k[1]]
	}
	return result 
}