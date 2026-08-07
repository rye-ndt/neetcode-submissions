func productExceptSelf(nums []int) []int {
	// special case: more than 1 number of 0 -> everything is 0, since there is always a 0 in the array
	// special case 2: 1 number 0 -> except for the position of the 0, everything else is 0 
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

	for _, num := range nums {
		appendNumber := total
		if numOfZero > 1 || (numOfZero == 1 && num != 0) {
			appendNumber = 0
		} else if numOfZero == 0 {
			appendNumber = total / num
		}

		result = append(result, appendNumber)
	}

	return result
}
