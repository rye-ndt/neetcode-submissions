func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	result := [][]int{}
	path := []int{}

	var backtrack func(path, choices []int) 

	backtrack = func(path, choices []int) {
		s := sum(path)

		if s == target {
			result = append(result, clone(path))
			return
		}

		if s > target {
			return 
		}

		for i, c := range choices {
			if i > 0 && choices[i] == choices[i-1] {
				continue
			}

			path = append(path, c)
			backtrack(path, choices[i+1:])
			path = path[:len(path)-1]
		}
	}

	backtrack(path, candidates)

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
