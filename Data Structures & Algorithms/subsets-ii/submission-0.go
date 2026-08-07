func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	result := [][]int{}
	path := []int{}

	var backtrack func(path, choices []int)

	backtrack = func(path, choices []int) {
		result = append(result, clone(path))

		for i, c := range choices {
			if i != 0 && choices[i] == choices[i-1] {
				continue
			}

			path = append(path, c)
			backtrack(path, choices[i+1:])
			path = path[:len(path)-1]
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