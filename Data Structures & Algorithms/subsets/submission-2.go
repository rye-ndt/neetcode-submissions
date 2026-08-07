func subsets(nums []int) [][]int {
	result := [][]int{}
	
	path := []int{}

	var backtrack func(path, choices []int) 

	backtrack = func(path, choices []int) {
		result = append(result, deepCopy(path))

		if len(choices) == 0 {
			return
		}

		for i, c := range choices {
			// try this one
			path = append(path, c)

			//recur
			backtrack(path, choices[i+1:])

			// try another one 
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
