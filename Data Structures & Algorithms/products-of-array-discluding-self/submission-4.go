func productExceptSelf(nums []int) []int {
	// special case: more than 1 number of 0 
	// special case 2: 1 number 0
	// base case: no 0

	numOfZero := 0
	total := 1
	result := []int{}

	for _, num := range nums {
		if num == 0 {
			numOfZero++
			continue
		}

		total *= num 
	}

	if numOfZero > 1 {
		for range nums {
			result = append(result, 0)
		}

		return result
	}

	if numOfZero == 1 {
		for _, num := range nums {
			if num == 0 {
				result = append(result, total)
				continue
			}

			result = append(result, 0)
		}

		return result
	}

	for _, num := range nums {
		result = append(result, total / num)
	}

	return result
}
