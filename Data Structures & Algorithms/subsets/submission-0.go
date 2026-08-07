func subsets(nums []int) [][]int {
	result := [][]int{}
	path := []int{}

	var backtrack func(index int) 

	backtrack = func(index int) {
		// is complete 
		// every path is valid
		// record
		result = append(result, deepCopy(path))

		// choices 
		// all numbers after this index are valid
		for i := index; i < len(nums); i++ {
			path = append(path, nums[i])
			backtrack(i+1)
			path = path[:len(path)-1] // undo before continue
		}
	}

	backtrack(0)
	
	return result
}

func deepCopy(orig []int) []int {
	clone := make([]int, len(orig))
	copy(clone, orig)

	return clone
}
