func combinationSum(nums []int, target int) [][]int {
    result := [][]int{}
	path := []int{}

	var backtrack func(path, choices []int) 

	backtrack = func(path, choices []int) {
		
		if sum(path) == target {
			result = append(result, clone(path))
			return
		}

		if sum(path) > target {
			return
		}

		for i, n := range choices {
			path = append(path, n)
			backtrack(path, choices[i:])
			path = path[:len(path)-1]
		}
	}

	backtrack(path, nums)

	return result
}

func sum(nums []int) int {
	total := 0

	for _, n := range nums {
		total += n
	}

	return total
}

func clone(a []int) []int {
	c := make([]int, len(a))

	copy(c, a)

	return c
}