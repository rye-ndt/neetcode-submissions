func productExceptSelf(nums []int) []int {
	total := 1
	zeroCount := 0

	for _, n := range nums {
		if n != 0 {
			total *= n
		} else {
			zeroCount++
		}
	}

	result := []int{}
	fmt.Println("processing ", total, zeroCount)

	for _, n := range nums {
		val := 0

		if zeroCount == 0 {
			val = total / n
		} else if zeroCount == 1 && n == 0 {
			val = total
		}

		result = append(result, val)
	}

	return result
}
