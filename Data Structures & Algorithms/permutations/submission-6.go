func permute(nums []int) [][]int {
	sort.Ints(nums)
	result := [][]int{}

	var backtrack func(path []int) 
	backtrack = func(path []int) {
		if len(path) == len(nums) { 
			result = append(result, append([]int{}, path...))
			return 
		}

		for i := 0; i < len(nums); i++ {
			if getSig(path)[nums[i]] { continue }
			path = append(path, nums[i])
			backtrack(path)
			path = path[:len(path)-1]
		}
	}

	backtrack([]int{})

	return result
}

func getSig(s []int) map[int]bool {
	result := map[int]bool{}
	for _, n := range s {
		result[n] = true 
	}
	return result
}