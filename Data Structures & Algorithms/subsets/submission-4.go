func subsets(nums []int) [][]int {
	result := [][]int{}

	var backtrack func(path []int, index int)
	backtrack = func(path []int, index int) {
		result = append(result, append([]int{}, path...))

		for i := index; i < len(nums); i++ {
			path = append(path, nums[i])
			backtrack(path, i+1)
			path = path[:len(path)-1]
		}
	}

	backtrack([]int{}, 0)

	return result
}