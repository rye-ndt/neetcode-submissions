func rotate(matrix [][]int)  {
	// swap row to col 

	for r := 0; r < len(matrix); r++ {
		// c := r+1 is the trick
		for c := r+1; c < len(matrix); c++ {
			matrix[r][c], matrix[c][r] = matrix[c][r], matrix[r][c]
		}
	}

	fmt.Println("after first step: ", matrix)

	for r := 0; r < len(matrix); r++ {
		left, right := 0, len(matrix)-1

		for left < right {
			matrix[r][left], matrix[r][right] = matrix[r][right], matrix[r][left]
			left++
			right--
		}
	}

	fmt.Println("second step: ", matrix)
}
