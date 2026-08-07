func searchMatrix2(matrix [][]int, target int) bool {
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

func searchMatrix(m [][]int, target int) bool {
	for _, s := range m {
		if target > s[len(s)-1] {
			continue
		}

		l := 0
		r := len(s)-1

		if target == s[l] || target == s[r] {
			return true
		}

		for l <= r {
			m := (l+r)/2

			if target > s[m] {
				l = m+1
				continue
			}
			
			if target < s[m] {
				r = m-1
				continue
			}

			return true
		}
	}
	return false
}