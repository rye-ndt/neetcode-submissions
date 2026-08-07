func permute(nums []int) [][]int {
	sort.Ints(nums)
	result := [][]int{}
	path := []int{}

	var backtrack func(path, choices []int) 

	backtrack = func(path, choices []int) {
		if len(path) == len(nums) {
			result = append(result, clone(path))
			return
		}

		for _, c := range choices {
			path = append(path, c)

			newChoices := []int{}

			for i := 0; i < len(nums); i++ {
				if indexOf(path, nums[i]) == -1 {
					newChoices = append(newChoices, nums[i])
				}
			}

			backtrack(path, newChoices)

			path = path[:len(path)-1]
		}
	}

	backtrack(path, nums)

	return result
}

func indexOf(arr []int, val int) int {
	result := -1

	for i, v := range arr {
		if v == val {
			result = i
		}
	}
	return result
}

func clone(a []int) []int {
	c := make([]int, len(a))

	copy(c, a)
	
	return c
}
