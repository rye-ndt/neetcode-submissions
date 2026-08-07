// tree with n nodes = acyclic + n-1 edges. this is the rule

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

	// start from node 0; it must go through all nodes
	visited := make([]bool, n)

	var dfs func(cur int) int
	dfs = func(cur int) int {
		if visited[cur] {
			return 0
		}

		visited[cur] = true

		counter := 1 // the current 

		for _, p := range mapper[cur] {
			counter += dfs(p)
		}

		return counter
	}

	return dfs(0) == n
}
