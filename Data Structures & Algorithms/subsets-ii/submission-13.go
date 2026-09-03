func subsetsWithDup(nums []int) [][]int {
	result := [][]int{}
	sigs := map[string]bool{}

	var backtrack func(path []int, curIndex int)
	backtrack = func(path []int, curIndex int) {
		cur := append([]int{}, path...)
		sort.Ints(cur)

		if sigs[signize(cur)] { return } 

		sigs[signize(cur)] = true
		result = append(result, cur)

		for i := curIndex+1; i < len(nums); i++ {
			backtrack(append(path, nums[i]), i)
		}
	}

	backtrack([]int{}, -1)

	return result
}

func signize(nums []int) string {
	total := ""
	for _, n := range nums { total += fmt.Sprint(n) }
	return total
}