func rotate(m [][]int) {
	l, r := 0, len(m)-1

	for l < r {
		m[l], m[r] = m[r], m[l]
		l++; r--
	}

	for i := 0; i < len(m); i++ {
		for j := i; j < len(m[0]); j++ {
			m[i][j], m[j][i] = m[j][i], m[i][j]
		}
	}
}