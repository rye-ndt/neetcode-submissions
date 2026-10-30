func setZeroes(matrix [][]int) {
	firstRowZ, firstColZ := false, false 

	for i := 0; i < len(matrix[0]); i++ {
		if matrix[0][i] == 0 { 
			firstRowZ = true 
			break 
		}
	}

	for i := 0; i < len(matrix); i++ {
		if matrix[i][0] == 0 {
			firstColZ = true
			break 
		}
	}

	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}

	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[0]); j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	if firstRowZ {
		for j := 0; j < len(matrix[0]); j++ { matrix[0][j] = 0 }
	}

	if firstColZ {
		for i := 0; i < len(matrix); i++ { matrix[i][0] = 0 }
	}
}