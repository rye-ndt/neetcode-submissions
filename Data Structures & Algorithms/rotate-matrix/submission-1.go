func rotate(matrix [][]int) {
	l, r := 0, len(matrix)-1

	for l < r {
		matrix[l], matrix[r] = matrix[r], matrix[l]
		l++
		r--
	}

	for i := 0; i < len(matrix); i++ {
		for j := i; j < len(matrix[i]); j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
}

/*
[1,2,3],
[4,5,6],
[7,8,9]

[7,8,9]
[4,5,6],
[1,2,3],

[7,4,1]
[8,5,2],
[9,6,3],
*/