func findOrder(n int, pre [][]int) []int {
	result := []int{}
	mapper := make([][]int, n)

	for _, p := range pre {
		a, b := p[0], p[1]

		mapper[a] = append(mapper[a], b)
	}

	freeze := make([]bool, n)
	done := make([]bool, n)

	var dfs func(cur int) bool 
	dfs = func(cur int) bool {
		if freeze[cur] {
			return false
		}

		if done[cur] {
			return true
		}

		freeze[cur] = true

		for _, p := range mapper[cur] {
			if !dfs(p) {
				return false
			}
		}

		freeze[cur] = false
		result = append(result, cur)
		done[cur] = true

		return true
	}

	for i := 0; i < n; i++ {
		if !dfs(i) {
			return []int{}
		}
	}

	return result
}
