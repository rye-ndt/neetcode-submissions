func searchMatrix(matrix [][]int, target int) bool {
	matrixLen := len(matrix)
	rowLen := len(matrix[0])

	l := 0
	r := matrixLen * rowLen - 1

	for l <= r {
		m := l + (r - l) / 2

		row := m / rowLen 
		col := m % rowLen 

		cur := matrix[row][col]

		if cur == target {
			return true
		}

		if cur > target {
			r = m-1
			continue
		}

		if cur < target {
			l = m+1
			continue
		}
	}

	return false
}
