func combinationSum(nums []int, target int) [][]int {
    result := [][]int{}

	var backtrack func(path []int, index int)
	backtrack = func(path []int, index int) {
		if sum(path) == target {
			result = append(result, append([]int{}, path...))
			return
		}	

		if sum(path) > target { return }

		for i := index; i < len(nums); i++ {
			path = append(path, nums[i])
			backtrack(path, i)
			path = path[:len(path)-1]
		}
	}

	backtrack([]int{}, 0)

	return result
}

func sum(v []int) int {
	result := 0
	for _, a := range v {
		result += a
	}
	return result
}