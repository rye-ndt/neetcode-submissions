func subsetsWithDup(nums []int) [][]int {
	result := [][]int{}
	sigs := map[string]bool{}

	var backtrack func(path []int, curIndex int)
	backtrack = func(path []int, curIndex int) {
		cur := append([]int{}, path...)
		sort.Ints(cur)
		sig := signize(cur)

		if found := sigs[sig]; !found {
			sigs[sig] = true
			result = append(result, cur)
		} else {
			return
		}

		for i := curIndex+1; i < len(nums); i++ {
			path = append(path, nums[i])
			backtrack(path, i)
			path = path[:len(path)-1]
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