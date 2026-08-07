type Q struct {
	C int 
	R int
}

func orangesRotting(grid [][]int) int {
	time := 0
	rows := len(grid)
	cols := len(grid[0])
	fresh := 0
	rottenQ := []Q{}

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			cur := grid[r][c]

			if cur == 1 {
				fresh++ 
			}

			if cur == 2 {
				rottenQ = append(rottenQ, Q{R: r, C: c})
			}
		}
	}

	directions := [][]int{
		{0, 1}, {1, 0}, {0, -1}, {-1, 0},
	}

	for len(rottenQ) > 0 && fresh > 0 {
		time++
		roundLen := len(rottenQ) 

		for i := 0; i < roundLen; i++ {
			cur := rottenQ[0]
			rottenQ = rottenQ[1:]

			for _, dir := range directions {
				nextR := cur.R + dir[0]
				nextC := cur.C + dir[1]
				rValid := nextR >= 0 && nextR < rows 
				cValid := nextC >= 0 && nextC < cols 

				if rValid && cValid && grid[nextR][nextC] == 1 {
					fresh--
					grid[nextR][nextC] = 2
					rottenQ = append(rottenQ, Q{R: nextR, C: nextC})
				} 
			}
		}
	}

	if fresh > 0 {
		return -1
	}

	return time
}
