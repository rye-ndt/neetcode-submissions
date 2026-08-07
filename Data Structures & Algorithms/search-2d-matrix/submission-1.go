func searchMatrix(matrix [][]int, target int) bool {
	// find the row first 
	// then use binary search on that row

	rowIndex := -1

	for index, r := range matrix {
		if r[0] <= target && r[len(r) - 1] >= target {
			rowIndex = index
		}
	}

	fmt.Println("row: ", rowIndex)
	if rowIndex == -1 {
		return false
	}

	row := matrix[rowIndex]

	l := 0
	r := len(row) - 1

	for l <= r {
		m := l + (r - l) / 2

		if row[m] == target || row[l] == target || row[r] == target {
			return true
		}

		if row[l] < target {
			l = m+1
			continue
		}

		if row[r] > target {
			r = m-1
			continue
		}
	}

	return false
}
