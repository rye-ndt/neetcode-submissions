func permute(nums []int) [][]int {
	sort.Ints(nums)
	result := [][]int{}
	path := []int{}
	store := map[int]bool{}

	var backtrack func(path, choices []int) 

	backtrack = func(path, choices []int) {
		if len(path) == len(nums) {
			result = append(result, clone(path))
			return
		}

		for _, c := range choices {
			path = append(path, c)

			store[c] = true

			newChoices := []int{}

			for i := 0; i < len(nums); i++ {
				if !store[nums[i]] {
					newChoices = append(newChoices, nums[i])
				}
			}

			backtrack(path, newChoices)

			path = path[:len(path)-1]

			store[c] = false
		}
	}

	backtrack(path, nums)

	return result
}

func clone(a []int) []int {
	c := make([]int, len(a))

	copy(c, a)
	
	return c
}
