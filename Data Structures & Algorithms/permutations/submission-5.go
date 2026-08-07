func permute(nums []int) [][]int {
	sort.Ints(nums)
	result := [][]int{}
	path := []int{}
	used := make([]bool, len(nums))

	var backtrack func() 

	backtrack = func() {
		if len(path) == len(nums) {
			result = append(result, clone(path))
			return
		}

		for i, c := range nums {
			if used[i] {
				continue
			}

			path = append(path, c)
			used[i] = true

			backtrack()

			path = path[:len(path)-1]
			used[i] = false
		}
	}

	backtrack()

	return result
}

func clone(a []int) []int {
	c := make([]int, len(a))

	copy(c, a)
	
	return c
}
