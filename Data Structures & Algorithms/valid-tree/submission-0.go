func validTree(n int, edges [][]int) bool {
	if len(edges) != n-1 {
		return false
	}
	
	mapper := make([][]int, n)

	for _, e := range edges {
		a, b := e[0], e[1]

		mapper[a] = append(mapper[a], b)
		mapper[b] = append(mapper[b], a)
	}

	freeze := make([]bool, n)

	var dfs func(cur int) int
	dfs = func(cur int) int {
		if freeze[cur] {
			return 0
		}

		freeze[cur] = true

		thisCycleReach := 1 // current node 
		for _, c := range mapper[cur] {
			thisCycleReach += dfs(c)
		}

		return thisCycleReach
	}

	return dfs(0) == n
}
