type Q struct {
	R int
	C int
}

func solve(board [][]byte) {
    rows := len(board)
	cols := len(board[0])
	q := []Q{}

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if board[r][c] == 'O' {
				q = append(q, Q{R: r, C: c})
			}
		}
	}

	directions := [][]int{
		{0, 1}, {1, 0}, {0, -1}, {-1, 0},
	}

	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		visited := map[Q]bool{} //key exists if node visited; val true if node is safe; else val false
		edgeFound := false
		
		if cur.R == 0 || cur.C == 0 {
			continue
		}

		var visit func(r, c int) 

		visit = func(r, c int) {
			rValid := r >= 0 && r < rows 
			cValid := c >= 0 && c < cols

			if !rValid || !cValid {
				return 
			}

			node := Q{R: r, C: c}

			if board[r][c] == 'X' || visited[node] {
				return
			}

 			if r == 0 || c == 0 || r == rows-1 || c == cols-1 {
				edgeFound = true
				return
			}

			visited[node] = true

			for _, dir := range directions {
				visit(node.R + dir[0], node.C + dir[1])
			}
		}

		visit(cur.R, cur.C)

		if !edgeFound {
			for node := range visited {
				board[node.R][node.C] = 'X'
			}
		}
	}
}
