func subsets(nums []int) [][]int {
	result := [][]int{}

	path := []int{}

	var backtrack func(path, choices []int)

	backtrack = func(path, choices []int) {
		// all choices are valid 
		result = append(result, deepCopy(path))

		if len(choices) == 0 {
			return
		}

		for i, c := range choices {
			path = append(path, c)
			backtrack(path, choices[i+1:])
			path = path[:len(path)-1]
		}
	}

	backtrack(path, nums)

	return result
}

func deepCopy(orig []int) []int {
	clone := make([]int, len(orig))
	copy(clone, orig)

	return clone
}
