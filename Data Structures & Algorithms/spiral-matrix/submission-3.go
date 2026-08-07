// right, down, left, up
func spiralOrder(matrix [][]int) []int {
	seen := map[[2]int]bool{}

	result := []int{}
	
	xy := []int{0, 0}

	rows, cols := len(matrix), len(matrix[0])

	directions := [][]int{
		{1, 0}, {0, 1}, {-1, 0}, {0, -1},
	}

	d := 0 

	for len(result) < rows * cols {
		r := xy[1]
		c := xy[0]

		result = append(result, matrix[r][c])
		seen[[2]int{r, c}] = true

		nextR := r + directions[d][1]
		nextC := c + directions[d][0]

		nextRValid := nextR >= 0 && nextR < rows 
		nextCValid := nextC >= 0 && nextC < cols 

		if !nextRValid || !nextCValid || seen[[2]int{nextR, nextC}] {
			d = (d+1) % 4

			nextR = r + directions[d][1]
			nextC = c + directions[d][0]
		}

		xy = []int{nextC, nextR}
	}

	return result
}
