// cycle detection 

func canFinish(n int, pre [][]int) bool {
	mapper := make([][]int, n)

	for _, p := range pre {
		a, b := p[0], p[1]

		mapper[a] = append(mapper[a], b)
	}

	freeze := make([]bool, n)

	var dfs func(cur int) bool
	dfs = func(cur int) bool {
		if freeze[cur] {
			return false
		} 

		freeze[cur] = true

		for _, d := range mapper[cur] {
			if !dfs(d) {
				return false
			}
		}

		freeze[cur] = false

		return true
	}

	for i := 0; i < n; i++ {
		if !dfs(i) {
			return false
		}
	}

	return true
}
